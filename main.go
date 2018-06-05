package main

import (
	"github.com/labstack/echo"
	"github.com/migueleliasweb/eltee/httphandlers"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/stats", httphandlers.GetStatsHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
