package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/insfractruture/ui/global"
	"github.com/safe_msvc_course/usecase/service"
)

type topicHandler struct {
	topic global.UITopicGlobal
}

func NewTopicHandler() global.UITopicGlobal {
	return &topicHandler{topic: service.NewTopicService()}
}

func (h *topicHandler) GetTopicFindAll(c *fiber.Ctx) error {
	return h.topic.GetTopicFindAll(c)
}

func (h *topicHandler) GetTopicFindById(c *fiber.Ctx) error {
	return h.topic.GetTopicFindById(c)
}

func (h *topicHandler) CreateTopic(c *fiber.Ctx) error {
	return h.topic.CreateTopic(c)
}

func (h *topicHandler) UpdateTopic(c *fiber.Ctx) error {
	return h.topic.UpdateTopic(c)
}

func (h *topicHandler) DeleteTopic(c *fiber.Ctx) error {
	return h.topic.DeleteTopic(c)
}
