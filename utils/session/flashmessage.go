package session

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetFlashMessage(c echo.Context) FlashMessage {
	session, err := Manager.store.Get(c.Request(), "flash-message")
	if err != nil {
		return FlashMessage{}
	}
	fm := session.Flashes()
	var flash FlashMessage
	if len(fm) > 0 {
		log.Info("FLASH MESSAGE ", fm[0])
		flash = fm[0].(FlashMessage)
	}
	if err := session.Save(c.Request(), c.Response()); err != nil {
		log.Fatal("ERROR GET FLASH MESSAGE ", err.Error())
	}
	return flash
}