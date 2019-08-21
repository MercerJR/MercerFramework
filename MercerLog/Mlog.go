package MercerLog

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var(
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Run     *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard,
		"[Mercer-TRACE]: ",
		log.Ldate|log.Ltime)

	Info = log.New(os.Stdout,
		"[Mercer-INFO]: ",
		log.Ldate|log.Ltime)

	Warning = log.New(os.Stdout,
		"[Mercer-WARNING]: ",
		log.Ldate|log.Ltime)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"[Mercer-ERROR]: ",
		log.Ldate|log.Ltime)

	Run = log.New(os.Stdout,
		"[Run]: ",
		log.Ldate|log.Ltime)
}
