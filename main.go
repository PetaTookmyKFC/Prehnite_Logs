package main

import (
	"log"

	"github.com/Petatookmykfc/Prehnite_logs/logger"
	"github.com/Petatookmykfc/Prehnite_logs/utils"
)

func main() {

	err := utils.CreateFile("./TESTFOLDER/Data.db")
	if err != nil {
		log.Panic(err)
	}

	Log := logger.CreateLogger("./TESTFOLDER")

}
