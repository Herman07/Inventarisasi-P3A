package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterDataKab struct {
	ID        string `gorm:"column:id_kab" json:"id_kab"`
	IDProv    string `gorm:"column:id_prov_fk" json:"id_prov_fk"`
	Kabupaten string `gorm:"column:nama_kab" json:"nama_kab"`
}

func (c *MasterDataKab) TableName() string {
	return "kabupaten"
}

func (c *MasterDataKab) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
