package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"skeleton-echo/models"
	"skeleton-echo/request"
	"skeleton-echo/services"
	"strconv"
)

type KecDataController struct {
	BaseFrontendController
	Controller
	service *services.KecDataService
}

func NewKecDataController(services *services.KecDataService) KecDataController {
	return KecDataController{
		service: services,
		BaseFrontendController: BaseFrontendController{
			Menu:        "Kecamatan Data",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}
func (c *KecDataController) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/inventaris/v1/master-data/kec",
	}
	return Render(ctx, "Home", "master-data/kecamatan/index", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}

func (c *KecDataController) Store(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/inventaris/v1/master-data/kec/add",
	}
	return Render(ctx, "Home", "master-data/kecamatan/add", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}
func (c *KecDataController) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := c.service.FindById(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/inventaris/v1/master-data/kec/update/:id",
	}
	dataKec := models.MasterDataKec{
		ID:        data.ID,
		Kecamatan: data.Kecamatan,
		IDKab:     data.IDKab,
	}
	return Render(ctx, "Home", "master-data/kecamatan/update", c.Menu, append(c.BreadCrumbs, breadCrumbs), dataKec)
}

func (c *KecDataController) GetDetail(ctx echo.Context) error {

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
	//var createdAt string
	var action string
	listOfData := make([]map[string]interface{}, len(data))
	for k, v := range data {
		action = `<a data-toggle="modal" data-target="#modal-update" class="btn btn-success btn-bold btn-upper" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fas fa-edit"></i></a>
		<a href="javascript:;" onclick="Delete('` + v.ID + `')" class="btn btn-danger btn-bold btn-upper" title="Delete" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fas fa-trash"></i></a>`
		//time := v.CreatedAt
		//createdAt = time.Format("2006-01-02")
		listOfData[k] = map[string]interface{}{
			"nama_kecamatan": v.Kecamatan,
			"id_kec":         v.ID,
			"action":         action,
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

func (c *KecDataController) AddData(ctx echo.Context) error {
	var entity request.KecReq

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	_, err := c.service.Create(entity)
	//entity.CreatedAt = time.Now()
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return ctx.Redirect(302, "/inventaris/v1/master-data/kec")
}

func (c *KecDataController) DoUpdate(ctx echo.Context) error {
	var entity request.KecReq
	id := ctx.Param("id_kec")
	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	data, err := c.service.UpdateById(id, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	fmt.Println(data)
	return ctx.Redirect(302, "/inventaris/v1/master-data/kec")
}

func (c *KecDataController) Delete(ctx echo.Context) error {
	id := ctx.Param("id_kec")

	err := c.service.Delete(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return c.Ok(ctx, nil)
}
