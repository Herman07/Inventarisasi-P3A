package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"skeleton-echo/models"
	"skeleton-echo/request"
	"skeleton-echo/services"
	"strconv"
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
	return Render(ctx, "Home", "p3a/index", c.Menu, append(c.BreadCrumbs, breadCrumbs),nil)
}

func (c *P3Controller) Add(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/admin/v1/inventaris/add",
	}
	return Render(ctx, "Home", "p3a/create", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}

func (c *P3Controller) GetDetail(ctx echo.Context) error {

	draw, err := strconv.Atoi(ctx.Request().URL.Query().Get("draw"))
	start, err := strconv.Atoi(ctx.Request().URL.Query().Get("start"))
	search := ctx.Request().URL.Query().Get("search[value]")
	length, err := strconv.Atoi(ctx.Request().URL.Query().Get("length"))
	order, err := strconv.Atoi(ctx.Request().URL.Query().Get("order[0][column]"))
	orderName := ctx.Request().URL.Query().Get("columns[" + strconv.Itoa(order) + "][name]")
	orderAscDesc := ctx.Request().URL.Query().Get("order[0][dir]")

	recordTotal, recordFiltered, data, err := c.service.QueryDatatable(search, orderAscDesc, orderName, length, start)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	var createdAt string
	var action string
	listOfData := make([]map[string]interface{}, len(data))
	for k, v := range data {
		action = `<a href="/inventaris/v1/update/` + (v.ID) + `" class="btn btn-success btn-bold btn-upper" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fas fa-edit"></i></a>
		<a href="javascript:;" onclick="Delete('` + v.ID + `')" class="btn btn-danger btn-bold btn-upper" title="Delete" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fas fa-trash"></i></a>`
		time := v.CreatedAt
		createdAt = time.Format("2006-01-02")
		listOfData[k] = map[string]interface{}{
			"id_p3a":              v.ID,
			"no_urut":             v.NoUrut,
			"nama_p3a":            v.NamaP3A,
			"jumlah_p3a":          v.JumlahP3A,
			"nama_daerah_irigasi": v.DaerahIrigasi,
			"luas_wilayah":        v.LuasWilayah,
			"luas_layanan_p3a":    v.LuasLayananP3A,
			"keterangan":          v.Keterangan,
			"created_at":          createdAt,
			"action":              action,
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
//func (c *P3Controller) GetData(ctx echo.Context) error {
//	dataReq := models.Inventaris{}
//
//	data, err := c.service.GetAll(dataReq)
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//
//	return c.Ok(ctx, data)
//}
//func (c *P3Controller) Detail(ctx echo.Context) error {
//	id := ctx.Param("id")
//	data, err := c.service.FindById(id)
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//	return c.Ok(ctx, data)
//}

func (c *P3Controller) AddData(ctx echo.Context) error {
	var entity request.RequestInventaris
	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}

	//Store Data Status Legal
	statusLegal, err := c.service.CreateStatusLegal(entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	// Store Data Kepengurusan
	pengurus , err := c.service.CreatePengurus(entity)
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
	p3a, err := c.service.CreateDataP3a(entity,statusLegal.ID,pengurus.ID,irigasi.ID,pertanian.ID)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	return c.Ok(ctx, p3a)
}

//func (c *P3Controller) Update(ctx echo.Context) error {
//	id := ctx.Param("id")
//	data, err := c.service.FindById(id)
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//
//	breadCrumbs := map[string]interface{}{
//		"menu": "Home",
//		"link": "/inventaris/v1/update/:id",
//	}
//	dataInventaris := models.Inventaris{
//		ID:             data.ID,
//		NoUrut:         data.NoUrut,
//		NamaP3A:        data.NamaP3A,
//		JumlahP3A:      data.JumlahP3A,
//		DaerahIrigasi:  data.DaerahIrigasi,
//		LuasWilayah:    data.LuasWilayah,
//		LuasLayananP3A: data.LuasLayananP3A,
//	}
//	return Render(ctx, "Home", "update", c.Menu, append(c.BreadCrumbs, breadCrumbs), dataInventaris)
//}
//
//func (c *P3Controller) DoUpdate(ctx echo.Context) error {
//	var entity request.RequestInventaris
//	id := ctx.Param("id")
//	if err := ctx.Bind(&entity); err != nil {
//		return ctx.JSON(400, echo.Map{"message": "error binding data"})
//	}
//	data, err := c.service.UpdateById(id, entity)
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//	fmt.Println(data)
//	return ctx.Redirect(302, "/inventaris/v1/admin")
//}

//func (c *P3Controller) Delete(ctx echo.Context) error {
//	id := ctx.Param("id")
//
//	err := c.service.Delete(id)
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//	return c.Ok(ctx, nil)
//}
