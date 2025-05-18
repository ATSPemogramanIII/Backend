package handler

import (
	"backendtourapp/model"
	"backendtourapp/repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetDestinasiByKode(c *fiber.Ctx) error {
    kode := c.Params("kode")

    destinasi, err := repository.GetDestinasiByKode(c.Context(), kode)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": fmt.Sprintf("Destinasi dengan Kode %s tidak ditemukan", kode),
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Berhasil mengambil data destinasi",
        "data":    destinasi,
    })
}



// GET all destinasi
func GetAllDestinasi(c *fiber.Ctx) error {
	data, err := repository.GetAllDestinasi(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data destinasi",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data destinasi",
		"data":    data,
	})
}

// GET destinasi by ID
func GetDestinasiByID(c *fiber.Ctx) error {
    idParam := c.Params("id")
    // Jangan convert dulu ke ObjectID

    data, err := repository.GetDestinasiByID(c.Context(), idParam)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": fmt.Sprintf("Destinasi dengan ID %s tidak ditemukan", idParam),
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Data destinasi ditemukan",
        "data":    data,
    })
}

// POST / insert destinasi
func InsertDestinasi(c *fiber.Ctx) error {
	var destinasi model.Destinasi
	if err := c.BodyParser(&destinasi); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data destinasi tidak valid",
		})
	}

	id, err := repository.InsertDestinasi(c.Context(), destinasi)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": fmt.Sprintf("Gagal menambahkan destinasi: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Destinasi berhasil ditambahkan",
		"id":      id,
	})
}



// PUT / update destinasi by ID
func UpdateDestinasi(c *fiber.Ctx) error {
    idParam := c.Params("id")
    // Jangan convert dulu ke ObjectID

    var update model.Destinasi
    if err := c.BodyParser(&update); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Format data destinasi tidak valid",
        })
    }

    updatedID, err := repository.UpdateDestinasi(c.Context(), idParam, update)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Data destinasi berhasil diperbarui",
        "id":      updatedID,
    })
}

// DELETE / delete destinasi by ID
func DeleteDestinasi(c *fiber.Ctx) error {
    idParam := c.Params("id")
    // Jangan convert dulu ke ObjectID

    deletedID, err := repository.DeleteDestinasi(c.Context(), idParam)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": fmt.Sprintf("Destinasi dengan ID %s tidak ditemukan: %v", idParam, err),
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Destinasi berhasil dihapus",
        "id":      deletedID,
    })
}
