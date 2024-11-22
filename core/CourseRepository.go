package core

import (
	"sync"

	constants "github.com/flabio/safe_constants"
	"github.com/safe_msvc_course/insfractruture/database"
	"github.com/safe_msvc_course/insfractruture/entities"
	"github.com/safe_msvc_course/insfractruture/ui/uicore"
)

func NewCourseRepository() uicore.UICourseCore {
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

func (c *OpenConnection) GetCourseFindAll() ([]entities.Course, error) {
	var courseEntities []entities.Course
	c.mux.Lock()
	defer c.mux.Unlock()
	result := c.connection.Preload("TypeCourse").Preload("Topic").Order(constants.DB_ORDER_DESC).Find(&courseEntities)
	return courseEntities, result.Error
}
func (c *OpenConnection) GetCourseFindById(id uint) (entities.Course, error) {
	var course entities.Course
	c.mux.Lock()
	defer c.mux.Unlock()
	result := c.connection.Where(constants.DB_EQUAL_ID, id).Find(&course)
	return course, result.Error
}

func (c *OpenConnection) CreateCourse(course entities.Course) (entities.Course, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	err := c.connection.Create(&course).Error
	return course, err
}
func (c *OpenConnection) UpdateCourse(id uint, course entities.Course) (entities.Course, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	err := c.connection.Where(constants.DB_EQUAL_ID, id).Updates(&course).Error
	return course, err
}

func (c *OpenConnection) DeleteCourse(id uint) (bool, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	var course entities.Course
	err := c.connection.Where(constants.DB_EQUAL_ID, id).Delete(&course).Error
	return err == nil, err
}

// add course and school
func (c *OpenConnection) AddSchoolToCourse(courseSchool entities.CourseSchool) (entities.CourseSchool, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	err := c.connection.Create(&courseSchool).Error
	return courseSchool, err
}

func (c *OpenConnection) GetCourseFindByIdSchoolAndIdCourse(idschool uint, idcourse uint) (bool, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	query := c.connection.Where("school_id", idschool).Where("course_id", idcourse).Find(&entities.CourseSchool{})
	if query.RowsAffected == 0 {
		return false, query.Error
	}
	return true, query.Error
}
func (c *OpenConnection) GetCourseFindCourseByIdSchool(idschool uint) ([]entities.CourseSchool, error) {
	var courseSchool []entities.CourseSchool
	c.mux.Lock()
	defer c.mux.Unlock()
	query := c.connection.Preload("Course.TypeCourse").Preload("Course.Topic").Where("school_id", idschool).Find(&courseSchool)
	return courseSchool, query.Error
}

// delete course and school
func (c *OpenConnection) DeleteCourseSchool(id uint) (bool, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	var course entities.CourseSchool
	err := c.connection.Where(constants.DB_EQUAL_ID, id).Delete(&course).Error
	return err == nil, err
}
func (c *OpenConnection) GetCourseSchoolFindAll() ([]entities.Course, error) {
	var courseEntities []entities.Course
	c.mux.Lock()
	defer c.mux.Unlock()
	result := c.connection.Preload("CourseSchools").Find(&courseEntities)
	return courseEntities, result.Error
}
func (c *OpenConnection) IsDuplicatedCourseName(id uint, name string) (bool, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	var course entities.Course
	query := c.connection.Where(constants.DB_EQUAL_NAME, name)
	if id > 0 {
		query = query.Where(constants.DB_DIFF_ID, id)
	}
	query = query.Find(&course)
	return query.RowsAffected == 1, query.Error
}
