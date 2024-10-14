package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/insfractruture/ui/global"
	"github.com/safe_msvc_course/usecase/service"
)

type courseHandler struct {
	course global.UICourseGlobal
}

func NewCourseHandler() global.UICourseGlobal {
	return &courseHandler{course: service.NewCourseService()}
}

func (h *courseHandler) GetCourseFindAll(c *fiber.Ctx) error {
	return h.course.GetCourseFindAll(c)
}

func (h *courseHandler) GetCourseFindById(c *fiber.Ctx) error {
	return h.course.GetCourseFindById(c)
}

func (h *courseHandler) CreateCourse(c *fiber.Ctx) error {
	return h.course.CreateCourse(c)
}

func (h *courseHandler) UpdateCourse(c *fiber.Ctx) error {
	return h.course.UpdateCourse(c)
}

func (h *courseHandler) DeleteCourse(c *fiber.Ctx) error {
	return h.course.DeleteCourse(c)
}
