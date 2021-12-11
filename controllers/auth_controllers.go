package controllers

import (
	"Inventarisasi-P3A/request"
	"Inventarisasi-P3A/services"
	"Inventarisasi-P3A/utils/session"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

type FrontAuthController struct {
	BaseFrontendController
	Controller
	service *services.AuthService
}

func NewAuthController(services *services.AuthService) FrontAuthController {
	return FrontAuthController{
		service: services,
		BaseFrontendController: BaseFrontendController{
			Menu:        "Login",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}
func (c *FrontAuthController) Index(ctx echo.Context) error {
	data, _ := session.Manager.Get(ctx,session.SessionId)
	if data == nil{
		return echoview.Render(ctx, http.StatusOK, "auth/login.html", echo.Map{
			"title":        "Login Page",
			"flashMessage": session.GetFlashMessage(ctx),
		})
	}
	return ctx.Redirect(http.StatusTemporaryRedirect, "/admin/v1/inventaris")

}

func (c *FrontAuthController) Login(ctx echo.Context) error {
	var dataReq request.LoginRequest

	if err := ctx.Bind(&dataReq); err != nil {
		return ctx.Redirect(http.StatusFound, "/admin/v1/inventaris")
	}
	data, err := c.service.Login(dataReq)
	if err != nil {
		return ctx.JSON(400, echo.Map{"message": "Gagal Login"})
	}
	userInfo := session.UserInfo{
		ID:       data.ID,
		TypeUser: data.TypeUser,
		Username: data.Username,
		Foto : data.Foto,
	}
	if err := session.Manager.Set(ctx, session.SessionId, &userInfo)
	err != nil {
		return ctx.Redirect(http.StatusFound, "/inventaris/v1/login")
	}
	return ctx.Redirect(http.StatusFound, "/admin/v1/inventaris")
}
func (c *FrontAuthController) Logout(ctx echo.Context) error {
	err := session.Manager.Delete(ctx, session.SessionId)
	if err != nil {
		return ctx.Redirect(302,  "/inventaris/v1/admin")
	}
	return ctx.Redirect(http.StatusFound,  "/login")
}
