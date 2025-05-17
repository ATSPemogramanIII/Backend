package handler

import (
	"backendtourapp/model"
	"backendtourapp/repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

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

func GetUlasanByID(c *fiber.Ctx) error {
	id := c.Params("id")
	ulasan, err := repository.GetUlasanByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Ulasan tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ulasan ditemukan",
		"data":    ulasan,
	})
}

func InsertUlasan(c *fiber.Ctx) error {
	var ulasan model.Ulasan
	if err := c.BodyParser(&ulasan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	insertedID, err := repository.InsertUlasan(c.Context(), ulasan)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Gagal menambahkan ulasan: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Ulasan berhasil ditambahkan",
		"id":      insertedID,
	})
}

func UpdateUlasan(c *fiber.Ctx) error {
	id := c.Params("id")
	var ulasan model.Ulasan
	if err := c.BodyParser(&ulasan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	updatedID, err := repository.UpdateUlasan(c.Context(), id, ulasan)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Gagal mengupdate ulasan: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ulasan berhasil diupdate",
		"id":      updatedID,
	})
}

func DeleteUlasan(c *fiber.Ctx) error {
	id := c.Params("id")
	deletedID, err := repository.DeleteUlasan(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Gagal menghapus ulasan: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ulasan berhasil dihapus",
		"id":      deletedID,
	})
}
