package models

import (
  "os"

  "book-organizer/db"
)

// Mongo server ip -> localhost -> 127.0.0.1 -> 0.0.0.0
var server = os.Getenv("DB_HOST")

// Database name
var databaseName = os.Getenv("DATABASE_NAME")

// Create a connection
var dbConnect = db.Connect()
