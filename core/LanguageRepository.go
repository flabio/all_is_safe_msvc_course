package core

import (
	"sync"

	constants "github.com/flabio/safe_constants"
	"github.com/safe_msvc_course/insfractruture/database"
	"github.com/safe_msvc_course/insfractruture/entities"
	"github.com/safe_msvc_course/insfractruture/ui/uicore"
)

func NewLanguageRepository() uicore.UILanguage {

	var (
		_OPEN *OpenConnection
		_ONCE sync.Once
	)
	_ONCE.Do(func() {
		_OPEN = &OpenConnection{
			connection: database.GetDatabaseInstance(),
		}
	})
	return _OPEN
}
func (c *OpenConnection) GetLanguageFindAll(begin int) ([]entities.Language, int64, error) {
	var languageEntities []entities.Language
	var countLanguage int64
	c.mux.Lock()
	defer c.mux.Unlock()
	result := c.connection.Offset(begin).Limit(5).Order(constants.DB_ORDER_DESC).Find(&languageEntities)
	c.connection.Table("languages").Count(&countLanguage)
	return languageEntities, countLanguage, result.Error
}
func (c *OpenConnection) GetLanguageFindById(id uint) (entities.Language, error) {
	var language entities.Language
	c.mux.Lock()
	defer c.mux.Unlock()
	result := c.connection.Where(constants.DB_EQUAL_ID, id).Find(&language)
	return language, result.Error
}
func (c *OpenConnection) CreateLanguage(language entities.Language) (entities.Language, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	err := c.connection.Create(&language).Error
	return language, err
}
func (c *OpenConnection) UpdateLanguageById(id uint, language entities.Language) (entities.Language, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	err := c.connection.Where(constants.DB_EQUAL_ID, id).Updates(&language).Error
	return language, err
}
func (c *OpenConnection) DeleteLanguageById(id uint) (bool, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	var language entities.Language
	err := c.connection.Where(constants.DB_EQUAL_ID, id).Delete(&language).Error
	return err == nil, err
}
func (c *OpenConnection) DuplicateLanguageName(id uint, name string) (bool, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	var language entities.Language
	query := c.connection.Where(constants.DB_EQUAL_NAME, name)
	if id > 0 {
		query = query.Where(constants.DB_DIFF_ID, id)
	}
	query = query.Find(&language)
	if query.RowsAffected == 1 {
		return true, query.Error
	}
	return false, query.Error
}
