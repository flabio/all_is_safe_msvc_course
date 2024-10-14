package dto

type CourseDTO struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	SchoolId uint   `json:"school_id"`
	Active   bool   `json:"active"`
}

type CourseResponseDTO struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	SchoolId   uint   `json:"school_id"`
	SchoolName string `json:"school_name"`
	Active     bool   `json:"active"`
}
