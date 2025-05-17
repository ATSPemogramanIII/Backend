package test

import (
	"backendtourapp/model"
	"backendtourapp/repository"
	"context"
	"fmt"
	"testing"
)

var ctx = context.TODO()

// ======= TEST PAKET WISATA =======
func TestInsertPaketWisata(t *testing.T) {
	paket := model.PaketWisata{
        ID:           "testpaket1",
        NamaPaket:    "Paket Test Wisata",
        Deskripsi:    "Deskripsi Paket Test",
        Harga:        1500000,
        DurasiHari:   3,
        TanggalMulai: "2025-05-17",
        Destinasi: []model.Destinasi{
            {Nama: "Pantai A", Lokasi: "Bali", Deskripsi: "Pantai cantik"},
        },
        DibuatPada: "2025-05-01",
    }

	insertedID, err := repository.InsertPaketWisata(ctx, paket)
	if insertedID == nil || err != nil {
		t.Errorf("InsertPaketWisata failed: %v", err)
	} else {
		fmt.Printf("Inserted PaketWisata with ID: %v\n", insertedID)
	}
}

func TestGetPaketWisataByID(t *testing.T) {
	id := "testpaket1"
	paket, err := repository.GetPaketWisataByID(ctx, id)
	if err != nil {
		t.Errorf("GetPaketWisataByID error: %v", err)
	} else if paket == nil || paket.ID != id {
		t.Errorf("Expected Paket ID %v, got %v", id, paket)
	} else {
		fmt.Printf("Got PaketWisata: %+v\n", paket)
	}
}

func TestGetAllPaketWisata(t *testing.T) {
	data, err := repository.GetAllPaketWisata(ctx)
	if err != nil {
		t.Errorf("GetAllPaketWisata error: %v", err)
	} else if len(data) == 0 {
		t.Errorf("No PaketWisata found")
	} else {
		fmt.Printf("Total PaketWisata: %d\n", len(data))
	}
}

func TestUpdatePaketWisata(t *testing.T) {
	id := "testpaket1"
	update := model.PaketWisata{
		// ID biasanya tetap sama, bisa kosong atau diisi sesuai kebutuhan
		ID:           id, 
		NamaPaket:    "Paket Test Wisata Updated",
		Deskripsi:    "Deskripsi Update",
		Harga:        1700000,
		DurasiHari:   4,
		TanggalMulai: "2025-05-20",
		Destinasi: []model.Destinasi{
			{
				Nama:      "Pantai Updated",
				Lokasi:    "Kota Update",
				Deskripsi: "Destinasi update",
			},
		},
		DibuatPada: "2025-05-01",
	}

	updatedID, err := repository.UpdatePaketWisata(ctx, id, update)
	if err != nil {
		t.Errorf("UpdatePaketWisata failed: %v", err)
	} else {
		fmt.Printf("Updated PaketWisata with ID: %v\n", updatedID)
	}
}


func TestDeletePaketWisata(t *testing.T) {
	id := "testpaket1"

	deletedID, err := repository.DeletePaketWisata(ctx, id)
	if err != nil {
		t.Errorf("DeletePaketWisata failed: %v", err)
	} else {
		fmt.Printf("Deleted PaketWisata with ID: %v\n", deletedID)
	}
}

// ======= TEST ULASAN =======
func TestInsertUlasan(t *testing.T) {
	ulasan := model.Ulasan{
		ID:          "testulasan1",
		IDPaket:     "testpaket1",
		NamaPengguna:"user_test",
		Rating:      4,
		Komentar:    "Ulasan Test",
	}

	insertedID, err := repository.InsertUlasan(ctx, ulasan)
	if insertedID == nil || err != nil {
		t.Errorf("InsertUlasan failed: %v", err)
	} else {
		fmt.Printf("Inserted Ulasan with ID: %v\n", insertedID)
	}
}

