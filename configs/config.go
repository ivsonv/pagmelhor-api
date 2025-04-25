package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	API_PORT   string `mapstructure:"API_PORT"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
	JWT_EXPIRY int    `mapstructure:"JWT_EXPIRY"`
	TokenAuth  *jwtauth.JWTAuth

	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	RABBITMQ_DEFAULT_USER string `mapstructure:"RABBITMQ_DEFAULT_USER"`
	RABBITMQ_DEFAULT_PASS string `mapstructure:"RABBITMQ_DEFAULT_PASS"`
	RABBITMQ_URL          string `mapstructure:"RABBITMQ_URL"`

	REDIS_URL string `mapstructure:"REDIS_URL"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		println("Error reading config file, will use environment variables")
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		println("Error unmarshalling config")
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWT_SECRET), nil)
	return cfg, nil
}
