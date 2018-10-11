package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ninjadotorg/uncommons-service/utils"
)

// DispatcherService : struct
type DispatcherService struct{}

// SignUp : create new ninja account
func (d DispatcherService) SignUp() {
	endpoint, found := utils.GetForwardingEndpoint("dispatcher")
	log.Println(endpoint, found)
	endpoint = fmt.Sprintf("%s/%s", endpoint, "user/sign-up")

	request, _ := http.NewRequest("POST", endpoint, nil)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("call sign-up failed ", err)
	} else {
		log.Println(response)
		log.Println("call sign-up success")
	}
}
