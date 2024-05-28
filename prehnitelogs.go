package prehnitelogs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

const MainLoc string = "Logs"

var LogLocks = sync.Map{}

type LogType int8

const (
	plain LogType = iota
	info
	warn
	danger
)

func Log(msg string)    { writeLog(plain, msg) }
func Info(msg string)   { writeLog(info, msg) }
func Warn(msg string)   { writeLog(warn, msg) }
func Danger(msg string) { writeLog(danger, msg) }

// type Log struct{}
func writeLog(t LogType, msg string) {
	// Get Folder - Module
	// Get File	  - Function
	// Get Time   - Date
	var folder string
	var file string
	var function string
	var time = time.Now().Format("2006-01-02 15:04:05")

	program, file, line, _ := runtime.Caller(2)
	fn := runtime.FuncForPC(program)

	function = fn.Name()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	wd = strings.ReplaceAll(wd, "\\", "/")

	file, _ = strings.CutPrefix(file, wd)

	folder = filepath.Dir(file)
	folder = strings.Replace(folder, "\\", "/", -1)

	function = filepath.Ext(function)
	function = function[1:]

	file = strings.Replace(file, folder+"/", "", 1)

	// fmt.Println(wd + folder + "/" + file + ":" + function + " - " + fmt.Sprint(line))
	// fmt.Println("folder : ", folder)
	// fmt.Println("file : ", file)
	// fmt.Println("function : ", function)
	// fmt.Println("line : ", line)

	// Make the path for the folder ->
	fileLoc := filepath.Join(wd, MainLoc, folder)
	err = os.MkdirAll(fileLoc, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}

	// Make the path for the file ->
	fullLocation := fileLoc + "/" + function + ".log"

	// Create the error string

	// Create the string to append
	var Log string = ""

	switch t {
	case plain:
		Log = "Log     >> "
	case info:
		Log = "Info    >> "
	case warn:
		Log = "Warning >> "
	case danger:
		Log = "Danger  >> "
	}

	Log = fmt.Sprintf("%s ( %s ) %s:%s >> %s\n", Log, time, file, fmt.Sprint(line), msg)
	fmt.Println(Log)

	lock, _ := LogLocks.LoadOrStore(fullLocation, &sync.Mutex{})
	lock.(*sync.Mutex).Lock()
	defer lock.(*sync.Mutex).Unlock()

	var ofile *os.File
	// check if the file exists

	ofile, err = os.OpenFile(fullLocation, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer ofile.Close()

	_, err = ofile.WriteString(Log)
	if err != nil {
		log.Fatal(err)
	}
}

func TestLogs() {
	// Loop 5 times to test the log functions
	for i := 0; i < 5; i++ {
		// Call the Log function with a message
		Log("TEST FUNCTION ( plain )  ")
		// Call the Info function with a message
		Info("TEST FUNCTION ( info )  ")
		// Call the Warn function with a message
		Warn("TEST FUNCTION ( warn )  ")
		// Call the Danger function with an error
		Danger("TEST FUNCTION ( danger )  ")
	}
}
