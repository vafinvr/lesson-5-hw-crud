package config

import (
	"fmt"
	"strings"
)

const (
	EnvPrefix = "" // Префикс для имен переменных

	listenAddr = EnvPrefix + "LISTEN_ADDR"

	dbHost     = EnvPrefix + "DB_HOST"
	dbUser     = EnvPrefix + "DB_USER"
	dbPassword = EnvPrefix + "DB_PASSWORD"
	dbBase     = EnvPrefix + "DB_BASE"
	dbNetwork  = EnvPrefix + "DB_NETWORK"
)

var instance *config

type DataSource interface {
	GetBoolValue(string) (bool, error)
	GetIntValue(string) (int, error)
	GetStringValue(string) (string, error)
}

func NewConfig(dataSource DataSource) (*config, error) {
	if instance != nil {
		return instance, nil
	}
	instance = &config{
		dataSource: dataSource,
		boolVars:   make(map[string]bool, 2),
		intVars:    make(map[string]int, 4),
		stringVars: make(map[string]string, 8),
		required: []string{
			dbHost,
			dbUser,
			dbBase,
		},
	}

	instance.stringVars[listenAddr] = "8080"

	instance.stringVars[dbPassword] = ""
	instance.stringVars[dbBase] = ""
	instance.stringVars[dbHost] = ""
	instance.stringVars[dbNetwork] = ""
	instance.stringVars[dbUser] = ""

	if err := instance.Load(); err != nil {
		return nil, err
	}

	return instance, nil
}

// config Конфигурация приложения
type config struct {
	dataSource DataSource        // Доступ к переменным
	boolVars   map[string]bool   // Параметры конфигурации типа bool
	intVars    map[string]int    // Параметры конфигурации типа int
	stringVars map[string]string // Параметры конфигурации типа string
	required   []string          // Обязательные параметры
}

// isRequired Проверка на обязательное поле конфигурации
func (c config) isRequired(name string) bool {
	for _, reqName := range c.required {
		if reqName == name {
			return true
		}
	}

	return false
}

// String Параметры конфигурации одной строкой
func (c config) String() string {
	res := ""

	for name, val := range c.boolVars {
		res += fmt.Sprintf("%s: %t\n", name, val)
	}
	for name, val := range c.intVars {
		res += fmt.Sprintf("%s: %d\n", name, val)
	}
	for name, val := range c.stringVars {
		if name != dbPassword {
			res += fmt.Sprintf("%s: %s\n", name, val)
		}
	}

	return res
}

// Load Загрузка параметров конфигурации
func (c *config) Load() error {
	loadErrors := make([]string, 0, 10)

	for name := range c.boolVars {
		val, err := c.dataSource.GetBoolValue(name)
		if err != nil {
			if c.isRequired(name) {
				loadErrors = append(loadErrors, err.Error())
			}
			continue
		}
		c.boolVars[name] = val
	}

	for name := range c.intVars {
		val, err := c.dataSource.GetIntValue(name)
		if err != nil {
			if c.isRequired(name) {
				loadErrors = append(loadErrors, err.Error())
			}
			continue
		}
		c.intVars[name] = val
	}

	for name := range c.stringVars {
		val, err := c.dataSource.GetStringValue(name)
		if err != nil {
			if c.isRequired(name) {
				loadErrors = append(loadErrors, err.Error())
			}
			continue
		}
		c.stringVars[name] = val
	}

	if len(loadErrors) > 0 {
		return fmt.Errorf("Ошибка загрузки конфигурации.\n%s\n", strings.Join(loadErrors, ".\n"))
	}

	return nil
}

func (c config) GetListenAddr() string {
	return c.stringVars[listenAddr]
}

func (c config) GetDbHost() string {
	return c.stringVars[dbHost]
}

func (c config) GetDbUser() string {
	return c.stringVars[dbUser]
}

func (c config) GetDbPassword() string {
	return c.stringVars[dbPassword]
}

func (c *config) GetDbBase() string {
	return c.stringVars[dbBase]
}

func (c *config) GetDbNetwork() string {
	return c.stringVars[dbNetwork]
}
