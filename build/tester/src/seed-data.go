package main

import (
	"converter"
	"fmt"
	"generator"
	"os"
	"strconv"
)

const EnvGenerateCities = "GENERATE_CITIES"
const EnvGenerateInterests = "GENERATE_INTERESTS"
const EnvGenerateUsers = "GENERATE_USERS" 
const EnvUserMinAge = "USER_MIN_AGE" 
const EnvUserMaxAge = "USER_MAX_AGE"
const EnvMaxUserInterests = "MAX_USER_INTERESTS"
const EnvSqlScriptsPath = "SQL_SCRIPTS_PATH"
const EnvBeforeSeedScript = "BEFORE_SEED_SCRIPT"
const EnvAfterSeedScript = "AFTER_SEED_SCRIPT"

type SeedDataParams struct {
	Cities int
	Interests int
	Users int
	MinAge int
	MaxAge int
	MaxUserInterests int
	BeforeDataSeedScript string
	AfterDataSeedScript string
}

func GetSeedDataParams() (seedDataParams SeedDataParams, err error) {

	envList := []string{
		EnvGenerateCities,
		EnvGenerateInterests,
		EnvGenerateUsers,
		EnvUserMinAge,
		EnvUserMaxAge,
		EnvMaxUserInterests,
		EnvSqlScriptsPath,
		EnvBeforeSeedScript,
		EnvAfterSeedScript,
	}

	stringMap := map[string]string{
		EnvSqlScriptsPath: "",
		EnvBeforeSeedScript: "",
		EnvAfterSeedScript: "",
	}

	intMap := map[string]int{
		EnvGenerateCities: 0,
		EnvGenerateInterests: 0,
		EnvGenerateUsers: 0,
		EnvUserMinAge: 0,
		EnvUserMaxAge: 0,
		EnvMaxUserInterests: 0,
	}

	for _, envName := range envList {
		envValue, envExists := os.LookupEnv(envName)
		if envExists == false {
			return seedDataParams, fmt.Errorf("envoirment variable %s not set", envName)
		}
		if _, isInt := intMap[envName]; isInt {
			intValue, err := strconv.Atoi(envValue)
			if err != nil {
				return seedDataParams, fmt.Errorf("envoirment variable %s is not integer", envName)
			}
			intMap[envName] = intValue
		}
		if _, isString := stringMap[envName]; isString {
			stringMap[envName] = envValue
		}
	}

	return SeedDataParams{
		Cities: intMap[EnvGenerateCities],
		Interests: intMap[EnvGenerateInterests],
		Users: intMap[EnvGenerateUsers],
		MinAge: intMap[EnvUserMinAge],
		MaxAge: intMap[EnvUserMaxAge],
		MaxUserInterests: intMap[EnvMaxUserInterests],
		BeforeDataSeedScript: stringMap[EnvSqlScriptsPath] + stringMap[EnvBeforeSeedScript],
		AfterDataSeedScript: stringMap[EnvSqlScriptsPath] + stringMap[EnvAfterSeedScript],
	}, nil
}

func SeedData(params SeedDataParams) {

	insertDataSql := len(fmt.Sprintln(
		params.BeforeDataSeedScript,
		"\n",
		converter.RowListToSqlInsertQuery(
			generator.CityRows(params.Cities)),
		"\n",
		converter.RowListToSqlInsertQuery(
			generator.InterestRows(params.Interests)),
		"\n",
		converter.RowListToSqlInsertQuery(
			generator.UserRows(params.Users, params.Cities, params.MinAge, params.MaxAge)),
		"\n",
		converter.RowListToSqlInsertQuery(
			generator.UserInterestRows(params.Users, params.Interests, params.MaxUserInterests)),
		"\n",
		params.AfterDataSeedScript,
	))

	fmt.Println(insertDataSql)
}
