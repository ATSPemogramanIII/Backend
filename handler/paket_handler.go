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

	// Validasi ID
	if paket.ID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID paket wisata harus diisi",
		})
	}

	// Validasi dan parsing tanggal
	parsedDate, err := time.Parse(time.RFC3339, paket.TanggalMulai)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Format tanggal tidak valid. Gunakan format ISO 8601 (contoh: 2025-06-01T00:00:00Z)",
		})
	}

	if parsedDate.Before(time.Now().Truncate(24 * time.Hour)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Tanggal mulai harus hari ini atau setelahnya",
		})
	}

	// Simpan ke database
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

	// Parse body ke struct PaketWisata
	if err := c.BodyParser(&paket); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	// Validasi tanggal mulai (harus format ISO dan tidak boleh di masa lalu)
	parsedDate, err := time.Parse(time.RFC3339, paket.TanggalMulai)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Format tanggal tidak valid. Gunakan format ISO 8601 (contoh: 2025-06-01T00:00:00Z)",
		})
	}

	if parsedDate.Before(time.Now().Truncate(24 * time.Hour)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Tanggal mulai harus hari ini atau setelahnya",
		})
	}

	// Lanjut update ke database
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
