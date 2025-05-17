package router

import (
	"backendtourapp/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {
	api := app.Group("/paket")

	// api.Get("/", handler.Homepage)
	api.Get("/", handler.GetAllPaketWisata)
	// api.Get("/mahasiswa", handler.GetAllMahasiswa)
	// api.Get("/mahasiswa/:npm", handler.GetMahasiswaByNPM)
	// api.Post("/mahasiswa", handler.InsertMahasiswa)
	// api.Put("/mahasiswa/:npm", handler.UpdateMahasiswa)
	// api.Delete("/mahasiswa/:npm", handler.DeleteMahasiswa)
}
