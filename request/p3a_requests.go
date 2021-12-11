package request

type RequestInventaris struct {
	IDProv                 string `json:"id_prov_fk" form:"id_prov_fk"`
	IDKab                  string `json:"id_kab_fk" form:"id_kab_fk"`
	IDKec                  string `json:"id_kec_fk" form:"id_kec_fk"`
	IDStatusLegal          string `json:"id_status_legal_fk"`
	IDPengurus             string `json:"id_kepengurusan_fk"`
	IDIrigasi              string `json:"id_teknik_irigasi_fk"`
	IDPertanian            string `json:"id_teknik_pertanian_fk"`
	NoUrut                 string `json:"no_urut" form:"no_urut"`
	NamaP3A                string `json:"nama_p3a" form:"nama_p3a"`
	JumlahP3A              string `json:"jumlah_p3a" form:"jumlah_p3a"`
	DaerahIrigasi          string `json:"nama_daerah_irigasi" form:"nama_daerah_irigasi"`
	LuasWilayah            string `json:"luas_wilayah" form:"luas_wilayah"`
	LuasLayananP3A         string `json:"luas_layanan_p3a" form:"luas_layanan_p3a"`
	Keterangan             string `json:"keterangan" form:"keterangan"`
	IDStatus               string `json:"id_status_legal"`
	TahunPembentukan       string `json:"tahun_pembentukan" form:"tahun_pembentukan"`
	LamTahunPembentukan    string `json:"lampiran_tahun_pembentukan" form:"lampiran_tahun_pembentukan"`
	DiketKplDaerah         string `json:"diket_kep_dc" form:"diket_kep_dc"`
	LamKplDesa             string `json:"lampiran_kep_dc" form:"lampiran_kep_dc"`
	SKBupati               string `json:"no_sk_bupati" form:"no_sk_bupati"`
	LamSKBupati            string `json:"lampiran_sk_bupati" form:"lampiran_sk_bupati"`
	AkteNotaris            string `json:"akte_notaris" form:"akte_notaris"`
	LamAkteNotaris         string `json:"lampiran_akte_notaris" form:"lampiran_akte_notaris"`
	NoPendaftaran          string `json:"no_pendaftaran" form:"no_pendaftaran"`
	LamPendaftaran         string `json:"lampiran_pendaftaran" form:"lampiran_pendaftaran"`
	IDPengurusan           string `json:"id_kepengurusan"`
	Ketua                  string `json:"ketua" form:"ketua"`
	Wakil                  string `json:"wakil" form:"wakil"`
	Sekretaris             string `json:"sekretaris" form:"sekretaris"`
	Bendahara              string `json:"bendahara" form:"bendahara"`
	SekTeknik              string `json:"sek_teknik" form:"sek_teknik"`
	SekOP                  string `json:"sek_op" form:"sek_op"`
	SekBisnis              string `json:"sek_bisnis" form:"sek_bisnis"`
	JumlahAnggota          string `json:"jumlah_anggota" form:"jumlah_anggota"`
	NoADRT                 string `json:"no_ad_art" form:"no_ad_art"`
	LampiranADRT           string `json:"lampiran_ad_art" form:"lampiran_ad_art"`
	Sekretariat            string `json:"sekretariat" form:"sekretariat"`
	LampiranSekretariat    string `json:"lampiran_sekretariat" form:"lampiran_sekretariat"`
	PresentasiPerempuanP3A string `json:"persentase_perempuan_p3a" form:"persentase_perempuan_p3a"`
	ArealTersier           string `json:"areal_tersier" form:"areal_tersier"`
	PengisianBuku          string `json:"pengisian_buku" form:"pengisian_buku"`
	Iuran                  string `json:"iuran" form:"iuran"`
	IDIrig                 string `json:"id_t_irigasi"`
	Operasi                string `json:"operasi" form:"operasi"`
	Partisipatif           string `json:"partisipatif" form:"partisipatif"`
	IDTani                 string `json:"id_t_pertanian"`
	PolaTanam              string `json:"pola_tanam" form:"pola_tanam"`
	UsahaTani              string `json:"usaha_tani" form:"usaha_tani"`
}

