// myapp/main.go
package main

import (
	logging "github.com/sadikshaik4324/LoggingRepo"
)

func main() {

	logging.InitLogger()

	
	logging.Info("This is an informational message.")
	logging.Warn("This is a warning message.")
	logging.Error("This is an error message.")
	logging.Debug("This is a debug message.")
	logging.Fatal("This is a fatal message.")
	logging.Panic("This is a panic message.")

	// logging.LogWithTime()

}
