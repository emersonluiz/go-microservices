package logger

import (
	"log"
	"os"
)

func SetLog(message string) {
	LOG_FILE := os.Getenv("LOG_FILE")

	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println(message)
}
