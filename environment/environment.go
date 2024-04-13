package environment

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Pop3 POP3Config
}

type POP3Config struct {
	Username   string
	Password   string
	Host       string
	Port       int
	TLSEnabled bool
}

func Get() Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	pop3TLSEnabledStr := os.Getenv("POP3_TLS_ENABLED")
	pop3TLSEnabled, err := strconv.ParseBool(pop3TLSEnabledStr)
	if err != nil {
		panic("Error parsing POP3_TLS_ENABLED config value")
	}

	pop3port := os.Getenv("POP3_PORT")
	port, err := strconv.Atoi(pop3port)
	if err != nil {
		panic("Error parsing POP3_PORT config value")
	}

	pop3config := POP3Config{
		Username:   os.Getenv("POP3_USERNAME"),
		Password:   os.Getenv("POP3_PASSWORD"),
		Host:       os.Getenv("POP3_HOST"),
		Port:       port,
		TLSEnabled: pop3TLSEnabled,
	}

	return Config{
		Pop3: pop3config,
	}
}
