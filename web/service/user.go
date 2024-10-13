package service

import (
	"x-ui-scratch/database"
	"x-ui-scratch/database/model"
	"x-ui-scratch/logger"

	"gorm.io/gorm"
)

type UserService struct{}

func (s *UserService) CheckUser(username string, password string, secret string) *model.User {
	db := database.GetDB()

	user := &model.User{}
	err := db.Model(model.User{}).
		Where("username = ? and password = ? and login_secret = ?", username, password, secret).
		First(user).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		// TODO
		logger.Info("check user err:", err)
		return nil
	}
	return user
}
