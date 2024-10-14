package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/insfractruture/ui/global"
	"github.com/safe_msvc_course/usecase/service"
)

type languageHanlder struct {
	language global.UILanguageGlobal
}

func NewLanguageHandler() global.UILanguageGlobal {
	return &languageHanlder{language: service.NewLanguageService()}
}

// CreateLanguage implements global.UILanguageGlobal.
func (l *languageHanlder) CreateLanguage(c *fiber.Ctx) error {
	return l.language.CreateLanguage(c)
}

// DeleteLanguageById implements global.UILanguageGlobal.
func (l *languageHanlder) DeleteLanguageById(c *fiber.Ctx) error {
	return l.language.DeleteLanguageById(c)
}

// GetLanguageFindAll implements global.UILanguageGlobal.
func (l *languageHanlder) GetLanguageFindAll(c *fiber.Ctx) error {
	return l.language.GetLanguageFindAll(c)
}

// GetLanguageFindById implements global.UILanguageGlobal.
func (l *languageHanlder) GetLanguageFindById(c *fiber.Ctx) error {
	return l.language.GetLanguageFindById(c)
}

// UpdateLanguageById implements global.UILanguageGlobal.
func (l *languageHanlder) UpdateLanguageById(c *fiber.Ctx) error {
	return l.language.UpdateLanguageById(c)
}
