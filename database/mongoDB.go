package database

import (
	"botota/models"
)

type MongoDB struct{
  ConnectionString string
}

func (db *MongoDB) Connect() {

}

func (db *MongoDB) CreateUser(user models.User) {

}

func (db *MongoDB) CreateQuestion(question models.Question) {

}

func (db *MongoDB) GetCurrentQuestion(u models.User) models.Question {
	q:= models.Question{}
	return q
}

func (db *MongoDB) GetFirstQuestion() models.Question {
	q:= models.Question{}
	return q
}

func (db *MongoDB) GetNextQuestion(q models.Question) models.Question {
	nq:= models.Question{}
	return nq
}

func (db *MongoDB) Close() {

}
