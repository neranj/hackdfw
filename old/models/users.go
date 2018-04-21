package models

import "time"

type User struct {
	Id  string `json: id`
	FirstName string
	LastName string
	Email string
	Password string
	CreatedOn time.Time
	VerificationToken string
	Status string
}


