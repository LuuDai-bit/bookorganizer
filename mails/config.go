package mails

import "os"

var from = os.Getenv("MAIL_EMAIL")
var password = os.Getenv("MAIL_PASSWORD")
