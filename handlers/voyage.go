package handlers

import (
	"fmt"
	"log"

	"github.com/couchbase/gocb/v2"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
)

func (h handler) CreateVoyage(c *fiber.Ctx) error {
	type CreateVoyageRequest struct {
		Fee   float32 `json:"fee"`
		From  string  `json:"from"`
		To    string  `json:"to"`
		Date  string  `json:"date"`
		BusID string  `json:"bus_id"`
	}

	json := new(CreateVoyageRequest)
	if err := c.BodyParser(json); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			`error`:   err.Error(),
		})
	}

	col := h.DB.Bucket("busmgmt").Collection("voyages")
	voyageID := guuid.New().String()
	var new = Voyage{
		ID:    voyageID,
		Fee:   json.Fee,
		From:  json.From,
		To:    json.To,
		Date:  json.Date,
		BusID: json.BusID,
	}

	_, err := col.Insert(voyageID, new, nil)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(200)
}

func (h handler) DeleteVoyage(c *fiber.Ctx) error {
	voyageID := c.Params("*1")
	deleteVoyageQueryParams := make(map[string]interface{}, 1)
	deleteVoyageQueryParams["voyage_id"] = voyageID
	query := "DELETE FROM `busmgmt`.`_default`.`voyages` WHERE id=$voyage_id"
	_, err := h.DB.Query(query, &gocb.QueryOptions{NamedParameters: deleteVoyageQueryParams})
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(200)
}

func (h handler) GetVoyage(c *fiber.Ctx) error {
	voyageId := c.Params("*1")

	queryParams := make(map[string]interface{}, 1)
	queryParams["voyage_id"] = voyageId

	getVoyage, err := h.DB.Bucket("busmgmt").Collection("voyages").Get(voyageId, nil)
	if err != nil {
		log.Fatal(err)
	}
	var voyage Voyage
	err = getVoyage.Content(&voyage)
	if err != nil {
		log.Fatal(err)
	}

	getBus, err := h.DB.Bucket("busmgmt").Collection("buses").Get(voyage.BusID, nil)
	if err != nil {
		log.Fatal(err)
	}
	var bus Bus
	err = getBus.Content(&bus)
	if err != nil {
		log.Fatal(err)
	}

	var provinces []Province

	getFromProvince, err := h.DB.Bucket("busmgmt").Collection("provinces").Get(voyage.From, nil)
	if err != nil {
		log.Fatal(err)
	}

	var province Province
	err = getFromProvince.Content(&province)
	if err != nil {
		log.Fatal(err)
	}

	provinces = append(provinces, province)

	getToProvince, err := h.DB.Bucket("busmgmt").Collection("provinces").Get(voyage.To, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = getToProvince.Content(&province)
	if err != nil {
		log.Fatal(err)
	}

	provinces = append(provinces, province)

	type Response struct {
		Bus      Bus
		Province []Province
		Voyage   Voyage
	}

	var response = Response{
		Bus:      bus,
		Province: provinces,
		Voyage:   voyage,
	}
	return c.JSON(response)
}

func (h handler) GetAllVoyages(c *fiber.Ctx) error {
	query := "SELECT * FROM `busmgmt`.`_default`.`buses`"
	busRows, err := h.DB.Query(query, &gocb.QueryOptions{})
	if err != nil {
		fmt.Println(err)
	}

	var buses []Bus

	type BusRow struct {
		Bus Bus `json:"buses"`
	}

	for busRows.Next() {
		var row BusRow
		if err := busRows.Row(&row); err != nil {
			fmt.Println(err)
		}
		buses = append(buses, row.Bus)
	}

	query = "SELECT * FROM `busmgmt`.`_default`.`provinces`"
	provinceRows, err := h.DB.Query(query, &gocb.QueryOptions{})
	if err != nil {
		fmt.Println(err)
	}

	var provinces []Province

	type ProvinceRow struct {
		Province Province `json:"provinces"`
	}

	for provinceRows.Next() {

		var row ProvinceRow
		if err := provinceRows.Row(&row); err != nil {
			fmt.Println(err)
		}
		provinces = append(provinces, row.Province)
	}

	query = "SELECT * FROM `busmgmt`.`_default`.`voyages`"
	voyageRows, err := h.DB.Query(query, &gocb.QueryOptions{})
	if err != nil {
		fmt.Println(err)
	}

	var voyages []VoyageResponse

	type VoyageRow struct {
		Voyage VoyageResponse `json:"voyages"`
	}

	for voyageRows.Next() {
		var row VoyageRow
		if err := voyageRows.Row(&row); err != nil {
			fmt.Println(err)
		}
		voyages = append(voyages, row.Voyage)
	}

	for i := 0; i < len(voyages); i++ {
		getBus, err := h.DB.Bucket("busmgmt").Collection("buses").Get(voyages[i].BusID, nil)
		if err != nil {
			log.Fatal(err)
		}
		err = getBus.Content(&voyages[i].Bus)
		if err != nil {
			log.Fatal(err)
		}
	}

	return c.JSON(fiber.Map{
		"bus":      buses,
		"province": provinces,
		"voyage":   voyages,
	})

}
