package core

import (
	"sync"

	constants "github.com/flabio/safe_constants"
	"github.com/safe_msvc_course/insfractruture/database"
	"github.com/safe_msvc_course/insfractruture/entities"
	"github.com/safe_msvc_course/insfractruture/ui/uicore"
)

func NewTopicRepository() uicore.UITopicCore {
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
func (c *OpenConnection) GetTopicFindAll() ([]entities.Topic, error) {
	var topicEntities []entities.Topic
	c.mux.Lock()
	defer c.mux.Unlock()
	result := c.connection.Preload("Course").Order(constants.DB_ORDER_DESC).Find(&topicEntities)
	return topicEntities, result.Error
}
func (c *OpenConnection) GetTopicByCoursoIdFindAll(courseId uint) ([]entities.Topic, error) {
	var topicEntities []entities.Topic
	c.mux.Lock()
	defer c.mux.Unlock()
	result := c.connection.Order(constants.DB_ORDER_DESC).Where("course_id=?", courseId).Find(&topicEntities)
	return topicEntities, result.Error
}
func (c *OpenConnection) GetTopicFindById(id uint) (entities.Topic, error) {
	var topic entities.Topic
	c.mux.Lock()
	defer c.mux.Unlock()
	result := c.connection.Where(constants.DB_EQUAL_ID, id).Find(&topic)
	return topic, result.Error
}
func (c *OpenConnection) CreateTopic(topic entities.Topic) (entities.Topic, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	err := c.connection.Create(&topic).Error
	return topic, err
}
func (c *OpenConnection) UpdateTopic(id uint, topic entities.Topic) (entities.Topic, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	err := c.connection.Where(constants.DB_EQUAL_ID, id).Updates(&topic).Error
	return topic, err
}
func (c *OpenConnection) DeleteTopic(id uint) (bool, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	var topic entities.Topic
	err := c.connection.Where(constants.DB_EQUAL_ID, id).Delete(&topic).Error
	return err == nil, err
}
