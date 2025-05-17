package handler

import (
	"backendtourapp/model"
	"backendtourapp/repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GET all ulasan
func GetAllUlasan(c *fiber.Ctx) error {
	data, err := repository.GetAllUlasan(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data ulasan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data ulasan",
		"data":    data,
	})
}

// GET ulasan by kode_paket
func GetUlasanByKodePaket(c *fiber.Ctx) error {
	kode := c.Params("kode")
	data, err := repository.GetUlasanByKodePaket(c.Context(), kode)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data ulasan tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data ulasan ditemukan",
		"data":    data,
	})
}

// POST / insert ulasan
func InsertUlasan(c *fiber.Ctx) error {
	var ulasan model.Ulasan
	if err := c.BodyParser(&ulasan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data ulasan tidak valid",
		})
	}

	id, err := repository.InsertUlasan(c.Context(), ulasan)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal menambahkan ulasan: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Ulasan berhasil ditambahkan",
		"id":      id,
	})
}

// PUT / update ulasan by ID
func UpdateUlasan(c *fiber.Ctx) error {
	id := c.Params("id")

	var update model.Ulasan
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format data ulasan tidak valid",
		})
	}

	updatedID, err := repository.UpdateUlasan(c.Context(), id, update)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ulasan berhasil diperbarui",
		"id":      updatedID,
	})
}


// DELETE / delete ulasan by ID
func DeleteUlasan(c *fiber.Ctx) error {
	id := c.Params("id")
	deletedID, err := repository.DeleteUlasan(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Ulasan dengan ID %s tidak ditemukan: %v", id, err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ulasan berhasil dihapus",
		"id":      deletedID,
	})
}
