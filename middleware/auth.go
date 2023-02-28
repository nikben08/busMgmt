package middleware

import (
	"busapp/config"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
)

func AccessCheck(userId int, accessLevel int, r *http.Request) bool {
	method := r.Method
	p := r.URL.Path

	if strings.Contains(p, "bus-definition") && method == "PUT" && accessLevel > 1 {
		return false
	}

	if strings.Contains(p, "bus") && (method == "POST" || method == "PUT") && accessLevel > 1 {
		return false
	}

	if strings.Contains(p, "voyage") && method == "POST" && accessLevel > 1 {
		return false
	}

	if (strings.Contains(p, "voyage") || strings.Contains(p, "model") || strings.Contains(p, "bus-definition")) && method == "GET" && accessLevel > 2 {
		return false
	}

	return true
}

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Token")
		secretKey := []byte(config.Config("JWT_SECRET_FOR_LOCAL"))
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, json.NewEncoder(w).Encode("Invalid token")
			}
			return secretKey, nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId, _ := strconv.Atoi(fmt.Sprint(claims["userId"]))
			accessLevel, _ := strconv.Atoi(fmt.Sprint(claims["accessLevel"]))
			access := AccessCheck(userId, accessLevel, r)
			fmt.Println("-----------")
			fmt.Println(claims["userId"])
			if !access {
				json.NewEncoder(w).Encode("Access denied")
			} else {
				next.ServeHTTP(w, r)
			}
		} else {
			json.NewEncoder(w).Encode("Invalid token")
		}
	})
}
