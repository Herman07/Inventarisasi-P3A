package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"skeleton-echo/request"
	"skeleton-echo/services"
)

type PengurusDataController struct {
	BaseFrontendController
	Controller
	service *services.PengurusDataService
}

func NewPengurusDataController(services *services.PengurusDataService) PengurusDataController {
	return PengurusDataController{
		service: services,
		BaseFrontendController: BaseFrontendController{
			Menu:        "Kepengurusan",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}
func (c *PengurusDataController) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Kepengurusan",
		"link": "/inventaris/v1/kepengurus",
	}
	return Render(ctx, "Home", "kepengurusan/index", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}

//func (c *KabDataController) Store(ctx echo.Context) error {
//	breadCrumbs := map[string]interface{}{
//		"menu": "Home",
//		"link": "/inventaris/v1/master-data/kab/add",
//	}
//	return Render(ctx, "Home", "master-data/kabupaten/add", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
//}
//func (c *KabDataController) Update(ctx echo.Context) error {
//	id := ctx.Param("id")
//	data, err := c.service.FindById(id)
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//
//	breadCrumbs := map[string]interface{}{
//		"menu": "Home",
//		"link": "/inventaris/v1/master-data/kab/update/:id",
//	}
//	dataKab := models.MasterDataKab{
//		ID:         data.ID,
//		Kabupaten:   data.Kabupaten,
//		IDProv:  data.IDProv,
//	}
//	return Render(ctx, "Home", "master-data/kabupaten/update", c.Menu, append(c.BreadCrumbs, breadCrumbs), dataKab)
//}
func (c *PengurusDataController) Store(ctx echo.Context) error {
	var entity request.PengurusReq

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	_, err := c.service.Create(entity)
	//entity.CreatedAt = time.Now()
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
//
//func (c *KabDataController) DoUpdate(ctx echo.Context) error {
//	var entity request.KabReq
//	id := ctx.Param("id_kab")
//	if err := ctx.Bind(&entity); err != nil {
//		return ctx.JSON(400, echo.Map{"message": "error binding data"})
//	}
//	data, err := c.service.UpdateById(id, entity)
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//	fmt.Println(data)
//	return ctx.Redirect(302, "/inventaris/v1/master-data/kab")
//}
//
//func (c *KabDataController) Delete(ctx echo.Context) error {
//	id := ctx.Param("id_kab")
//
//	err := c.service.Delete(id)
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//	return c.Ok(ctx,nil)
//}
