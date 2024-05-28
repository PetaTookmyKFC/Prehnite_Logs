package logger

import "path/filepath"

func CreateLogger(directory string) *Logger {

	// Check if folder exists

	return &Logger{
		Directory: directory,
		LogPools:  make(map[string]Module),
	}
}

func (l *Logger) GetModule(name string) Module {
	// Check if the module is already loaded
	if _, ok := l.LogPools[name]; ok {
		return l.LogPools[name]
	}
	path := filepath.Join(l.Directory, name)
	// Create the module
	l.LogPools[name] = Module{
		fileloc: path,
		Ledger:  make(map[string]Ledger),
	}
	return l.LogPools[name]
}

func (m *Module) CreateLedger(Name string) *Ledger {

	// Check if the ledger exists
	var led *Ledger
	if led, ok := m.Ledger[Name]; ok {
		return &led
	}
	// Create the ledger
	led = &Ledger{}
	return led
}
