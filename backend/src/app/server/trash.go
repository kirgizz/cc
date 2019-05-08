package server

import (

	"app/models"
	"app/services"

	"encoding/json"
	"github.com/ivahaev/go-logger"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)


//HTTP hadnlers
var AddArticle = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var u models.User
	err := r.ParseForm()
	if err != nil {
		logger.Error(err)
		w.Write([]byte(err.Error()))
	}
	logger.Info("coo")
	logger.Info(r.Cookies())
	ssid, err := r.Cookie("ssid")
	if err != nil {
		logger.Error(err)
		//Is this status code is correct???
		w.WriteHeader(403)
		w.Write([]byte("Wrong auth token"))
		return
	}

	services.GetInstanceDB().Raw("SELECT * from users WHERE id in (select user_id from credentials where ssid = ?)", ssid.Value).Scan(&u)
	status, result := models.SaveArticlenDb(r.Form.Get("name"), r.Form.Get("body"), u.Id)
	w.WriteHeader(status)
	w.Write([]byte(result))
})

var GetArticle = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	parsedRequestURI := strings.Split(r.RequestURI, "/")
	rawId := parsedRequestURI[len(parsedRequestURI)-1]
	id, err := strconv.Atoi(rawId)
	if err != nil {
		//maybe redirect will be better???
		logger.Error(err)
		w.WriteHeader(404)
		w.Write([]byte("not found"))
		return
	}
	article := models.GetArticleById(id)
	//a := models.GetArticleById()
	//this is baaaaad stuff!!!!
	jsonArticle, err := json.Marshal(article)
	if err != nil {
		logger.Error(err)
	}
	//lw.WriteHeader(200)
	logger.Info(jsonArticle)
	//w.Write([]byte(jsonArticle))
	http.Redirect(w,r, "http://c-c.ru/post.html", http.StatusMovedPermanently)
	//w.Write([]byte(jsonArticle))
})

var CalculateRating = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	//is it correct error answer??
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		w.Write([]byte("wrong data in request"))
	}
	var msg struct{
		Article string `json: "article"`
		Direction int `json: direction`
	}
	err = json.Unmarshal(b, &msg)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		w.Write([]byte("cant parse json"))
	}

	err, rating := models.CalculateRating(msg.Article, msg.Direction)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		w.Write([]byte("something goes wrong"))
	}
	w.WriteHeader(200)
	logger.Info(rating)
	w.Write([]byte(strconv.Itoa(rating)))
	//w.Write([]byte(string(rating)))

})

func Cleaner(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	logger.Info(msg)

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}


