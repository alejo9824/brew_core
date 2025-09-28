package config

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Port      string `mapstructure:"port"`
		JWTSecret string `mapstructure:"jwtSecret"`
	} `mapstructure:"server"`
	DB struct {
		ConnectionString string `mapstructure:"connectionString"`
	} `mapstructure:"db"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config.dev")
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // Permite sobrescribir con variables de entorno (ej. DB_CONNECTIONSTRING)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
