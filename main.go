package main

import (
	"api48hours/auth"
	"api48hours/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	repository.Start(mysql.Config{
		User:   "root",
		Passwd: "pass",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "database",
	})

	r := chi.NewRouter()

	auth.SetRoutes(r)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3210", r)
}
