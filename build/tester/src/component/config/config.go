package config

import (
	"component/handler"
	"fmt"
	"os"
)

func GetEnv(envName string) (envValue string, err error) {
	envValue, envExists := os.LookupEnv(envName)
	if envExists == false {
		err = fmt.Errorf("envoirment variable %s not set", envName)
		handler.ErrorLog(err)
	}
	return envValue, nil
}