func TestGetUlasanByID(t *testing.T) {
	id := "testulasan1"
	ulasan, err := repository.GetUlasanByID(ctx, id)
	if err != nil {
		t.Errorf("GetUlasanByID error: %v", err)
	} else if ulasan == nil || ulasan.ID != id {
		t.Errorf("Expected Ulasan ID %v, got %v", id, ulasan)
	} else {
		fmt.Printf("Got Ulasan: %+v\n", ulasan)
	}
}

func TestGetAllUlasan(t *testing.T) {
	data, err := repository.GetAllUlasan(ctx)
	if err != nil {
		t.Errorf("GetAllUlasan error: %v", err)
	} else if len(data) == 0 {
		t.Errorf("No Ulasan found")
	} else {
		fmt.Printf("Total Ulasan: %d\n", len(data))
	}
}

func TestUpdateUlasan(t *testing.T) {
	id := "testulasan1"
	update := model.Ulasan{
		IDPaket:      "testpaket1",
		NamaPengguna: "user_test_updated",
		Rating:       5,
		Komentar:     "Ulasan Update",
	}

	updatedID, err := repository.UpdateUlasan(ctx, id, update)
	if err != nil {
		t.Errorf("UpdateUlasan failed: %v", err)
	} else {
		fmt.Printf("Updated Ulasan with ID: %v\n", updatedID)
	}
}

func TestDeleteUlasan(t *testing.T) {
	id := "testulasan1"

	deletedID, err := repository.DeleteUlasan(ctx, id)
	if err != nil {
		t.Errorf("DeleteUlasan failed: %v", err)
	} else {
		fmt.Printf("Deleted Ulasan with ID: %v\n", deletedID)
	}
}

// ======= TEST PEMESANAN =======
func TestInsertPemesanan(t *testing.T) {
	pemesanan := model.Pemesanan{
		ID:          "testpesan1",
		IDPaket:     "testpaket1",
		NamaPemesan: "pemesan_test",
		TanggalPesan:"2025-05-17",
		JumlahOrang: 2,
	}

	insertedID, err := repository.InsertPemesanan(ctx, pemesanan)
	if insertedID == nil || err != nil {
		t.Errorf("InsertPemesanan failed: %v", err)
	} else {
		fmt.Printf("Inserted Pemesanan with ID: %v\n", insertedID)
	}
}

func TestGetPemesananByID(t *testing.T) {
	id := "testpesan1"
	pemesanan, err := repository.GetPemesananByID(ctx, id)
	if err != nil {
		t.Errorf("GetPemesananByID error: %v", err)
	} else if pemesanan == nil || pemesanan.ID != id {
		t.Errorf("Expected Pemesanan ID %v, got %v", id, pemesanan)
	} else {
		fmt.Printf("Got Pemesanan: %+v\n", pemesanan)
	}
}

func TestGetAllPemesanan(t *testing.T) {
	data, err := repository.GetAllPemesanan(ctx)
	if err != nil {
		t.Errorf("GetAllPemesanan error: %v", err)
	} else if len(data) == 0 {
		t.Errorf("No Pemesanan found")
	} else {
		fmt.Printf("Total Pemesanan: %d\n", len(data))
	}
}

func TestUpdatePemesanan(t *testing.T) {
	id := "testpesan1"
	update := model.Pemesanan{
		IDPaket:     "testpaket1",
		NamaPemesan: "pemesan_updated",
		TanggalPesan:"2025-06-01",
		JumlahOrang: 3,
	}

	updatedID, err := repository.UpdatePemesanan(ctx, id, update)
	if err != nil {
		t.Errorf("UpdatePemesanan failed: %v", err)
	} else {
		fmt.Printf("Updated Pemesanan with ID: %v\n", updatedID)
	}
}

func TestDeletePemesanan(t *testing.T) {
	id := "testpesan1"

	deletedID, err := repository.DeletePemesanan(ctx, id)
	if err != nil {
		t.Errorf("DeletePemesanan failed: %v", err)
	} else {
		fmt.Printf("Deleted Pemesanan with ID: %v\n", deletedID)
	}
}
