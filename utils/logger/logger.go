package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func Log(httpStatus int, message string, method string, url string, service string, errorFlag bool) {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	if errorFlag {
		log.SetFormatter(&log.TextFormatter{DisableColors: false, FullTimestamp: true})
		log.SetOutput(os.Stdout)
		logError(httpStatus, message, method, url, service)

		fileOut := openFile(service)
		log.SetFormatter(&log.TextFormatter{DisableColors: true, FullTimestamp: true})
		log.SetOutput(fileOut)
		logError(httpStatus, message, method, url, service)
	}

	log.SetFormatter(&log.TextFormatter{DisableColors: false, FullTimestamp: true})
	log.SetOutput(os.Stdout)
	logINFO(httpStatus, message, method, url, service)

	fileOut := openFile(service)
	log.SetFormatter(&log.TextFormatter{DisableColors: true, FullTimestamp: true})
	log.SetOutput(fileOut)
	logINFO(httpStatus, message, method, url, service)
}

func logINFO(httpStatus int, message string, method string, url string, service string) {
	log.WithFields(log.Fields{
		"method":      method,
		"http_status": httpStatus,
		"service":     service,
		"url":         url,
	}).Info(message)
}

func logError(httpStatus int, message string, method string, url string, service string) {
	log.WithFields(log.Fields{
		"method":      method,
		"http_status": httpStatus,
		"service":     service,
		"url":         url,
	}).Error(message)
}

func openFile(service string) *os.File {
	filepath := "./logs/" + service + ".txt"
	fileOut, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if os.IsNotExist(err) {
		fileOut, err = os.Create(filepath)
		if err != nil {
			log.WithFields(log.Fields{
				"http_status": 0,
				"service":     "logger",
			}).Error(err.Error())
		}
	} else if err != nil {
		log.WithFields(log.Fields{
			"http_status": 0,
			"service":     "logger",
		}).Error(err.Error())
	}
	return fileOut

}
