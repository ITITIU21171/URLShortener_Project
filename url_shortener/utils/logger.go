package utils

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "URL_SHORTENER: ", log.LstdFlags|log.Lshortfile)
