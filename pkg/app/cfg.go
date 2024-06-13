package app

import (
	"os"
	"strconv"
)

const (
	defPort = 8080
)

type Config struct {
	Port int
}

func LoadConfig() (*Config, error) {
	return &Config{
		Port: func(port int) int {
			val, ok := os.LookupEnv("PORT")
			if !ok {
				return port
			}
			port, _ = strconv.Atoi(val)

			return port
		}(defPort),
	}, nil
}
