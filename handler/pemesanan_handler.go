package handler

import (
	"backendtourapp/model"
	"backendtourapp/repository"
	"fmt"
	"net/mail"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllPemesanan(c *fiber.Ctx) error {
	data, err := repository.GetAllPemesanan(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data pemesanan",
			"status":  fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data pemesanan",
		"data":    data,
		"status":  fiber.StatusOK,
	})
}

func GetPemesananByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "ID tidak valid",
			"status": fiber.StatusBadRequest,
		})
	}

	pemesanan, err := repository.GetPemesananByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Pemesanan tidak ditemukan",
			"status":  fiber.StatusNotFound,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Pemesanan ditemukan",
		"data":    pemesanan,
		"status":  fiber.StatusOK,
	})
}

func InsertPemesanan(c *fiber.Ctx) error {
	var pemesanan model.Pemesanan
	if err := c.BodyParser(&pemesanan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Data tidak valid",
			"status": fiber.StatusBadRequest,
		})
	}

	// Validasi jumlah orang minimal 1
	if pemesanan.JumlahOrang < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Jumlah orang minimal 1",
			"status": fiber.StatusBadRequest,
		})
	}

	// Validasi email sederhana
	if _, err := mail.ParseAddress(pemesanan.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Email tidak valid",
			"status": fiber.StatusBadRequest,
		})
	}

	// Set tanggal pesan ke waktu saat ini (otomatis)
	pemesanan.TanggalPesan = time.Now()

	// Jika status kosong, set default 'pending'
	if pemesanan.Status == "" {
		pemesanan.Status = "pending"
	}

	insertedID, err := repository.InsertPemesanan(c.Context(), pemesanan)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  fmt.Sprintf("Gagal menambahkan pemesanan: %v", err),
			"status": fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Pemesanan berhasil ditambahkan",
		"id":      insertedID.Hex(), // Convert ObjectID ke string hex agar enak dibaca client
		"status":  fiber.StatusCreated,
	})
}

func UpdatePemesanan(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "ID tidak valid",
			"status": fiber.StatusBadRequest,
		})
	}

	var pemesanan model.Pemesanan
	if err := c.BodyParser(&pemesanan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Data tidak valid",
			"status": fiber.StatusBadRequest,
		})
	}

	// Validasi jumlah orang minimal 1
	if pemesanan.JumlahOrang < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Jumlah orang minimal 1",
			"status": fiber.StatusBadRequest,
		})
	}

	// Validasi email sederhana
	if _, err := mail.ParseAddress(pemesanan.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Email tidak valid",
			"status": fiber.StatusBadRequest,
		})
	}

	updatedID, err := repository.UpdatePemesanan(c.Context(), id, pemesanan)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  fmt.Sprintf("Gagal mengupdate pemesanan: %v", err),
			"status": fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Pemesanan berhasil diupdate",
		"id":      updatedID.Hex(),
		"status":  fiber.StatusOK,
	})
}

func DeletePemesanan(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "ID tidak valid",
			"status": fiber.StatusBadRequest,
		})
	}

	deletedID, err := repository.DeletePemesanan(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  fmt.Sprintf("Gagal menghapus pemesanan: %v", err),
			"status": fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Pemesanan berhasil dihapus",
		"id":      deletedID.Hex(),
		"status":  fiber.StatusOK,
	})
}
