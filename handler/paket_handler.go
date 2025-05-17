package handler

import (
	"backendtourapp/repository"

	"github.com/gofiber/fiber/v2"
)

func GetAllPaket(c *fiber.Ctx) error {
	data, err := repository.GetAllPaket(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Gagal mengambil data paket wisata",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,                      
		"message": "Berhasil mengambil data paket wisata", 
		"data":    data,                                
	})
}