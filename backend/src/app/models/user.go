package models

import (
	//	"app/auth"
	"app/services"
	"github.com/ivahaev/go-logger"
	"net/http"
	"time"
)

//`gorm:"not null;unique"`

type User struct {
	Id        	int      	`gorm:"primary_key:true"`
	BirthDate 	time.Time 	`gorm:"column:birthdate"`
	CreatedAt 	time.Time	`gorm:"column:created_at`
	Rating   	int 		`gorm:"column:rating"`
	Name     	string 		`gorm:"column:name"`
	NickName 	string 		`gorm:"column:nickname;unique" sql:"DEFAULT:NULL"`
	Email    	string 		`gorm:"column:email;not null;unique"`
	Avatar   	string 		`gorm:"column:avatar"`
	Status   	string 		`gorm:"column:status"`
	About		string 		`gorm:"column:about"`
	Password	string 		`gorm:"column:password;not null"`
}

func (m *Model)CreateTableUsers() {
	services.GetInstanceDB().CreateTable(&User{})
}


func (u User) RegisterUser() int{
	statusCode := http.StatusOK
	db := services.GetInstanceDB()
	db.Create(&u)
	return statusCode
}

func UpdateUser(name, nickname string) error {
	db := services.GetInstanceDB()
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Model(&User{}).Where("nickname = ?", nickname).Update("name", name).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func DeleteUser(nickname string) error {
	db := services.GetInstanceDB()
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Model(&User{}).Where("nickname = ?", nickname).Update("status", "deleted").Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

func GetUserByNickname(nickname string) *User {
	var u User
	err := services.GetInstanceDB().Where("nickname = ?", nickname).First(&u).Error
	if err != nil {
		logger.Error(err)
	}
	return &u
}

func GetUsersByEmail(email string) *User {
	var u User
	err := services.GetInstanceDB().Where("email = ?", email).First(&u).Error
	if err != nil {
		logger.Error(err)
	}
	return &u
}

func GetUserById(user_id int) *User {
	var u User
	err := services.GetInstanceDB().Where("id = ?", user_id).First(&u).Error
	if err != nil {
		logger.Error(err)
	}
	return &u
}
