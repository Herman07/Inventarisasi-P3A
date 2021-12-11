package main

import (
	"encoding/gob"
	echotemplate "github.com/foolin/goview/supports/echoview-v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"Inventarisasi-P3A/config"
	middlewareFunc "Inventarisasi-P3A/middleware"
	"Inventarisasi-P3A/models"
	"Inventarisasi-P3A/routes"
	"Inventarisasi-P3A/utils/session"
)

func main() {
	e := echo.New()

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.Render(http.StatusNotFound, "auth/404.html", nil)
	}

	e.Renderer = echotemplate.Default()

	// load ENV
	var err = godotenv.Load()
	if err != nil {
		log.Fatal("ERROR ", err)
	}
	gob.Register(session.UserInfo{})
	gob.Register(session.FlashMessage{})
	gob.Register(models.Users{})
	gob.Register(map[string]interface{}{})

	// Load static dashboard
	e.Static("/style", "static/assets")
	// Load static auth
	e.Static("/login/template", "static/auth")
	e.Static("/static/image", "static/image")

	//DB Connected
	db := config.Landscape()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	session.Manager = session.NewSessionManager(middlewareFunc.NewCookieStore())

	routes.Api(e,db)

	//Port
	e.Logger.Fatal(e.Start(":"+os.Getenv("APP_PORT")))

}
