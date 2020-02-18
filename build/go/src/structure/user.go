package structure

const SexNotSelected = "не указан"
const SexMale = "мужской"
const SexFemale = "женский"

type User struct {
	Id int					`json:"id"`
	FirstName string		`json:"firstName"`
	LastName string			`json:"lastName"`
	Age int					`json:"age"`
	Sex string				`json:"sex"`
	City City				`json:"city"`
	InterestList []Interest	`json:"interestList"`
}
