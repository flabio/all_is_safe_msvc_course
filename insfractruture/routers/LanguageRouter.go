package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/handler"
)

var (
	handlerLanguage = handler.NewLanguageHandler()
)

func NewLanguageRouter(app *fiber.App) {
	api := app.Group("/api/language")
	//api.Use(middleware.ValidateToken)
	api.Get("/", func(c *fiber.Ctx) error {
		return handlerLanguage.GetLanguageFindAll(c)
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return handlerLanguage.CreateLanguage(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return handlerLanguage.UpdateLanguageById(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return handlerLanguage.DeleteLanguageById(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return handlerLanguage.GetLanguageFindById(c)
	})
}
