package models

type P3AModels struct {
	IDProv                 string `gorm:"column:id_prov_fk" json:"id_prov_fk" form:"id_prov_fk" param:"id_prov_fk"`
	IDKab                  string `gorm:"column:id_kab_fk" json:"id_kab_fk" form:"id_kab_fk" param:"id_kab_fk"`
	IDKec                  string `gorm:"column:id_kec_fk" json:"id_kec_fk" form:"id_kec_fk" param:"id_kec_fk"`
	NamaProv               string `gorm:"column:nama_prov" json:"nama_prov" form:"nama_prov" param:"nama_prov"`
	NamaKab                string `gorm:"column:nama_kab" json:"nama_kab" form:"nama_kab" param:"nama_kab"`
	NamaKec                string `gorm:"column:nama_kecamatan" json:"nama_kecamatan" form:"nama_kecamatan" param:"nama_kecamatan"`
	IDP3A                  string `gorm:"column:id_p3a" json:"id_p3a" form:"id_p3a" param:"id_p3a"`
	NoUrut                 string `gorm:"column:no_urut" json:"no_urut" form:"no_urut" param:"no_urut"`
	NamaP3A                string `gorm:"column:nama_p3a" json:"nama_p3a" form:"nama_p3a" param:"nama_p3a"`
	JumlahP3A              string `gorm:"column:jumlah_p3a" json:"jumlah_p3a" form:"jumlah_p3a" param:"jumlah_p3a"`
	DaerahIrigasi          string `gorm:"column:nama_daerah_irigasi" json:"nama_daerah_irigasi" form:"nama_daerah_irigasi" param:"nama_daerah_irigasi"`
	LuasWilayah            string `gorm:"column:luas_wilayah" json:"luas_wilayah" form:"luas_wilayah" param:"luas_wilayah"`
	LuasLayananP3A         string `gorm:"column:luas_layanan_p3a" json:"luas_layanan_p3a" form:"luas_layanan_p3a" param:"luas_layanan_p3a"`
	Keterangan             string `gorm:"column:keterangan" json:"keterangan" form:"keterangan" param:"keterangan"`
	IDStatus               string `gorm:"column:id_status_legal" json:"id_status_legal" param:"id_status_legal"`
	TahunPembentukan       string `gorm:"column:tahun_pembentukan" json:"tahun_pembentukan" form:"tahun_pembentukan" param:"tahun_pembentukan"`
	LamTahunPembentukan    string `gorm:"column:lampiran_tahun_pembentukan" json:"lampiran_tahun_pembentukan" form:"lampiran_tahun_pembentukan" param:"lampiran_tahun_pembentukan"`
	DiketKplDaerah         string `gorm:"column:diket_kep_dc" json:"diket_kep_dc" form:"diket_kep_dc" param:"diket_kep_dc"`
	LamKplDesa             string `gorm:"column:lampiran_kep_dc" json:"lampiran_kep_dc" form:"lampiran_kep_dc" param:"lampiran_kep_dc"`
	SKBupati               string `gorm:"column:no_sk_bupati" json:"no_sk_bupati" form:"no_sk_bupati" param:"no_sk_bupati"`
	LamSKBupati            string `gorm:"column:lampiran_sk_bupati" json:"lampiran_sk_bupati" form:"lampiran_sk_bupati" param:"lampiran_sk_bupati"`
	AkteNotaris            string `gorm:"column:akte_notaris" json:"akte_notaris" form:"akte_notaris" param:"akte_notaris"`
	LamAkteNotaris         string `gorm:"column:lampiran_akte_notaris" json:"lampiran_akte_notaris" form:"lampiran_akte_notaris" param:"lampiran_akte_notaris"`
	NoPendaftaran          string `gorm:"column:no_pendaftaran" json:"no_pendaftaran" form:"no_pendaftaran" param:"no_pendaftaran"`
	LamPendaftaran         string `gorm:"column:lampiran_pendaftaran" json:"lampiran_pendaftaran" form:"lampiran_pendaftaran" param:"lampiran_pendaftaran"`
	IDPengurusan           string `gorm:"column:id_kepengurusan" json:"id_kepengurusan" param:"id_kepengurusan"`
	Ketua                  string `gorm:"column:ketua" json:"ketua" form:"ketua" param:"ketua"`
	Wakil                  string `gorm:"column:wakil" json:"wakil" form:"wakil" param:"wakil"`
	Sekretaris             string `gorm:"column:sekretaris" json:"sekretaris" form:"sekretaris" param:"sekretaris"`
	Bendahara              string `gorm:"column:bendahara" json:"bendahara" form:"bendahara" param:"bendahara"`
	SekTeknik              string `gorm:"column:sek_teknik" json:"sek_teknik" form:"sek_teknik" param:"sek_teknik"`
	SekOP                  string `gorm:"column:sek_op" json:"sek_op" form:"sek_op" param:"sek_op"`
	SekBisnis              string `gorm:"column:sek_bisnis" json:"sek_bisnis" form:"sek_bisnis" param:"sek_bisnis"`
	JumlahAnggota          string `gorm:"column:jumlah_anggota" json:"jumlah_anggota" form:"jumlah_anggota" param:"jumlah_anggota"`
	NoADRT                 string `gorm:"column:no_ad_art" json:"no_ad_art" form:"no_ad_art" param:"no_ad_art"`
	LampiranADRT           string `gorm:"column:lampiran_ad_art" json:"lampiran_ad_art" form:"lampiran_ad_art" param:"lampiran_ad_art"`
	Sekretariat            string `gorm:"column:sekretariat" json:"sekretariat" form:"sekretariat" param:"sekretariat"`
	LampiranSekretariat    string `gorm:"column:lampiran_sekretariat" json:"lampiran_sekretariat" form:"lampiran_sekretariat" param:"lampiran_sekretariat"`
	PresentasiPerempuanP3A string `gorm:"column:persentase_perempuan_p3a" json:"persentase_perempuan_p3a" form:"persentase_perempuan_p3a" param:"persentase_perempuan_p3a"`
	ArealTersier           string `gorm:"column:areal_tersier" json:"areal_tersier" form:"areal_tersier" param:"areal_tersier"`
	PengisianBuku          string `gorm:"column:pengisian_buku" json:"pengisian_buku" form:"pengisian_buku" param:"pengisian_buku"`
	Iuran                  string `gorm:"column:iuran" json:"iuran" form:"iuran" param:"iuran"`
	IDIrig                 string `gorm:"column:id_t_irigasi" json:"id_t_irigasi" param:"id_t_irigasi"`
	Operasi                string `gorm:"column:operasi" json:"operasi" form:"operasi" param:"operasi"`
	Partisipatif           string `gorm:"column:partisipatif" json:"partisipatif" form:"partisipatif" param:"partisipatif"`
	IDTani                 string `gorm:"column:id_t_pertanian" json:"id_t_pertanian" param:"id_t_pertanian"`
	PolaTanam              string `gorm:"column:pola_tanam" json:"pola_tanam" form:"pola_tanam" param:"pola_tanam"`
	UsahaTani              string `gorm:"column:usaha_tani" json:"usaha_tani" form:"usaha_tani" param:"usaha_tani"`
}
