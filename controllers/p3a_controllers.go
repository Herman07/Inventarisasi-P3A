package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/xuri/excelize/v2"
	"io"
	"net/http"
	"os"
	"Inventarisasi-P3A/models"
	"Inventarisasi-P3A/request"
	"Inventarisasi-P3A/services"
	"Inventarisasi-P3A/utils/session"
	"strconv"
	"time"
)

type P3Controller struct {
	BaseFrontendController
	Controller
	service *services.P3Service
}

func NewP3Controller(services *services.P3Service) P3Controller {
	return P3Controller{
		service: services,
		BaseFrontendController: BaseFrontendController{
			Menu:        "Dashboard",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}
func (c *P3Controller) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/admin/v1/inventaris",
	}
	ses, _ := session.Manager.Get(ctx, session.SessionId)
	dataSes, _ := json.Marshal(ses)
	var data session.UserInfo
	userInfo := session.UserInfo{
		ID:       data.ID,
		TypeUser: data.TypeUser,
	}
	_ = json.Unmarshal(dataSes, &userInfo)
	return Render(ctx, "Home", "p3a/index", c.Menu, append(c.BreadCrumbs, breadCrumbs), userInfo)
}

func (c *P3Controller) Add(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/admin/v1/inventaris/add",
	}
	return Render(ctx, "Home", "p3a/create", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}

func (c *P3Controller) GetDetail(ctx echo.Context) error {
	ses, _ := session.Manager.Get(ctx, session.SessionId)
	dataSes, _ := json.Marshal(ses)
	var data1 session.UserInfo
	userInfo := session.UserInfo{
		ID:       data1.ID,
		TypeUser: data1.TypeUser,
	}
	_ = json.Unmarshal(dataSes, &userInfo)

	draw, err := strconv.Atoi(ctx.Request().URL.Query().Get("draw"))
	start, err := strconv.Atoi(ctx.Request().URL.Query().Get("start"))
	search := ctx.Request().URL.Query().Get("search[value]")
	length, err := strconv.Atoi(ctx.Request().URL.Query().Get("length"))
	order, err := strconv.Atoi(ctx.Request().URL.Query().Get("order[0][column]"))
	orderName := ctx.Request().URL.Query().Get("columns[" + strconv.Itoa(order) + "][name]")
	//orderAscDesc := ctx.Request().URL.Query().Get("order[0][dir]")

	recordTotal, recordFiltered, data, err := c.service.QueryDatatable(search, "DESC", orderName, length, start)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	//var createdAt string
	var action string
	listOfData := make([]map[string]interface{}, len(data))
	for k, v := range data {
		if userInfo.TypeUser != "2" {
			action = `
		<a href="/admin/v1/inventaris/update/` + (v.IDP3A) + `" class="btn btn-primary" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fa fa-edit"></i></a>
		<a href="/admin/v1/inventaris/detail/` + (v.IDP3A) + `" class="btn btn-primary" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fa fa-eye"></i></a>
		<button onclick="Delete('` + v.IDP3A + `')" class="btn btn-danger" title="Delete" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fa fa-trash"></i></button>`
		} else {
			action = `<a href="/admin/v1/inventaris/detail/` + (v.IDP3A) + `" class="btn btn-primary" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fa fa-eye"></i></a>`
		}
		listOfData[k] = map[string]interface{}{
			"id_p3a":                   v.IDP3A,
			"no_urut":                  v.NoUrut,
			"nama_p3a":                 v.NamaP3A,
			"jumlah_p3a":               v.JumlahP3A,
			"nama_daerah_irigasi":      v.DaerahIrigasi,
			"luas_wilayah":             v.LuasWilayah,
			"luas_layanan_p3a":         v.LuasLayananP3A,
			"keterangan":               v.Keterangan,
			"nama_prov":                v.NamaProv,
			"nama_kab":                 v.NamaKab,
			"nama_kecamatan":           v.NamaKec,
			"tahun_pembentukan":        v.TahunPembentukan,
			"diket_kep_dc":             v.DiketKplDaerah,
			"sk_bupati":                v.SKBupati,
			"akte_notaris":             v.AkteNotaris,
			"no_pendaftaran":           v.NoPendaftaran,
			"ketua":                    v.Ketua,
			"wakil":                    v.Wakil,
			"sekretaris":               v.Sekretaris,
			"bendahara":                v.Bendahara,
			"sek_op":                   v.SekOP,
			"sek_bisnis":               v.SekBisnis,
			"sek_teknik":               v.SekTeknik,
			"jumlah_anggota":           v.JumlahAnggota,
			"no_ad_art":                v.NoADRT,
			"sekretariat":              v.Sekretariat,
			"persentase_perempuan_p3a": v.PresentasiPerempuanP3A,
			"areal_tersier":            v.ArealTersier,
			"pengisian_buku":           v.PengisianBuku,
			"iuran":                    v.Iuran,
			"operasi":                  v.Operasi,
			"partisipatif":             v.Partisipatif,
			"pola_tanam":               v.PolaTanam,
			"usaha_tani":               v.UsahaTani,
			"action":                   action,
		}
	}
	result := models.ResponseDatatable{
		Draw:            draw,
		RecordsTotal:    recordTotal,
		RecordsFiltered: recordFiltered,
		Data:            listOfData,
	}
	return ctx.JSON(http.StatusOK, &result)
}

