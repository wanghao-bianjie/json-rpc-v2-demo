package conf

import (
	"bytes"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server Server `mapstructure:"server"`
		Mysql  Mysql  `mapstructure:"mysql"`
	}

	Server struct {
		Port         string `mapstructure:"port"`
		Prod         bool   `mapstructure:"prod"`
		DataCenterId int64  `mapstructure:"data_center_id"`
		WorkerId     int64  `mapstructure:"worker_id"`
	}

	Mysql struct {
		Host     string
		Port     int
		Username string
		Password string
		Database string
		Charset  string
		TimeZone string `mapstructure:"time_zone"`
	}
)

func ReadConfig(data []byte) (*Config, error) {
	v := viper.New()
	v.SetConfigType("toml")
	reader := bytes.NewReader(data)
	err := v.ReadConfig(reader)
	if err != nil {
		return nil, err
	}
	var conf Config
	if err := v.Unmarshal(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
