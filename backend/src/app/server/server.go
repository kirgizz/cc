package server

import (
	"app/models"
	"app/utils"
	"github.com/ivahaev/go-logger"
	"net/http"
)

const (
	SSID = "ssid"
	REMEMBER = "remember"
	EMAIL = "email"
	PASS = "password"
	DOMAIN = "c-c.ru"
	PATH = "/"
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

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	logger.Info(r)
	logger.Info(r.Cookies())
	err := r.ParseForm()
	if err != nil {
		logger.Crit(err)
	}
	w.Header().Set("header_name", "header_value")
	r.Header.Set("name", "value")

	ssidCookie := http.Cookie{
		Name:   SSID,
		Value:  "ssid",
		Domain: "localhost",
		Path:   PATH,
		MaxAge: 36000}
	http.SetCookie(w, &ssidCookie)
	w.Header().Set("header_name", "header_value")
	r.Header.Set("name", "value")

	w.Write([]byte("Hello from not implemented function"))
})

func (s Server) getPublications(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	//if err != nil {
	//	logger.Error(err)
	//	http.Error(w, http.StatusText(422), 422)
	//	return
	//}
	logger.Info("Get publications function")
}

func (s Server) registerUser(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusInternalServerError
	err := r.ParseForm()
	if err != nil {
		logger.Error(err)
		statusCode = http.StatusInternalServerError
	}
	var u models.User
	u.Email = r.FormValue("email")
	u.NickName = r.FormValue("nickname")
	u.Password = utils.CreateHash(r.FormValue("password"))
	statusCode = u.RegisterUser()
	w.WriteHeader(statusCode)
}
