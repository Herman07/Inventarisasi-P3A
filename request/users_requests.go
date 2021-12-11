package request

type UsersReq struct {
	ID           string `json:"id_user" form:"id_user"`
	Nama         string `json:"nama" form:"nama"`
	TglLahir     string `json:"tgl_lahir" form:"tgl_lahir"`
	Alamat       string `json:"alamat" form:"alamat"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin"`
	NoTlp        string `json:"no_telepon" form:"no_telepon"`
	Email        string `json:"email" form:"email"`
	Username     string `json:"username" form:"username"`
	Password     string `json:"password" form:"password"`
	Foto         string `json:"foto" form:"foto"`
	TypeUser     string `json:"type_users" form:"type_users"`
}
