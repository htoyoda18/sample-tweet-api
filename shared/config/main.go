package config

import "github.com/kelseyhightower/envconfig"

type ConfigurationMode string

const (
	ConfigurationModeByDebug = ConfigurationMode("debug")
	ConfigurationModeByTest  = ConfigurationMode("test")
)

type Configuration struct {
	SecretKey  string            `default:"secret key"`
	DBHost     string            `default:"db"`
	DBPort     string            `default:"3306"`
	DBName     string            `default:"attivita_development"`
	DBNameTest string            `default:"attivita_test"`
	DBUser     string            `default:"root"`
	DBUserTest string            `default:"root"`
	DBPass     string            `default:"hogehoge"`
	DBPassTest string            `default:"hogehoge"`
	APIHost    string            `default:"http://localhost:8080"`
	WEBHost    string            `default:"http://localhost:8081"`
	Mode       ConfigurationMode `default:"debug"`
}

func InitConfiguration() (Configuration, error) {
	var conf Configuration
	if err := envconfig.Process("test", &conf); err != nil {
		return conf, err
	}

	return conf, nil
}
