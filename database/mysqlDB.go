package database

import (
	"botota/models"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

// Define the database configuration variables
// TODO: Get them from an environment config file
const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "botota"
	DB_USER = "<< MySql username >>"
	DB_PASS = "<< MySql password >>"

	USERS_TABLE = "users"
	QUESTIONS_TABLE = "questions"
)

type MySqlDB struct{
	conn *sql.DB;
}

func (db *MySqlDB) Connect() {
	// Format the connection string
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	con, err := sql.Open("mysql", dsn)

	db.conn = con

	checkError(err)
}

// Insert a new user into the database
func (db MySqlDB) CreateUser(u models.User) {

	createUsersTableIfNotExists(db.conn)

	_, err := db.conn.Exec("insert into " + USERS_TABLE + " (uuid) values (?)", u.Uuid)
	checkError(err);
}

// Insert a new question into the database
func (db MySqlDB) CreateQuestion(question models.Question) {
	createQuestionsTableIfNotExists(db.conn)

	_, err := db.conn.Exec("insert into " + QUESTIONS_TABLE + " (text, nextQuestion) values (?, ?)", question.Text, question.NextQuestionId)
	checkError(err);
}

// Close the open database connection
func (db MySqlDB) Close() {
	db.conn.Close()
}


func createUsersTableIfNotExists(conn *sql.DB) {
	if _, err := conn.Exec("DESCRIBE " + USERS_TABLE); err != nil {
		// MySql error: table doesn't exist
		var createUserTableStatements = []string {
			`CREATE DATABASE IF NOT EXISTS ` + DB_NAME + `DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
			`USE ` + DB_NAME + `;`,
			`CREATE TABLE IF NOT EXISTS ` + USERS_TABLE + ` (
				id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
				uuid VARCHAR(255) NOT NULL,
				destination VARCHAR(255) NULL,
				startDate VARCHAR(255) NULL,
				endDate VARCHAR(255) NULL,
				budget VARCHAR(255) NULL,
			)`,
		}

		createTable(conn, createUserTableStatements)
	}
}


func createQuestionsTableIfNotExists(conn *sql.DB) {
	if _, err := conn.Exec("DESCRIBE " + QUESTIONS_TABLE); err != nil {
		// MySql error: table doesn't exist
		var createQuestionsTableStatements = []string {
			`CREATE DATABASE IF NOT EXISTS ` + DB_NAME + ` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
			`USE ` + DB_NAME + `;`,
			`CREATE TABLE IF NOT EXISTS ` + QUESTIONS_TABLE + ` (
				id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
				text TEXT NULL,
				nextQuestion INT UNSIGNED,
				FOREIGN KEY(nextQuestion) REFERENCES ` + QUESTIONS_TABLE +`(id) ON DELETE CASCADE
			)`,
		}

		createTable(conn, createQuestionsTableStatements)
	}
}


func createTable(conn *sql.DB, statements []string) error {
	// execute all the statements in the table creation statements
	for _, stmt := range statements {
		_, err := conn.Exec(stmt)
		checkError(err)
	}
	return nil
}

func checkError(err error) {
	if(err != nil) {
		log.Fatal(err);
	}
}
