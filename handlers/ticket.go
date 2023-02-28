package handlers

import (
	"busapp/models"
	"fmt"
	"log"

	"github.com/couchbase/gocb/v2"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
)

func (h handler) BuyTicket(c *fiber.Ctx) error {
	type BuyTicketRequest struct {
		No       int    `json:"no"`
		Sex      int    `json:"sex"`
		VoyageID string `json:"voyage_id"`
	}

	json := new(BuyTicketRequest)
	if err := c.BodyParser(json); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			`error`:   err.Error(),
		})
	}

	var voyage Voyage

	getBus, err := h.DB.Bucket("busmgmt").Collection("voyages").Get(json.VoyageID, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = getBus.Content(&voyage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(voyage.BusID)
	col := h.DB.Bucket("busmgmt").Collection("bus_seats")
	busSeatID := guuid.New().String()
	var bus_seat = BusSeat{
		ID:       busSeatID,
		BusID:    voyage.BusID,
		VoyageID: json.VoyageID,
		SeatNo:   json.No,
		Sex:      json.Sex,
	}

	col.Insert(busSeatID, bus_seat, nil)

	return c.JSON(200)
}

func (h handler) BuyTicketPage(c *fiber.Ctx) error {
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

		// Get Props
		getPropsQueryParams := make(map[string]interface{}, 1)
		getPropsQueryParams["bus_id"] = voyages[i].BusID
		query = "SELECT * FROM `busmgmt`.`_default`.`bus_properties` WHERE bus_id=$bus_id"
		busPropertyRows, err := h.DB.Query(query, &gocb.QueryOptions{NamedParameters: getPropsQueryParams})
		if err != nil {
			fmt.Println(err)
		}

		for busPropertyRows.Next() {
			type BusPropertyRow struct {
				BusProperty BusProperty `json:"bus_properties"`
			}
			var row BusPropertyRow
			if err := busPropertyRows.Row(&row); err != nil {
				fmt.Println(err)
			}
			var a models.BusProperty = models.BusProperty{ID: row.BusProperty.ID, BusID: row.BusProperty.BusID}
			voyages[i].Bus.Property = append(voyages[i].Bus.Property, a)
		}

		// Get Seats

		getSeatsQueryParams := make(map[string]interface{}, 1)
		getSeatsQueryParams["bus_id"] = voyages[i].BusID
		query = "SELECT * FROM `busmgmt`.`_default`.`bus_seats` WHERE bus_id=$bus_id"
		busSeatRows, err := h.DB.Query(query, &gocb.QueryOptions{NamedParameters: getSeatsQueryParams})
		if err != nil {
			fmt.Println(err)
		}

		for busSeatRows.Next() {
			type BusSeatRow struct {
				BusSeat BusSeat `json:"bus_seats"`
			}
			var row BusSeatRow
			if err := busSeatRows.Row(&row); err != nil {
				fmt.Println(err)
			}

			var a models.BusSeat = models.BusSeat{ID: row.BusSeat.ID, BusID: row.BusSeat.BusID, VoyageID: row.BusSeat.VoyageID, SeatNo: row.BusSeat.SeatNo, Sex: row.BusSeat.Sex}
			voyages[i].Bus.Seat = append(voyages[i].Bus.Seat, a)
		}

	}

	return c.JSON(fiber.Map{
		"bus":      buses,
		"province": provinces,
		"voyage":   voyages,
	})
}
