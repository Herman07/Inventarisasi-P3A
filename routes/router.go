package routes

import (
	"Inventarisasi-P3A/config"
	"Inventarisasi-P3A/middleware"
	"Inventarisasi-P3A/utils/session"
	"encoding/json"
	"github.com/foolin/goview"
	echotemplate "github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

func Api(e *echo.Echo, db *gorm.DB) {

	mv := echotemplate.NewMiddleware(goview.Config{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/master",
		Funcs: template.FuncMap{
			"session": func(ctx echo.Context) session.UserInfo {
				ses, _ := session.Manager.Get(ctx, session.SessionId)
				dataSes, _ := json.Marshal(ses)
				var data session.UserInfo
				userInfo := session.UserInfo{
					ID: data.ID,
					Username: data.Username,
					TypeUser: data.TypeUser,
					Foto : data.Foto,
				}
				_ = json.Unmarshal(dataSes, &userInfo)
				return userInfo
			},
		},
	})
	authorizationMiddleware := middleware.NewAuthorizationMiddleware(db)

	adminGroup := e.Group("/admin",mv, middleware.SessionMiddleware(session.Manager))
	g := adminGroup.Group("/v1",mv, authorizationMiddleware.AuthorizationMiddleware([]string{"1","2"}))


	e.GET("/", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	})
	authController := config.InjectAuthController(db)
	e.GET("/login", authController.Index)
	e.POST("/do-login", authController.Login)
	e.POST("/logout", authController.Logout)

	{
		invenGroup := g.Group("/inventaris",mv,authorizationMiddleware.AuthorizationMiddleware([]string{"1","2"}))
		dashboardController := config.InjectDashboardController(db)
		invenGroup.GET("", dashboardController.Index)
		invenGroup.GET("/add", dashboardController.Add)
		invenGroup.POST("/store", dashboardController.AddData)
		invenGroup.GET("/generate", dashboardController.GenerateExcel)
		invenGroup.GET("/tables", dashboardController.GetDetail)
		invenGroup.DELETE("/delete/:id", dashboardController.Delete)
		invenGroup.GET("/update/:id", dashboardController.Update)
		invenGroup.GET("/detail/:id", dashboardController.Detail)
		invenGroup.POST("/do-update/:id", dashboardController.DoUpdate)

	}


	m := g.Group("/master-data", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	provController := config.InjectMasterController(db)
	p := m.Group("/provinsi", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	p.GET("", provController.Index)
	p.GET("/table", provController.GetDetail)
	p.GET("/update/:id", provController.Update)
	e.POST("/prov/doupdate/:id", provController.DoUpdate)
	e.POST("/prov/add", provController.AddData)
	e.POST("/prov/:id",provController.FindByID)
	e.DELETE("/prov/:id/delete",provController.Delete)


	kabController := config.InjectKabController(db)
	kb := m.Group("/kab", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	kb.GET("", kabController.Index)
	kb.GET("/table", kabController.GetDetail)
	kb.GET("/update/:id", kabController.Update)
	e.POST("/kab/doupdate/:id", kabController.DoUpdate)
	e.POST("/kab/addkab", kabController.AddData)
	e.POST("kab/:id",kabController.FindByID)
	e.DELETE("/kab/:id/delete",kabController.Delete)

	kecController := config.InjectKecController(db)
	kc := m.Group("/kec", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	kc.GET("", kecController.Index)
	kc.GET("/table", kecController.GetDetail)
	kc.GET("/update/:id", kecController.Update)
	e.POST("/kec/doupdate/:id", kecController.DoUpdate)
	e.POST("/kec/addkec", kecController.AddData)
	e.POST("/kec/:id",kecController.FindByID)
	e.DELETE("/kec/:id/delete",kecController.Delete)

	userController := config.InjectUserController(db)
	u := g.Group("/user", authorizationMiddleware.AuthorizationMiddleware([]string{"1","2"}))
	u.GET("/create", userController.Index)
	u.GET("/profile/:id",userController.Profile)
	u.GET("/profile/update/:id",userController.UpdateProfile)
	u.POST("/update_profile/:id",userController.DoUpdateProfile)
	u.POST("/adduser",userController.AddData)
	u.GET("/table", userController.GetDetail)
	u.GET("",userController.TableUser)
	u.GET("/update/:id", userController.Update)
	u.POST("/updateuser/:id", userController.DoUpdate)
	u.DELETE("/delete/:id",userController.Delete)
}