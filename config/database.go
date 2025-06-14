package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBName             = "tour"
	destinasi = "destinasi"
	paket_wisata = "paket_wisata"
	ulasan      = "ulasan"
	pemesanan   = "pemesanan"
)

var MongoString string = os.Getenv("MONGODBSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}