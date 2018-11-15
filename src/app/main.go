package main

import (
	"app/migrations"
	"app/server"
	"github.com/gorilla/mux"
	"github.com/ivahaev/go-logger"
	"io/ioutil"
	"net/http"
)

//createDataase()
//addUsers()
//addArticles()
//fmt.Print(models.GetArticlesByRating(0, "="))
//models.GetArticlesByUserNickname("john")

func OpenIndexHtml() string {
	b, err := ioutil.ReadFile("/home/evgeniy.sergeev/stuff/culture-city-golang/frontend/index.html")
	if err != nil {
		logger.Error(err)
	}
	return string(b)

}

const migrate = true


//LOGGING???
//ADD INTERFACES???

func main() {
	if migrate == true {
		migrations.CreateDBStruct()
	}

	r := mux.NewRouter()

	// Страница по умолчанию для нашего сайта это простой html
	// Статику (картинки, скрипти, стили) будем раздавать
	// по определенному роуту /static/{file}


	r.HandleFunc("/api/login", server.Login).Methods("POST")
	r.HandleFunc("/api/logout", server.NotImplemented).Methods("DELETE")

	r.HandleFunc("/api/register", server.Register).Methods("POST")

	r.HandleFunc("/api/getArticles", server.GetArticles).Methods("GET")
	r.HandleFunc("/api/find", server.NotImplemented).Methods("GET")

	r.HandleFunc("/api/addArticle", server.NotImplemented).Methods("POST")
	r.HandleFunc("/api/updateArticle", server.NotImplemented).Methods("POST")
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("/home/evgeniy.sergeev/stuff/culture-city-golang/frontend/")))

	logger.Info("Started")
	logger.Crit(http.ListenAndServe(":8080", nil))

}
