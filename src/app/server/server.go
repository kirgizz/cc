package server

import (
	"github.com/ivahaev/go-logger"
	"net/http"
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