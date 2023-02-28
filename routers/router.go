package routers

import (
	"busapp/handlers"

	"github.com/couchbase/gocb/v2"
	"github.com/gofiber/fiber/v2"
)

func Initalize(app *fiber.App, DB *gocb.Cluster) {
	h := handlers.New(DB)

	public := app.Group("/api")
	public.Post("/register", h.Register)
	public.Post("/login", h.Login)
	public.Get("/buy-ticket", h.BuyTicketPage)
	public.Post("/buy-ticket", h.BuyTicket)

	busApp := app.Group("/api")
	// Access level = 1
	busApp.Post("/bus", h.CreateBus)
	busApp.Put("/bus", h.UpdateBus)
	busApp.Post("/voyage", h.CreateVoyage)
	busApp.Delete("/voyage/*", h.DeleteVoyage)

	// Aceess level <= 2
	busApp.Get("/bus-definition", h.BusDefinition)
	busApp.Get("/model/*", h.GetBusModels)
	busApp.Get("/voyage", h.GetAllVoyages)

}
