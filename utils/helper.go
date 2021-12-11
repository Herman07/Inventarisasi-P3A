package utils

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
	"Inventarisasi-P3A/models"
	"Inventarisasi-P3A/utils/session"
	"time"
)

func GenerateTimeNow() string {
	//fetching current time
	loc, _ := time.LoadLocation("Asia/Jakarta")
	currentTime := time.Now().In(loc).Format(time.RFC3339)
	//differnce between pastdate and current date
	return currentTime
}

func ItemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)
	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}
func GetUserInfoFromContext(ctx echo.Context, db *gorm.DB) (userModel models.Users, err error) {
	result, err := session.Manager.Get(ctx, session.SessionId)
	if err != nil {
		panic(err)
	}
	userInfo := result.(session.UserInfo)
	err = db.Model(models.Users{}).Preload(clause.Associations).Where("id_usergroup =?", userInfo.ID).First(&userModel).Error
	return
}