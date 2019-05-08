package models

import (
	"app/services"
)

type Rubrics struct {
	Id   int                       `json:"id";gorm:"primary_key:true"`
	Name string                    `json:"name";gorm:"column:name"`
	Publications []*Publication    `json:"-";gorm:"many2many:publication_rubrics;"`
}

func (m *Model) CreateTableRubrics() {
	services.GetInstanceDB().CreateTable(&Rubrics{})
}

func GetRubrics() ([]Rubrics, error) {
	var rubrics []Rubrics
	err := services.GetInstanceDB().Find(&rubrics).Error
	return rubrics, err
}
