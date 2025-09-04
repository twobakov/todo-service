package routes

import (
	"todo-service/internal/handlers"
	"todo-service/internal/services"
	"todo-service/internal/storage/repository"
	"todo-service/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func InitRoutes(conn *pgx.Conn) *fiber.App {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})
	app.Use(logger.LoggerMiddleware())

	repo := repository.NewTasksRepository(conn)
	service := services.NewTasksService(repo)
	handler := handlers.NewTasksHandler(service)

	api := app.Group("/api")
	{
		api.Post("/tasks", handler.CreateTask)
		api.Get("/tasks", handler.GetTasks)
		api.Put("/tasks/:id", handler.UpdateTask)
		api.Delete("/tasks/:id", handler.DeleteTask)
	}

	return app
}
