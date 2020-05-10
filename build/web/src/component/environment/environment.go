package environment

import (
	"component/handler"
	"fmt"
	"os"
)

func Get(name string) (value string, err error){
	value, envExists := os.LookupEnv(name)
	if envExists == false {
		err = fmt.Errorf("envoirment variable %s not set", name)
	}
	handler.ErrorLog(err)
	return value, err
}

func Set(name, value string) (err error) {
	err = os.Setenv(name, value)
	handler.ErrorLog(err)
	return err
}