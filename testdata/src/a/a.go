package hello

import (
	go_format "fmt"
	"log"
	structLog "log/slog"
)

func Hello() {
	go_format.Println("Hello, World!")
	structLog.Info("Hello, World!")
	log.Println("Hello, World!")
}
