package config

import "time"

type Config struct {
	Port int
	Env  string
	DB   struct {
		MaxOpenConns string
		MaxIdleTime  time.Duration
	}
}
