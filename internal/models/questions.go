// Filename: internal/models/questions.go
package models

import (
	"database/sql"
	"time"
)

//lets model the questions table
type Question struct{
	QuestionID int64
	Body string
	CreatedAt time.Time
}

// Setup dependecy injection 
type QuestionModel struct{
	DB * sql.DB
}

// Write sql code to access the database 
//TODO
