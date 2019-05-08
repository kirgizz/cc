package server

import (
	"app/models"
	"encoding/json"
	"github.com/ivahaev/go-logger"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	SSID = "ssid"
	PASS = "password"
	DOMAIN = "c-c.ru"
	PATH = "/"
	MAXAGE = 3600
)

type Server struct {
	Address string
	Port string
	m models.Model
}

type Message struct {
	Article   string    `json:"article"`
	Direction int 		`json:"direction"`
}

func (s Server) NotImplemented(w http.ResponseWriter, r *http.Request) {
	logger.Info(r)
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Not implemented"))
}

func (s Server) registerUser(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusInternalServerError
	err := r.ParseForm()
	if err != nil {
		logger.Error(err)
		statusCode = http.StatusInternalServerError
	}
	var u models.User
	var data struct {
		Email string    `json:"email"`
		Nickname string `json:"nickaname"`
		Password string `json:"password"`
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Error(err)
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		logger.Error(err)
	}
	statusCode,err = u.RegisterUser(data.Email, data.Nickname, data.Password)
	w.WriteHeader(statusCode)
}

func (s Server) getPublications(w http.ResponseWriter, r *http.Request) {

}

func (s Server) getRubrics(w http.ResponseWriter, r *http.Request) {
	//var r []models.Rubrics
	statusCode := http.StatusOK
	rubrics, err := models.GetRubrics()
	if err != nil {
		logger.Error(err)
		statusCode = http.StatusInternalServerError
	}

	b, err := json.Marshal(rubrics)
	if err != nil {
		logger.Error(err)
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(b))
}

func (s Server) savePublication(w http.ResponseWriter, r *http.Request) {
	var c models.Credentials
	//var p models.Publication

	var publication struct {
		Rubrics []struct {
			Id int `json:"id"`
			Name string `jon:"name"`
		}
		Tags []struct {
			Label string `json:"label"`
			Value string `json:"value"`
		}
		Title string `json:"title"`
		Body string `json:"body"`
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Error(err)
	}
	err = json.Unmarshal(b, &publication)
	if err != nil {
		logger.Error(err)
	}

	cookie, err := r.Cookie("ssid")
	if err != nil {
		logger.Error(err)
	}
	err = c.GetUser(cookie.Value)
	if err != nil {
		logger.Error(err)
	}





	logger.Info(publication)
}

func (s Server) login(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusInternalServerError
	var cookie http.Cookie
	var c models.Credentials
	var data struct{
		Email string       `json:"email"`
		Password string    `json:"password"`
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Error(err)
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		logger.Error(err)
	}
	statusCode, err = c.CheckCredentials(data.Email, data.Password)

	if statusCode == http.StatusOK {
		cookie = http.Cookie{
			Name:   SSID,
			Value:  c.Ssid,
			Domain: DOMAIN,
			Path:   PATH,
			MaxAge: MAXAGE}
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(statusCode)
}


func (s Server) checkSession(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusForbidden
	var c models.Credentials
	var session struct {
		Ssid string `json:"ssid"`
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Error(err)
	}
	err = json.Unmarshal(b, &session)
	if err != nil {
		logger.Error(err)
	}

	if err := c.GetUser(session.Ssid); err != nil{
		statusCode = http.StatusInternalServerError
		logger.Error(err)
	}
	if c.UserId != 0 {
		statusCode = http.StatusOK

	}
	w.WriteHeader(statusCode)
}


func (s Server) logout(w http.ResponseWriter, r *http.Request) {
	var c models.Credentials
	var cookie http.Cookie
	var data struct {
		Ssid string `json:"ssid"`
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Error(err)
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		logger.Error(err)
	}

	c.GetUser(data.Ssid)
	logger.Info(string(b))

	cookie = http.Cookie{
		Name:   SSID,
		Value:  c.Ssid,
		Domain: DOMAIN,
		Path:   PATH,
		MaxAge: 0,
		Expires: time.Now(),
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusInternalServerError)
}
