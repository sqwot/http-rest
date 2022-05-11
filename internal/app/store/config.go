package store

type Config struct {
	DatabseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
