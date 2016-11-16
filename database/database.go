package database

import (
	"botota/models"
)

type Database interface {
	Connect()
	CreateUser(u models.User)
	CreateQuestion(q models.Question)
	GetUser(uuid string) models.User
	UpdateUser(u models.User)
	GetCurrentQuestion(u models.User) models.Question
	GetFirstQuestion() models.Question
	GetNextQuestion(q models.Question) models.Question
	Close()
}
