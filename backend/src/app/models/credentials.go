package models

import (
	"app/services"
	"app/utils"
	"github.com/ivahaev/go-logger"
	"net/http"
)


type Credentials struct {
	Id        	int       	`gorm:"primary_key:true"`
	UserId     	int       	`gorm:"column:user_id;unique"`
	Ssid	   	string		`gorm:"column:ssid"`
	Password	string 		`gorm:"column:password;not null;" sql:"DEFAULT:NULL"`
}


func (m *Model)CreateTableCredentials() {
	services.GetInstanceDB().CreateTable(&Credentials{})
}



func (c *Credentials) GetUserCredentials(id int) error {
	return services.GetInstanceDB().Where("user_id = ?", id).Find(&c).Error
}

func (c *Credentials) UpdateCredentials() error {
	c.Ssid = utils.RandString(16)
	return services.GetInstanceDB().Model(&c).Where("user_id = ?", c.UserId).Update("ssid", c.Ssid).Error
}

func (c *Credentials) CheckCredentials(email, password string) (int, error){
	//TODO return error codes not http
	var u User
	var err error
	if err = u.GetUser(email); err != nil {
		logger.Error(err)
		return http.StatusNotFound, err
	}
	if err = c.GetUserCredentials(u.Id); err != nil {
		logger.Error(err)
		return http.StatusNotFound, err
	}

	if c.Password != utils.CreateHash(password) {
		logger.Error(c)
		return http.StatusForbidden, nil
	}

	if err = c.UpdateCredentials(); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, err
}

func (c *Credentials) GetUserSession(user_id int) (error){
	err := services.GetInstanceDB().Where("user_id = ?", user_id).Find(&c).Error
	if err != nil {
		return err
	}
	return err
}


func (c *Credentials) GetUser(ssid string) (error) {
	err := services.GetInstanceDB().Where("ssid = ?", ssid).Find(&c).Error
	if err != nil {
		return err
	}
	return err
}
