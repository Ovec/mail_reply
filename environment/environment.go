package environment

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Pop3  ConnectionConfig
	Smtp  ConnectionConfig
	Email EmailConfig
}

type ConnectionConfig struct {
	Username   string
	Password   string
	Host       string
	Port       int
	TLSEnabled bool
}

type EmailConfig struct {
	Sender string
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

	pop3config := ConnectionConfig{
		Username:   os.Getenv("POP3_USERNAME"),
		Password:   os.Getenv("POP3_PASSWORD"),
		Host:       os.Getenv("POP3_HOST"),
		Port:       port,
		TLSEnabled: pop3TLSEnabled,
	}

	smtpPort := os.Getenv("SMTP_PORT")
	port, err = strconv.Atoi(smtpPort)
	if err != nil {
		panic("Error parsing SMTP_PORT config value")
	}

	smtpTLSEnabledStr := os.Getenv("POP3_TLS_ENABLED")
	smtpTLSEnabled, err := strconv.ParseBool(smtpTLSEnabledStr)
	if err != nil {
		panic("Error parsing SMTP_TLS_ENABLED config value")
	}

	smtpConfig := ConnectionConfig{
		Username:   os.Getenv("SMTP_USERNAME"),
		Password:   os.Getenv("SMTP_PASSWORD"),
		Host:       os.Getenv("SMTP_HOST"),
		Port:       port,
		TLSEnabled: smtpTLSEnabled,
	}

	emailSender := os.Getenv("EMAIL_SENDER")
	if err != nil {
		panic("Error parsing EMAIL_SENDER config value")
	}

	return Config{
		Pop3:  pop3config,
		Smtp:  smtpConfig,
		Email: EmailConfig{Sender: emailSender},
	}
}
