package config

import "github.com/spf13/viper"

type Config struct {
	Workers    int // Количество воркеров
	Ttl        int64
	ServerPort string
}

func Init() (*Config, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return &Config{}, err
	}

	cfg := &Config{
		Workers:    viper.GetInt("workers"),
		Ttl:        viper.GetInt64("ttl"),
		ServerPort: viper.GetString("server.port"),
	}

	return cfg, nil
}
