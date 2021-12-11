package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"Inventarisasi-P3A/models"
	"Inventarisasi-P3A/request"
	"Inventarisasi-P3A/services"
	"strconv"
)

type ProvDataController struct {
	BaseFrontendController
	Controller
	service *services.ProvDataService
}

func NewProvDataController(services *services.ProvDataService) ProvDataController {
	return ProvDataController{
		service: services,
		BaseFrontendController: BaseFrontendController{
			Menu:        "Master Data",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}
func (c *ProvDataController) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/inventaris/v1/master-data",
	}
	return Render(ctx, "Home", "master-data/provinsi/index", c.Menu, append(c.BreadCrumbs, breadCrumbs),nil)
}
func (c *ProvDataController) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := c.service.FindById(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/inventaris/v1/update/:id",
	}
	dataProv := models.MasterDataProvinsi{
		ID:         data.ID,
		Provinsi:   data.Provinsi,

	}
	return Render(ctx, "Home", "master-data/provinsi/update", c.Menu, append(c.BreadCrumbs, breadCrumbs), dataProv)
}

func (c *ProvDataController) GetDetail(ctx echo.Context) error {

	draw, err := strconv.Atoi(ctx.Request().URL.Query().Get("draw"))
	start, err := strconv.Atoi(ctx.Request().URL.Query().Get("start"))
	search := ctx.Request().URL.Query().Get("search[value]")
	length, err := strconv.Atoi(ctx.Request().URL.Query().Get("length"))
	order, err := strconv.Atoi(ctx.Request().URL.Query().Get("order[0][column]"))
	orderName := ctx.Request().URL.Query().Get("columns[" + strconv.Itoa(order) + "][name]")
	orderAscDesc := ctx.Request().URL.Query().Get("order[0][dir]")

	recordTotal, recordFiltered, data ,err := c.service.QueryDatatable(search,orderAscDesc, orderName, length, start)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	//var createdAt string
	var action string
	listOfData := make([]map[string]interface{}, len(data))
	for k, v := range data {
		action = `<a href="/admin/v1/master-data/provinsi/update/` + (v.ID) + `" class="btn btn-success" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fa fa-edit"></i></a>
		<a onclick="Deleted('` + v.ID + `')" class="btn btn-danger" title="Delete" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fa fa-trash"></i></a>`
		//time := v.CreatedAt
		//createdAt = time.Format("2006-01-02")
		listOfData[k] = map[string]interface{}{
			"nama_prov":    v.Provinsi,
			"id_prov":          v.ID,
			"action": action,
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

func (c *ProvDataController) AddData(ctx echo.Context) error {
	var entity request.ProvinsiReq

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	_, err := c.service.Create(entity)
	//entity.CreatedAt = time.Now()
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return ctx.Redirect(302, "/admin/v1/master-data/provinsi")
}

func (c *ProvDataController) DoUpdate(ctx echo.Context) error {
	var entity request.ProvinsiReq
	id := ctx.Param("id")
	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	data, err := c.service.UpdateById(id, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	fmt.Println(data)
	return ctx.Redirect(302, "/admin/v1/master-data/provinsi")
}

func (c *ProvDataController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return c.Ok(ctx,nil)
}

func (c *ProvDataController) FindByID(ctx echo.Context) error {
	id := ctx.Param("id")

	data , err := c.service.Find(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return c.Ok(ctx,data)
}