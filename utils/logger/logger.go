package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	serviceName   string
	formatStdOut  = &log.TextFormatter{DisableColors: false, FullTimestamp: true}
	formatFileOut = &log.TextFormatter{DisableColors: true, FullTimestamp: true}
)

func Init(service string) {
	serviceName = service
	createINotExist(service)
}

func Log(httpStatus int, message string, method string, url string, errorFlag bool) {
	fields := log.Fields{
		"method":      method,
		"http_status": httpStatus,
		"service":     serviceName,
		"url":         url,
	}
	if errorFlag {
		log.SetFormatter(formatStdOut)
		log.SetOutput(os.Stdout)
		log.WithFields(fields).Error(message)

		fileOut := openFile(serviceName)
		defer fileOut.Close()
		log.SetFormatter(formatFileOut)
		log.SetOutput(fileOut)
		log.WithFields(fields).Error(message)
	}

	log.SetFormatter(formatStdOut)
	log.SetOutput(os.Stdout)
	log.WithFields(fields).Info(message)

	fileOut := openFile(serviceName)
	defer fileOut.Close()
	log.SetFormatter(formatFileOut)
	log.SetOutput(fileOut)
	log.WithFields(fields).Info(message)
}

func openFile(service string) *os.File {
	filepath := "./logs/" + service + ".txt"
	fileOut, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.WithFields(log.Fields{
			"http_status": 0,
			"service":     "logger",
		}).Error(err.Error())
	}
	return fileOut

}

func createINotExist(service string) {
	filepath := "./logs/" + service + ".txt"
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		_, err = os.Create(filepath)
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
}
