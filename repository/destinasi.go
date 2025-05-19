package repository

import (
	"backendtourapp/config"
	"backendtourapp/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDestinasiByKode(ctx context.Context, kode string) (*model.Destinasi, error) {
    collection := config.MongoConnect(config.DBName).Collection("destinasi")
    
    var destinasi model.Destinasi
    err := collection.FindOne(ctx, bson.M{"kode_destinasi": kode}).Decode(&destinasi)
    if err != nil {
        return nil, err
    }

    return &destinasi, nil
}

func InsertDestinasi(ctx context.Context, destinasi model.Destinasi) (interface{}, error) {
	collection := config.MongoConnect(config.DBName).Collection("destinasi")

	// Cek apakah destinasi dengan nama sama sudah ada 
	filter := bson.M{"nama": destinasi.Nama}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		fmt.Printf("InsertDestinasi - Count: %v\n", err)
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("destinasi dengan nama %v sudah terdaftar", destinasi.Nama)
	}

	result, err := collection.InsertOne(ctx, destinasi)
	if err != nil {
		fmt.Printf("InsertDestinasi - Insert: %v\n", err)
		return nil, err
	}

	return result.InsertedID, nil
}

func GetDestinasiByID(ctx context.Context, id string) (*model.Destinasi, error) {
	collection := config.MongoConnect(config.DBName).Collection("destinasi")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("ID tidak valid")
	}

	filter := bson.M{"_id": objectID}
	var destinasi model.Destinasi

	err = collection.FindOne(ctx, filter).Decode(&destinasi)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("GetDestinasiByID error: %v", err)
	}
	return &destinasi, nil
}

func GetAllDestinasi(ctx context.Context) ([]model.Destinasi, error) {
	collection := config.MongoConnect(config.DBName).Collection("destinasi")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("GetAllDestinasi - Find:", err)
		return nil, err
	}

	var data []model.Destinasi
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Println("GetAllDestinasi - Decode:", err)
		return nil, err
	}

	return data, nil
}

func UpdateDestinasi(ctx context.Context, id string, update model.Destinasi) (string, error) {
	collection := config.MongoConnect(config.DBName).Collection("destinasi")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", fmt.Errorf("ID tidak valid")
	}

	filter := bson.M{"_id": objectID}

	// Jangan update _id agar tidak error
	updateData := bson.M{
		"kode_destinasi": update.KodeDestinasi,
		"nama":      update.Nama,
		"lokasi":    update.Lokasi,
		"deskripsi": update.Deskripsi,
	}

	result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": updateData})
	if err != nil {
		fmt.Printf("UpdateDestinasi: %v\n", err)
		return "", err
	}
	if result.ModifiedCount == 0 {
		return "", fmt.Errorf("tidak ada data destinasi yang diupdate untuk ID %v", id)
	}
	return id, nil
}

func DeleteDestinasi(ctx context.Context, id string) (string, error) {
	collection := config.MongoConnect(config.DBName).Collection("destinasi")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", fmt.Errorf("ID tidak valid")
	}

	filter := bson.M{"_id": objectID}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Printf("DeleteDestinasi: %v\n", err)
		return "", err
	}
	if result.DeletedCount == 0 {
		return "", fmt.Errorf("tidak ada data destinasi yang dihapus untuk ID %v", id)
	}
	return id, nil
}
