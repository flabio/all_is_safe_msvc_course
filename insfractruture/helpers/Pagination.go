package helpers

import (
	"strconv"

	constants "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
)

func Pagination(c *fiber.Ctx) (int, int) {

	pageParam := c.Query(constants.PAGE)
	if pageParam == "" {
		return 1, 0
	}
	page, _ := strconv.Atoi(c.Query(constants.PAGE))
	if page < 1 {
		return 1, 0
	}
	begin := (constants.LiMIT * page) - constants.LiMIT
	return page, begin
}
