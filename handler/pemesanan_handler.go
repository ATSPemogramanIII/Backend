package handler

import (
	"backendtourapp/model"
	"backendtourapp/repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetAllPemesanan(c *fiber.Ctx) error {
	data, err := repository.GetAllPemesanan(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data pemesanan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data pemesanan",
		"data":    data,
	})
}

func GetPemesananByID(c *fiber.Ctx) error {
	id := c.Params("id")
	pemesanan, err := repository.GetPemesananByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Pemesanan tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Pemesanan ditemukan",
		"data":    pemesanan,
	})
}

func InsertPemesanan(c *fiber.Ctx) error {
	var pemesanan model.Pemesanan
	if err := c.BodyParser(&pemesanan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	insertedID, err := repository.InsertPemesanan(c.Context(), pemesanan)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Gagal menambahkan pemesanan: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Pemesanan berhasil ditambahkan",
		"id":      insertedID,
	})
}

func UpdatePemesanan(c *fiber.Ctx) error {
	id := c.Params("id")
	var pemesanan model.Pemesanan
	if err := c.BodyParser(&pemesanan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	updatedID, err := repository.UpdatePemesanan(c.Context(), id, pemesanan)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Gagal mengupdate pemesanan: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Pemesanan berhasil diupdate",
		"id":      updatedID,
	})
}

func DeletePemesanan(c *fiber.Ctx) error {
	id := c.Params("id")
	deletedID, err := repository.DeletePemesanan(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Gagal menghapus pemesanan: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Pemesanan berhasil dihapus",
		"id":      deletedID,
	})
}
