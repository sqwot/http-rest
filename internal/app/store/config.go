package store

type Config struct {
	DatabseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		DatabseURL: "postgres://sqwot:windows@localhost/restapi_dev?sslmode=disable",
	}
}
