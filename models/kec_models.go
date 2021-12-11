package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterDataKec struct {
	ID        string `gorm:"column:id_kec" json:"id_kec"`
	IDKab     string `gorm:"column:id_kab_fk" json:"id_kab_fk"`
	Kecamatan string `gorm:"column:nama_kecamatan" json:"nama_kecamatan"`
}

func (c *MasterDataKec) TableName() string {
	return "kecamatan"
}

func (c *MasterDataKec) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}