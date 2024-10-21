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

type CourseService struct {
	UiCourse uicore.UICourseCore
}

func NewCourseService() global.UICourseGlobal {
	return &CourseService{UiCourse: core.NewCourseRepository()}
}

func (s *CourseService) GetCourseFindAll(c *fiber.Ctx) error {
	results, err := s.UiCourse.GetCourseFindAll()
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

func (s *CourseService) GetCourseFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.UiCourse.GetCourseFindById(uint(id))

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
func (s *CourseService) CreateCourse(c *fiber.Ctx) error {
	var courseCreate entities.Course

	dataDto, msgError := ValidateCourse(0, s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(dataDto).To(&courseCreate)
	result, err := s.UiCourse.CreateCourse(courseCreate)
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

func (s *CourseService) UpdateCourse(c *fiber.Ctx) error {
	var updatedCourse entities.Course
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.UiCourse.GetCourseFindById(uint(id))
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

	courseDto, msgError := ValidateCourse(uint(id), s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(courseDto).To(&updatedCourse)
	user, err := s.UiCourse.UpdateCourse(uint(id), updatedCourse)
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
func (s *CourseService) DeleteCourse(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	courseFindById, err := s.UiCourse.GetCourseFindById(uint(id))
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
	result, err := s.UiCourse.DeleteCourse(uint(id))
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

// course with school
func (s *CourseService) GetCourseSchoolFindAll(c *fiber.Ctx) error {
	results, err := s.UiCourse.GetCourseSchoolFindAll()
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
func (s *CourseService) AddSchoolToCourse(c *fiber.Ctx) error {
	var courseCreate entities.CourseSchool

	dataDto, msgError := ValidateCourseWitnSchool(0, s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(dataDto).To(&courseCreate)
	result, err := s.UiCourse.AddSchoolToCourse(courseCreate)
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

func (s *CourseService) DeleteCourseSchool(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))

	result, err := s.UiCourse.DeleteCourseSchool(uint(id))
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

//AddSchoolToCourse
