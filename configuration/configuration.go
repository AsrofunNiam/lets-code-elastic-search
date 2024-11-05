package configuration

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	// connection to db
	Port     string `mapstructure:"PORT"`
	PortDB   string `mapstructure:"PORT_DB"`
	Host     string `mapstructure:"HOST_DB"`
	Password string `mapstructure:"PASSWORD_DB"`
	User     string `mapstructure:"USER_DB"`
	Db       string `mapstructure:"DATABASE_DB"`

	// connection to elastic
	ElasticHost     string `mapstructure:"ELASTIC_HOST"`
	ElasticPort     string `mapstructure:"ELASTIC_PORT"`
	ElasticUser     string `mapstructure:"ELASTIC_USER"`
	ElasticPassword string `mapstructure:"ELASTIC_PASSWORD"`
}

func LoadConfig() (config Configuration, err error) {
	viper.SetConfigFile("./configuration/.env")
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
