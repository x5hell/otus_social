package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func testPageSingleThead(url string) error {
	httpClient := http.Client{}
	domain := "http://localhost:8000/"
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

func BenchmarkUserProfileList(b *testing.B) {
	errorCounter := 0
	successCounter := 0
	for i := 0; i < b.N; i++ {
		if testPageSingleThead("user-profile-list") != nil {
			errorCounter++
		} else {
			successCounter++
		}
	}
	fmt.Printf("success: %d, errors: %d", successCounter, errorCounter)
}
