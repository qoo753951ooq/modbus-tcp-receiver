package util

import (
	"log"
	"os"
)

const logFolder = "log"

func Writelog(message string) {

	logFolderExist, err := PathExist(logFolder)

	if !logFolderExist {
		os.Mkdir(logFolder, 0666)
	}

	logName := CombineString(logFolder, "/", GetDateNow(), ".log")

	file, err := os.OpenFile(logName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println(message)
}
