package entities

import "time"

type CourseSchool struct {
	Id        uint       `gorm:"primary_key:auto_increment" json:"id"`
	SchoolId  uint       `gorm:"type:int" json:"school_id"`
	CourseId  uint       `gorm:"null" json:"course_id"`
	Course    Course     `gorm:"foreignkey:CourseId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"course"`
	Active    bool       `gorm:"type:boolean" json:"active"`
	CreatedAt time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
}
