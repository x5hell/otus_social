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

func getSiteDomain() string {
	return "http://" +
		os.ExpandEnv("$WEB_SITE_CONTAINER_NAME") +
		":" +
		os.ExpandEnv("$WEB_SITE_INTERNAL_PORT") + "/"
}

func GetSiteIp() string {
	ipList, err := net.LookupIP(os.ExpandEnv("$WEB_SITE_CONTAINER_NAME"))

	if err != nil {
		handler.ErrorLog(err)
		return ""
	}
	ip := ipList[0]
	return "http://" +
		ip.String() +
		":" +
		os.ExpandEnv("$WEB_SITE_INTERNAL_PORT")
}

func TestPageSingleThead(url string) error {
	httpClient := http.Client{Timeout: requestTimeout}
	fullUrl := getSiteDomain() + url
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

func ApplyFixture(seedDataParams generator.SeedDataParams, hostname string) {
	err := generator.GenerateFixture(seedDataParams)
	if err != nil {
		log.Fatal(err)
	}
	err = fixture.Apply(seedDataParams.FixtureGeneratedScript, hostname)
	if err != nil {
		log.Fatal(err)
	}
}

func AddIndexes(seedDataParams generator.SeedDataParams, hostname string) {
	err := fixture.Apply(seedDataParams.AddIndexScript, hostname)
	if err != nil {
		log.Fatal(err)
	}
}

func RemoveIndexes(seedDataParams generator.SeedDataParams, hostname string) {
	err := fixture.Apply(seedDataParams.RemoveIndexScript, hostname)
	if err != nil {
		log.Fatal(err)
	}
}