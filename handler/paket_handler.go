package handler

import (
	"backendtourapp/model"
	"backendtourapp/repository"
	"fmt"
	"time"

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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal mengambil data: %v", err),
			"status":  fiber.StatusInternalServerError,
		})
	}

	// Tambahan pengecekan kalau data tidak ditemukan (paket == nil)
	if paket == nil {
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

	// Validasi ID harus diisi
	if paket.ID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID paket wisata harus diisi",
		})
	}

	// Validasi tanggal mulai tidak boleh di masa lalu
	if paket.TanggalMulai.Before(time.Now().Truncate(24 * time.Hour)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Tanggal mulai harus hari ini atau setelahnya",
		})
	}

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

	// Validasi tanggal mulai tidak boleh di masa lalu
	if paket.TanggalMulai.Before(time.Now().Truncate(24 * time.Hour)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Tanggal mulai harus hari ini atau setelahnya",
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
