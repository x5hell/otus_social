package model

import (
	"component/validation"
	"repository"
	"structure"
)

const searchFieldNameInterests = "interests"
const searchFieldNameCity = "city"
const searchUserLimit = 1000
const SexMale = "male"
const SexFemale = "female"

func Search(requestStruct repository.UserSearchRequest) (userList []structure.User, fieldErrors map[string]error) {
	validationResult := validateSearchRequest(requestStruct)
	if validationResult.ValidationResult {
		userProfileList, _ := searchUsersProfileListData(requestStruct, searchUserLimit)
		return userProfileList, nil
	} else {
		return userList, validationResult.FieldErrors
	}
}

func validateSearchRequest(requestStruct repository.UserSearchRequest) (result validation.FieldValidationResult) {
	fieldAliasList := GetSearchFieldAliasList()
	result.ValidationResult, result.FieldErrors = validation.ValidateStructure(requestStruct, fieldAliasList)
	result.ValidationResult =
		result.ValidationResult &&
			validation.ValidateInterests(requestStruct.Interests, searchFieldNameInterests, &result) &&
			validation.ValidateCity(requestStruct.City, searchFieldNameCity, &result)
	return result
}

func searchUsersProfileListData(requestStruct repository.UserSearchRequest, limit int) (userList []structure.User, err error) {
	lastUsers, err := repository.SearchUsers(requestStruct, limit)
	if err != nil {
		return userList, err
	}
	userProfileList, err := GetUserProfileListData(lastUsers)
	if err != nil {
		return userList, err
	}
	return UserProfileListToUserList(userProfileList), nil
}

func UserProfileListToUserList(userProfileList UserProfileList) (userList []structure.User) {
	for _, userEntity := range userProfileList.Users {
		user := structure.User{}
		user.Id = userEntity.ID
		user.FirstName = userEntity.FirstName
		user.LastName = userEntity.LastName
		user.Age = userEntity.Age
		if userEntity.Sex.Valid {
			switch userEntity.Sex.String {
				case SexMale:
					user.Sex = structure.SexMale
					break
				case SexFemale:
					user.Sex = structure.SexFemale
					break
			}
		} else {
			user.Sex = structure.SexNotSelected
		}
		if userEntity.CityId.Valid {
			user.City = structure.City{
				Id: int(userEntity.CityId.Int64),
			}
			user.City.Name = userProfileList.CityList[user.City.Id].Name
		}
		userInterestList, interestListExists := userProfileList.UserIdToInterestList[user.Id]
		if interestListExists {
			for _, userInterest := range userInterestList {
				interest := structure.Interest{
					Id:   userInterest.ID,
					Name: userInterest.Name,
				}
				user.InterestList = append(user.InterestList, interest)
			}
		}
		userList = append(userList, user)
	}
	return userList
}


func GetSearchFieldAliasList() map[string]string {
	return map[string]string{
		"searchFieldNameCity": searchFieldNameCity,
	}
}