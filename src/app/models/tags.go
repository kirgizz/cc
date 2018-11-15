package models

import "app/services"

type tags struct {
	id   int `gorm:"primary_key:true"`
	name string
}

func CreateTableTags() {
	services.GetInstanceDB().CreateTable(&tags{})
}
