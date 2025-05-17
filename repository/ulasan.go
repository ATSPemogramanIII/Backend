package repository

import (
	"backendtourapp/config"
	"backendtourapp/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertUlasan(ctx context.Context, ulasan model.Ulasan) (interface{}, error) {
	collection := config.MongoConnect(config.DBName).Collection("ulasan")

	result, err := collection.InsertOne(ctx, ulasan)
	if err != nil {
		fmt.Printf("InsertUlasan - Insert: %v\n", err)
		return nil, err
	}
	return result.InsertedID, nil
}

func GetUlasanByKodePaket(ctx context.Context, kodePaket string) ([]model.Ulasan, error) {
	collection := config.MongoConnect(config.DBName).Collection("ulasan")
	filter := bson.M{"kode_paket": kodePaket}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println("GetUlasanByKodePaket - Find:", err)
		return nil, err
	}

	var data []model.Ulasan
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Println("GetUlasanByKodePaket - Decode:", err)
		return nil, err
	}
	return data, nil
}

func GetAllUlasan(ctx context.Context) ([]model.Ulasan, error) {
	collection := config.MongoConnect(config.DBName).Collection("ulasan")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("GetAllUlasan - Find:", err)
		return nil, err
	}

	var data []model.Ulasan
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Println("GetAllUlasan - Decode:", err)
		return nil, err
	}
	return data, nil
}

func UpdateUlasan(ctx context.Context, id interface{}, updatedData model.Ulasan) (interface{}, error) {
	collection := config.MongoConnect(config.DBName).Collection("ulasan")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedData}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Printf("UpdateUlasan: %v\n", err)
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("ulasan dengan ID tersebut tidak ditemukan")
	}
	return id, nil
}

func DeleteUlasan(ctx context.Context, id interface{}) (interface{}, error) {
	collection := config.MongoConnect(config.DBName).Collection("ulasan")
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Printf("DeleteUlasan: %v\n", err)
		return nil, err
	}
	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("ulasan dengan ID tersebut tidak ditemukan")
	}
	return id, nil
}
