package main

import (
	"go-clean/middleware"

	program "github.com/didintri196/go-mode"
	"github.com/gin-gonic/gin"
)

func main() {
	// flag.Parse()
	config := program.Configuration{
		Mode: program.DebugMode,
		Tipe: "fmt",
	}
	program.SetMode(config)
	gin.SetMode(gin.DebugMode)

	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()

	middleware.Middleware()
}
