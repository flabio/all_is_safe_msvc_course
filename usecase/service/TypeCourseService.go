package service

import (
	"net/http"
	"strconv"

	utils "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/core"
	"github.com/safe_msvc_course/insfractruture/entities"
	"github.com/safe_msvc_course/insfractruture/ui/global"
	"github.com/safe_msvc_course/insfractruture/ui/uicore"

	"github.com/ulule/deepcopier"
)

type TypeCourseService struct {
	UiTypeCourse uicore.UITypeCourseCore
}

func NewTypeCourseService() global.UITypeCourseGlobal {
	return &TypeCourseService{UiTypeCourse: core.GetTypeCourseInstance()}
}
func (s *TypeCourseService) GetTypeCourseFindAll(c *fiber.Ctx) error {
	results, err := s.UiTypeCourse.GetTypeCourseFindAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS: fiber.StatusOK,
		utils.DATA:   results,
	})
}
func (s *TypeCourseService) GetTypeCourseFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.UiTypeCourse.GetTypeCourseFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	if result.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS: fiber.StatusOK,
		utils.DATA:   result,
	})
}
func (s *TypeCourseService) CreateTypeCourse(c *fiber.Ctx) error {
	var courseCreate entities.TypeCourse
	dataDto, msgError := ValidateTypeCourse(0, s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(dataDto).To(&courseCreate)
	result, err := s.UiTypeCourse.CreateTypeCourse(courseCreate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_CREATE,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.CREATED,
		utils.DATA:    result,
	})
}
func (s *TypeCourseService) UpdateTypeCourse(c *fiber.Ctx) error {
	var updatedCourse entities.TypeCourse
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.UiTypeCourse.GetTypeCourseFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	if result.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	courseDto, msgError := ValidateTypeCourse(uint(id), s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(courseDto).To(&updatedCourse)
	user, err := s.UiTypeCourse.UpdateTypeCourse(uint(id), updatedCourse)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_UPDATE,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.UPDATED,
		utils.DATA:    user,
	})
}
func (s *TypeCourseService) DeleteTypeCourse(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	courseFindById, err := s.UiTypeCourse.GetTypeCourseFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	if courseFindById.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	result, err := s.UiTypeCourse.DeleteTypeCourse(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_DELETE,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.REMOVED,
		utils.DATA:    result,
	})
}
