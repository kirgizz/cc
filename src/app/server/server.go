package server

import (
	"app/auth"
	"app/models"
	"app/services"
	"encoding/json"
	"fmt"
	"github.com/ivahaev/go-logger"
	"net/http"
	"time"
)

var H string

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		logger.Crit(err)
	}

	logger.Info(r.Form)
	w.Write([]byte("Hello from not implemented function"))
})

var GetArticles = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	a := models.GetArticles()
	b, err := json.Marshal(a)
	if err != nil {
		logger.Error(err)
	}
	w.Write([]byte(b))
})


var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	//Check form values?
	status, result := auth.Login(r.Form.Get("name"),r.Form.Get("email"), r.Form.Get("password"), false )
	w.WriteHeader(status)
	w.Write([]byte(result))

})


//RegisterUser(nickname, email, password, name string, birthday time.Time)

var Register = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var birthday time.Time
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	//Check form values?
	birthday, err = time.Parse("2006-01-02T15:04:05.000Z", r.Form.Get("birthday"))
	status, result := auth.RegisterUser(r.Form.Get("nickname"),r.Form.Get("email"), r.Form.Get("password"),r.Form.Get("name"), birthday)
	w.WriteHeader(status)
	w.Write([]byte(result.Error()))
})

var AddArticle = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var u models.User
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	//where get current user???
	services.GetInstanceDB().Where("email = ?", "la-la-la").Find(&u)
	picturePath := "testPath"
	status, result := models.AddArticle(r.Form.Get("name"), r.Form.Get("body"), picturePath, u.Id)
	w.WriteHeader(status)
	w.Write([]byte(result.Error()))
	
})





//func IncidentGet(w http.ResponseWriter, r *http.Request) {
//	result := models.Result{Status: 404}
//	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
//	incident, ok := models.IncidentResponceGetByID(id)

//	if ok == true {
//		result.Status = 200
//		result.Data = incident
//	}

//	jsonAnswer, _ := json.Marshal(result)
//	w.WriteHeader(result.Status)
//	io.WriteString(w, string(jsonAnswer))
//}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}