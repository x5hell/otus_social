package entity

import "database/sql"

type User struct {
	ID int
	Login string
	Password string
	FirstName string
	LastName string
	Age int
	Sex sql.NullString
	CityId sql.NullInt64
}