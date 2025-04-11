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
func BadRequestHandler(message string, c echo.Context) error {
	errorMessage := ErrorMsgType{
		status:  400,
		code:    "bad_request",
		message: message,
	}
	log.Printf("Error message: %s\nError code: %s", errorMessage.message, errorMessage.code)
	return c.String(errorMessage.status, message)
}
