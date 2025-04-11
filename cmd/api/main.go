package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rzmsv/golang-project/config"
	"github.com/rzmsv/golang-project/internal/db"
	"github.com/rzmsv/golang-project/internal/handler"
	"github.com/rzmsv/golang-project/internal/repository"
	v1_route "github.com/rzmsv/golang-project/internal/routes/v1"
	"github.com/rzmsv/golang-project/internal/service"
	"github.com/rzmsv/golang-project/internal/utils"
	"log"
)

func main() {
	e := echo.New()

	/* -------------------------------------------------------------------------- */
	/*                                   Configs                                  */
	/* -------------------------------------------------------------------------- */
	configs := config.Configs(e)
	/* -------------------------------------------------------------------------- */
	/*                                    Utils                                   */
	/* -------------------------------------------------------------------------- */
	utils.Validator(&e.Validator)
	/* -------------------------------------------------------------------------- */
	/*                                  Databases                                 */
	/* -------------------------------------------------------------------------- */
	db.InitMongoDB(configs)
	/* -------------------------------------------------------------------------- */
	/*                                 Repository                                 */
	/* -------------------------------------------------------------------------- */
	productRepo := repository.NewProductRepository(configs)
	/* -------------------------------------------------------------------------- */
	/*                                   Service                                  */
	/* -------------------------------------------------------------------------- */
	productService := service.NewProductService(productRepo)
	/* -------------------------------------------------------------------------- */
	/*                                   Handler                                  */
	/* -------------------------------------------------------------------------- */
	productHandler := handler.NewProductHandler(productService)
	/* -------------------------------------------------------------------------- */
	/*                                   Routes                                   */
	/* -------------------------------------------------------------------------- */
	v1_route.NewProductRoutes(productHandler, configs.E).ProductV1ApiRoutes()

	log.Printf("App connect to port: %s", configs.AppConfig("APP_PORT"))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", configs.AppConfig("APP_PORT"))))
}
