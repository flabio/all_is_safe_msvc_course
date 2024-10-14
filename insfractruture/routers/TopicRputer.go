package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/handler"
)

var (
	handlerTopic = handler.NewTopicHandler()
)

func NewTopicRouter(app *fiber.App) {
	api := app.Group("/api/topic")
	//api.Use(middleware.ValidateToken)
	api.Get("/", func(c *fiber.Ctx) error {
		return handlerTopic.GetTopicFindAll(c)
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return handlerTopic.CreateTopic(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return handlerTopic.UpdateTopic(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return handlerTopic.DeleteTopic(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return handlerTopic.GetTopicFindById(c)
	})
}