func (c *P3Controller) AddData(ctx echo.Context) error {
	var entity request.RequestInventaris
	if err := ctx.Bind(&entity); err != nil {
		log.Error("[Error] ", err)
		return ctx.JSON(500, echo.Map{"message": "error binding data"})
	}

	name := []string{"lampiran_tahun_pembentukan", "lampiran_kep_dc", "lampiran_sk_bupati", "lampiran_akte_notaris", "lampiran_pendaftaran", "lampiran_ad_art", "lampiran_sekretariat"}
	var namaFile []string
	var prefixFile []string
	fmt.Println("Nama File : ", name)
	if name != nil {
		for i := range name {
			file, err := ctx.FormFile(name[i])
			if err == nil {
				src, _ := file.Open()
				defer src.Close()

				// Destination
				t := time.Now().UnixNano()
				nf := name[i] + "_" + strconv.FormatInt(t, 10) + "_" + file.Filename
				nama := "static/image/" + nf
				dst, _ := os.Create(nama)
				defer dst.Close()

				// Copy
				_, err = io.Copy(dst, src)
				if err != nil {
					log.Error("[Error] ", err)
					return c.InternalServerError(ctx, err)
				}
				prefixFile = append(prefixFile, name[i])
				i++
				namaFile = append(namaFile, nf)			}
		}
	}

	//Store Data Status Legal
	statusLegal, err := c.service.CreateStatusLegal(entity, namaFile, prefixFile)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	// Store Data Kepengurusan
	pengurus, err := c.service.CreatePengurus(entity, namaFile, prefixFile)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	//Store Data Teknik Irigasi
	irigasi, err := c.service.CreateIrigasi(entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	//Store Data Teknik Pertanian
	pertanian, err := c.service.CreatePertanian(entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	//Store Data to Table p3a
	_, err = c.service.CreateDataP3a(entity, statusLegal.ID, pengurus.ID, irigasi.ID, pertanian.ID)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	time.Sleep(1 * time.Second)
	return ctx.Redirect(http.StatusFound, "/admin/v1/inventaris")
}

func (c *P3Controller) GenerateExcel(ctx echo.Context) error {
	// Get Data Export
	data, err := c.service.GetDataExport()
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	f := excelize.NewFile()
	_, _ = f.NewConditionalStyle("center")
	center, _ := f.NewStyle(`{"alignment":{"horizontal":"center"},"font":{"italic":true},"border": [
			{
				"type": "left",
				"color": "202020",
				"style": 5
			},
			{
				"type": "top",
				"color": "202020",
				"style": 5
			},
			{
				"type": "bottom",
				"color": "202020",
				"style": 5
			},
			{
				"type": "right",
				"color": "202020",
				"style": 5
			}]}`)
	columncolor, _ := f.NewStyle(`
		{"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true,"italic":true},
		"fill":{"type":"pattern","color":["#20FF00"],"pattern":1},
		"border": [
			{
				"type": "left",
				"color": "202020",
				"style": 5
			},
			{
				"type": "top",
				"color": "202020",
				"style": 5
			},
			{
				"type": "bottom",
				"color": "202020",
				"style": 5
			},
			{
				"type": "right",
				"color": "202020",
				"style": 5
			}]}`)
	columncolor1, _ := f.NewStyle(`
		{"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true,"italic":true},
		"fill":{"type":"pattern","color":["#FF0000"],"pattern":1},
		"border": [
			{
				"type": "left",
				"color": "202020",
				"style": 5
			},
			{
				"type": "top",
				"color": "202020",
				"style": 5
			},
			{
				"type": "bottom",
				"color": "202020",
				"style": 5
			},
			{
				"type": "right",
				"color": "202020",
				"style": 5
			}]}`)
	columncolor2, _ := f.NewStyle(`
		{"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true,"italic":true},
		"fill":{"type":"pattern","color":["#0095FF"],"pattern":1},
		"border": [
			{
				"type": "left",
				"color": "202020",
				"style": 5
			},
			{
				"type": "top",
				"color": "202020",
				"style": 5
			},
			{
				"type": "bottom",
				"color": "202020",
				"style": 5
			},
			{
				"type": "right",
				"color": "202020",
				"style": 5
			}]}`)
	f.SetCellStyle("Sheet1", "A1", "H3", columncolor)
	f.SetCellStyle("Sheet1", "I1", "M3", columncolor1)
	f.SetCellStyle("Sheet1", "N1", "AA3", columncolor2)
	f.SetCellStyle("Sheet1", "AB1", "AF3", columncolor)
	f.SetColWidth("Sheet1", "A", "AF", 20)
	in := 4
	for range data {
		in++
	}
	number := strconv.Itoa(in)
	column := fmt.Sprintf("AF%s", number)
	_ = f.SetCellStyle("Sheet1", "A4", column, center)

	// Create a new sheet.
	index := f.NewSheet("Data Export")
	// Set value of a cell.
	_ = f.SetCellValue("Sheet1", "A1", "No")
	_ = f.SetCellValue("Sheet1", "B1", "Provinsi/Kabupaten")
	_ = f.SetCellValue("Sheet1", "C1", "Kecamatan")
	_ = f.SetCellValue("Sheet1", "D1", "Nama Daerah Imigrasi")
	_ = f.SetCellValue("Sheet1", "E1", "Luas Wilayah (Ha)")
	_ = f.SetCellValue("Sheet1", "F1", "Jumlah P3A")
	_ = f.SetCellValue("Sheet1", "G1", "Nama P3A")
	_ = f.SetCellValue("Sheet1", "H1", "Luas Layanan P3A (Ha)")

	// Status Legal
	_ = f.SetCellValue("Sheet1", "I1", "Status Legal P3A")
	_ = f.SetCellValue("Sheet1", "I2", "Tahun Pembentukan")
	_ = f.SetCellValue("Sheet1", "J2", "Diketahui Kepala Desa / Camat")
	_ = f.SetCellValue("Sheet1", "K2", "SK Bupati")
	_ = f.SetCellValue("Sheet1", "L2", "Akte Notaris")
	_ = f.SetCellValue("Sheet1", "M2", "Terdaftar di Pengadilan Negeri / Kemenkum HAM")
	_ = f.MergeCell("Sheet1", "I1", "M1")

	// Kepengurusan
	_ = f.SetCellValue("Sheet1", "N1", "Kepengurusan")
	_ = f.SetCellValue("Sheet1", "N2", "Ketua")
	_ = f.SetCellValue("Sheet1", "O2", "Wakil Ketua")
	_ = f.SetCellValue("Sheet1", "P2", "Sekertaris")
	_ = f.SetCellValue("Sheet1", "Q2", "Bendahara")
	_ = f.SetCellValue("Sheet1", "R2", "Seksi (L/P)")
	_ = f.SetCellValue("Sheet1", "U2", "Jumlah Anggota")
	_ = f.SetCellValue("Sheet1", "N3", "L/P")
	_ = f.SetCellValue("Sheet1", "O3", "L/P")
	_ = f.SetCellValue("Sheet1", "P3", "L/P")
	_ = f.SetCellValue("Sheet1", "Q3", "L/P")
	_ = f.SetCellValue("Sheet1", "R3", "Teknik")
	_ = f.SetCellValue("Sheet1", "S3", "O & P")
	_ = f.SetCellValue("Sheet1", "T3", "Bisnis")
	_ = f.MergeCell("Sheet1", "N1", "U1")
	_ = f.MergeCell("Sheet1", "R2", "T2")
	_ = f.MergeCell("Sheet1", "U2", "U3")

	_ = f.SetCellValue("Sheet1", "V1", "AD/ART")
	_ = f.SetCellValue("Sheet1", "W1", "Sekertariat")
	_ = f.SetCellValue("Sheet1", "X1", "Persentase (%) perempuan")
	_ = f.SetCellValue("Sheet1", "Y1", "areal tersier (ha)")
	_ = f.SetCellValue("Sheet1", "Z1", "Pengisian Buku")
	_ = f.SetCellValue("Sheet1", "AA1", "Iuran")

	//Teknik irigasi
	_ = f.SetCellValue("Sheet1", "AB1", "Teknik Irigasi")
	_ = f.SetCellValue("Sheet1", "AB2", "Operasi")
	_ = f.SetCellValue("Sheet1", "AC2", "Partisipatif")
	_ = f.MergeCell("Sheet1", "AB1", "AC1")
	_ = f.MergeCell("Sheet1", "AB2", "AB3")
	_ = f.MergeCell("Sheet1", "AC2", "AC3")

	//Teknik Pertanian
	_ = f.SetCellValue("Sheet1", "AD1", "Teknik Pertanian")
	_ = f.SetCellValue("Sheet1", "AD2", "Pola Tanam")
	_ = f.SetCellValue("Sheet1", "AE2", "Usaha Tani")
	_ = f.MergeCell("Sheet1", "AD1", "AE1")
	_ = f.MergeCell("Sheet1", "AD2", "AD3")
	_ = f.MergeCell("Sheet1", "AE2", "AE3")

	// Keterangan
	_ = f.SetCellValue("Sheet1", "AF1", "Keterangan")
	_ = f.MergeCell("Sheet1", "AF1", "AF3")

	// ROW MERGE
	for i := 0; i < 9; i++ {
		_ = f.MergeCell("Sheet1", string(rune('A'-1+i))+"1", string(rune('A'-1+i))+"3")
	}

	for i := 0; i < 6; i++ {
		_ = f.MergeCell("Sheet1", string(rune('I'-1+i))+"2", string(rune('I'-1+i))+"3")
	}

	for i := 0; i < 5; i++ {
		_ = f.MergeCell("Sheet1", string(rune('W'-1+i))+"1", string(rune('W'-1+i))+"3")
	}
	_ = f.MergeCell("Sheet1", "AA1", "AA3")

	for i, v := range data {
		_ = f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+4), i+1)
		i = i + 4
		_ = f.SetCellValue("Sheet1", "B"+strconv.Itoa(i), v.NamaProv+"/"+v.NamaKab)
		_ = f.SetCellValue("Sheet1", "C"+strconv.Itoa(i), v.NamaKec)
		_ = f.SetCellValue("Sheet1", "D"+strconv.Itoa(i), v.DaerahIrigasi)
		_ = f.SetCellValue("Sheet1", "E"+strconv.Itoa(i), v.LuasWilayah)
		_ = f.SetCellValue("Sheet1", "F"+strconv.Itoa(i), v.JumlahP3A)
		_ = f.SetCellValue("Sheet1", "G"+strconv.Itoa(i), v.NamaP3A)
		_ = f.SetCellValue("Sheet1", "H"+strconv.Itoa(i), v.LuasLayananP3A)
		_ = f.SetCellValue("Sheet1", "I"+strconv.Itoa(i), v.TahunPembentukan)
		_ = f.SetCellValue("Sheet1", "J"+strconv.Itoa(i), v.LamKplDesa)
		_ = f.SetCellValue("Sheet1", "K"+strconv.Itoa(i), v.SKBupati)
		_ = f.SetCellValue("Sheet1", "L"+strconv.Itoa(i), v.AkteNotaris)
		_ = f.SetCellValue("Sheet1", "M"+strconv.Itoa(i), v.NoPendaftaran)
		_ = f.SetCellValue("Sheet1", "N"+strconv.Itoa(i), v.Ketua)
		_ = f.SetCellValue("Sheet1", "O"+strconv.Itoa(i), v.Wakil)
		_ = f.SetCellValue("Sheet1", "P"+strconv.Itoa(i), v.Sekretaris)
		_ = f.SetCellValue("Sheet1", "Q"+strconv.Itoa(i), v.Bendahara)
		_ = f.SetCellValue("Sheet1", "R"+strconv.Itoa(i), v.SekTeknik)
		_ = f.SetCellValue("Sheet1", "S"+strconv.Itoa(i), v.SekOP)
		_ = f.SetCellValue("Sheet1", "T"+strconv.Itoa(i), v.SekBisnis)
		_ = f.SetCellValue("Sheet1", "U"+strconv.Itoa(i), v.JumlahAnggota)
		_ = f.SetCellValue("Sheet1", "V"+strconv.Itoa(i), v.NoADRT)
		_ = f.SetCellValue("Sheet1", "W"+strconv.Itoa(i), v.Sekretariat)
		_ = f.SetCellValue("Sheet1", "X"+strconv.Itoa(i), v.PresentasiPerempuanP3A)
		_ = f.SetCellValue("Sheet1", "Y"+strconv.Itoa(i), v.ArealTersier)
		_ = f.SetCellValue("Sheet1", "Z"+strconv.Itoa(i), v.PengisianBuku)
		_ = f.SetCellValue("Sheet1", "AA"+strconv.Itoa(i), v.Iuran)
		_ = f.SetCellValue("Sheet1", "AB"+strconv.Itoa(i), v.Operasi)
		_ = f.SetCellValue("Sheet1", "AC"+strconv.Itoa(i), v.Partisipatif)
		_ = f.SetCellValue("Sheet1", "AD"+strconv.Itoa(i), v.PolaTanam)
		_ = f.SetCellValue("Sheet1", "AE"+strconv.Itoa(i), v.UsahaTani)
		_ = f.SetCellValue("Sheet1", "AF"+strconv.Itoa(i), v.Keterangan)

	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	t := time.Now()
	name := "Report - " + t.Format("2006-01-02") + ".xlsx"
	if err := f.SaveAs(name); err != nil {
		fmt.Println(err)
	}
	//Delete File
	defer os.Remove(name)

	return ctx.File(name)
}

func (c *P3Controller) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := c.service.FindById(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/inventaris/v1/update/:id",
	}

	return Render(ctx, "Home", "p3a/update", c.Menu, append(c.BreadCrumbs, breadCrumbs), data)
}

