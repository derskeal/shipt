package main

import (
	"shipt/app/config"
	"shipt/database/seeds"
)

func main() {
	config.InitDatabase()
	config.MigrateDatabase()
	seeds.Seed()
}
