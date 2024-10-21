package global

import "github.com/gofiber/fiber/v2"

type UILanguageGlobal interface {
	GetLanguageFindAll(c *fiber.Ctx) error
	GetLanguageFindById(c *fiber.Ctx) error
	CreateLanguage(c *fiber.Ctx) error
	UpdateLanguageById(c *fiber.Ctx) error
	DeleteLanguageById(c *fiber.Ctx) error
}
