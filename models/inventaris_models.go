package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Inventaris struct {
	ID             string    `gorm:"column:id_p3a" json:"id_p3a"`
	IDProv         string    `gorm:"column:id_prov_fk" json:"id_prov_fk"`
	IDKab          string    `gorm:"column:id_kab_fk" json:"id_kab_fk"`
	IDKec          string    `gorm:"column:id_kec_fk" json:"id_kec_fk"`
	IDStatusLegal  string    `gorm:"column:id_status_legal_fk" json:"id_status_legal_fk"`
	IDPengurus     string    `gorm:"column:id_kepengurusan_fk" json:"id_kepengurusan_fk"`
	IDIrigasi      string    `gorm:"column:id_teknik_irigasi_fk" json:"id_teknik_irigasi_fk"`
	IDPertanian    string    `gorm:"column:id_teknik_pertanian_fk" json:"id_teknik_pertanian_fk"`
	NoUrut         string    `gorm:"column:no_urut" json:"no_urut"`
	NamaP3A        string    `gorm:"column:nama_p3a" json:"nama_p3a"`
	JumlahP3A      string    `gorm:"column:jumlah_p3a" json:"jumlah_p3a"`
	DaerahIrigasi  string    `gorm:"column:nama_daerah_irigasi" json:"nama_daerah_irigasi"`
	LuasWilayah    string    `gorm:"column:luas_wilayah" json:"luas_wilayah"`
	LuasLayananP3A string    `gorm:"column:luas_layanan_p3a" json:"luas_layanan_p3a"`
	Keterangan     string    `gorm:"column:keterangan" json:"keterangan"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
}

func (c *Inventaris) TableName() string {
	return "data_p3a"
}

// BeforeCreate - Lifecycle callback - Generate UUID before persisting
func (c *Inventaris) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()

	return
}
