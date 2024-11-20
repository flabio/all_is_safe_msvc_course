package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/handler"
	"github.com/safe_msvc_course/insfractruture/middleware"
)

var (
	handlerTypeCourse = handler.NewTypeCourseHandler()
)

func NewTypeCourseRouter(app *fiber.App) {
	api := app.Group("/api/type_course")
	api.Use(middleware.ValidateToken)
	api.Get("/", func(c *fiber.Ctx) error {
		return handlerTypeCourse.GetTypeCourseFindAll(c)
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return handlerTypeCourse.CreateTypeCourse(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return handlerTypeCourse.UpdateTypeCourse(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return handlerTypeCourse.DeleteTypeCourse(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return handlerTypeCourse.GetTypeCourseFindById(c)
	})
}
