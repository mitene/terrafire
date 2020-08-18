package utils

import log "github.com/sirupsen/logrus"

func LogError(err error) {
	if err != nil {
		log.Error(err)
	}
}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogDefer(f func() error) {
	LogError(f())
}
