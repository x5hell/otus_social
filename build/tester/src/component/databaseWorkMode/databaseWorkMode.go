package databaseWorkMode

import (
	"component/handler"
	"encoding/json"
	"fmt"
	"net/http"
	"test/helper"
)


type Response struct {
	Ok string `json:"ok,omitempty"`
	Error string `json:"error,omitempty"`
}

func Set(workMode string) (err error) {
	setDbWorkModeUrl := helper.GetSiteIp() + "/search?first-name="
	response, err := http.Get(setDbWorkModeUrl + workMode)
	handler.ErrorLog(err)
	if err != nil {
		return err
	}
	jsonResponse := Response{}
	err = json.NewDecoder(response.Body).Decode(&jsonResponse)
	handler.ErrorLog(err)
	if err != nil {
		return err
	}
	if jsonResponse.Error != "" {
		return fmt.Errorf(jsonResponse.Error)
	}
	return nil
}