package config
type Config struct {
	ServerPort string
	DatabasePath string
}

func Load() Config {
	return Config{
		ServerPort: ":8080",
		DatabasePath: "data/app.db",
	}
}