package generator

import "math/rand"

func Age(minAge int, maxAge int) int {
	return minAge + rand.Intn(maxAge-minAge)
}
