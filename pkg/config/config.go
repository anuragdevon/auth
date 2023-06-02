package config

import "github.com/spf13/viper"

type Config struct {
	Port         string `mapstructure:"PORT"`
	DbHost       string `mapstructure:"DB_HOST"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
	DbUser       string `mapstructure:"DB_USER"`
	DbPassword   string `mapstructure:"DB_PASSWORD"`
	DbName       string `mapstructure:"DB_NAME"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("/auth_svc/pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
