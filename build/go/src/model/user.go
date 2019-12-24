package model

type User struct {
	ID int `db:"id"`
	Login string `db:"login"`
	Password string `db:"password"`
	FirstName string `db:"first_name"`
	LastName string `db:"last_name"`
	Age int `db:"age"`
	Sex int `db:"sex"`
	CityId int `db:"city_id"`
}