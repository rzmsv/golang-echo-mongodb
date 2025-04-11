package core

import (
	"github.com/labstack/echo/v4"
)

func RequestHandler(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		incomingFunction := fn(c)
		status := c.Response().Status
		if status >= 400 {
			return c.JSON(status, map[string]string{
				"dsada": "dsadas",
			})
		} else {
			return incomingFunction
		}
	}
}
