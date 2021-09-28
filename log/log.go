package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

const logsDir = "logs"

func Init(logFileName string, writeToConsole bool) {
	if len(logFileName) > 64 {
		panic("Too long log file name")
	}

	var logStream io.Writer
	createLogsDir()
	fStream, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logsDir, logFileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	if writeToConsole {
		logStream = io.MultiWriter(os.Stdout, fStream)
	} else {
		logStream = io.MultiWriter(fStream)
	}

	Info = log.New(logStream, "[INFO]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix)
	Warning = log.New(logStream, "[WARNING]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix)
	Error = log.New(logStream, "[ERROR]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix)
}

func createLogsDir() {
	if _, err := os.Stat(logsDir); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(logsDir, 0666)
			if err != nil {
				panic(err.Error())
			}
		} else {
			panic(err.Error())
		}
	}
}