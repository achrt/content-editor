package handler

import (
	"content-editor/cmd/editor/application"
	"content-editor/internal/domain/models"
	"content-editor/internal/domain/repositories"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	catalogue repositories.Catalogue
}

func New(app *application.App) *Handler {
	return &Handler{catalogue: app.Catalogue}
}

func (h *Handler) Route(r fiber.Router) {
	r.Get("/catalogue/:id", h.Read)
	r.Post("/catalogue", h.Create)
}

func (h *Handler) Read(c *fiber.Ctx) error {
	id, err := h.paramInt(c, "id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	cat, err := h.catalogue.Find(id)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(cat)
}

func (h *Handler) Create(c *fiber.Ctx) error {

	catalogue := models.Catalogue{}
	if err := c.BodyParser(&catalogue); err != nil {
		return fiber.ErrBadRequest
	}

	// TODO: дописать метод создания каталога

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) paramInt(c *fiber.Ctx, name string) (id int, err error) {
	if id, err = strconv.Atoi(c.Params(name)); err != nil {
		err = fmt.Errorf("invalid param %s: %w", name, err)
		return
	}
	if id <= 0 {
		err = fmt.Errorf("invalid param %s", name)
	}
	return
}
