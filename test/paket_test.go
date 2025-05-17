package test

// import (
// 	"backendtourapp/model"
// 	"backendtourapp/repository"
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"
// )

// var ctx = context.TODO()

// func TestInsertPaketWisata(t *testing.T) {
// 	paket := model.PaketWisata{
// 		ID:           "TEST123",
// 		NamaPaket:    "Paket Uji",
// 		Deskripsi:    "Deskripsi untuk pengujian",
// 		Harga:        999999,
// 		DurasiHari:   4,
// 		TanggalMulai: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
// 		Destinasi: []model.Destinasi{
// 			{
// 				Nama:      "Gunung Testing",
// 				Lokasi:    "Uji Lokasi",
// 				Deskripsi: "Tempat untuk testing",
// 			},
// 		},
// 	}

// 	insertedID, err := repository.InsertPaketWisata(ctx, paket)
// 	if insertedID == nil {
// 		fmt.Printf("InsertPaketWisata failed: %v\n", err)
// 	} else {
// 		fmt.Printf("Inserted PaketWisata with ID: %v\n", insertedID)
// 	}
// }

// func TestGetPaketWisataByID(t *testing.T) {
// 	id := "TEST123"
// 	paket, err := repository.GetPaketWisataByID(ctx, id)
// 	if err != nil {
// 		t.Errorf("Terjadi kesalahan saat mengambil PaketWisata: %v", err)
// 	} else if paket == nil {
// 		t.Errorf("Data PaketWisata tidak ditemukan untuk ID: %v", id)
// 	} else {
// 		fmt.Printf("Data PaketWisata ditemukan: %+v\n", paket)
// 	}
// }

// func TestGetAllPaketWisata(t *testing.T) {
// 	all, err := repository.GetAllPaketWisata(ctx)
// 	if len(all) == 0 {
// 		fmt.Printf("Tidak ada PaketWisata ditemukan: %v\n", err)
// 	} else {
// 		fmt.Printf("Total PaketWisata: %d\n", len(all))
// 		for _, p := range all {
// 			fmt.Printf("%+v\n", p)
// 		}
// 	}
// }

// func TestUpdatePaketWisata(t *testing.T) {
// 	id := "TEST123"
// 	newData := model.PaketWisata{
// 		NamaPaket:    "Paket Uji Update",
// 		Deskripsi:    "Deskripsi telah diperbarui",
// 		Harga:        888888,
// 		DurasiHari:   5,
// 		TanggalMulai: time.Now().Add(48 * time.Hour).Format(time.RFC3339),
// 		Destinasi: []model.Destinasi{
// 			{
// 				Nama:      "Pantai Uji Update",
// 				Lokasi:    "Lokasi Baru",
// 				Deskripsi: "Destinasi telah diperbarui",
// 			},
// 		},
// 	}

// 	updatedID, err := repository.UpdatePaketWisata(ctx, id, newData)
// 	if err != nil {
// 		t.Errorf("UpdatePaketWisata failed: %v", err)
// 	} else {
// 		fmt.Printf("Updated PaketWisata with ID: %v\n", updatedID)
// 	}
// }

// func TestDeletePaketWisata(t *testing.T) {
// 	id := "TEST123"
// 	deletedID, err := repository.DeletePaketWisata(ctx, id)
// 	if err != nil {
// 		t.Errorf("DeletePaketWisata failed: %v", err)
// 	} else {
// 		fmt.Printf("Deleted PaketWisata with ID: %v\n", deletedID)
// 	}
// }
