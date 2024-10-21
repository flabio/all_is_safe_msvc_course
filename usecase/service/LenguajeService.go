package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/core"
	"github.com/safe_msvc_course/insfractruture/entities"
	"github.com/safe_msvc_course/insfractruture/helpers"
	"github.com/safe_msvc_course/insfractruture/ui/global"
	"github.com/safe_msvc_course/insfractruture/ui/uicore"
	"github.com/safe_msvc_course/insfractruture/utils"

	"github.com/safe_msvc_course/usecase/dto"
	"github.com/ulule/deepcopier"
)

type LanguageService struct {
	Uilenguage uicore.UILanguage
}

func NewLanguageService() global.UILanguageGlobal {
	return &LanguageService{Uilenguage: core.NewLanguageRepository()}
}

func (s *LanguageService) GetLanguageFindAll(c *fiber.Ctx) error {
	page, begin := helpers.Pagination(c)
	results, total, err := s.Uilenguage.GetLanguageFindAll(begin)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:      http.StatusOK,
		utils.DATA:        results,
		utils.TOTAL_COUNT: total,
		utils.PAGE_COUNT:  page,
		utils.BEGIN:       begin,
	})
}
func (s *LanguageService) GetLanguageFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	resultLenguageById, err := s.Uilenguage.GetLanguageFindById(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	if resultLenguageById.Id == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusBadRequest,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   resultLenguageById,
	})
}
func (s *LanguageService) CreateLanguage(c *fiber.Ctx) error {
	var lenguageCreate entities.Language
	dataDto, msgError := ValidateLanguage(0, s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(dataDto).To(&lenguageCreate)
	result, err := s.Uilenguage.CreateLanguage(lenguageCreate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_CREATE,
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS:  http.StatusCreated,
		utils.MESSAGE: utils.CREATED,
		utils.DATA:    result,
	})
}

// UpdateLenguageById implements global.UILenguageGlobal.
func (s *LanguageService) UpdateLanguageById(c *fiber.Ctx) error {
	var updateLanguage entities.Language
	id, _ := strconv.Atoi(c.Params(utils.ID))
	languageExistById, err := s.Uilenguage.GetLanguageFindById(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ID_NO_EXIST,
			utils.DATA:    utils.ERROR_QUERY,
		})
	}
	if languageExistById.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}

	languageDto, errMsg := ValidateLanguage(uint(id), s, c)
	if errMsg != utils.EMPTY {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: errMsg,
		})
	}
	deepcopier.Copy(languageDto).To(&updateLanguage)
	data, err := s.Uilenguage.UpdateLanguageById(uint(languageExistById.Id), updateLanguage)
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

// DeleteLenguageById implements global.UILenguageGlobal.
func (s *LanguageService) DeleteLanguageById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	languageExistById, err := s.Uilenguage.GetLanguageFindById(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ID_NO_EXIST,
			utils.DATA:    utils.ERROR_QUERY,
		})
	}
	if languageExistById.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	data, err := s.Uilenguage.DeleteLanguageById(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_DELETE,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.REMOVED,
		utils.DATA:    data,
	})
}
func ValidateLanguage(id uint, s *LanguageService, c *fiber.Ctx) (dto.LanguageDTO, string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("controlando el panic", r)
		}
	}()
	var lenguageDto dto.LanguageDTO
	var msg string = utils.EMPTY
	body := c.Body()

	var dataMap map[string]interface{}

	err := json.Unmarshal([]byte(body), &dataMap)
	if err != nil {
		msg = err.Error()
	}

	msg = helpers.ValidateField(dataMap)
	if msg != utils.EMPTY {
		return dto.LanguageDTO{}, msg
	}
	helpers.MapToStruct(dataMap, &lenguageDto)
	if msg != utils.EMPTY {
		return dto.LanguageDTO{}, msg
	}
	nameReadyExist, _ := s.Uilenguage.DuplicateLanguageName(uint(lenguageDto.Id), lenguageDto.Name)
	log.Println(nameReadyExist)
	if nameReadyExist {
		return lenguageDto, utils.NAME_ALREADY_EXIST
	}
	return lenguageDto, msg
}
