package handlers

import (
	"fmt"

	"github.com/couchbase/gocb/v2"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
)

func (h handler) BusDefinition(c *fiber.Ctx) error {
	type Reponse struct {
		Brands     []Brand
		Types      []Type
		Properties []Property
	}

	var response Reponse

	// Brands
	brandsQuery := "SELECT * FROM `busmgmt`.`_default`.`brands`;"
	bransRows, _ := h.DB.Query(brandsQuery, &gocb.QueryOptions{})

	type BrandRow struct {
		Brand Brand `json:"brands"`
	}

	for bransRows.Next() {
		var row BrandRow
		if err := bransRows.Row(&row); err != nil {
			fmt.Println(err)
		}
		response.Brands = append(response.Brands, row.Brand)
	}

	// Types
	typesQuery := "SELECT * FROM `busmgmt`.`_default`.`types`;"
	typesRows, _ := h.DB.Query(typesQuery, &gocb.QueryOptions{})

	type TypeRow struct {
		Type Type `json:"types"`
	}

	for typesRows.Next() {
		var row TypeRow
		if err := typesRows.Row(&row); err != nil {
			fmt.Println(err)
		}
		response.Types = append(response.Types, row.Type)
	}

	// Properties
	propertiesQuery := "SELECT * FROM `busmgmt`.`_default`.`properties`;"
	propertiesRows, _ := h.DB.Query(propertiesQuery, &gocb.QueryOptions{})

	type PropertyRow struct {
		Property Property `json:"properties"`
	}

	for propertiesRows.Next() {
		var row PropertyRow
		if err := propertiesRows.Row(&row); err != nil {
			fmt.Println(err)
		}
		response.Properties = append(response.Properties, row.Property)

	}
	return c.JSON(response)
}

func (h handler) GetBusModels(c *fiber.Ctx) error {
	brandId := c.Params("*1")
	queryParams := make(map[string]interface{}, 1)
	queryParams["brand_id"] = brandId
	query := "SELECT * FROM `busmgmt`.`_default`.`bus_models` WHERE brand_id = $brand_id;"
	busModelsrows, err := h.DB.Query(query, &gocb.QueryOptions{NamedParameters: queryParams})
	if err != nil {
		fmt.Println(err)
	}

	type Response struct {
		Model []BusModel
	}
	var response Response

	type BusModelRow struct {
		BusModel BusModel `json:"bus_models"`
	}

	for busModelsrows.Next() {
		var row BusModelRow
		if err := busModelsrows.Row(&row); err != nil {
			fmt.Println(err)
		}
		response.Model = append(response.Model, row.BusModel)
	}

	return c.JSON(response)
}

func (h handler) CreateBus(c *fiber.Ctx) error {
	type PropertyRequest struct {
		ID string `json:"id"`
	}

	type CreateBusRequest struct {
		PlateNumber   string            `json:"plate_number"`
		Model         string            `json:"model"`
		NumberOfSeats int               `json:"number_of_seats"`
		Type          string            `json:"type"`
		Properties    []PropertyRequest `json:"properties"`
	}

	json := new(CreateBusRequest)
	if err := c.BodyParser(json); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			`error`:   err.Error(),
		})
	}

	col := h.DB.Bucket("busmgmt").Collection("buses")
	busId := guuid.New().String()
	var new = Bus{
		ID:            busId,
		PlateNumber:   json.PlateNumber,
		NumberOfSeats: json.NumberOfSeats,
		Model:         json.Model,
		Type:          json.Type,
	}

	_, err := col.Insert(busId, new, nil)
	if err != nil {
		fmt.Println(err)
	}

	bus_properties_col := h.DB.Bucket("busmgmt").Collection("bus_properties")

	for i := 0; i < len(json.Properties); i++ {
		var bus_property = BusProperty{
			BusID: busId,
			ID:    json.Properties[i].ID,
		}

		_, err := bus_properties_col.Insert(guuid.New().String(), bus_property, nil)
		if err != nil {
			fmt.Println(err)
		}
	}

	return c.JSON(200)
}

func (h handler) UpdateBus(c *fiber.Ctx) error {
	type PropertyRequest struct {
		ID string `json:"id"`
	}

	type UpdateBusRequest struct {
		PlateNumber   string            `json:"plate_number"`
		ModelID       string            `json:"model_id"`
		NumberOfSeats int               `json:"number_of_seats"`
		TypeID        string            `json:"type_id"`
		Properties    []PropertyRequest `json:"properties"`
	}

	json := new(UpdateBusRequest)
	if err := c.BodyParser(json); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			`error`:   err.Error(),
		})
	}

	getBusQueryParams := make(map[string]interface{}, 1)
	getBusQueryParams["plate_number"] = json.PlateNumber

	query := "SELECT * FROM `busmgmt`.`_default`.`buses` WHERE plate_number=$plate_number;"

	busesRows, err := h.DB.Query(query, &gocb.QueryOptions{NamedParameters: getBusQueryParams})
	if err != nil {
		fmt.Println(err)
	}

	type BusRow struct {
		Bus Bus `json:"buses"`
	}

	var busID string

	for busesRows.Next() {
		var row BusRow
		if err := busesRows.Row(&row); err != nil {
			fmt.Println(err)
		}
		busID = row.Bus.ID
	}

	updtBusQueryParams := make(map[string]interface{}, 1)
	updtBusQueryParams["bus_id"] = busID
	updtBusQueryParams["plate_number"] = json.PlateNumber
	updtBusQueryParams["model_id"] = json.ModelID
	updtBusQueryParams["number_of_seats"] = json.NumberOfSeats
	updtBusQueryParams["type_id"] = json.TypeID

	query = "UPDATE `busmgmt`.`_default`.`buses`" +
		"SET plate_number=$plate_number, model_id=$model_id, number_of_seats=$number_of_seats, type_id=$type_id" +
		" WHERE id=$bus_id;"

	_, err = h.DB.Query(query, &gocb.QueryOptions{NamedParameters: updtBusQueryParams})
	if err != nil {
		fmt.Println(err)
	}

	query = "DELETE FROM  `busmgmt`.`_default`.`bus_properties` WHERE bus_id=$bus_id;"
	_, err = h.DB.Query(query, &gocb.QueryOptions{NamedParameters: updtBusQueryParams})
	if err != nil {
		fmt.Println(err)
	}

	bus_properties_col := h.DB.Bucket("busmgmt").Collection("bus_properties")
	for i := 0; i < len(json.Properties); i++ {
		var bus_property = BusProperty{
			BusID: busID,
			ID:    json.Properties[i].ID,
		}

		_, err := bus_properties_col.Insert(guuid.New().String(), bus_property, nil)
		if err != nil {
			fmt.Println(err)
		}
	}

	return c.JSON(200)
}
