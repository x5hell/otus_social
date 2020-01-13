package model

import (
	"component/controllerResponse"
	"component/handler"
	"component/validation"
	"database/sql"
	"fmt"
)

func FormFieldServerErrorWithRollback(
	transaction *sql.Tx,
	validationResult validation.FieldValidationResult,
	err error,
	errorFieldName string,
) map[string]error {
	if transaction != nil {
		handler.ErrorLog(transaction.Rollback())
	}
	return FormFieldServerError(validationResult, err, errorFieldName)
}

func FormFieldServerError(
	validationResult validation.FieldValidationResult,
	err error,
	errorFieldName string,
) map[string]error {
	validationResult.FieldErrors[errorFieldName] = fmt.Errorf(controllerResponse.ServerErrorMessage)
	handler.ErrorLog(err)
	return validationResult.FieldErrors
}