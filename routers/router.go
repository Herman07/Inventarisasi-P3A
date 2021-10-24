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
	vGroup := e.Group("/inventaris", middleware.SessionMiddleware(session.Manager))
	authorizationMiddleware := middleware.NewAuthorizationMiddleware(db)

	e.GET("/", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	})
	authController := config.InjectAuthController(db)
	e.GET("/login", authController.Index)
	e.POST("/do-login", authController.Login)
	e.POST("/logout", authController.Logout)

	dashboardController := config.InjectDashboardController(db)
	g := vGroup.Group("/v1", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	g.GET("/admin", dashboardController.Index)
	g.GET("/add", dashboardController.Add)
	g.POST("/create", dashboardController.AddData)
	g.GET("/tables", dashboardController.GetDetail)
	g.GET("/table", dashboardController.GetData)
	g.GET("/table/:id", dashboardController.Detail)
	g.GET("/update/:id", dashboardController.Update)
	g.POST("/do-update/:id", dashboardController.DoUpdate)
	g.DELETE("/delete/:id", dashboardController.Delete)

	m := g.Group("/master-data", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	masterController := config.InjectMasterController(db)
	p := m.Group("/provinsi", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	p.GET("", masterController.Index)
	p.GET("/add", masterController.Store)
	p.GET("/table", masterController.GetDetail)
	p.GET("/update/:id", masterController.Update)
	p.POST("/update/:id", masterController.DoUpdate)
	p.POST("/create", masterController.AddData)

	kabController := config.InjectKabController(db)
	kb := m.Group("/kab", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	kb.GET("", kabController.Index)
	kb.GET("/add", kabController.Store)
	kb.GET("/table", kabController.GetDetail)
	kb.GET("/update/:id", kabController.Update)
	kb.POST("/update/:id", kabController.DoUpdate)
	kb.POST("/create", kabController.AddData)

	kecController := config.InjectKecController(db)
	kc := m.Group("/kec", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	kc.GET("", kecController.Index)
	kc.GET("/add", kecController.Store)
	kc.GET("/table", kecController.GetDetail)
	kc.GET("/update/:id", kecController.Update)
	kc.POST("/update/:id", kecController.DoUpdate)
	kc.POST("/create", kecController.AddData)

	statusController := config.InjectStatusController(db)
	//kc := m.Group("/kec", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	m.GET("/status-legal", statusController.Index)
	m.GET("/status-legal/tables", statusController.GetDetail)
	m.GET("/status-legal/add", statusController.Store)
	m.GET("/status-legal/update",statusController.Update)
	//kc.POST("/update/:id",kecController.DoUpdate)
	//kc.POST("/create",kecController.AddData)

	pengurusController := config.InjectPengurusController(db)
	//kc := m.Group("/kec", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	m.GET("/kepengurus", pengurusController.Index)
	//kc.GET("/add",kecController.Store)
	//kc.GET("/table",kecController.GetDetail)
	//kc.GET("/update/:id",kecController.Update)
	//kc.POST("/update/:id",kecController.DoUpdate)
	//kc.POST("/create",kecController.AddData)

	irigasiController := config.InjectIrigasiController(db)
	//kc := m.Group("/kec", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	m.GET("/tk-irigasi", irigasiController.Index)
	//kc.GET("/add",kecController.Store)
	m.GET("/tk-irigasi/table",irigasiController.GetDetail)
	m.GET("/tk-irigasi/update/:id",irigasiController.Update)
	m.POST("/tk-irigasi/update/:id",irigasiController.DoUpdate)
	m.POST("/tk-irigasi/create",irigasiController.AddData)

	taniController := config.InjectPertanianController(db)
	//kc := m.Group("/kec", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	m.GET("/tk-tani", taniController.Index)
	//kc.GET("/add",kecController.Store)
	//kc.GET("/table",kecController.GetDetail)
	//kc.GET("/update/:id",kecController.Update)
	//kc.POST("/update/:id",kecController.DoUpdate)
	//kc.POST("/create",kecController.AddData)

	usersController := config.InjectUsersController(db)
	//kc := m.Group("/kec", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	m.GET("/users", usersController.Index)
	//kc.GET("/add",kecController.Store)
	//kc.GET("/table",kecController.GetDetail)
	//kc.GET("/update/:id",kecController.Update)
	//kc.POST("/update/:id",kecController.DoUpdate)
	//kc.POST("/create",kecController.AddData)
}