package config

import (
	"flag"
	"fmt"
)

type Config struct {
	Workers int // Количество воркеров

}

func Init() (*Config, error) {

	workers := flag.Int("workers", 5, "an int")

	flag.Parse()

	fmt.Println(*workers)

	cfg := &Config{
		Workers: *workers,
	}

	return cfg, nil
}
