package test

import (
	"generator"
	"math/rand"
	"structure"
	"testing"
	"time"
)

func BenchmarkCity(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		generator.City()
	}
}

func BenchmarkInterest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generator.Interest(i)
	}
}

func BenchmarkSex(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		generator.Sex()
	}
}

func BenchmarkLogin(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		generator.Login(i + 1)
	}
}

func BenchmarkAge(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		generator.Age(i, i + 60)
	}
}

func getSexList() []structure.NullString {
	return []structure.NullString{
		{Valid:true, String:generator.SexMale},
		{Valid:true, String:generator.SexFemale},
		{Valid:false, String:generator.SexMale},
		{Valid:false, String:generator.SexFemale},
	}
}

func BenchmarkFirstName(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sexList := getSexList()
	for i := 0; i < b.N; i++ {
		sex := sexList[i % 4]
		generator.FirstName(sex)
	}
}

func BenchmarkLastName(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sexList := getSexList()
	for i := 0; i < b.N; i++ {
		sex := sexList[i % 4]
		generator.LastName(sex)
	}
}