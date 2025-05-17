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

func InsertUlasan(ctx context.Context, ulasan model.Ulasan) (insertedID interface{}, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.UlasanCollection)

	// Cek duplikasi ulasan oleh pengguna yang sama untuk paket yang sama (optional)
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

func GetUlasanByID(ctx context.Context, id string) (*model.Ulasan, error) {
	collection := config.MongoConnect(config.DBName).Collection(config.UlasanCollection)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("ID tidak valid")
	}

	filter := bson.M{"_id": objID}

	var ulasan model.Ulasan
	err = collection.FindOne(ctx, filter).Decode(&ulasan)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("Gagal mengambil data ulasan: %v", err)
	}

	return &ulasan, nil
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

func UpdateUlasan(ctx context.Context, id string, update model.Ulasan) (string, error) {
	collection := config.MongoConnect(config.DBName).Collection(config.UlasanCollection)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", fmt.Errorf("ID tidak valid")
	}

	filter := bson.M{"_id": objID}

	// Update hanya field yang boleh diubah, hindari overwrite _id
	updateData := bson.M{
		"$set": bson.M{
			"id_paket":      update.IDPaket,
			"nama_pengguna": update.NamaPengguna,
			"rating":        update.Rating,
			"komentar":      update.Komentar,
			"tanggal":       update.Tanggal,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		fmt.Printf("UpdateUlasan: %v\n", err)
		return "", err
	}
	if result.MatchedCount == 0 {
		return "", fmt.Errorf("Data ulasan dengan ID %v tidak ditemukan", id)
	}
	if result.ModifiedCount == 0 {
		return "", fmt.Errorf("Tidak ada data yang diupdate untuk ID %v", id)
	}

	return id, nil
}

func DeleteUlasan(ctx context.Context, id string) (string, error) {
	collection := config.MongoConnect(config.DBName).Collection(config.UlasanCollection)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", fmt.Errorf("ID tidak valid")
	}

	filter := bson.M{"_id": objID}
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
