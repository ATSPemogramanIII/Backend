package repository

import (
	"backendtourapp/config"
	"backendtourapp/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertPemesanan(ctx context.Context, pemesanan model.Pemesanan) (insertedID interface{}, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.PemesananCollection)

	// Optional: cek duplikasi jika mau
	filter := bson.M{
		"id_paket":      pemesanan.IDPaket,
		"nama_pemesan":  pemesanan.NamaPemesan,
		"tanggal_pesan": pemesanan.TanggalPesan,
	}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		fmt.Printf("InsertPemesanan - Count: %v\n", err)
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("Pemesanan untuk paket %v oleh %v pada tanggal %v sudah ada", pemesanan.IDPaket, pemesanan.NamaPemesan, pemesanan.TanggalPesan)
	}

	insertResult, err := collection.InsertOne(ctx, pemesanan)
	if err != nil {
		fmt.Printf("InsertPemesanan - Insert: %v\n", err)
		return nil, err
	}

	return insertResult.InsertedID, nil
}

func GetPemesananByID(ctx context.Context, id string) (pemesanan *model.Pemesanan, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.PemesananCollection)
	filter := bson.M{"_id": id}

	err = collection.FindOne(ctx, filter).Decode(&pemesanan)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("Gagal mengambil data pemesanan: %v", err)
	}

	return pemesanan, nil
}

func GetAllPemesanan(ctx context.Context) ([]model.Pemesanan, error) {
	collection := config.MongoConnect(config.DBName).Collection(config.PemesananCollection)
	filter := bson.M{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println("GetAllPemesanan (Find):", err)
		return nil, err
	}

	var data []model.Pemesanan
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Println("GetAllPemesanan (Decode):", err)
		return nil, err
	}

	return data, nil
}

func UpdatePemesanan(ctx context.Context, id string, update model.Pemesanan) (updatedID string, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.PemesananCollection)

	filter := bson.M{"_id": id}
	updateData := bson.M{"$set": update}

	result, err := collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		fmt.Printf("UpdatePemesanan: %v\n", err)
		return "", err
	}
	if result.ModifiedCount == 0 {
		return "", fmt.Errorf("Tidak ada data yang diupdate untuk ID %v", id)
	}

	return id, nil
}

func DeletePemesanan(ctx context.Context, id string) (deletedID string, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.PemesananCollection)

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Printf("DeletePemesanan: %v\n", err)
		return "", err
	}
	if result.DeletedCount == 0 {
		return "", fmt.Errorf("Tidak ada data yang dihapus untuk ID %v", id)
	}

	return id, nil
}
