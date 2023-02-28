package handlers

import (
	"busapp/config"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/couchbase/gocb/v2"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"

	"github.com/golang-jwt/jwt"
)

func (h handler) Register(c *fiber.Ctx) error {
	type RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json := new(RegisterRequest)
	if err := c.BodyParser(json); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			`error`:   err.Error(),
		})
	}

	hash := generateHash(json.Password)
	userId := guuid.New().String()

	col := h.DB.Bucket("busmgmt").Collection("users")
	var new = User{
		ID:          userId,
		Username:    json.Username,
		Hash:        hash,
		AccessLevel: "3",
	}

	_, err := col.Insert(userId, new, nil)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Error while creation",
		})
	}

	getResult, err := col.Get(userId, nil)
	if err != nil {
		log.Fatal(err)
	}

	var inUser User
	err = getResult.Content(&inUser)
	if err != nil {
		log.Fatal(err)
	}

	token, _ := generateJWT(inUser.Username, "3")

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "User successfuly created",
		"User ID": inUser.ID,
		"Token":   token,
	})
}

func (h handler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json := new(LoginRequest)
	if err := c.BodyParser(json); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	queryParams := make(map[string]interface{}, 1)
	queryParams["username"] = json.Username

	query := "SELECT * FROM `busmgmt`.`_default`.`users` WHERE username=$username;"
	rows, err := h.DB.Query(query, &gocb.QueryOptions{NamedParameters: queryParams})
	if err != nil {
		fmt.Println(err)
	}

	type Row struct {
		User User `json:"users"`
	}

	for rows.Next() {
		var row Row
		if err = rows.Row(&row); err != nil {
			fmt.Println(err)
		}
		if row.User.Hash == generateHash(json.Password) {
			token, _ := generateJWT(row.User.ID, row.User.AccessLevel)

			return c.JSON(fiber.Map{
				"Data":    token,
				"Status":  "success",
				"Message": "Login Succesfully",
			})

		}
	}

	return c.JSON("Fail")
}

func generateHash(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func generateJWT(userId string, accessLevel string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":      userId,
		"accessLevel": accessLevel,
	})
	secretKey := []byte(config.Config("JWT_SECRET_FOR_LOCAL"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString, err
}
