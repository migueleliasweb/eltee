package httphandlers

import (
	"net/http"

	"github.com/labstack/echo"
)

//GetStatsHandler Returns data for the /stat endpoint
func GetStatsHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
