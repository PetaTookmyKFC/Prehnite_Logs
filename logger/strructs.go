package logger

type Logger struct {
	Directory string
	// LogPools are used to seperate the logs
	// These should be used to seperate the logs based on modules
	LogPools map[string]Module
}

// Modules are used to manage the folders that each ledger is in.
type Module struct {
	fileloc string
	Ledger  map[string]Ledger
}

// Ledger is used mannage the logs. ( Inside a single file )
type Ledger struct {
	Fileloc string
	Logs    []string
}
