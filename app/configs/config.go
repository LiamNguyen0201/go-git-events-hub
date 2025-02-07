package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// DotEnvVariable -> get .env
func DotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("(DotEnvVariable) No .env file found. Using system environment variables.")
	}

	return os.Getenv(key)
}

func DotEnvVariableWithDefault(key string, defaultValue string) string {
	var stringVal = DotEnvVariable(key)
	if stringVal == "" {
		return defaultValue
	}
	return stringVal
}

func DotEnvVariableBool(key string, defaultValue bool) bool {
	var stringVal = DotEnvVariable(key)
	if stringVal == "" {
		return defaultValue
	}
	result, err := strconv.ParseBool(stringVal)
	if err != nil {
		return defaultValue
	}
	return result
}

func DotEnvVariableInt(key string, defaultValue int64) int64 {
	var stringVal = DotEnvVariable(key)
	if stringVal == "" {
		return defaultValue
	}
	result, err := strconv.ParseInt(stringVal, 10, 64)
	if err != nil {
		return defaultValue
	}
	return result
}

var (
	ApplicationName = "git-events-hub"
	DatabasePath    = "events.db"
	GitLabAPIURL    = "https://gitlab.com/api/v4/events"
	GitLabToken     = DotEnvVariable("GITLAB_TOKEN") // Set this as an env variable
	Environment     = DotEnvVariableWithDefault("ENVIRONMENT", "development")
	LogWithLogstash = DotEnvVariableBool("LOG_WITH_LOGSTASH", false)
	LogstashAddress = DotEnvVariableWithDefault("LOG_LOGSTASH_ADDRESS", "localhost:5044")
	LogWithLoki     = DotEnvVariableBool("LOG_WITH_LOKI", false)
	LokiAddress     = DotEnvVariableWithDefault("LOG_LOKI_ADDRESS", "http://localhost:3100")
	LokiUsername    = DotEnvVariableWithDefault("LOG_LOKI_USERNAME", "admin")
	LokiPassword    = DotEnvVariableWithDefault("LOG_LOKI_PASSWORD", "secretpassword")
	PollInterval    = 30 // Polling interval in seconds
	Port            = DotEnvVariableWithDefault("PORT", "8080")
	WebhookURL      = DotEnvVariable("WEBHOOK_URL") // URL to forward events (e.g., Jenkins)
)
