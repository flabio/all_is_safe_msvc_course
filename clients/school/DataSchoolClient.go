package school

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/safe_msvc_course/clients"
)

func DataSchoolClient(data *http.Request, err error) (SchoolClient, string) {
	var schoolClient SchoolClient
	var dataMessage clients.MessageClient
	var msg string
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()
	if err != nil {
		msg = err.Error()
	}

	clientHttp := &http.Client{}
	resp, err := clientHttp.Do(data)
	if err != nil {
		msg = err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		msg = err.Error()
	}
	err = json.Unmarshal(body, &schoolClient)
	if err != nil {
		msg = err.Error()
	}
	if schoolClient.Id == 0 {
		json.Unmarshal(body, &dataMessage)
		msg = dataMessage.Message
	}

	return schoolClient, msg
}
