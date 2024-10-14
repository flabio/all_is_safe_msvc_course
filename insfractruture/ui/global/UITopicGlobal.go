package global

import "github.com/gofiber/fiber/v2"

type UITopicGlobal interface {
	GetTopicFindAll(c *fiber.Ctx) error
	GetTopicFindById(c *fiber.Ctx) error
	CreateTopic(c *fiber.Ctx) error
	UpdateTopic(c *fiber.Ctx) error
	DeleteTopic(c *fiber.Ctx) error
}
