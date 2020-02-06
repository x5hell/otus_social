package test

import (
	"component/fixture"
	"fmt"
	"generator"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func testPageSingleThead(url string) error {
	httpClient := http.Client{}
	domain := "http://social_go:8001/"
	fullUrl := domain + url
	response, err := httpClient.Get(fullUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return  nil
}

func getSeedDataParams() (seedDataParams generator.SeedDataParams) {
	seedDataParams, err := generator.GetSeedDataParams()
	if err != nil {
		log.Fatal(err)
	}
	return seedDataParams
}

func addIndexes(seedDataParams generator.SeedDataParams) {
	fixture.Apply(seedDataParams.AddIndexScript)
}

func removeIndexes(seedDataParams generator.SeedDataParams) {
	fixture.Apply(seedDataParams.RemoveIndexScript)
}

func applyFixture(seedDataParams generator.SeedDataParams) {
	err := generator.GenerateFixture(seedDataParams)
	if err != nil {
		log.Fatal(err)
	}
	err = fixture.Apply(seedDataParams.FixtureGeneratedScript)
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkUserProfileListWith10000UsersWithoutIndexes(b *testing.B) {
	seedDataParams := getSeedDataParams()
	seedDataParams.Users = 10000
	removeIndexes(seedDataParams)
	applyFixture(seedDataParams)
	errorCounter := 0
	successCounter := 0
	for i := 0; i < b.N; i++ {
		if testPageSingleThead("user-profile-list") != nil {
			errorCounter++
		} else {
			successCounter++
		}
	}
	fmt.Printf("\nsuccess: %d, errors: %d\n", successCounter, errorCounter)
	addIndexes(seedDataParams)
}

func BenchmarkUserProfileListWith10000UsersWithIndexes(b *testing.B) {
	seedDataParams := getSeedDataParams()
	seedDataParams.Users = 10000
	applyFixture(seedDataParams)
	errorCounter := 0
	successCounter := 0
	for i := 0; i < b.N; i++ {
		if testPageSingleThead("user-profile-list") != nil {
			errorCounter++
		} else {
			successCounter++
		}
	}
	fmt.Printf("\nsuccess: %d, errors: %d\n", successCounter, errorCounter)
}
