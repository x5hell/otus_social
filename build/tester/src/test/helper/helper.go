package helper

import (
	"component/fixture"
	"component/handler"
	"generator"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

const requestTimeout = 10 * time.Second

func getSocialDomain() string {
	return "http://" +
		os.ExpandEnv("$SOCIAL_SITE_CONTAINER_NAME") +
		":" +
		os.ExpandEnv("$SOCIAL_SITE_INTERNAL_PORT") + "/"
}

func GetSocialIp() string {
	ipList, err := net.LookupIP(os.ExpandEnv("$SOCIAL_SITE_CONTAINER_NAME"))

	if err != nil {
		handler.ErrorLog(err)
		return ""
	}
	ip := ipList[0]
	return "http://" +
		ip.String() +
		":" +
		os.ExpandEnv("$SOCIAL_SITE_INTERNAL_PORT")
}

func TestPageSingleThead(url string) error {
	httpClient := http.Client{Timeout: requestTimeout}
	fullUrl := getSocialDomain() + url
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