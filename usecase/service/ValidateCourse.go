package service

import (
	"encoding/json"
	"log"

	constants "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/insfractruture/helpers"
	"github.com/safe_msvc_course/usecase/dto"
)

func ValidateCourse(id uint, s *CourseService, c *fiber.Ctx) (dto.CourseDTO, string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("controlando el panic", r)
		}
	}()
	var courseDto dto.CourseDTO
	var msg string = constants.EMPTY
	body := c.Body()

	var dataMap map[string]interface{}
	err := json.Unmarshal([]byte(body), &dataMap)
	if err != nil {
		msg = err.Error()
	}

	msgValid := helpers.ValidateField(dataMap)
	if msgValid != constants.EMPTY {
		return dto.CourseDTO{}, msgValid
	}

	helpers.MapToStruct(dataMap, &courseDto)
	msgReq := helpers.ValidateRequired(courseDto)
	if msgReq != constants.EMPTY {
		return dto.CourseDTO{}, msgReq
	}
	IsNameExist, _ := s.UiCourse.IsDuplicatedCourseName(courseDto.Id, courseDto.Name)
	if IsNameExist {
		msg = constants.NAME_ALREADY_EXIST
	}
	return courseDto, msg
}

func ValidateCourseWitnSchool(id uint, s *CourseService, c *fiber.Ctx) (dto.CourseSchoolDTO, string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("controlando el panic", r)
		}
	}()
	var courseDto dto.CourseSchoolDTO
	var msg string = constants.EMPTY
	body := c.Body()

	var dataMap map[string]interface{}
	err := json.Unmarshal([]byte(body), &dataMap)
	if err != nil {
		msg = err.Error()
	}

	msgValid := helpers.ValidateCourseSchoolField(dataMap)
	if msgValid != constants.EMPTY {
		return dto.CourseSchoolDTO{}, msgValid
	}

	helpers.MapToStruct(dataMap, &courseDto)
	msgReq := helpers.ValidateCourseWithSchoolRequired(courseDto)
	if msgReq != constants.EMPTY {
		return dto.CourseSchoolDTO{}, msgReq
	}
	exist, err := s.UiCourse.GetCourseFindByIdSchoolAndIdCourse(courseDto.SchoolId, courseDto.CourseId)
	log.Println(exist, err)
	if err != nil || exist {
		msg = "The course is already assigned to the school"
	}
	return courseDto, msg
}
