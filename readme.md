# Prehnite log

This is a very basic log system. It will write to the console and write to a log file within the direct `/Logs`. I don't know what I'm doing use at your own risk.


Writes logs into files. Files are stored in folders ( The folders are a direct copy of the source code locations ). For example the function `func MagicFunc ( )` in the file `/wd/magic/test.go ` would log to `/Logs/magic/MagicFunc.log`
The log will also fmt.println the log message into the console

## Error Parts

#### Log Level
This is how important is the log. This can have multiple values. I would sugest using them as follows. The default values include.
* Log - This is the base, this should just be used to note unimportant information
* Info - This is the next level, this is not that important but may be required. Such as optional data.
* Warning - The next level, the program can continue but this should be looked into. An example of this could be a enviroment variable that isn't quite correct. Such as a url thats missing `localhost`
* Danger - This is the highest level. This should be used when a log.Fatal or log.Panic is triggered. The program couldn't continue.

#### Time
This is the time that the log is triggered. ( YYYY-MM-DD HH-MM-SS )

#### Location
This is the location where the log was triggered. this includes the file and the line in the file.

#### Message
This is the message that is send into the log function. ( This must be a string! )

