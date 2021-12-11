package session

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

const SessionId = "id"

var Manager *ConfigSession

type UserInfo struct {
	Username string `json:"user_name" form:"user_name"`
	TypeUser string `json:"type_users" form:"type_users"`
	ID       string `json:"id_usergroup" form:"id_usergroup"`
	Foto     string `json:"foto" form:"foto"`
}

type FlashMessage struct {
	Type    string
	Message string
	Data    interface{}
}

type ConfigSession struct {
	store    *sessions.CookieStore
	valueKey string
}

func NewSessionManager(store *sessions.CookieStore) *ConfigSession {
	s := new(ConfigSession)
	s.valueKey = "data"
	s.store = store

	return s
}

func (s *ConfigSession) Get(c echo.Context, name string) (interface{}, error) {
	session, err := s.store.Get(c.Request(), name)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, nil
	}
	if val, ok := session.Values[s.valueKey]; ok {
		return val, nil
	} else {
		return nil, nil
	}
}

func (s *ConfigSession) Set(c echo.Context, name string, value interface{}) error {
	session, _ := s.store.Get(c.Request(), name)
	session.Values[s.valueKey] = value

	err := session.Save(c.Request(), c.Response())
	return err
}

func (s *ConfigSession) Delete(c echo.Context, name string) error {
	session, err := s.store.Get(c.Request(), name)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	return session.Save(c.Request(), c.Response())
}

func (s *ConfigSession) GetWithKeyValues(c echo.Context, name string, keyValue string) (interface{}, error) {
	session, err := s.store.Get(c.Request(), name)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, nil
	}
	if val, ok := session.Values[keyValue]; ok {
		return val, nil
	} else {
		return nil, nil
	}
}
