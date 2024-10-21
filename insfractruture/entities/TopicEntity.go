package entities

import "time"

type Topic struct {
	Id        uint       `gorm:"primary_key:auto_increment" json:"id"`
	Title     string     `gorm:"type:text;not null" json:"title"`
	TimeHours string     `gorom:"type:varchar(50);not null" json:"time_hours"`
	CourseId  uint       `gorm:"null" json:"course_id"`
	Course    Course     `gorm:"foreignkey:CourseId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"course"`
	Active    bool       `gorm:"type:boolean" json:"active"`
	CreatedAt time.Time  `gorm:"<-:created_at"`
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)"`
	// Relationships

}
