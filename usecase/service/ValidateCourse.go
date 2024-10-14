package service

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/insfractruture/helpers"
	"github.com/safe_msvc_course/insfractruture/utils"
	"github.com/safe_msvc_course/usecase/dto"
)

func ValidateCourse(id uint, s *CourseService, c *fiber.Ctx) (dto.CourseDTO, string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("controlando el panic", r)
		}
	}()
	var courseDto dto.CourseDTO
	var msg string = utils.EMPTY
	body := c.Body()

	var dataMap map[string]interface{}
	err := json.Unmarshal([]byte(body), &dataMap)
	if err != nil {
		msg = err.Error()
	}

	msgValid := helpers.ValidateField(dataMap)
	if msgValid != utils.EMPTY {
		return dto.CourseDTO{}, msgValid
	}

	helpers.MapToStruct(dataMap, &courseDto)

	msgReq := helpers.ValidateRequired(courseDto)
	if msgReq != utils.EMPTY {
		return dto.CourseDTO{}, msgReq
	}
	IsNameExist, _ := s.UiCourse.IsDuplicatedCourseName(courseDto.Id, courseDto.Name)
	if IsNameExist {
		msg = utils.NAME_ALREADY_EXIST
	}

	return courseDto, msg
}
