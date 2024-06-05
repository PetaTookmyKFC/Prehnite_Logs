package prehnitelogs

import (
	"errors"
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

var GroupByFunc bool = true

var LogLocks = sync.Map{}

type LogFunc func(msg string) (genMessage error)
type LogType int8

const (
	plain LogType = iota
	info
	warn
	danger
	custom
)

var logtypes map[LogType]string = map[LogType]string{
	plain:  "Log",
	info:   "Info",
	warn:   "Warn",
	danger: "Danger",
	custom: "- ??? -",
}

var cLogs map[string]string = map[string]string{}

func Log(msg string) (genMessage error)    { return writeLog(plain, msg) }
func Info(msg string) (genMessage error)   { return writeLog(info, msg) }
func Warn(msg string) (genMessage error)   { return writeLog(warn, msg) }
func Danger(msg string) (genMessage error) { return writeLog(danger, msg) }

func AddType(name string, prefix string) {
	cLogs[name] = prefix
}

// This is a method so you dont have to keep writing in the type of log you want and you can just save it into a variable
func GetCustomLogMethod(name string) LogFunc {
	return func(msg string) (genMessage error) {
		return writeCustomLog(name, msg)
	}
}

// This method is useful for custom types, these must be added prior to being used. ( Only the prefix can be changed -- to the end of the `>>` characters)
func CustomLog(LogTypeName string, msg string) (genMessage error) {
	return writeCustomLog(LogTypeName, msg)
}

func writeCustomLog(logType string, msg string) error {
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
	// Make the path for the folder ->
	fileLoc := filepath.Join(wd, MainLoc, folder)
	err = os.MkdirAll(fileLoc, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	var fullLocation string
	if !GroupByFunc {
		// Make the path for the file ->
		fullLocation = fileLoc + "/" + function + ".log"
	} else {
		fullLocation = fileLoc + ".log"
	}

	// Create the error string

	// Create the string to append
	var Log string = ""

	// Check logtype exists
	if _, ok := cLogs[logType]; !ok {
		Log = "Undefined ??? >> "
	} else {
		Log = cLogs[logType]
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

	return errors.New(Log)
}

// type Log struct{}
func writeLog(t LogType, msg string) error {
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

	// switch t {
	// case plain:
	// 	Log = "Log     >> "
	// case info:
	// 	Log = "Info    >> "
	// case warn:
	// 	Log = "Warning >> "
	// case danger:
	// 	Log = "Danger  >> "
	// }

	// Check logtype exists
	if _, ok := logtypes[t]; !ok {
		Log = "Log ??? >> "
	} else {
		Log = fmt.Sprintf("%s  >>", logtypes[t])
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

	return errors.New(Log)
}

func RunTestLogs() {
	// Loop 5 times to test the log functions
	for i := 0; i < 5; i++ {
		// Call the Log function with a message
		err := Log("TEST FUNCTION ( plain )  ")
		fmt.Println(err)

		// Call the Info function with a message
		err = Info("TEST FUNCTION ( info )  ")
		fmt.Println(err)
		// Call the Warn function with a message
		err = Warn("TEST FUNCTION ( warn )  ")
		fmt.Println(err)

		// Call the Danger function with an error
		err = Danger("TEST FUNCTION ( danger )  ")
		fmt.Println(err)

	}

	Info("--- MAIN TEST ---")
	Info("--- MAIN TEST ---")
	Info("--- MAIN TEST ---")

	fmt.Println("Test Script")

	AddType("TEST", " - TEST : ")

	err := Log("SimpleLog!")

	fmt.Println("Error - Simple ", err)

	err = CustomLog("TEST", "Wow much custom logging!")
	fmt.Println("Error - Custom ", err)

	// Testing the Make Custom functions
	AddType("Multiple", " - Multiple - So Lazy! : ")
	l := GetCustomLogMethod("Multiple")
	err = l("Wow much custom logging!")
	fmt.Println("Error - Custom MAGIC! ", err)
	err = l("MORE CUSTOM LOGGING!")
	fmt.Println("Error - MORE MAGIC! ", err)
}
