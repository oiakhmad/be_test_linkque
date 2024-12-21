package env

import (
	"errors"
	"os"
	"strconv"
	"strings"

	port "be_test_linkque/core/port/utils"

	_ "github.com/joho/godotenv/autoload"
)

type env struct{}

func GetInstance() port.Env {
	return &env{}
}

func (e *env) SanitizeEnv(envName string) (string, error) {
	if len(envName) == 0 {
		return "", errors.New("Environment Variable Name Should Not Empty")
	}

	retValue := strings.TrimSpace(os.Getenv(envName))
	if len(retValue) == 0 {
		return "", errors.New("Environment Variable '" + envName + "' Has an Empty Value")
	}

	return retValue, nil
}

func (e *env) GetEnvString(envName string) (string, error) {
	envValue, err := e.SanitizeEnv(envName)
	if err != nil {
		return "", err
	}

	return envValue, nil
}

func (e *env) GetEnvBool(envName string) (bool, error) {
	envValue, err := e.SanitizeEnv(envName)
	if err != nil {
		return false, err
	}

	retValue, err := strconv.ParseBool(envValue)
	if err != nil {
		return false, err
	}

	return retValue, nil
}

func (e *env) GetEnvInt(envName string) (int, error) {
	envValue, err := e.SanitizeEnv(envName)
	if err != nil {
		return 0, err
	}

	retValue, err := strconv.ParseInt(envValue, 0, 0)
	if err != nil {
		return 0, err
	}

	return int(retValue), nil
}

func (e *env) GetEnvFloat32(envName string) (float32, error) {
	envValue, err := e.SanitizeEnv(envName)
	if err != nil {
		return 0, err
	}

	retValue, err := strconv.ParseFloat(envValue, 32)
	if err != nil {
		return 0, err
	}

	return float32(retValue), nil
}

func (e *env) GetEnvFloat64(envName string) (float64, error) {
	envValue, err := e.SanitizeEnv(envName)
	if err != nil {
		return 0, err
	}

	retValue, err := strconv.ParseFloat(envValue, 64)
	if err != nil {
		return 0, err
	}

	return float64(retValue), nil
}
