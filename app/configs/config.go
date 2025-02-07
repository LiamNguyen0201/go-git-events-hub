package configs

import (
	"log"
	"os"

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

var (
	GitLabAPIURL = "https://gitlab.com/api/v4/events"
	GitLabToken  = DotEnvVariable("GITLAB_TOKEN") // Set this as an env variable
	WebhookURL   = DotEnvVariable("WEBHOOK_URL")  // URL to forward events (e.g., Jenkins)
	PollInterval = 30                             // Polling interval in seconds
	Port         = DotEnvVariable("PORT")
	DatabasePath = "events.db"
)
