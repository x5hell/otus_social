package validation

import (
	"component/controllerResponse"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type FieldValidation struct {
	FieldName			string
	FieldValue	 	 	string
	ValidationName		string
	ValidationParameter string
}

type Validator interface{
	Validate() error
}

type required struct { *FieldValidation }
type symbolsMin struct { *FieldValidation }
type symbolsMax struct { *FieldValidation }
type regex struct { *FieldValidation }
type isInt struct { *FieldValidation }
type digitMin struct { *FieldValidation }
type digitMax struct { *FieldValidation }

func (field required) Validate() error {
	if len(field.FieldValue) == 0 {
		return fmt.Errorf("field %s required", field.FieldName)
	}
	return nil
}

func (field symbolsMin) Validate() error {
	min, err := strconv.Atoi(field.ValidationParameter)
	if err != nil {
		return fmt.Errorf("field %s :: %s", field.FieldName, err.Error())
	}
	if len(field.FieldValue) < min {
		return fmt.Errorf(
			"field %s contains %d symbols (min %d symbols required)",
			field.FieldName,
			len(field.FieldValue),
			min)
	}
	return nil
}

func (field symbolsMax) Validate() error {
	max, err := strconv.Atoi(field.ValidationParameter)
	if err != nil {
		return fmt.Errorf("field %s :: %s", field.FieldName, err.Error())
	}
	if len(field.FieldValue) > max {
		return fmt.Errorf(
			"field %s contains %d symbols (max %d symbols required)",
			field.FieldName,
			len(field.FieldValue),
			max)
	}
	return nil
}

func (field regex) Validate() error {
	regexpValidation, err := regexp.Compile(field.ValidationParameter)
	if err != nil {
		return fmt.Errorf("field %s :: %s", field.FieldName, err.Error())
	}
	if regexpValidation.MatchString(field.FieldValue) == false {
		return fmt.Errorf(
			"field %s does not match regular expression: %s",
			field.FieldName,
			field.ValidationParameter)
	}
	return nil
}

func (field isInt) Validate() error {
	_, err := strconv.Atoi(field.FieldValue)
	if err != nil {
		return fmt.Errorf("field %s :: %s", field.FieldName, "is not int")
	}
	return nil
}

func (field digitMin) Validate() error {
	min, err := strconv.Atoi(field.ValidationParameter)
	if err != nil {
		return fmt.Errorf("field %s :: %s", field.FieldName, err.Error())
	}
	fieldValue, err := strconv.Atoi(field.FieldValue)
	if err != nil {
		return fmt.Errorf("field %s :: %s", field.FieldName, "is not a digit")
	}
	if fieldValue < min {
		return fmt.Errorf(
			"field %s is %d (min %d required)",
			field.FieldName,
			fieldValue,
			min)
	}
	return nil
}

func (field digitMax) Validate() error {
	max, err := strconv.Atoi(field.ValidationParameter)
	if err != nil {
		return fmt.Errorf("field %s :: %s", field.FieldName, err.Error())
	}
	fieldValue, err := strconv.Atoi(field.FieldValue)
	if err != nil {
		return fmt.Errorf("field %s :: %s", field.FieldName, "is not a digit")
	}
	if fieldValue > max {
		return fmt.Errorf(
			"field %s is %d (max %d required)",
			field.FieldName,
			fieldValue,
			max)
	}
	return nil
}


func (field FieldValidation) getValidator() (result Validator) {
	switch field.ValidationName {
		case "required":
			return required{&field}
		case "symbolsMin":
			return symbolsMin{&field}
		case "symbolsMax":
			return symbolsMax{&field}
		case "regex":
			return regex{&field}
		case "isInt":
			return isInt{&field}
		case "digitMin":
			return digitMin{&field}
		case "digitMax":
			return digitMax{&field}
	}
	return nil
}

func ValidateStructure(requestStructure interface{}) (result bool, fieldErrors map[string]error) {
	errors := make(map[string]error)
	typeList := reflect.TypeOf(requestStructure)
	valueList := reflect.ValueOf(requestStructure)
	fieldsCount := typeList.NumField()
	for fieldNumber := 0; fieldNumber < fieldsCount; fieldNumber++ {
		fieldType := typeList.Field(fieldNumber)
		fieldValue := valueList.Field(fieldNumber)
		fieldName := fieldType.Tag.Get(controllerResponse.FieldTagName)
		fieldValidationList := getFieldValidationList(fieldType, fieldValue)
		err := validateField(fieldValidationList)
		if err != nil {
			errors[fieldName] = err
		}
	}
	return len(errors) == 0, errors
}

func validateField(fieldValidationList []FieldValidation) error {
	for _, fieldValidation := range fieldValidationList {
		validator := fieldValidation.getValidator()
		if validator != nil {
			err := fieldValidation.getValidator().Validate()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getFieldValidationList(fieldType reflect.StructField, fieldValue reflect.Value) []FieldValidation {
	var fieldValidationList []FieldValidation
	validationRules := fieldType.Tag.Get(controllerResponse.FieldTagValidation)
	fieldName := fieldType.Tag.Get(controllerResponse.FieldTagName)
	validationRulesParts := strings.Split(validationRules, ",")
	for _, validationRule := range validationRulesParts {
		fieldValidation := FieldValidation{}
		fieldValidation.FieldName = fieldName
		fieldValidation.FieldValue = fieldValue.String()
		rulePartList := strings.Split(validationRule, "=")
		if len(rulePartList) >= 1 {
			fieldValidation.ValidationName = rulePartList[0]
		}
		if len(rulePartList) >= 2 {
			fieldValidation.ValidationParameter = rulePartList[1]
		}
		fieldValidationList = append(fieldValidationList, fieldValidation)
	}
	return fieldValidationList
}

