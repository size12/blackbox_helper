package main

import "flag"

type Config struct {
	AccrualSystemAddress string
}

func GetConfig() Config {
	var cfg Config

	flag.StringVar(&cfg.AccrualSystemAddress, "r", "http://127.0.0.1:8080", "Accrual system address")
	flag.Parse()

	return cfg
}
