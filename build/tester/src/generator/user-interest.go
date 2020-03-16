package generator

import (
	"math/rand"
	"strconv"
	"structure"
)

func UserInterestRows(maxUserId, maxInterestId, maxUserInterests int) (result []structure.Table) {
	for userId := 1; userId <= maxUserId; userId++ {
		userInterests := rand.Intn(maxUserInterests)
		interestMap := make(map[int]bool)
		for userInterest := 1; userInterest <= userInterests; userInterest++ {
			userInterestId := rand.Intn(maxInterestId) + 1
			interestMap[userInterestId] = true
		}
		for interestId, valid := range interestMap {
			userInterestRow := structure.UserInterest{
				UserId: structure.NullString{
					Valid:true,
					String:strconv.Itoa(userId),
				},
				InterestId: structure.NullString{
					Valid: valid,
					String: strconv.Itoa(interestId),
				},
			}
			result = append(result, userInterestRow)
		}
	}
	return result
}
