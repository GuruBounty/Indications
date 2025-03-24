package config

import "github.com/spf13/viper"

type Config struct {
	DB     Postgres
	Server Server `mapstructure:"server"`
}
type Server struct {
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}
type Postgres struct {
	Port    int    `mapstructure:"port"`
	Db      string `mapstructure:"db"`
	Host    string `mapstructure:"host"`
	User    string `mapstructure:"username"`
	Pass    string `mapstructure:"password"`
	SSLMode string `mapstructure:"sslmode"`
}

func New(folder, file string) (*Config, error) {
	cfg := new(Config)

	viper.AddConfigPath(folder)
	viper.SetConfigName(file)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
