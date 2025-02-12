package databases

import (
	"git_events_hub/models"
	"git_events_hub/utils"

	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// Initialize the SQLite database
func InitDB(logrusLogger *logrus.Logger) {
	utils.LogInfo("(InitDB) Starting ...")

	var err error

	// GORM logger using Logrus
	newLogger := logger.New(
		logrusLogger, // Logrus as GORM logger
		logger.Config{
			SlowThreshold: time.Second, // Log queries slower than 1s
			LogLevel:      logger.Info, // Log all SQL queries
			Colorful:      false,
		},
	)

	db, err = gorm.Open(sqlite.Open("events.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		utils.LogFatal("Failed to connect to database:", err)
	}

	// Auto-migrate schema
	db.AutoMigrate(&models.GitLabEvent{})
	db.AutoMigrate(&models.GitLabProject{})
	db.AutoMigrate(&models.Workflow{})
}
