package logging

import (
	"log"
	"os"
	"sync"
)

var instance *log.Logger
var once sync.Once

// WriteMessage writes a message to the log
func WriteMessage(msgType string, msg string) {

	if instance == nil {
		once.Do(func() {
			instance = log.New(os.Stdout,
				"FOODAPI: ",
				log.Ldate|log.Ltime)
		})
	}

	instance.Printf("%s %s", msgType, msg)
}
