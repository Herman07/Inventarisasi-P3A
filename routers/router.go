package routers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"skeleton-echo/config"
	"skeleton-echo/middleware"
	"skeleton-echo/utils/session"
)

func Api(e *echo.Echo, db *gorm.DB) {
	vGroup := e.Group("/inventaris",middleware.SessionMiddleware(session.Manager))
	authorizationMiddleware := middleware.NewAuthorizationMiddleware(db)

	e.GET("/", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect,"/login")
	})
	authController := config.InjectAuthController(db)
	e.GET("/login", authController.Index)
	e.POST("/do-login", authController.Login)
	e.POST("/logout", authController.Logout)

	dashboardController := config.InjectDashboardController(db)
	g := vGroup.Group("/v1", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	g.GET("/admin", dashboardController.Index)
	g.GET("/add", dashboardController.Add)
	g.POST("/create",dashboardController.AddData)
	g.GET("/tables",dashboardController.GetDetail)
	g.GET("/table",dashboardController.GetData)
	g.GET("/table/:id",dashboardController.Detail)
	g.GET("/update/:id", dashboardController.Update)
	g.POST("/do-update/:id",dashboardController.DoUpdate)
	g.DELETE("/delete/:id",dashboardController.Delete)


	m := g.Group("/master-data", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	masterController := config.InjectMasterController(db)
	p := m.Group("/provinsi", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	p.GET("",masterController.Index)
	p.GET("/add",masterController.Store)
	p.GET("/table",masterController.GetDetail)


}