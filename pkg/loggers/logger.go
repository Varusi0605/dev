package loggers

import (
	"log"
	"os"
)

var (
	InfoLog  *log.Logger
	FatalLog *log.Logger
	ErrorLog *log.Logger
	WarnLog  *log.Logger
)

func ForLogs() {
	file, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		log.Panic("Error while opeaning file")
	}

	InfoLog = log.New(file, "INFO:", log.LstdFlags|log.Llongfile)
	FatalLog = log.New(file, "FATAL:", log.LstdFlags|log.Llongfile)
	ErrorLog = log.New(file, "ERROR:", log.LstdFlags|log.Llongfile)
	WarnLog = log.New(file, "WARNING:", log.LstdFlags|log.Llongfile)
}