func (c *P3Controller) DoUpdate(ctx echo.Context) error {
	var entity request.UpdateInventaris
	id := ctx.Param("id")
	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	name := []string{"lampiran_tahun_pembentukan", "lampiran_kep_dc", "lampiran_sk_bupati", "lampiran_akte_notaris", "lampiran_pendaftaran", "lampiran_ad_art", "lampiran_sekretariat"}
	var namaFile []string
	var prefixFile []string
	fmt.Println("Nama File : ", name)
	if name != nil {
		for i := range name {
			file, err := ctx.FormFile(name[i])
			if err == nil {
				src, _ := file.Open()
				defer src.Close()

				// Destination
				t := time.Now().UnixNano()
				nf := name[i] + "_" + strconv.FormatInt(t, 10) + "_" + file.Filename
				nama := "static/image/" + nf
				dst, _ := os.Create(nama)
				defer dst.Close()

				// Copy
				_, err = io.Copy(dst, src)
				if err != nil {
					log.Error("[Error] ", err)
					return c.InternalServerError(ctx, err)
				}
				prefixFile = append(prefixFile, name[i])
				i++
				namaFile = append(namaFile, nf)			}
		}
	}
	// Update Data Status Legal
	_, err := c.service.UpdateStatusLegal(entity.IDStatus, entity, namaFile,prefixFile)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	// Update Data Kepengurusan
	_, err = c.service.UpdatePengurus(entity.IDPengurusan, entity,namaFile, prefixFile)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	// Update Data Teknik Irigasi
	_, err = c.service.UpdateIrigasi(entity.IDIrig, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	// Update Data Teknik Pertanian
	_, err = c.service.UpdatePertanian(entity.IDTani, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	//Update Data to Table p3a
	_, err = c.service.UpdateById(id, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	time.Sleep(1 * time.Second)
	return ctx.Redirect(http.StatusFound, "/admin/v1/inventaris")
}

func (c *P3Controller) Delete(ctx echo.Context) error {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return c.Ok(ctx, nil)
}
func (c *P3Controller) Detail(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := c.service.FindById(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/inventaris/v1/update/:id",
	}

	return Render(ctx, "Home", "p3a/detail", c.Menu, append(c.BreadCrumbs, breadCrumbs), data)
}
