package models

import "app/services"

type tags struct {
	Id   int     `gorm:"primary_key:true"`
	Name string  `gorm:"column:name"`
	Publications []*Publication    `gorm:"many2many:publication_tags;"`
}

func (m *Model) CreateTableTags() {
	services.GetInstanceDB().CreateTable(&tags{})
}
