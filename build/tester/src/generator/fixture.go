package generator

import (
	"component/file"
	"component/handler"
	"converter"
	"fmt"
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
const EnvRemoveIndexScript = "REMOVE_INDEX_SCRIPT"
const EnvFixtureGeneratedScript = "GENERATED_FIXTURE_SCRIPT"

type SeedDataParams struct {
	Cities               	int
	Interests            	int
	Users                	int
	MinAge               	int
	MaxAge               	int
	MaxUserInterests     	int
	BeforeDataSeedScript 	string
	AfterDataSeedScript  	string
	RemoveIndexScript 	   	string
	FixtureGeneratedScript 	string
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
		EnvRemoveIndexScript,
		EnvFixtureGeneratedScript,
	}

	stringMap := map[string]string{
		EnvSqlScriptsPath:   "",
		EnvBeforeSeedScript: "",
		EnvAfterSeedScript:  "",
		EnvRemoveIndexScript: "",
		EnvFixtureGeneratedScript: "",
	}

	intMap := map[string]int{
		EnvGenerateCities:    0,
		EnvGenerateInterests: 0,
		EnvGenerateUsers:     0,
		EnvUserMinAge:        0,
		EnvUserMaxAge:        0,
		EnvMaxUserInterests:  0,
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
		Cities:               intMap[EnvGenerateCities],
		Interests:            intMap[EnvGenerateInterests],
		Users:                intMap[EnvGenerateUsers],
		MinAge:               intMap[EnvUserMinAge],
		MaxAge:               intMap[EnvUserMaxAge],
		MaxUserInterests:     intMap[EnvMaxUserInterests],
		BeforeDataSeedScript: stringMap[EnvSqlScriptsPath] + stringMap[EnvBeforeSeedScript],
		AfterDataSeedScript:  stringMap[EnvSqlScriptsPath] + stringMap[EnvAfterSeedScript],
		RemoveIndexScript:  stringMap[EnvSqlScriptsPath] + stringMap[EnvRemoveIndexScript],
		FixtureGeneratedScript:  stringMap[EnvSqlScriptsPath] + stringMap[EnvFixtureGeneratedScript],
	}, nil
}

func GenerateFixture(params SeedDataParams) error {
	beforeDataSeed, err := file.GetContent(params.BeforeDataSeedScript)
	if err != nil {
		handler.ErrorLog(err)
		return err
	}
	afterDataSeed, err := file.GetContent(params.AfterDataSeedScript)
	if err != nil {
		handler.ErrorLog(err)
		return err
	}
	removeIndex, err := file.GetContent(params.RemoveIndexScript)
	if err != nil {
		handler.ErrorLog(err)
		return err
	}

	err = file.WriteList(
		params.FixtureGeneratedScript,
		[]string{
			beforeDataSeed + "\n",
			removeIndex + "\n",
			converter.RowListToSqlInsertQueryList(CityRows(params.Cities)) + "\n",
			converter.RowListToSqlInsertQueryList(InterestRows(params.Interests)) + "\n",
			converter.RowListToSqlInsertQueryList(
				UserRows(params.Users, params.Cities, params.MinAge, params.MaxAge)) + "\n",
			converter.RowListToSqlInsertQueryList(
				UserInterestRows(params.Users, params.Interests, params.MaxUserInterests)) + "\n",
			afterDataSeed,
		})

	if err != nil {
		return err
	}
	return nil
}
