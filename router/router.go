package router

import (
	"backendtourapp/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {
	api := app.Group("/tour")

	api.Get("/", handler.GetAllPaketWisata)
	api.Get("/paket:id", handler.GetPaketWisataByID)        
	api.Post("/paket", handler.InsertPaketWisata)        
	api.Put("/paket:id", handler.UpdatePaketWisata)         
	api.Delete("/paket:id", handler.DeletePaketWisata) 

	//Ulasan
	api.Get("/ulasan", handler.GetAllUlasan)
	api.Get("/ulasan:id", handler.GetUlasanByID)        
	api.Post("/ulasan", handler.InsertUlasan)        
	api.Put("/ulasan:id", handler.UpdateUlasan)         
	api.Delete("/ulasan:id", handler.DeleteUlasan)

	//Pemesanan
	api.Get("/pemesanan", handler.GetAllPemesanan)
	api.Get("/pemesanan:id", handler.GetPemesananByID)        
	api.Post("/pemesanan", handler.InsertPemesanan)        
	api.Put("/pemesanan:id", handler.UpdatePemesanan)         
	api.Delete("/pemesanan:id", handler.DeletePemesanan) 
}
