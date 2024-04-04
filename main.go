package main

import (
	"Banking/app"
	"Banking/logger"
)


func main() {
	logger.Info("Strting the application")
	app.Start()
}

