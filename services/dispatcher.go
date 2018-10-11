package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ninjadotorg/uncommons-service/utils"
)

// DispatcherService : struct
type DispatcherService struct{}

// SignUp : create new ninja account
func (d DispatcherService) SignUp() (payload string, res bool) {
	endpoint, _ := utils.GetForwardingEndpoint("dispatcher")
	endpoint = fmt.Sprintf("%s/%s", endpoint, "user/sign-up")

	request, _ := http.NewRequest("POST", endpoint, nil)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("call sign-up failed ", err)
		res = false
		return
	}
	b, _ := ioutil.ReadAll(response.Body)

	var data map[string]interface{}
	json.Unmarshal(b, &data)

	status, ok := data["status"]
	message, _ := data["message"]

	if ok && (float64(1) == status) {
		rData := data["data"].(map[string]interface{})
		payload = rData["passpharse"].(string)
		res = true
	} else {
		fmt.Println(message)
		res = false
	}
	return
}
