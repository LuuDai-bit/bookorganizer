package main

import (
	"book-organizer/jobs"
	"book-organizer/migrates"
	"book-organizer/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cron := new(jobs.Cron)
	go cron.RunSchedule()

	indexMigration := new(migrates.IndexMigration)
	indexMigration.Run()

	server.Init()
}
