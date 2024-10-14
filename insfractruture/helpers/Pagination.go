package helpers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/safe_msvc_course/insfractruture/utils"
)

func Pagination(c *fiber.Ctx) (int, int) {

	pageParam := c.Query(utils.PAGE)
	if pageParam == "" {
		return 1, 0
	}
	page, _ := strconv.Atoi(c.Query(utils.PAGE))
	if page < 1 {
		return 1, 0
	}
	begin := (utils.LiMIT * page) - utils.LiMIT
	return page, begin
}
