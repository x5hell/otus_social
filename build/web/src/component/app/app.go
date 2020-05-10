package app

import (
	"component/environment"
	"fmt"
)

const WorkModeEnv = "APP_WORK_MODE"

type WorkMode string

const (
	WorkModeDev WorkMode = "dev"
	WorkModeTest WorkMode = "test"
	WorkModeProd WorkMode = "prod"
)

func GetWorkMode() (result WorkMode, err error) {
	workMode, err := environment.Get(WorkModeEnv)
	if err != nil {
		return result, err
	}
	err = ValidateWorkMode(workMode)
	if err != nil {
		return result, err
	}
	return WorkMode(workMode), nil
}

func ValidateWorkMode(workMode string) (err error) {
	switch workMode {
		case string(WorkModeDev):
		case string(WorkModeTest):
		case string(WorkModeProd):
			return nil
		default:
			return fmt.Errorf(
				"ivalid WorkMode %s expected: %s, %s, %s ",
				workMode,
				WorkModeDev,
				WorkModeTest,
				WorkModeProd,
			)
	}
	return nil
}