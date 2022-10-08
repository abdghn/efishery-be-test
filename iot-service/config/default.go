package config

import "github.com/spf13/viper"

type Config struct {
	DBUri             string `mapstructure:"MONGODB_LOCAL_URI"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	DBPort            string `mapstructure:"DB_PORT"`
	FeederServicePort string `mapstructure:"FEEDER_SERVICE_PORT"`
	IotServicePort    string `mapstructure:"IOT_SERVICE_PORT"`
	RabbitMQUser      string `mapstructure:"RABBITMQ_USER"`
	RabbitMQPass      string `mapstructure:"RABBITMQ_PASS"`
	RabbitMQPort      string `mapstructure:"RABBITMQ_PORT"`
	RabbitMQUI        string `mapstructure:"RABBITMQ_UI"`
	RabbitMQHost      string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQProtocol  string `mapstructure:"RABBITMQ_PROTOCOL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigType("./.env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
