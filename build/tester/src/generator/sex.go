package generator

import (
	"math/rand"
	"structure"
)

const SexMale = "male"
const SexFemale = "female"

func Sex() structure.NullString {
	num := rand.Intn(4)
	switch num {
		case 0: return structure.NullString{Valid:true, String: SexMale}
		case 1: return structure.NullString{Valid:true, String: SexFemale}
		case 2: return structure.NullString{Valid:false, String: SexMale}
		default: return structure.NullString{Valid:false, String: SexFemale}
	}
}