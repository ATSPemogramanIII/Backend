package repository

import (
	"backendtourapp/config"
	"backendtourapp/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Menambahkan data paket wisata
func InsertPaketWisata(ctx context.Context, paket model.PaketWisata) (insertedID interface{}, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.PaketWisataCollection)

	// Cek apakah ID sudah ada
	filter := bson.M{"_id": paket.ID}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		fmt.Printf("InsertPaketWisata - Count: %v\n", err)
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("ID paket wisata %v sudah terdaftar", paket.ID)
	}

	insertResult, err := collection.InsertOne(ctx, paket)
	if err != nil {
		fmt.Printf("InsertPaketWisata - Insert: %v\n", err)
		return nil, err
	}

	return insertResult.InsertedID, nil
}

// Mengambil data berdasarkan ID
func GetPaketWisataByID(ctx context.Context, id string) (paket *model.PaketWisata, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.PaketWisataCollection)
	filter := bson.M{"_id": id}

	var result model.PaketWisata
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("Gagal mengambil data paket wisata: %v", err)
	}

	return &result, nil
}

// Mengambil semua data paket wisata
func GetAllPaketWisata(ctx context.Context) ([]model.PaketWisata, error) {
	collection := config.MongoConnect(config.DBName).Collection(config.PaketWisataCollection)
	filter := bson.M{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println("GetAllPaketWisata (Find):", err)
		return nil, err
	}

	var data []model.PaketWisata
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Println("GetAllPaketWisata (Decode):", err)
		return nil, err
	}

	return data, nil
}

// Update data paket wisata berdasarkan ID
func UpdatePaketWisata(ctx context.Context, id string, update model.PaketWisata) (updatedID string, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.PaketWisataCollection)

	filter := bson.M{"_id": id}
	updateData := bson.M{
		"$set": bson.M{
			"nama_paket":    update.NamaPaket,
			"deskripsi":     update.Deskripsi,
			"harga":         update.Harga,
			"durasi_hari":   update.DurasiHari,
			"tanggal_mulai": update.TanggalMulai,
			"destinasi":     update.Destinasi,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		fmt.Printf("UpdatePaketWisata: %v\n", err)
		return "", err
	}
	if result.ModifiedCount == 0 {
		return "", fmt.Errorf("Tidak ada data yang diupdate untuk ID %v", id)
	}

	return id, nil
}

// Hapus data paket wisata berdasarkan ID
func DeletePaketWisata(ctx context.Context, id string) (deletedID string, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.PaketWisataCollection)

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Printf("DeletePaketWisata: %v\n", err)
		return "", err
	}
	if result.DeletedCount == 0 {
		return "", fmt.Errorf("Tidak ada data yang dihapus untuk ID %v", id)
	}

	return id, nil
}
