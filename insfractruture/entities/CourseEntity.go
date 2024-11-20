package entities

import "time"

type Course struct {
	Id            uint           `gorm:"primary_key:auto_increment" json:"id"`
	Name          string         `gorm:"type:varchar(150);not null" json:"name"`
	Active        bool           `gorm:"type:boolean" json:"active"`
	TypeCourseId  uint           `gorm:"null" json:"course_id"`
	TypeCourse    TypeCourse     `gorm:"foreignkey:TypeCourseId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"type_course"`
	CreatedAt     time.Time      `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt     *time.Time     `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
	CourseSchools []CourseSchool `gorm:"foreignkey:CourseId" json:"course_schools"`
}
