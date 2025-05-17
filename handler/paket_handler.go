package handler

import (
	"backendtourapp/model"
	"backendtourapp/repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GET all paket wisata
func GetAllPaketWisata(c *fiber.Ctx) error {
	data, err := repository.GetAllPaketWisata(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data paket wisata",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data paket wisata",
		"data":    data,
	})
}

// GET by kode_paket
func GetPaketWisataByKode(c *fiber.Ctx) error {
	kode := c.Params("kode")
	data, err := repository.GetPaketWisataByKode(c.Context(), kode)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data ditemukan",
		"data":    data,
	})
}

// POST / insert data
func InsertPaketWisata(c *fiber.Ctx) error {
	var paket model.PaketWisata
	if err := c.BodyParser(&paket); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data tidak valid",
		})
	}

	id, err := repository.InsertPaketWisata(c.Context(), paket)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal menambahkan paket: %v", err),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Paket wisata berhasil ditambahkan",
		"id":      id,
	})
}

// PUT / update data
func UpdatePaketWisata(c *fiber.Ctx) error {
	kode := c.Params("kode")
	var update model.PaketWisata
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format data tidak valid",
		})
	}

	kodeUpdated, err := repository.UpdatePaketWisata(c.Context(), kode, update)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data berhasil diperbarui",
		"kode":    kodeUpdated,
	})
}

// DELETE / hapus data
func DeletePaketWisata(c *fiber.Ctx) error {
	kode := c.Params("kode")
	kodeDeleted, err := repository.DeletePaketWisata(c.Context(), kode)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Data dengan kode %s tidak ditemukan: %v", kode, err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data berhasil dihapus",
		"kode":    kodeDeleted,
	})
}
