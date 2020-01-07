package model

import (
	"html/template"
)

func GetTemplateFunctions() template.FuncMap {
	var templateFunctions = template.FuncMap{}
	templateFunctions["Authorized"] = Authorized
	return templateFunctions
}
