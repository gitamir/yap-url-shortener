package config

type Options struct {
	Host string
	ResolvedHost string
}

func NewConfig(host, resolvedHost string) *Options {
	return &Options{
		Host: host,
		ResolvedHost: resolvedHost,
	}
}