package config

import "os"

type Config struct {
	ServerPort   string
	DatabasePath string
}

func Load() Config {
	var port = os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	var databasePath = os.Getenv("DATABASE_PATH")
	if databasePath == "" {
		databasePath = "data/app.db"
	}
	return Config{
		ServerPort:   port,
		DatabasePath: databasePath,
	}
}
