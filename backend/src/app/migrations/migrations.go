package migrations

import (
	"app/models"
)

func CreateDBStruct() {
	models.CreateTableAuthors()
	models.CreateTableArticles()
	models.CreateTableComments()
	models.CreateTableCredentials()
	//models.CreateTableTags()
}
