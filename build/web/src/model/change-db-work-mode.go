package model

import (
	"component/app"
	"component/database"
	"component/environment"
	"component/validation"
	"fmt"
)

const WorkModeFieldName = "workMode"

type ChangeDbWorkModeRequest struct {
	WorkMode string `name:"workMode" validation:"regex=^(useReplica|masterOnly)$"`
}

func ChangeDbWorkMode(requestStruct ChangeDbWorkModeRequest) (err error) {
	err = validateChangeDbWorkModeRequest(requestStruct)
	if err == nil {
		err := environment.Set(database.WorkModeEnv, requestStruct.WorkMode)
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

func validateChangeDbWorkModeRequest(requestStruct ChangeDbWorkModeRequest) (err error) {
	validationResult, fieldErrors := validation.ValidateStructure(requestStruct, map[string]string{})
	fmt.Println(validationResult, fieldErrors)
	if validationResult {
		err = validateWorkMode()
		if err != nil {
			return err
		}
		return nil
	}
	return fieldErrors[WorkModeFieldName]
}

func validateWorkMode() (err error) {
	workMode, err := app.GetWorkMode()
	if err != nil {
		return err
	}
	switch workMode {
		case app.WorkModeDev:
		case app.WorkModeTest:
			return nil
	    default:
			return fmt.Errorf(
				"operaion avalable only in WorkMode: %s, %s",
				app.WorkModeDev,
				app.WorkModeTest,
			)
	}
	return nil
}