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
	DSN  string
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
		DSN: "postgres://lnk:123456@localhost:5432/lnkdb",
	}, nil
}
