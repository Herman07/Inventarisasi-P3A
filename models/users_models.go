package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           string `gorm:"column:id_user" json:"id_user"`
	Nama         string `gorm:"column:nama" json:"nama"`
	TglLahir     string `gorm:"column:tgl_lahir" json:"tgl_lahir"`
	Alamat       string `gorm:"column:alamat" json:"alamat"`
	JenisKelamin string `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	NoTlp        string `gorm:"column:no_telepon" json:"no_telepon"`
	Email        string `gorm:"column:email" json:"email"`
}

func (c *User) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()

	return
}

type UserUpdate struct {
	ID           string `gorm:"column:id_user" json:"id_user"`
	Nama         string `gorm:"column:nama" json:"nama"`
	TglLahir     string `gorm:"column:tgl_lahir" json:"tgl_lahir"`
	Alamat       string `gorm:"column:alamat" json:"alamat"`
	JenisKelamin string `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	NoTlp        string `gorm:"column:no_telepon" json:"no_telepon"`
	Email        string `gorm:"column:email" json:"email"`
	IDAkun       string `gorm:"column:id_usergroup" json:"id_usergroup"`
	Username     string `gorm:"column:username" json:"username"`
	Password     string `gorm:"column:password" json:"password""`
	Foto         string `gorm:"column:foto" json:"foto"`
	TypeUser     string `gorm:"column:type_users" json:"type_users"`
}