type UpdateInventaris struct {
	IDProv                 string  `json:"id_prov_fk" form:"id_prov_fk"`
	IDKab                  string  `json:"id_kab_fk" form:"id_kab_fk"`
	IDKec                  string  `json:"id_kec_fk" form:"id_kec_fk"`
	IDStatusLegal          string  `json:"id_status_legal_fk" form:"id_status_legal_fk"`
	IDPengurus             string  `json:"id_kepengurusan_fk" form:"id_kepengurusan_fk"`
	IDIrigasi              string  `json:"id_teknik_irigasi_fk" form:"id_teknik_irigasi_fk"`
	IDPertanian            string  `json:"id_teknik_pertanian_fk" form:"id_teknik_pertanian_fk"`
	IDP3A                  string  `json:"id_p3a" form:"id_p3a"`
	NoUrut                 string  `json:"no_urut" form:"no_urut"`
	NamaP3A                string  `json:"nama_p3a" form:"nama_p3a"`
	JumlahP3A              string  `json:"jumlah_p3a" form:"jumlah_p3a"`
	DaerahIrigasi          string  `json:"nama_daerah_irigasi" form:"nama_daerah_irigasi"`
	LuasWilayah            string  `json:"luas_wilayah" form:"luas_wilayah"`
	LuasLayananP3A         string  `json:"luas_layanan_p3a" form:"luas_layanan_p3a"`
	Keterangan             string  `json:"keterangan" form:"keterangan"`
	IDStatus               string  `json:"id_status_legal" form:"id_status_legal"`
	TahunPembentukan       string  `json:"tahun_pembentukan" form:"tahun_pembentukan"`
	LamTahunPembentukan    *string `json:"lampiran_tahun_pembentukan" form:"lampiran_tahun_pembentukan"`
	DiketKplDaerah         string  `json:"diket_kep_dc" form:"diket_kep_dc"`
	LamKplDesa             *string `json:"lampiran_kep_dc" form:"lampiran_kep_dc"`
	SKBupati               string  `json:"no_sk_bupati" form:"no_sk_bupati"`
	LamSKBupati            *string `json:"lampiran_sk_bupati" form:"lampiran_sk_bupati"`
	AkteNotaris            string  `json:"akte_notaris" form:"akte_notaris"`
	LamAkteNotaris         *string `json:"lampiran_akte_notaris" form:"lampiran_akte_notaris"`
	NoPendaftaran          string  `json:"no_pendaftaran" form:"no_pendaftaran"`
	LamPendaftaran         *string `json:"lampiran_pendaftaran" form:"lampiran_pendaftaran"`
	IDPengurusan           string  `json:"id_kepengurusan" form:"id_kepengurusan"`
	Ketua                  string  `json:"ketua" form:"ketua"`
	Wakil                  string  `json:"wakil" form:"wakil"`
	Sekretaris             string  `json:"sekretaris" form:"sekretaris"`
	Bendahara              string  `json:"bendahara" form:"bendahara"`
	SekTeknik              string  `json:"sek_teknik" form:"sek_teknik"`
	SekOP                  string  `json:"sek_op" form:"sek_op"`
	SekBisnis              string  `json:"sek_bisnis" form:"sek_bisnis"`
	JumlahAnggota          string  `json:"jumlah_anggota" form:"jumlah_anggota"`
	NoADRT                 string  `json:"no_ad_art" form:"no_ad_art"`
	LampiranADRT           *string `json:"lampiran_ad_art" form:"lampiran_ad_art"`
	Sekretariat            string  `json:"sekretariat" form:"sekretariat"`
	LampiranSekretariat    *string `json:"lampiran_sekretariat" form:"lampiran_sekretariat"`
	PresentasiPerempuanP3A string  `json:"persentase_perempuan_p3a" form:"persentase_perempuan_p3a"`
	ArealTersier           string  `json:"areal_tersier" form:"areal_tersier"`
	PengisianBuku          string  `json:"pengisian_buku" form:"pengisian_buku"`
	Iuran                  string  `json:"iuran" form:"iuran"`
	IDIrig                 string  `json:"id_t_irigasi" form:"id_t_irigasi"`
	Operasi                string  `json:"operasi" form:"operasi"`
	Partisipatif           string  `json:"partisipatif" form:"partisipatif"`
	IDTani                 string  `json:"id_t_pertanian" form:"id_t_pertanian"`
	PolaTanam              string  `json:"pola_tanam" form:"pola_tanam"`
	UsahaTani              string  `json:"usaha_tani" form:"usaha_tani"`
}
