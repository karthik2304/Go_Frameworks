package mongo

import (
	"fmt"

	"github.com/stphivos/todo-api-go-grpc/models"
	"gopkg.in/mgo.v2"
)


type Handler struct {
	*mgo.Session
}


func NewHandler(config *models.Config) (*Handler, error) {
	session, err := mgo.Dial(fmt.Sprintf("mongodb://%v:%v", config.Database.Host, config.Database.Port))
	handler := &Handler{
		Session: session,
	}
	return handler, err
}


func (db *Handler) GetTodos() ([]models.Todo, error) {
	session := db.getSession()
	defer session.Close()

	todos := []models.Todo{}
	err := session.DB("TodosDB").C("todos").Find(nil).All(&todos)

	return todos, err
}

func (db *Handler) getSession() *mgo.Session {
	return db.Session.Copy()
}
