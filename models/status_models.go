package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusLegal struct {
	ID                  string `gorm:"column:id_status_legal" json:"id_status_legal"`
	TahunPembentukan    string `gorm:"column:tahun_pembentukan" json:"tahun_pembentukan"`
	LamTahunPembentukan string `gorm:"column:lampiran_tahun_pembentukan" json:"lampiran_tahun_pembentukan"`
	LamKplDesa          string `gorm:"column:lampiran_kep_dc" json:"lampiran_kep_dc"`
	DiketKplDaerah      string `gorm:"column:diket_kep_dc" json:"diket_kep_dc"`
	SKBupati            string `gorm:"column:no_sk_bupati" json:"no_sk_bupati"`
	LamSKBupati         string `gorm:"column:lampiran_sk_bupati" json:"lampiran_sk_bupati"`
	AkteNotaris         string `gorm:"column:akte_notaris" json:"akte_notaris"`
	LamAkteNotaris      string `gorm:"column:lampiran_akte_notaris" json:"lampiran_akte_notaris"`
	NoPendaftaran       string `gorm:"column:no_pendaftaran" json:"no_pendaftaran"`
	LamPendaftaran      string `gorm:"column:lampiran_pendaftaran" json:"lampiran_pendaftaran"`
}

func (c *StatusLegal) TableName() string {
	return "status_legal"
}
func (c *StatusLegal) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()

	return
}
