package core

import "github.com/labstack/echo/v4"

func ResponseOk(message interface{}, c echo.Context) error {
	responseType := ResponseOkType{
		status:  200,
		code:    "OK",
		message: message,
	}
	return c.JSON(responseType.status, map[string]interface{}{
		"status":  responseType.status,
		"code":    responseType.code,
		"message": responseType.message,
	})
}
func Created(message string, c echo.Context) error {
	responseType := ResponseOkType{
		status:  201,
		code:    "Created",
		message: message,
	}
	return c.JSON(responseType.status, map[string]interface{}{
		"status":  responseType.status,
		"code":    responseType.code,
		"message": responseType.message,
	})
}
