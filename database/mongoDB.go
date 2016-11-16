package database

import (
	"botota/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	URL = "localhost:27017"
	DB = "botota"
	USERS_COLLECTION = "users"
	QUESTIONS_COLLECTION = "questions"
)

type MongoDB struct{
	session *mgo.Session
}

func (db *MongoDB) Connect() {
	s, err := mgo.Dial(URL)
	checkError(err)

	db.session = s
}

func (db *MongoDB) CreateUser(u models.User) {
	u.CurrentQuestionId = 1
	
	c := db.session.DB(DB).C(USERS_COLLECTION)
	err := c.Insert(u)

	checkError(err)
}

func (db *MongoDB) CreateQuestion(q models.Question) {
	c := db.session.DB(DB).C(QUESTIONS_COLLECTION)
	err := c.Insert(q)

	checkError(err)
}

func (db *MongoDB) GetUser(uuid string) models.User {
	c := db.session.DB(DB).C(USERS_COLLECTION)

	res := models.User{}
	err := c.Find(bson.M{"uuid": uuid}).One(&res)
	checkError(err)

	return res
}

func (db *MongoDB) UpdateUser(u models.User) {
	c := db.session.DB(DB).C(USERS_COLLECTION)

	colQuerier := bson.M{"uuid": u.Uuid}
	change := bson.M{"$set": bson.M{"destination": u.Destination,
		"startdate": u.StartDate,
		"enddate": u.EndDate,
		"budget": u.Budget,
		"currentquestionid": u.CurrentQuestionId}}

	err := c.Update(colQuerier, change)

	checkError(err)
}

func (db *MongoDB) GetCurrentQuestion(u models.User) models.Question {
	c := db.session.DB(DB).C(QUESTIONS_COLLECTION)

	res := models.Question{}
	err := c.Find(bson.M{"id": u.CurrentQuestionId}).One(&res)
	checkError(err)

	return res
}

func (db *MongoDB) GetFirstQuestion() models.Question {
	c := db.session.DB(DB).C(QUESTIONS_COLLECTION)

	res := models.Question{}
	err := c.Find(bson.M{"id": 1}).One(&res)
	checkError(err)

	return res
}

func (db *MongoDB) GetNextQuestion(q models.Question) models.Question {
	c := db.session.DB(DB).C("questions")

	res := models.Question{}
	err := c.Find(bson.M{"id": q.NextQuestionId}).One(&res)
	checkError(err)

	return res
}

func (db *MongoDB) Close() {
	db.session.Close()
}
