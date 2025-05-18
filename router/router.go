package router

import (
	"backendtourapp/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", handler.Homepage)

	// Destinasi routes
api.Get("/destinasi", handler.GetAllDestinasi)
api.Get("/destinasi/:kode", handler.GetDestinasiByKode)
api.Get("/destinasi/:id", handler.GetDestinasiByID)
api.Post("/destinasi", handler.InsertDestinasi)
api.Put("/destinasi/:id", handler.UpdateDestinasi)
api.Delete("/destinasi/:id", handler.DeleteDestinasi)

api.Get("/paket/destinasi", handler.GetPaketWithDestinasi)

	// Paket wisata routes
	api.Get("/paket", handler.GetAllPaketWisata)
	api.Get("/paket/:kode", handler.GetPaketWisataByKode)
	api.Post("/paket", handler.InsertPaketWisata)
	api.Put("/paket/:kode", handler.UpdatePaketWisata)
	api.Delete("/paket/:kode", handler.DeletePaketWisata)

	// Ulasan routes
	api.Get("/ulasan", handler.GetAllUlasan)
	api.Get("/ulasan/kodepaket/:kode", handler.GetUlasanByKodePaket) // lebih spesifik
	api.Post("/ulasan", handler.InsertUlasan)
	api.Put("/ulasan/:id", handler.UpdateUlasan)
	api.Delete("/ulasan/:id", handler.DeleteUlasan)

	// Pemesanan routes
	api.Get("/pemesanan", handler.GetAllPemesanan)
	api.Get("/pemesanan/kode/:kode", handler.GetPemesananByKode)
	api.Post("/pemesanan", handler.InsertPemesanan)
	api.Put("/pemesanan/:id", handler.UpdatePemesanan)
	api.Delete("/pemesanan/:id", handler.DeletePemesanan)
}
