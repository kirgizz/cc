package models

import (
	"app/services"
)


type Credentials struct {
	Id        	int       	`gorm:"primary_key:true"`
	UserId     	int       	`gorm:"column:user_id"`
	Password   	string		`gorm:"column:password;not null;unique"`
	Ssid	   	string		`gorm:"column:ssid"`
}

func CreateTableCredentials() {
	services.GetInstanceDB().CreateTable(&Credentials{})
	services.GetInstanceDB().Model(&Comment{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}