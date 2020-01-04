package controller

import (
	"component/template"
	"fmt"
	"model"
	"net/http"
	"repository"
)

type TemplateData struct {
	CityList []model.City
	InterestList []model.Interest
}

func RegistrationForm(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template.OpenGuestTemplate("registration-form.html")
	if err != nil {
		fmt.Fprintf(response, "error: %v", err)
	} else {
		data, err := getTemplateData()
		if err != nil {
			fmt.Fprintf(response, "error: %v", err)
		} else {
			htmlTemplate.ExecuteTemplate(response, template.LayoutName, data)
		}
	}
}

func getTemplateData() (data TemplateData, err error) {
	data.CityList, err = repository.GetAllCities()
	if err != nil {
		return data, err
	}
	data.InterestList, err = repository.GetAllInterests()
	return data, err
}