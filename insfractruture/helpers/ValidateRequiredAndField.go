package helpers

import (
	"github.com/safe_msvc_course/insfractruture/utils"
	"github.com/safe_msvc_course/usecase/dto"
)

func ValidateRequired(course dto.CourseDTO) string {
	var msg string = utils.EMPTY
	if course.Name == utils.EMPTY {
		msg = utils.NAME_IS_REQUIRED
	}

	return msg
}

func ValidateField(value map[string]interface{}) string {
	var msg string = utils.EMPTY
	if value[utils.NAME] == nil {
		msg = utils.NAME_FIELD_IS_REQUIRED
	}
	return msg
}

func ValidateRequiredTopic(topic dto.TopicDTO) string {
	var msg string = utils.EMPTY
	if topic.Title == utils.EMPTY {
		msg = utils.TITLE_IS_REQUIRED
	}
	if topic.CourseId == 0 {
		msg = utils.COURSE_ID_IS_REQUIRED
	}
	if topic.TimeHours == utils.EMPTY {
		msg = utils.TIME_HOURS_IS_REQUIRED
	}
	return msg
}

func ValidateFieldTopic(value map[string]interface{}) string {
	var msg string = utils.EMPTY
	if value[utils.TITLE] == nil {
		msg = utils.TITLE_FIELD_IS_REQUIRED
	}
	if value[utils.TIME_HOURS] == nil {
		msg = utils.TIME_HOURS_IS_FIELD_REQUIRED
	}
	if value[utils.COURSE_ID] == nil {
		msg = utils.COURSE_ID_IS_FIELD_REQUIRED
	}

	return msg
}

func ValidateCourseSchoolField(value map[string]interface{}) string {
	var msg string = utils.EMPTY
	if value[utils.COURSE_ID] == nil {
		msg = utils.COURSE_ID_IS_FIELD_REQUIRED
	}
	if value[utils.SCHOOL_ID] == nil {
		msg = utils.SCHOOL_ID_IS_FIELD_REQUIRED
	}
	return msg
}

func ValidateCourseWithSchoolRequired(course dto.CourseSchoolDTO) string {
	var msg string = utils.EMPTY
	if course.CourseId == 0 {
		msg = utils.COURSE_ID_IS_REQUIRED
	}
	if course.SchoolId == 0 {
		msg = utils.SCHOOL_ID_IS_REQUIRED
	}
	return msg
}
