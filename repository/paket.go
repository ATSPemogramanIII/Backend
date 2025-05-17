package repository

import (
	"backendtourapp/config"
	"backendtourapp/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllPaket(ctx context.Context) ([]model.PaketWisata, error) {
	collection := config.MongoConnect(config.DBName).Collection(config.PaketCollection)
	filter := bson.M{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println("GetAllPaket (Find):", err)
		return nil, err
	}

	var data []model.PaketWisata
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Println("GetAllPaket (Decode):", err)
		return nil, err
	}

	return data, nil
}