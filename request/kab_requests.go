package request

type KabReq struct {
	ID   string `json:"id_kab" form:"id_kab"`
	IDProv   string `json:"id_prov_fk" form:"id_prov_fk"`
	Nama string `json:"nama_kab" form:"nama_kab" param:"nama_kab"`
}
