package structure

import (
	"reflect"
)

const TableFieldTagName = "field"

type NullString struct {
	Valid bool
	String string
}

type Table interface {
	GetTableName() string
}

type City struct {
	ID NullString `field:"id"`
	Name NullString `field:"name"`
}
func (t City) GetTableName() string { return "city" }

type Interest struct {
	ID NullString `field:"id"`
	Name NullString `field:"name"`
}
func (t Interest) GetTableName() string { return "interest" }

type User struct {
	ID NullString `field:"id"`
	Login NullString `field:"login"`
	Password NullString `field:"password"`
	FirstName NullString `field:"first_name"`
	LastName NullString `field:"last_name"`
	Age NullString `field:"age"`
	Sex NullString `field:"sex"`
	CityId NullString `field:"city_id"`
}
func (t User) GetTableName() string { return "user" }

type UserInterest struct {
	UserId NullString `field:"user_id"`
	InterestId NullString `field:"interest_id"`
}
func (t UserInterest) GetTableName() string { return "user_interest" }

type StructList struct {
	ElementList []interface{}
}

func GetTablePropertyList(item Table) (result []string) {
	typeList := reflect.TypeOf(item)
	fieldsCount := typeList.NumField()
	for fieldNumber := 0; fieldNumber < fieldsCount; fieldNumber++ {
		fieldType := typeList.Field(fieldNumber)
		result = append(result, fieldType.Name)
	}
	return result
}

func GetTableFieldList(item Table) (result []string) {
	typeList := reflect.TypeOf(item)
	fieldsCount := typeList.NumField()
	for fieldNumber := 0; fieldNumber < fieldsCount; fieldNumber++ {
		fieldType := typeList.Field(fieldNumber)
		fieldName := fieldType.Tag.Get(TableFieldTagName)
		result = append(result, fieldName)
	}
	return result
}