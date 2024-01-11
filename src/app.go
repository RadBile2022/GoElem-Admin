package src

import (
	"elementary-admin/config"
	"elementary-admin/registries"
	"elementary-admin/src/http/controllers/api"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"strings"
)

// InitServer TODO : InitRepository, InitUseCase
func InitServer(cfg config.Config) Server {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: false,
		Prefork:           false,
	})

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Static("/", "./client2")

	repo := registries.InitRepository(cfg)
	useCase := registries.InitUseCase(repo, cfg)

	return &server{
		httpServer: app,
		cfg:        cfg,
		ucr:        useCase,
	}
}

type Server interface {
	Run()
}

func (di *server) Run() {
	routeGroup := di.httpServer.Group("/api")

	/* ROUTE API */
	authController := api.InitAdminController(di.ucr.Admin())
	authController.Groups(routeGroup)

	// TODO --- write more ---

	port := di.cfg.Port()
	fmt.Println(di.httpServer.Listen(fmt.Sprintf(":%s", port)))
}

type server struct {
	httpServer *fiber.App
	cfg        config.Config
	ucr        registries.UseCaseRegistry
}
