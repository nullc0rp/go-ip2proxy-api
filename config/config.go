package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DBUSERNAME string
	DBPASSWORD string
	DBPORT     string
	DBHOST     string
	DBNAME     string
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("./config/%s_config.json", env)
	gonfig.GetConf(fileName, &configuration)
	return configuration
}
