package tool

import (
	"log"
)

func Info(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func Warn(format string, args ...interface{}) {
	log.Printf("[WARN] "+format, args...)
}


func Error(err error, msg string) bool {
	if err != nil {
		log.Printf("[ERROR] %s: %v", msg, err)
		return true
	}
	return false
}

func Errorf(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}

