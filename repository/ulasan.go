package repository

import (
	"backendtourapp/config"
	"backendtourapp/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertUlasan(ctx context.Context, ulasan model.Ulasan) (insertedID interface{}, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.UlasanCollection)

	// Cek apakah sudah ada ulasan untuk paket dan pengguna yg sama (optional)
	filter := bson.M{
		"id_paket":      ulasan.IDPaket,
		"nama_pengguna": ulasan.NamaPengguna,
	}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		fmt.Printf("InsertUlasan - Count: %v\n", err)
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("Ulasan untuk paket %v oleh %v sudah ada", ulasan.IDPaket, ulasan.NamaPengguna)
	}

	insertResult, err := collection.InsertOne(ctx, ulasan)
	if err != nil {
		fmt.Printf("InsertUlasan - Insert: %v\n", err)
		return nil, err
	}

	return insertResult.InsertedID, nil
}

func GetUlasanByID(ctx context.Context, id string) (ulasan *model.Ulasan, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.UlasanCollection)
	filter := bson.M{"_id": id}

	err = collection.FindOne(ctx, filter).Decode(&ulasan)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("Gagal mengambil data ulasan: %v", err)
	}

	return ulasan, nil
}

func GetAllUlasan(ctx context.Context) ([]model.Ulasan, error) {
	collection := config.MongoConnect(config.DBName).Collection(config.UlasanCollection)
	filter := bson.M{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println("GetAllUlasan (Find):", err)
		return nil, err
	}

	var data []model.Ulasan
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Println("GetAllUlasan (Decode):", err)
		return nil, err
	}

	return data, nil
}

func UpdateUlasan(ctx context.Context, id string, update model.Ulasan) (updatedID string, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.UlasanCollection)

	filter := bson.M{"_id": id}
	updateData := bson.M{"$set": update}

	result, err := collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		fmt.Printf("UpdateUlasan: %v\n", err)
		return "", err
	}
	if result.ModifiedCount == 0 {
		return "", fmt.Errorf("Tidak ada data yang diupdate untuk ID %v", id)
	}

	return id, nil
}

func DeleteUlasan(ctx context.Context, id string) (deletedID string, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.UlasanCollection)

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Printf("DeleteUlasan: %v\n", err)
		return "", err
	}
	if result.DeletedCount == 0 {
		return "", fmt.Errorf("Tidak ada data yang dihapus untuk ID %v", id)
	}

	return id, nil
}
