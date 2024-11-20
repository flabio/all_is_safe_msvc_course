package global

import "github.com/gofiber/fiber/v2"

type UITypeCourseGlobal interface {
	GetTypeCourseFindAll(c *fiber.Ctx) error
	GetTypeCourseFindById(c *fiber.Ctx) error
	CreateTypeCourse(c *fiber.Ctx) error
	UpdateTypeCourse(c *fiber.Ctx) error
	DeleteTypeCourse(c *fiber.Ctx) error
}
