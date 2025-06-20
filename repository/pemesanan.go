package repository

import (
	"backendtourapp/config"
	"backendtourapp/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertPemesanan(ctx context.Context, pemesanan model.Pemesanan) (interface{}, error) {
	collection := config.MongoConnect(config.DBName).Collection("pemesanan")

	// Validasi field wajib
	if pemesanan.NamaPemesan == "" || pemesanan.Email == "" || pemesanan.KodePaket == "" {
		return nil, fmt.Errorf("nama pemesan, email, dan kode paket tidak boleh kosong")
	}
	if pemesanan.JumlahOrang <= 0 {
		return nil, fmt.Errorf("jumlah orang harus lebih dari 0")
	}

	// Set tanggal pesan jika belum diset
	if pemesanan.TanggalPesan.IsZero() {
		pemesanan.TanggalPesan = time.Now()
	}

	// Validasi status
	validStatus := map[string]bool{
		"pending":    true,
		"dikonfirmasi":    true,
		"dibatalkan": true,
	}
	if _, ok := validStatus[pemesanan.Status]; !ok {
		return nil, fmt.Errorf("status pemesanan tidak valid: gunakan 'pending', 'dikonfirmasi', atau 'dibatalkan'")
	}

	// Insert ke database
	result, err := collection.InsertOne(ctx, pemesanan)
	if err != nil {
		log.Printf("InsertPemesanan - Insert error: %v\n", err)
		return nil, err
	}

	return result.InsertedID, nil
}

func GetPemesananByKode(ctx context.Context, kode string) ([]model.Pemesanan, error) {
	collection := config.MongoConnect(config.DBName).Collection("pemesanan")
	filter := bson.M{"kode_paket": kode}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println("GetPemesananByKode - Find:", err)
		return nil, err
	}

	var data []model.Pemesanan
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Println("GetPemesananByKode - Decode:", err)
		return nil, err
	}
	return data, nil
}

func GetAllPemesanan(ctx context.Context) ([]model.Pemesanan, error) {
	collection := config.MongoConnect(config.DBName).Collection("pemesanan")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("GetAllPemesanan - Find:", err)
		return nil, err
	}

	var data []model.Pemesanan
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Println("GetAllPemesanan - Decode:", err)
		return nil, err
	}
	return data, nil
}

func UpdatePemesanan(ctx context.Context, id interface{}, pemesanan model.Pemesanan) (interface{}, error) {
	collection := config.MongoConnect(config.DBName).Collection("pemesanan")
	filter := bson.M{"_id": id}

	// Jangan update _id agar tidak error, jadi buat update tanpa _id
	updateData := bson.M{
		"nama_pemesan":  pemesanan.NamaPemesan,
		"email":         pemesanan.Email,
		"nomor_telepon": pemesanan.NomorTelepon,
		"kode_paket":    pemesanan.KodePaket,
		"jumlah_orang":  pemesanan.JumlahOrang,
		"tanggal_pesan": pemesanan.TanggalPesan,
		"status":        pemesanan.Status,
	}

	update := bson.M{"$set": updateData}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("pemesanan dengan ID tersebut tidak ditemukan")
	}
	return id, nil
}


func DeletePemesanan(ctx context.Context, id interface{}) (interface{}, error) {
	collection := config.MongoConnect(config.DBName).Collection("pemesanan")
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Printf("DeletePemesanan: %v\n", err)
		return nil, err
	}
	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("pemesanan dengan ID tersebut tidak ditemukan")
	}
	return id, nil
}
