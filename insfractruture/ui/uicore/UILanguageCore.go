package uicore

import "github.com/safe_msvc_course/insfractruture/entities"

type UILanguage interface {
	GetLanguageFindAll(begin int) ([]entities.Language, int64, error)
	GetLanguageFindById(id uint) (entities.Language, error)
	CreateLanguage(data entities.Language) (entities.Language, error)
	UpdateLanguageById(id uint, data entities.Language) (entities.Language, error)
	DeleteLanguageById(id uint) (bool, error)
	DuplicateLanguageName(id uint, name string) (bool, error)
}
