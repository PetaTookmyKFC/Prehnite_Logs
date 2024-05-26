package main

import (
	"log"

	"github.com/Petatookmykfc/Prehnite_logs/utils"
)

func main() {

	err := utils.CreateFile("./TESTFOLDER/Data.db")
	if err != nil {
		log.Panic(err)
	}

}

func CreateLogger(directory string) *Logger {

	// Check if folder exists

	return &Logger{
		Directory: directory,
		LogPools:  make(map[string]Binder),
	}
}

type Logger struct {
	Directory string
	// LogPools are used to seperate the logs
	// These should be used to seperate the logs based on modules
	LogPools map[string]Binder
}

// Binders are used to seperate the logs based on the a submethod.
type Binder struct {
	Ledger map[string]Ledger
}

type Ledger struct {
}
