package router

import (
	"backendtourapp/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", handler.Homepage)
	
	// Paket wisata routes
	api.Get("/paket", handler.GetAllPaketWisata)
	api.Get("/paket/:kode", handler.GetPaketWisataByKode)
	api.Post("/paket", handler.InsertPaketWisata)
	api.Put("/paket/:kode", handler.UpdatePaketWisata)
	api.Delete("/paket/:kode", handler.DeletePaketWisata)

	// Ulasan routes
	api.Get("/ulasan", handler.GetAllUlasan)
	api.Get("/ulasan/:id", handler.GetUlasanByKodePaket)
	api.Post("/ulasan", handler.InsertUlasan)
	api.Put("/ulasan/:id", handler.UpdateUlasan)
	api.Delete("/ulasan/:id", handler.DeleteUlasan)

	// Pemesanan routes
	api.Get("/pemesanan", handler.GetAllPemesanan)
	api.Get("/pemesanan/email/:email", handler.GetPemesananByEmail)
	api.Post("/pemesanan", handler.InsertPemesanan)
	api.Put("/pemesanan/:id", handler.UpdatePemesanan)
	api.Delete("/pemesanan/:id", handler.DeletePemesanan)
}
