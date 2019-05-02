package models

import "app/services"

type tags struct {
	Id   int     `gorm:"primary_key:true"`
	Name string  `gorm:"column:name"`
	Articles []*article    `gorm:"many2many:article_tags;"`
}

func (m *Model) CreateTableTags() {
	services.GetInstanceDB().CreateTable(&tags{})
}
