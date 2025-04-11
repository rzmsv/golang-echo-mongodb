package config

import (
	dotenv "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rzmsv/golang-project/core"
)

type AppConfigs struct {
	E *echo.Echo
}

func Configs(e *echo.Echo) *AppConfigs {
	return &AppConfigs{
		E: e,
	}
}
func (c *AppConfigs) AppConfig(envKey string) string {
	envs, err := dotenv.Read(".env")
	if err != nil {
		core.PanicHandler(err.Error())
	}
	if len(envs) <= 0 {
		core.PanicHandler("Env file was Empty!")
	}

	configsList := make(map[string]string)
	for key, value := range envs {
		configsList[key] = value
	}
	return configsList[envKey]
}
