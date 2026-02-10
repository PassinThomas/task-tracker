package main

import(
	"log/slog"
	"task/cmd"
)
	
func main() {
	slog.Info("Hello task-cli")
	cmd.Execute()
}

