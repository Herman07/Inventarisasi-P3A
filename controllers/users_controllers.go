package controllers

import (
	"Inventarisasi-P3A/models"
	"Inventarisasi-P3A/request"
	"Inventarisasi-P3A/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type UsersDataController struct {
	BaseFrontendController
	Controller
	service *services.UsersDataService
}

func NewUsersDataController(services *services.UsersDataService) UsersDataController {
	return UsersDataController{
		service: services,
		BaseFrontendController: BaseFrontendController{
			Menu:        "Users",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}
func (c *UsersDataController) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Users",
		"link": "/admin/v1/users",
	}
	return Render(ctx, "Home", "users/create", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}
func (c *UsersDataController) Profile(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := c.service.FindById(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	breadCrumbs := map[string]interface{}{
		"menu": "Users",
		"link": "/admin/v1/users",
	}
	return Render(ctx, "Home", "users/profile", c.Menu, append(c.BreadCrumbs, breadCrumbs), data)
}
func (c *UsersDataController) UpdateProfile(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := c.service.FindById(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	breadCrumbs := map[string]interface{}{
		"menu": "Users",
		"link": "/admin/v1/user/profile/update",
	}
	return Render(ctx, "Home", "users/update_profile", c.Menu, append(c.BreadCrumbs, breadCrumbs), data)
}



func (c *UsersDataController) TableUser(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Users",
		"link": "/admin/v1/users/table",
	}
	return Render(ctx, "Home", "users/index", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}
func (c *UsersDataController) GetDetail(ctx echo.Context) error {

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
		action = `<a href="/admin/v1/user/update/` + (v.ID) + `" class="btn btn-success" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fa fa-edit"></i></a>
		<a onclick="DeleteUser('` + v.ID + `')" class="btn btn-danger" title="Delete" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fa fa-trash"></i></a>`
		listOfData[k] = map[string]interface{}{
			"id_user":       v.ID,
			"nama":          v.Nama,
			"alamat":        v.Alamat,
			"tgl_lahir":     v.TglLahir,
			"no_telepon":    v.NoTlp,
			"jenis_kelamin": v.JenisKelamin,
			"email":         v.Email,
			"action":        action,
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
func (c *UsersDataController) Detail(ctx echo.Context) error {
	id := ctx.Param("id")
	_, err := c.service.FindById(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	breadCrumbs := map[string]interface{}{
		"menu": "Users",
		"link": "/admin/v1/users/detail",
	}
	data := models.User{}
	return Render(ctx, "Home", "users/detail", c.Menu, append(c.BreadCrumbs, breadCrumbs), data)
}

func (c *UsersDataController) AddData(ctx echo.Context) error {
	var entity request.UsersReq

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	user, err := c.service.Create(entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	_, err = c.service.CreateAkun(entity, user.ID)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return ctx.Redirect(http.StatusFound, "/admin/v1/user")
}

func (c *UsersDataController) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := c.service.FindById(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/inventaris/v1/user/update/:id",
	}

	return Render(ctx, "Home", "users/update", c.Menu, append(c.BreadCrumbs, breadCrumbs), data)
}
func (c *UsersDataController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return c.Ok(ctx, nil)
}

func (c *UsersDataController) DoUpdate(ctx echo.Context) error {
	var entity request.UsersReq
	id := ctx.Param("id")
	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	//Update User
	_, err := c.service.UpdateUser(id, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	//Update Akun
	_, err = c.service.UpdateAkun(entity.ID, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return ctx.Redirect(302, "/admin/v1/user")
}

func (c *UsersDataController) DoUpdateProfile(ctx echo.Context) error {
	var entity request.UsersReq
	id := ctx.Param("id")
	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	//Update User
	_, err := c.service.UpdateUser(id, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	//Update Foto
	file, _ := ctx.FormFile("foto")
	if file != nil {
		src, _ := file.Open()
		defer src.Close()

		// Destination
		t := time.Now().UnixNano()
		nf := entity.Nama + "_" + strconv.FormatInt(t, 10) + "_" + file.Filename
		nama := "static/image/" + nf
		dst, _ := os.Create(nama)
		defer dst.Close()

		// Copy
		_, err := io.Copy(dst, src)
		if err != nil {
			log.Error("[Error] ", err)
			return c.InternalServerError(ctx, err)
		}
		//Update Akun
		_, err = c.service.UpdateFoto(entity.ID,nf)
		if err != nil {
			return c.InternalServerError(ctx, err)
		}
	}
	//Update Akun
	_, err = c.service.UpdateFoto(entity.ID,"")
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	return ctx.Redirect(302, "/admin/v1/user/profile/"+entity.ID)
}