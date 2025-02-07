package databases

import (
	"git_events_hub/models"
	"git_events_hub/utils"
	"log"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// Initialize the SQLite database
func InitDB() {
	log.Println("(InitDB) Starting ...")

	var err error
	db, err = gorm.Open(sqlite.Open("events.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // ✅ Logs all SQL queries
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate schema
	db.AutoMigrate(&models.GitLabEvent{})
}

// Check if an event already exists
func EventExists(eventID int64) bool {
	var count int64
	db.Model(&models.GitLabEvent{}).Where("id = ?", eventID).Count(&count)
	return count > 0
}

func GetEventByID(eventID int64) (*models.GitLabEvent, error) {
	var event models.GitLabEvent

	// Query event by ID
	err := db.First(&event, eventID).Error
	if err != nil {
		return nil, err // Return error if event is not found
	}

	return &event, nil // Return event
}

func GetEvents(projectID string, startDate string, endDate string, page int, limit int) (*[]models.GitLabEvent, int64) {
	log.Println("(GetEvents) Page: " + strconv.Itoa(page))
	log.Println("(GetEvents) Limit: " + strconv.Itoa(limit))

	var events []models.GitLabEvent
	var total int64

	offset := (page - 1) * limit // Pagination logic

	// Query builder
	query := db.Model(&models.GitLabEvent{})
	if projectID != "" {
		query = query.Where("project_id = ?", utils.StringToNumber(projectID))
	}
	if startDate != "" && endDate != "" {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// Count total results
	query.Count(&total)

	// Fetch paginated results
	query.Limit(limit).Offset(offset).Find(&events)

	return &events, total // Return event
}

// Save a new event
func SaveEvent(event models.GitLabEvent) {
	db.Create(&event)
}
