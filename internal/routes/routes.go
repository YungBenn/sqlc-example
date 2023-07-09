package routes

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"sqlc-example/internal/todo"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *pgxpool.Pool) {
	repo := todo.NewRepository(db)
	handler := todo.NewHandler(repo, db)

	router := app.Group("api/v1/todos")

	router.Post("/", handler.Create)
	router.Get("/", handler.GetAllTodos)
	router.Get("/:id", handler.GetTodo)
	router.Put("/:id", handler.UpdateTodo)
	router.Delete("/:id", handler.DeleteTodo)
}