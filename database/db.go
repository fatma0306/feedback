package database

import (
	"database/sql"

	"github.com/fatma0306/feedback/models"
)

// GetDB cheks the connection with the database
func GetDB(db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "db_feedback"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}

// ReviewModel
type ReviewModel struct {
	Db *sql.DB
}

// GetReview to find all the review
func (reviewModel ReviewModel) GetReview() (review []models.Review, err error) {
	rows, err := reviewModel.Db.Query("SELECT * FROM review")
	if err != nil {
		return nil, err
	} else {
		var reviews []models.Review
		for rows.Next() {
			var id int
			var content string
			var state bool
			var customer int
			//return the nth word from the character
			err2 := rows.Scan(&id, &content, &state, &customer)
			if err2 != nil {
				return nil, err2
			} else {
				review := models.Review{
					ReviewID:      id,
					ReviewContent: content,
					ReviewState:   state,
					CustomerID:    customer,
				}
				//function appends elements to the end of a slice
				reviews = append(reviews, review)
			}
		}
		return reviews, nil
	}
}
