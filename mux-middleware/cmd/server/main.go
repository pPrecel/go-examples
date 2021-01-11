package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
)

const (
	userKey = "user"
	userVal = "admin"
	passKey = "pass"
	passVal = "1234"
)

func AuthHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Info("Hello from AuthHandler")
	user := r.URL.Query().Get(userKey)
	pass := r.URL.Query().Get(passKey)
	log.WithFields(log.Fields{
		userKey: user,
		passKey: pass,
	}).Info("URL values")

	if user != userVal || pass != passVal {
		http.Error(rw, "Unauthorized", 401)
		return
	}

	r = r.WithContext(context.WithValue(r.Context(), userKey, user))
	next.ServeHTTP(rw, r)
}

func InfoHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Info("Hello from HelloHandler")
	next.ServeHTTP(rw, r)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Request received for user: %s", r.Context().Value(userKey).(string))
	w.Write([]byte("looks ok :)"))
}

func main() {
	log.Info("Starting server...")
	server := http.NewServeMux()
	server.HandleFunc("/hello", MainHandler)

	n := negroni.New()
	n.UseFunc(AuthHandler)
	n.UseFunc(InfoHandler)
	n.UseHandler(server)
	n.UseFunc(InfoHandler)

	log.Info("Listen And Serve...")
	log.Fatal(http.ListenAndServe(":8080", n))
}
