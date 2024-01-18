package migrates

import (
	"os"

	"book-organizer/db"
)

// Database name
var databaseName = os.Getenv("DATABASE_NAME")

// Create a connection
var dbConnect = db.Connect()
