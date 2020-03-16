package generator

import "strconv"

func Login(seed int) string {
	return "login" + strconv.Itoa(seed)
}