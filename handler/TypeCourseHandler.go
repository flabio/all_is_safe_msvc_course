package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/insfractruture/ui/global"
	"github.com/safe_msvc_course/usecase/service"
)

type TypeCourseHandler struct {
	typeCourse global.UITypeCourseGlobal
}

func NewTypeCourseHandler() global.UITypeCourseGlobal {
	return &TypeCourseHandler{typeCourse: service.NewTypeCourseService()}
}
func (h *TypeCourseHandler) GetTypeCourseFindAll(c *fiber.Ctx) error {
	return h.typeCourse.GetTypeCourseFindAll(c)
}
func (h *TypeCourseHandler) GetTypeCourseFindById(c *fiber.Ctx) error {
	return h.typeCourse.GetTypeCourseFindById(c)
}
func (h *TypeCourseHandler) CreateTypeCourse(c *fiber.Ctx) error {
	return h.typeCourse.CreateTypeCourse(c)
}
func (h *TypeCourseHandler) UpdateTypeCourse(c *fiber.Ctx) error {
	return h.typeCourse.UpdateTypeCourse(c)
}
func (h *TypeCourseHandler) DeleteTypeCourse(c *fiber.Ctx) error {
	return h.typeCourse.DeleteTypeCourse(c)
}
