package main

import (
	"gopkg.in/macaron.v1"
	"log"
	r "macaron-api/routers"
	"net/http"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())


	m.Get("/login", r.Login)
	m.Post("/register", r.CreateProfile)

	m.Group("" , func() {
		m.Get("/", r.Home)
		m.Get("/profiles", r.ReadProfile)
	}, r.CheckAuth)
	
	log.Fatal(http.ListenAndServe(":4000", m))
}
