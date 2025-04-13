package core

import (
	"github.com/labstack/echo/v4"
	"log"
)

func PanicHandler(message string) {
	errorMessage := ErrorMsgType{
		status:  0,
		code:    "Panic",
		message: message,
	}
	log.Fatalf("Error message: %s\nError code: %s (It's mean panic Error!)", errorMessage.message, errorMessage.code)
	panic(message)
}

/* ------------------------------- 500 ERRORS ------------------------------- */
func InternalServerError(message string, c echo.Context) error {
	errorMessage := ErrorMsgType{
		status:  500,
		code:    "internal_server_error",
		message: message,
	}
	log.Printf("Error message: %s\nError code: %s", errorMessage.message, errorMessage.code)
	return c.JSON(errorMessage.status, map[string]interface{}{
		"status":  errorMessage.status,
		"code":    errorMessage.code,
		"message": errorMessage.message,
	})
}

/* ------------------------------- 400 ERRORS ------------------------------- */
func BadRequestHandler(message string, c echo.Context) error {
	errorMessage := ErrorMsgType{
		status:  400,
		code:    "bad_request",
		message: message,
	}
	log.Printf("Error message: %s\nError code: %s", errorMessage.message, errorMessage.code)
	return c.JSON(errorMessage.status, map[string]interface{}{
		"status":  errorMessage.status,
		"code":    errorMessage.code,
		"message": errorMessage.message,
	})
}
func NotFound(c echo.Context) error {
	errorMessage := ErrorMsgType{
		status:  404,
		code:    "not_found",
		message: struct{}{},
	}
	log.Printf("Error message: %s\nError code: %s", errorMessage.message, errorMessage.code)
	return c.JSON(errorMessage.status, map[string]interface{}{
		"status":  errorMessage.status,
		"code":    errorMessage.code,
		"message": errorMessage.message,
	})
}
