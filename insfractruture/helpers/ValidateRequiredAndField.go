package helpers

import (
	constants "github.com/flabio/safe_constants"
	"github.com/safe_msvc_course/usecase/dto"
)

func ValidateRequired(course dto.CourseDTO) string {
	var msg string = constants.EMPTY
	if course.Name == constants.EMPTY {
		msg = constants.NAME_IS_REQUIRED
	}

	return msg
}

func ValidateField(value map[string]interface{}) string {
	var msg string = constants.EMPTY
	if value[constants.NAME] == nil {
		msg = constants.NAME_FIELD_IS_REQUIRED
	}
	return msg
}

func ValidateRequiredTopic(topic dto.TopicDTO) string {
	var msg string = constants.EMPTY
	if topic.Title == constants.EMPTY {
		msg = constants.TITLE_IS_REQUIRED
	}
	if topic.CourseId == 0 {
		msg = constants.COURSE_ID_IS_REQUIRED
	}
	if topic.TimeHours == constants.EMPTY {
		msg = constants.TIME_HOURS_IS_REQUIRED
	}
	return msg
}

func ValidateFieldTopic(value map[string]interface{}) string {
	var msg string = constants.EMPTY
	if value[constants.TITLE] == nil {
		msg = constants.TITLE_FIELD_IS_REQUIRED
	}
	if value[constants.TIME_HOURS] == nil {
		msg = constants.TIME_HOURS_IS_FIELD_REQUIRED
	}
	if value[constants.COURSE_ID] == nil {
		msg = constants.COURSE_ID_IS_FIELD_REQUIRED
	}

	return msg
}

func ValidateCourseSchoolField(value map[string]interface{}) string {
	var msg string = constants.EMPTY
	if value[constants.COURSE_ID] == nil {
		msg = constants.COURSE_ID_IS_FIELD_REQUIRED
	}
	if value[constants.SCHOOL_ID] == nil {
		msg = constants.SCHOOL_ID_IS_FIELD_REQUIRED
	}
	return msg
}

func ValidateCourseWithSchoolRequired(course dto.CourseSchoolDTO) string {
	var msg string = constants.EMPTY
	if course.CourseId == 0 {
		msg = constants.COURSE_ID_IS_REQUIRED
	}
	if course.SchoolId == 0 {
		msg = constants.SCHOOL_ID_IS_REQUIRED
	}
	return msg
}
