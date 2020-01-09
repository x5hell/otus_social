package controllerResponse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

const FieldTagName = "name"
const FieldTagValidation = "validation"

const ServerErrorMessage = "ошибка на стороне сервера"
const SessionExpiredMessage = "сессия закончилась войдите снова"

type ErrorMessage struct {
	Error string `json:"error"`
}

type OkMessage struct {
	Ok string `json:"ok"`
}

type ErrorForm struct {
	Error map[string]string `json:"error"`
}

func JsonOkMessage(okMessage string, response http.ResponseWriter) {
	result, _ := json.Marshal(OkMessage{Ok: okMessage})
	fmt.Fprintf(response, string(result))
}

func JsonErrorMessage(errorMessage string, response http.ResponseWriter) {
	result, _ := json.Marshal(ErrorMessage{Error: errorMessage})
	fmt.Fprintf(response, string(result))
}

func JsonFormError(errorList map[string]error, response http.ResponseWriter) {
	errorMessages := make(map[string]string)
	for fieldName, errorMessage := range errorList {
		errorMessages[fieldName] = errorMessage.Error()
	}
	result, _ := json.Marshal(ErrorForm{Error: errorMessages})
	fmt.Fprintf(response, string(result))
}

func FillStructureFromRequest(request *http.Request, requestStructure interface{}, fieldAliasList map[string]string) {
	typeList := reflect.TypeOf(requestStructure).Elem()
	valueList := reflect.ValueOf(requestStructure).Elem()
	fieldsCount := typeList.NumField()
	for fieldNumber := 0; fieldNumber < fieldsCount; fieldNumber++ {
		fieldType := typeList.Field(fieldNumber)
		fieldValue := valueList.Field(fieldNumber)
		fieldName := fieldType.Tag.Get(FieldTagName)
		if fieldNameAlias, fieldNameAliasExists := fieldAliasList[fieldName] ; fieldNameAliasExists {
			fieldName = fieldNameAlias
		}
		switch fieldType.Type.String() {
			case "string":
				formValue := request.Form.Get(fieldName)
				fieldValue.SetString(formValue)
				break
			default:
				formElement := request.Form[fieldName]
				formValue := reflect.ValueOf(formElement)
				fieldValue.Set(formValue)
				break
		}
	}
}

func ParseRequest (response http.ResponseWriter, request *http.Request, method string,
	callback func(response http.ResponseWriter, request *http.Request)){
		switch request.Method {
		case method:
			err := request.ParseForm()
			if err != nil {
				JsonErrorMessage("ошибка разбора запроса", response)
			} else {
				callback(response, request)
			}
			break
		default:
			JsonErrorMessage(
				fmt.Sprintf("запрос типа '%s' не поддерживается", request.Method),
				response)
		}
}