package database

import (
	"botota/models"
)

type MySqlDB struct{
	ConnectionString string
}

func (mysql MySqlDB) Connect() {

}

func (mysql MySqlDB) CreateUser(user models.User) {

}

func (mysql MySqlDB) CreateQuestion(question models.Question) {

}
