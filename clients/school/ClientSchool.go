package school

import (
	"net/http"
	"strconv"

	constants "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
)

func MsvcSchoolFindId(id uint, c *fiber.Ctx) (SchoolClient, string) {
	data, err := http.NewRequest(constants.GET, constants.MSVC_SCHOOL_URL+"/"+strconv.Itoa(int(id)), nil)
	if err != nil {
		return SchoolClient{}, err.Error()
	}
	data.Header.Set(constants.AUTHORIZATION, c.Get(constants.AUTHORIZATION))
	school, msg := DataSchoolClient(data, err)
	return school, msg
}

// TODO: Implement other MSVC School API calls here. For example, GetSchoolFindByProviderNumber, etc.
// For each new API call, create a new function with the appropriate method (GET, POST, PUT, DELETE) and parameters.
// Make sure to handle the response data and any potential errors appropriately.
// Return the requested data or an error message.
// For example:
// func MsvcSchoolFindByProviderNumber(id uint, providerNumber string)(School, string){
//     data,err:=http.NewRequest(constants.GET, constants.MSVC_SCHOOL_URL+"/"+strconv
