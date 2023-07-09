package todo

import (
	"net/http"
	"sqlc-example/pkg/sqlc"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	repo Repository
	db   *pgxpool.Pool
}

func NewHandler(repo Repository, db *pgxpool.Pool) *Handler {
	return &Handler{repo, db}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req sqlc.CreateTodoParams
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	ctx := c.Context()
	todo, err := h.repo.CreateTodo(ctx, req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "New todo created",
		"data": todo,
	})
}

func (h *Handler) GetAllTodos(c *fiber.Ctx) error {
	ctx :=  c.Context()
	todos, err := h.repo.ListTodos(ctx)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Success get list todos",
		"data": todos,
	})
}

func (h *Handler) GetTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	
	ctx := c.Context()

	todo, err := h.repo.GetTodo(ctx, int32(id))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Success get todo",
		"data": todo,
	})
}

func (h *Handler) UpdateTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	
	var req sqlc.UpdateTodoParams
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	ctx := c.Context()

	req.ID = int32(id)
	todo, err := h.repo.UpdateTodo(ctx, req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Success update todo",
		"data": todo,
	})
}

func (h *Handler) DeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	ctx := c.Context()

	err = h.repo.DeleteTodo(ctx, int32(id))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Success delete todo",
	})
}
