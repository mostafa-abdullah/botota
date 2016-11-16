package database

import (
	"botota/models"
)

type Database interface {
	Connect()
	CreateUser(u models.User)
	CreateQuestion(q models.Question)
	GetCurrentQuestion(u models.User) models.Question
	GetFirstQuestion() models.Question
	GetNextQuestion(q models.Question) models.Question
	Close()
}
