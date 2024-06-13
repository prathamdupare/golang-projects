package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prathamdupare/golang-webserver/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request) (p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.NewObjectId()

	u := models.User{}

	uc.Session.DB("mongo-golang").C("users")
}
