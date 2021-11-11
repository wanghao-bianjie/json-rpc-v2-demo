package conf

type (
	Config struct {
		Server Server `mapstructure:"server"`
		Mysql  Mysql  `mapstructure:"mysql"`
	}

	Server struct {
		Port string `mapstructure:"port"`
		Prod bool   `mapstructure:"prod"`
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
