package auth

import (
	"app/models"
	"app/services"
	"app/utils"
	"github.com/ivahaev/go-logger"
)


//This registration is bad stuff. After commit new user is users create new transaction, witch is bad
func RegisterUser(email, nickname, password string) (int, string) {


	var u models.User
	var c models.Credentials

	u.Email = email
	u.NickName = nickname

	//change function to create password hash
	passwordHash := utils.CreateHash(password)

	db := services.GetInstanceDB()

	tx := db.Begin()
	if tx.Error != nil {
		logger.Error(tx.Error)
		return 500, tx.Error.Error()
	}
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		logger.Error(err)
		return 500, err.Error()
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		logger.Error(err)
		return 500, tx.Commit().Error.Error()
	}

	err := db.Where("email = ?", email).First(&u).Error
	if err != nil {
		logger.Error(err)
		return 500, err.Error()
	}

	//check password cerrtions
	c.Password = passwordHash
	c.UserId = u.Id

	tx = db.Begin()
	if tx.Error != nil {
		logger.Error(err)
		return 500, tx.Error.Error()
	}
	//rollback dont rollback transaction
	if err := tx.Create(&c).Error; err != nil {
		tx.Rollback()
		logger.Error(err)
		return 500, err.Error()
	}
	if tx.Commit().Error != nil {
		logger.Error(err)
		tx.Rollback()
		return 500, tx.Commit().Error.Error()
	}
	//CREATE SSID hash for user
	return 200, "User registered"
}