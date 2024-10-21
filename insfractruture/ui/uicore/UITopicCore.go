package uicore

import (
	"github.com/safe_msvc_course/insfractruture/entities"
)

type UITopicCore interface {
	GetTopicFindAll() ([]entities.Topic, error)
	GetTopicByCoursoIdFindAll(courseId uint) ([]entities.Topic, error)
	GetTopicFindById(id uint) (entities.Topic, error)
	CreateTopic(topic entities.Topic) (entities.Topic, error)
	UpdateTopic(id uint, topic entities.Topic) (entities.Topic, error)
	DeleteTopic(id uint) (bool, error)
}
