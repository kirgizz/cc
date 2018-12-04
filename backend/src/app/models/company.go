package models

import "app/services"

type Company struct {
	Name 			string
	Advertising		string
}


func CreateTableCompany() {
	services.GetInstanceDB().CreateTable(&Company{})
}