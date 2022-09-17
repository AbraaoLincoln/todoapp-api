package log

import "log"

func Info(message string) {
	log.Println("[ INFO ]", message)
}

func Warn(message string) {
	log.Println("[ WARN ]", message)
}

func Error(message string) {
	log.Println("[ ERROR ]", message)
}
