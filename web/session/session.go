package session

import (
	"x-ui-scratch/database/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	loginUser   = "LOGIN_USER"
	defaultPath = "/"
)

func IsLogin(c *gin.Context) bool {
	return GetLoginUser(c) != nil
}

func GetLoginUser(c *gin.Context) *model.User {
	s := sessions.Default(c)
	obj := s.Get(loginUser)
	if obj == nil {
		return nil
	}
	user, ok := obj.(model.User)
	if !ok {
		return nil
	}
	return &user
}

func SetMaxAge(c *gin.Context, maxAge int) error {
	s := sessions.Default(c)
	s.Options(sessions.Options{
		Path:     defaultPath,
		MaxAge:   maxAge,
		HttpOnly: true,
	})
	return s.Save()
}

func SetLoginUser(c *gin.Context, user *model.User) error {
	s := sessions.Default(c)
	s.Set(loginUser, user)
	return s.Save()
}
