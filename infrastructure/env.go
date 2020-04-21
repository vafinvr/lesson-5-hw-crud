package infrastructure

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type Env struct{}

// GetBoolValue получает булевое значение из переменной окружения
// name Наименование переменной
func (e Env) GetBoolValue(name string) (bool, error) {
	value, found := os.LookupEnv(strings.ToUpper(name))
	if found {
		res, err := strconv.ParseBool(value)
		if err != nil {
			return false, errors.New("Ошибка конвертации в bool " + err.Error())
		}

		return res, nil
	}

	return false, errors.New(name + " не задана в окружении")
}

// GetIntValue получает целочисленное значение из переменной окружения
// name Наименование переменной
func (e Env) GetIntValue(name string) (int, error) {
	value, found := os.LookupEnv(strings.ToUpper(name))
	if found {
		res, err := strconv.Atoi(value)
		if err != nil {
			return 0, errors.New("Ошибка конвертации в int " + err.Error())
		}

		return res, nil
	}

	return 0, errors.New(name + " не задана в окружении")
}

// GetStringValue получает строчное значение из переменной окружения
// name Наименование переменной
func (e Env) GetStringValue(name string) (string, error) {
	value, found := os.LookupEnv(strings.ToUpper(name))
	if found {
		return value, nil
	}

	return "", errors.New(name + " не задана в окружении")
}
