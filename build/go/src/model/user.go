package model

type User struct {
	ID int
	Login string
	Password string
	FirstName string
	LastName string
	Age int
	Sex int
	City City
	InterestList []Interest
}