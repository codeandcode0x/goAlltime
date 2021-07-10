package model

import (
	"gin-scaffold/db"
	"os"
)

// init
func init() {
	if runMode := os.Getenv("RUN_MODE"); runMode == "testing" {
		// TO-DO
	} else {
		AutoMigrateAll()
	}
}

// Migrate Model
func AutoMigrateAll() {
	_ = db.Conn.Table("users").AutoMigrate(&User{})
}
