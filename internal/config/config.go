package config

import (
	"flag"
	"os"
)

type Options struct {
	Host         string
	ResolvedHost string
}

var flagHost string
var flagResolvedHost string

func NewConfig() *Options {
	flag.StringVar(&flagHost, "a", "localhost:8080", "server address")
	flag.StringVar(&flagResolvedHost, "b", "http://localhost:8080", "generated url address")
	flag.Parse()

	if envServerAddr := os.Getenv("SERVER_ADDRESS"); envServerAddr != "" {
		flagHost = envServerAddr
	}

	if envBaseAddr := os.Getenv("BASE_URL"); envBaseAddr != "" {
		flagResolvedHost = envBaseAddr
	}

	return &Options{
		Host:         flagHost,
		ResolvedHost: flagResolvedHost,
	}
}
