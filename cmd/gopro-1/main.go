package main

import (
	"log/slog"
	"time"

	"github.com/kongra/gopro-1/internal/routes"
)

func main() {
	x := routes.SomeValue
	slog.Info("Hello", "value", x, "dupa", "jaros")

	<-time.After(3 * time.Second)
	slog.Info("Some text")
}
