package config

import (
	"shipt/app/models"
)

// not for production usage. Should be replaced with a standard, version-controlled, migration solution instead
// auto-migration does not add useful database constraints
func MigrateDatabase() {
	println("Migrating shipment...")
	DB.AutoMigrate(&models.Shipment{})
	println("Migrating country...")
	DB.AutoMigrate(&models.Country{})
	println("Done.")
}
