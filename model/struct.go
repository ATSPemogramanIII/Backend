package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Destinasi struct {
	Nama      string `json:"nama" bson:"nama"`
	Lokasi    string `json:"lokasi" bson:"lokasi"`
	Deskripsi string `json:"deskripsi" bson:"deskripsi"`
}

type PaketWisata struct {
	ID           string      `json:"_id" bson:"_id"`
	NamaPaket    string      `json:"nama_paket" bson:"nama_paket"`
	Deskripsi    string      `json:"deskripsi" bson:"deskripsi"`
	Harga        int         `json:"harga" bson:"harga"`
	DurasiHari   int         `json:"durasi_hari" bson:"durasi_hari"`
	TanggalMulai time.Time   `json:"tanggal_mulai" bson:"tanggal_mulai"`
	Destinasi    []Destinasi `json:"destinasi" bson:"destinasi"`
}

type Ulasan struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IDPaket      string             `json:"id_paket" bson:"id_paket"` // tetap string, referensi ke PaketWisata.ID
	NamaPengguna string             `json:"nama_pengguna" bson:"nama_pengguna"`
	Rating       int                `json:"rating" bson:"rating"` // 1-5
	Komentar     string             `json:"komentar" bson:"komentar"`
	Tanggal      time.Time          `json:"tanggal" bson:"tanggal"`
}

type Pemesanan struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	NamaPemesan  string             `json:"nama_pemesan" bson:"nama_pemesan"`
	Email        string             `json:"email" bson:"email"`
	NomorTelepon string             `json:"nomor_telepon" bson:"nomor_telepon"`
	IDPaket      string             `json:"id_paket" bson:"id_paket"`
	JumlahOrang  int                `json:"jumlah_orang" bson:"jumlah_orang"`
	TanggalPesan time.Time          `json:"tanggal_pesan" bson:"tanggal_pesan"`
	Status       string             `json:"status" bson:"status"` // pending, confirmed, cancelled
}
