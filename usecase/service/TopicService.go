package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	utils "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/core"
	"github.com/safe_msvc_course/insfractruture/entities"
	"github.com/safe_msvc_course/insfractruture/helpers"
	"github.com/safe_msvc_course/insfractruture/ui/global"
	"github.com/safe_msvc_course/insfractruture/ui/uicore"
	"github.com/safe_msvc_course/usecase/dto"

	"github.com/ulule/deepcopier"
)

type TopicService struct {
	UiTopic uicore.UITopicCore
}

func NewTopicService() global.UITopicGlobal {
	return &TopicService{UiTopic: core.NewTopicRepository()}
}

func (s *TopicService) GetTopicFindAll(c *fiber.Ctx) error {
	results, err := s.UiTopic.GetTopicFindAll()
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
func (s *TopicService) GetTopicByCoursoIdFindAll(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	results, err := s.UiTopic.GetTopicByCoursoIdFindAll(uint(id))

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
func (s *TopicService) GetTopicFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.UiTopic.GetTopicFindById(uint(id))
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
func (s *TopicService) CreateTopic(c *fiber.Ctx) error {
	var courseCreate entities.Topic

	dataDto, msgError := ValidateTopic(0, s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(dataDto).To(&courseCreate)
	result, err := s.UiTopic.CreateTopic(courseCreate)
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

func (s *TopicService) UpdateTopic(c *fiber.Ctx) error {
	var updatedCourse entities.Topic
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.UiTopic.GetTopicFindById(uint(id))
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

	courseDto, msgError := ValidateTopic(uint(id), s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(courseDto).To(&updatedCourse)
	data, err := s.UiTopic.UpdateTopic(uint(id), updatedCourse)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_UPDATE,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.UPDATED,
		utils.DATA:    data,
	})
}
func (s *TopicService) DeleteTopic(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	courseFindById, err := s.UiTopic.GetTopicFindById(uint(id))
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
	result, err := s.UiTopic.DeleteTopic(uint(id))
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

func ValidateTopic(id uint, s *TopicService, c *fiber.Ctx) (dto.TopicDTO, string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("controlando el panic", r)
		}
	}()
	var topicDto dto.TopicDTO
	var msg string = utils.EMPTY
	body := c.Body()

	var dataMap map[string]interface{}
	err := json.Unmarshal([]byte(body), &dataMap)

	if err != nil {
		msg = err.Error()
	}

	msgValid := helpers.ValidateFieldTopic(dataMap)
	if msgValid != utils.EMPTY {
		return dto.TopicDTO{}, msgValid
	}

	helpers.MapToStructTopic(&topicDto, dataMap)

	msgReq := helpers.ValidateRequiredTopic(topicDto)
	if msgReq != utils.EMPTY {
		return dto.TopicDTO{}, msgReq
	}
	return topicDto, msg
}
