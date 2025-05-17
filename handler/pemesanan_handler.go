package handler

import (
	"backendtourapp/model"
	"backendtourapp/repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GET all pemesanan
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

// GET pemesanan by kode pemesan
func GetPemesananByKode(c *fiber.Ctx) error {
	kode := c.Params("kode")
	data, err := repository.GetPemesananByKode(c.Context(), kode)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data pemesanan tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data pemesanan ditemukan",
		"data":    data,
	})
}


// POST / insert pemesanan
func InsertPemesanan(c *fiber.Ctx) error {
	var pemesanan model.Pemesanan
	if err := c.BodyParser(&pemesanan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data pemesanan tidak valid",
		})
	}

	id, err := repository.InsertPemesanan(c.Context(), pemesanan)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal menambahkan pemesanan: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Pemesanan berhasil ditambahkan",
		"id":      id,
	})
}

// PUT / update pemesanan by ID
func UpdatePemesanan(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID tidak valid",
		})
	}

	var update model.Pemesanan
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format data pemesanan tidak valid",
		})
	}

	updatedID, err := repository.UpdatePemesanan(c.Context(), objectID, update)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Pemesanan berhasil diperbarui",
		"id":      updatedID,
	})
}


// DELETE / delete pemesanan by ID
func DeletePemesanan(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID tidak valid",
		})
	}

	deletedID, err := repository.DeletePemesanan(c.Context(), objectID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Pemesanan dengan ID %s tidak ditemukan: %v", idParam, err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Pemesanan berhasil dihapus",
		"id":      deletedID,
	})
}
