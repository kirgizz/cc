package models

import (
	"app/services"
)


type Credentials struct {
	Id        	int       	`gorm:"primary_key:true"`
	UserId     	int       	`gorm:"column:user_id"`
	//ADD DEPEND USER BY PASSWORD (could delete credential of user_id)
	Password   	string		`gorm:"column:password;not null;unique" sql:"DEFAULT:NULL"`
	Ssid	   	string		`gorm:"column:ssid"`
}


func CreateTableCredentials() {
	services.GetInstanceDB().CreateTable(&Credentials{})
	services.GetInstanceDB().Model(&Comment{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}

