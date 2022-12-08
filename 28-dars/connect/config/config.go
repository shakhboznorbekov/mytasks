package config

type Config struct {
	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string
}

func Load() Config {
	var cfg Config
	cfg.PostgresHost = "localhost"
	cfg.PostgresDatabase = "korzinka"
	cfg.PostgresUser = "admin"
	cfg.PostgresPassword = "2711"
	cfg.PostgresPort = "5432"

	return cfg
}
