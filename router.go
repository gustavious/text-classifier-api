package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux


func routers() *chi.Mux {
	router.Options("/message", OptionsPreflight)
	router.Post("/message", ProcessMessage)
	return router
}

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
}
