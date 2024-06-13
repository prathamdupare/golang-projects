package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prathamdupare/golang-webserver/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.getUser)
	r.POST("/user", uc.createUser)
	r.DELETE("/user/:id", uc.deleteUser)

	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	return s
}
