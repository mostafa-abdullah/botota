package database

import (
	"botota/models"
)

type MongoDB struct{
  ConnectionString string
}

func (mongo MongoDB) Connect() {

}

func (mongo MongoDB) CreateUser(user models.User) {

}

func (mongo MongoDB) CreateQuestion(question models.Question) {

}

func (mongo MongoDB) Close() {

}
