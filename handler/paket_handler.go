package handler

import (
	"backendtourapp/model"
	"backendtourapp/repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetAllPaketWisata(c *fiber.Ctx) error {
	data, err := repository.GetAllPaketWisata(c.Context())
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

func GetPaketWisataByID(c *fiber.Ctx) error {
	id := c.Params("id")

	paket, err := repository.GetPaketWisataByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Paket wisata tidak ditemukan",
			"status":  fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Paket wisata ditemukan",
		"data":    paket,
		"status":  fiber.StatusOK,
	})
}

func InsertPaketWisata(c *fiber.Ctx) error {
	var paket model.PaketWisata

	if err := c.BodyParser(&paket); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	// Validasi ID unik bisa ditambahkan di repository
	insertedID, err := repository.InsertPaketWisata(c.Context(), paket)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": fmt.Sprintf("Gagal menambahkan paket wisata: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Paket wisata berhasil ditambahkan",
		"id":      insertedID,
		"status":  fiber.StatusCreated,
	})
}

func UpdatePaketWisata(c *fiber.Ctx) error {
	id := c.Params("id")
	var paket model.PaketWisata

	if err := c.BodyParser(&paket); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	updatedID, err := repository.UpdatePaketWisata(c.Context(), id, paket)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Gagal mengupdate data: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Paket wisata berhasil diupdate",
		"id":      updatedID,
		"status":  fiber.StatusOK,
	})
}

func DeletePaketWisata(c *fiber.Ctx) error {
	id := c.Params("id")

	deletedID, err := repository.DeletePaketWisata(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Gagal menghapus paket wisata: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Paket wisata berhasil dihapus",
		"id":      deletedID,
		"status":  fiber.StatusOK,
	})
}
