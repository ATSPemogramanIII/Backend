package repository

import (
	"backendtourapp/config"
	"backendtourapp/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertPaketWisata(ctx context.Context, paket model.PaketWisata) (interface{}, error) {
	collection := config.MongoConnect(config.DBName).Collection("paket_wisata")

	// Cek apakah kode_paket sudah ada
	filter := bson.M{"kode_paket": paket.KodePaket}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		fmt.Printf("InsertPaket - Count: %v\n", err)
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("kode paket %v sudah terdaftar", paket.KodePaket)
	}

	result, err := collection.InsertOne(ctx, paket)
	if err != nil {
		fmt.Printf("InsertPaket - Insert: %v\n", err)
		return nil, err
	}
	return result.InsertedID, nil
}

func GetPaketWisataByKode(ctx context.Context, kode string) (*model.PaketWisata, error) {
	collection := config.MongoConnect(config.DBName).Collection("paket_wisata")
	filter := bson.M{"kode_paket": kode}

	var paket model.PaketWisata
	err := collection.FindOne(ctx, filter).Decode(&paket)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("GetPaketByKode error: %v", err)
	}
	return &paket, nil
}

func GetAllPaketWisata(ctx context.Context) ([]model.PaketWisata, error) {
	collection := config.MongoConnect(config.DBName).Collection("paket_wisata")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("GetAllPaket - Find:", err)
		return nil, err
	}

	var data []model.PaketWisata
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Println("GetAllPaket - Decode:", err)
		return nil, err
	}

	return data, nil
}

func UpdatePaketWisata(ctx context.Context, kode string, update model.PaketWisata) (string, error) {
	collection := config.MongoConnect(config.DBName).Collection("paket_wisata")
	filter := bson.M{"kode_paket": kode}
	updateData := bson.M{"$set": update}

	result, err := collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		fmt.Printf("UpdatePaket: %v\n", err)
		return "", err
	}
	if result.ModifiedCount == 0 {
		return "", fmt.Errorf("tidak ada data yang diupdate untuk kode paket %v", kode)
	}
	return kode, nil
}

func DeletePaketWisata(ctx context.Context, kode string) (string, error) {
	collection := config.MongoConnect(config.DBName).Collection("paket_wisata")
	filter := bson.M{"kode_paket": kode}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Printf("DeletePaket: %v\n", err)
		return "", err
	}
	if result.DeletedCount == 0 {
		return "", fmt.Errorf("tidak ada data yang dihapus untuk kode paket %v", kode)
	}
	return kode, nil
}
