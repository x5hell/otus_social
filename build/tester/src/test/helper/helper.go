package helper

import (
	"component/fixture"
	"generator"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const domain = "http://social_go:8001/"
const requestTimeout = 10 * time.Second

func TestPageSingleThead(url string) error {
	httpClient := http.Client{Timeout: requestTimeout}
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
	return nil
}

func GetSeedDataParams() (seedDataParams generator.SeedDataParams) {
	seedDataParams, err := generator.GetSeedDataParams()
	if err != nil {
		log.Fatal(err)
	}
	return seedDataParams
}

func ApplyFixture(seedDataParams generator.SeedDataParams) {
	err := generator.GenerateFixture(seedDataParams)
	if err != nil {
		log.Fatal(err)
	}
	err = fixture.Apply(seedDataParams.FixtureGeneratedScript)
	if err != nil {
		log.Fatal(err)
	}
}

func AddIndexes(seedDataParams generator.SeedDataParams) {
	err := fixture.Apply(seedDataParams.AddIndexScript)
	if err != nil {
		log.Fatal(err)
	}
}

func RemoveIndexes(seedDataParams generator.SeedDataParams) {
	err := fixture.Apply(seedDataParams.RemoveIndexScript)
	if err != nil {
		log.Fatal(err)
	}
}

