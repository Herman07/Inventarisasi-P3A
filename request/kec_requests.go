package request

type KecReq struct {
	ID   string `json:"id_kec" form:"id_kec"`
	IDKab string `json:"id_kab_fk" form:"id_kab_fk" param:"id_kab_fk"`
	Nama string `json:"nama_kecamatan" form:"nama_kecamatan" param:"nama_kecamatan"`
}
