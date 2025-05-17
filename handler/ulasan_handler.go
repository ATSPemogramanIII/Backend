package handler

import (
	"backendtourapp/model"
	"backendtourapp/repository"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllUlasan(c *fiber.Ctx) error {
	data, err := repository.GetAllUlasan(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data ulasan",
			"status":  fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data ulasan",
		"data":    data,
		"status":  fiber.StatusOK,
	})
}

func GetUlasanByID(c *fiber.Ctx) error {
	id := c.Params("id")
	ulasan, err := repository.GetUlasanByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fmt.Sprintf("Gagal mengambil ulasan: %v", err),
			"status":  fiber.StatusInternalServerError,
		})
	}
	if ulasan == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Ulasan tidak ditemukan",
			"status":  fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ulasan ditemukan",
		"data":    ulasan,
		"status":  fiber.StatusOK,
	})
}

func InsertUlasan(c *fiber.Ctx) error {
	var ulasan model.Ulasan
	if err := c.BodyParser(&ulasan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Data tidak valid",
			"status": fiber.StatusBadRequest,
		})
	}

	// Validasi rating
	if ulasan.Rating < 1 || ulasan.Rating > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Rating harus antara 1 sampai 5",
			"status": fiber.StatusBadRequest,
		})
	}

	// Validasi required fields
	if ulasan.IDPaket == "" || ulasan.NamaPengguna == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "IDPaket dan NamaPengguna tidak boleh kosong",
			"status": fiber.StatusBadRequest,
		})
	}

	// Set tanggal sekarang secara otomatis
	ulasan.Tanggal = time.Now()

	insertedID, err := repository.InsertUlasan(c.Context(), ulasan)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  fmt.Sprintf("Gagal menambahkan ulasan: %v", err),
			"status": fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Ulasan berhasil ditambahkan",
		"id":      insertedID,
		"status":  fiber.StatusCreated,
	})
}

func UpdateUlasan(c *fiber.Ctx) error {
	id := c.Params("id")
	var ulasan model.Ulasan
	if err := c.BodyParser(&ulasan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Data tidak valid",
			"status": fiber.StatusBadRequest,
		})
	}

	// Validasi rating
	if ulasan.Rating < 1 || ulasan.Rating > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Rating harus antara 1 sampai 5",
			"status": fiber.StatusBadRequest,
		})
	}

	// Validasi required fields
	if ulasan.IDPaket == "" || ulasan.NamaPengguna == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "IDPaket dan NamaPengguna tidak boleh kosong",
			"status": fiber.StatusBadRequest,
		})
	}

	updatedID, err := repository.UpdateUlasan(c.Context(), id, ulasan)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  fmt.Sprintf("Gagal mengupdate ulasan: %v", err),
			"status": fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ulasan berhasil diupdate",
		"id":      updatedID,
		"status":  fiber.StatusOK,
	})
}

func DeleteUlasan(c *fiber.Ctx) error {
	id := c.Params("id")
	deletedID, err := repository.DeleteUlasan(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  fmt.Sprintf("Gagal menghapus ulasan: %v", err),
			"status": fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ulasan berhasil dihapus",
		"id":      deletedID,
		"status":  fiber.StatusOK,
	})
}
