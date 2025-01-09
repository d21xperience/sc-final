package helper

import "log"

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorLog(err error, msgError string) {
	if err != nil {
		log.Fatalf(msgError, err)
	}
}
