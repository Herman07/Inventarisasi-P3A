package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"net/http"
	"Inventarisasi-P3A/utils"
)

type authorizationConfig struct {
	db      *gorm.DB
	Skipper middleware.Skipper
}

func NewAuthorizationMiddleware(db *gorm.DB) *authorizationConfig {
	return &authorizationConfig{
		db: db,
		Skipper: func(context echo.Context) bool {
			return false
		},
	}
}

func (m *authorizationConfig) AuthorizationMiddleware(roles []string) echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if m.Skipper(context) {
				return handlerFunc(context)
			}

			//check role
			userModel, err := utils.GetUserInfoFromContext(context, m.db)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return echo.NewHTTPError(404, "you must login before access this resource")
				}
				return echo.NewHTTPError(500, "error get user from context ")
			}
			if !utils.ItemExists(roles, userModel.TypeUser) {
				return echo.NewHTTPError(http.StatusForbidden, "this user role don't have permission to access this resource ")
			}else {
				return handlerFunc(context)
			}

		}
	}
}
