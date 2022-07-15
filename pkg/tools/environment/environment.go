package environment

import (
	"os"
	"strconv"
)

func GetStringFromEnv(key, defaultValue string) (string, error) {
	if value, ok := os.LookupEnv(key); ok {
		return value, nil
	}
	return defaultValue, nil
}

func GetIntFromEnv(key string) (int, error) {
	strValue, err := GetStringFromEnv(key, "")
	if err != nil {
		return 0, nil
	}
	value, err := strconv.Atoi(strValue)

	if err != nil {
		return 0, err
	}
	return value, nil

}

func GetBoolFromEnv(key string) (bool, error) {
	strValue, err := GetStringFromEnv(key, "")
	if err != nil {
		return false, nil
	}
	value, err := strconv.ParseBool(strValue)
	if err != nil {
		return false, err
	}
	return value, nil
}
