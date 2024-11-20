package uicore

import (
	"github.com/safe_msvc_course/insfractruture/entities"
)

type UITypeCourseCore interface {
	GetTypeCourseFindAll() ([]entities.TypeCourse, error)
	GetTypeCourseFindById(id uint) (entities.TypeCourse, error)
	IsDuplicatedTypeCourseName(id uint, name string) (bool, error)
	CreateTypeCourse(course entities.TypeCourse) (entities.TypeCourse, error)
	UpdateTypeCourse(id uint, course entities.TypeCourse) (entities.TypeCourse, error)
	DeleteTypeCourse(id uint) (bool, error)
}
