package core

import (
	"sync"

	var_db "github.com/flabio/safe_var_db"
	"github.com/safe_msvc_course/insfractruture/database"
	"github.com/safe_msvc_course/insfractruture/entities"
	"github.com/safe_msvc_course/insfractruture/ui/uicore"
)

var (
	_OPEN *OpenConnection
	_ONCE sync.Once
)

func GetTypeCourseInstance() uicore.UITypeCourseCore {
	_ONCE.Do(func() {
		_OPEN = &OpenConnection{
			connection: database.GetDatabaseInstance(),
		}
	})
	return _OPEN
}
func (db *OpenConnection) GetTypeCourseFindAll() ([]entities.TypeCourse, error) {
	var typeCourse []entities.TypeCourse
	db.mux.Lock()
	defer db.mux.Unlock()
	result := db.connection.Order(var_db.DB_ORDER_DESC).Find(&typeCourse)
	return typeCourse, result.Error
}

func (db *OpenConnection) GetTypeCourseFindById(id uint) (entities.TypeCourse, error) {
	var state entities.TypeCourse
	db.mux.Lock()
	defer db.mux.Unlock()
	result := db.connection.Where(var_db.DB_EQUAL_ID, id).First(&state)
	return state, result.Error
}

func (db *OpenConnection) CreateTypeCourse(state entities.TypeCourse) (entities.TypeCourse, error) {
	db.mux.Lock()
	defer db.mux.Unlock()
	err := db.connection.Create(&state).Error
	return state, err
}

func (db *OpenConnection) UpdateTypeCourse(id uint, state entities.TypeCourse) (entities.TypeCourse, error) {
	db.mux.Lock()
	defer db.mux.Unlock()
	err := db.connection.Where(var_db.DB_EQUAL_ID, id).Updates(&state).Error
	return state, err
}
func (db *OpenConnection) DeleteTypeCourse(id uint) (bool, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	result := db.connection.Where(var_db.DB_EQUAL_ID, id).Delete(&entities.TypeCourse{})
	return result.RowsAffected > 0, result.Error
}
func (db *OpenConnection) IsDuplicatedTypeCourseName(id uint, name string) (bool, error) {
	var state entities.TypeCourse
	db.mux.Lock()
	defer db.mux.Unlock()
	query := db.connection.Where(var_db.DB_EQUAL_NAME, name)
	if id > 0 {
		query = query.Where(var_db.DB_DIFF_ID, id)
	}
	result := query.First(&state)
	return result.RowsAffected > 0, result.Error
}
