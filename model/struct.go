package model

import "go.mongodb.org/mongo-driver/bson/primitive"


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
    TanggalMulai string      `json:"tanggal_mulai" bson:"tanggal_mulai"`
    Destinasi    []Destinasi `json:"destinasi" bson:"destinasi"`
    DibuatPada   string      `json:"dibuat_pada" bson:"dibuat_pada"`
}

type Ulasan struct {
	ID           string `json:"_id" bson:"_id"`
	IDPaket      string `json:"id_paket" bson:"id_paket"`         // Referensi ke PaketWisata
	NamaPengguna string `json:"nama_pengguna" bson:"nama_pengguna"`
	Rating       int    `json:"rating" bson:"rating"`             // Rating 1-5
	Komentar     string `json:"komentar" bson:"komentar"`
	Tanggal      string `json:"tanggal" bson:"tanggal"`           // Format: YYYY-MM-DD
}

type Pemesanan struct {
	ID            string `json:"_id" bson:"_id"`
	NamaPemesan   string `json:"nama_pemesan" bson:"nama_pemesan"`
	Email         string `json:"email" bson:"email"`
	NomorTelepon  string `json:"nomor_telepon" bson:"nomor_telepon"`
	IDPaket       string `json:"id_paket" bson:"id_paket"`               // Referensi ke PaketWisata
	JumlahOrang   int    `json:"jumlah_orang" bson:"jumlah_orang"`       // Min 1
	TanggalPesan  string `json:"tanggal_pesan" bson:"tanggal_pesan"`     // Format: YYYY-MM-DD
	Status        string `json:"status" bson:"status"`                   // ex: pending, confirmed, cancelled
}
type Mahasiswa struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Nama       string             `bson:"nama" json:"nama"`
	NPM        int             	  `bson:"npm" json:"npm"`
	Prodi      string             `bson:"prodi" json:"prodi"`
	Fakultas   string             `bson:"fakultas" json:"fakultas"`
	Alamat     Alamat             `bson:"alamat" json:"alamat"`
	Minat      []string           `bson:"minat" json:"minat"`
	MataKuliah []MataKuliah       `bson:"mata_kuliah" json:"mata_kuliah"`
}

type Alamat struct {
	Jalan     string `bson:"jalan" json:"jalan"`
	Kelurahan string `bson:"kelurahan" json:"kelurahan"`
	Kota      string `bson:"kota" json:"kota"`
}

type MataKuliah struct {
	Kode  string `bson:"kode" json:"kode"`
	Nama  string `bson:"nama" json:"nama"`
	Nilai int    `bson:"nilai" json:"nilai"`
}
