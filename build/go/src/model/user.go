package model

type User struct {
	ID int
	Login string
	Password string
	FirstName string
	LastName string
	Age int
	Sex string
	City City
	InterestList []Interest
}