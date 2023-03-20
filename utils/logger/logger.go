package logger

import "log"

func Log(httpStatus int, message string, method string, url string) {
	log.Println(httpStatus, method, url, message)
}
