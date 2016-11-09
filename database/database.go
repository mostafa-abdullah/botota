package database

import (
	"botota/models"
)

type Database interface {
	Connect()
	CreateUser(user models.User)
	CreateQuestion(question models.Question)
}
