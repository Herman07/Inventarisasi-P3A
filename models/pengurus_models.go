package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pengurus struct {
	ID                     string `gorm:"column:id_kepengurusan" json:"id_kepengurusan"`
	Ketua                  string `gorm:"column:ketua" json:"ketua"`
	Wakil                  string `gorm:"column:wakil" json:"wakil"`
	Sekretaris             string `gorm:"column:sekretaris" json:"sekretaris"`
	Bendahara              string `gorm:"column:bendahara" json:"bendahara"`
	SekTeknik              string `gorm:"column:sek_teknik" json:"sek_teknik"`
	SekOP                  string `gorm:"column:sek_op" json:"sek_op"`
	SekBisnis              string `gorm:"column:sek_bisnis" json:"sek_bisnis"`
	JumlahAnggota          string `gorm:"column:jumlah_anggota" json:"jumlah_anggota"`
	NoADRT                 string `gorm:"column:no_ad_art" json:"no_ad_art"`
	LampiranADRT           string `gorm:"column:lampiran_ad_art" json:"lampiran_ad_art"`
	Sekretariat            string `gorm:"column:sekretariat" json:"sekretariat"`
	LampiranSekretariat    string `gorm:"column:lampiran_sekretariat" json:"lampiran_sekretariat"`
	PresentasiPerempuanP3A string `gorm:"column:persentase_perempuan_p3a" json:"persentase_perempuan_p3a"`
	ArealTersier           string `gorm:"column:areal_tersier" json:"areal_tersier"`
	PengisianBuku          string `gorm:"column:pengisian_buku" json:"pengisian_buku"`
	Iuran                  string `gorm:"column:iuran" json:"iuran"`
}

func (c *Pengurus) TableName() string {
	return "kepengurusan"
}

// BeforeCreate - Lifecycle callback - Generate UUID before persisting
func (c *Pengurus) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()

	return
}
