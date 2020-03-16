package generator

import (
	"strconv"
	"structure"
)

func UserRows(users int, maxCityId int, minAge int, maxAge int) (result []structure.Table) {
	for i := 1; i <= users; i++ {
		userSex := Sex()
		userRow := structure.User{
			ID: structure.NullString{Valid:true, String:strconv.Itoa(i)},
			Login: structure.NullString{Valid:true,String: Login(i)},
			Password: structure.NullString{Valid:true, String: "e10adc3949ba59abbe56e057f20f883e"},
			FirstName: structure.NullString{Valid:true, String: FirstName(userSex)},
			LastName: structure.NullString{Valid:true, String: LastName(userSex)},
			Age: structure.NullString{Valid:true, String: strconv.Itoa(Age(minAge, maxAge))},
			Sex: userSex,
			CityId: generateCityId(maxCityId),
		}
		result = append(result, userRow)
	}
	return result
}
