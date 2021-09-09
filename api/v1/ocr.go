package v1

import (
	"api_server/pkg/httplib"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GeneralBasic(c *gin.Context) {
	accessTokenUrl := "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id="
	req := httplib.Get(accessTokenUrl)
	req.Param("grant_type", "client_credentials")
	req.Param("client_id", "4PUsW94eWt1HzAc9WkmEEweq")
	req.Param("client_secret", "7xYN6kYltwtZHMNdCw8NT4vGQbwpF4wR")

	str, err := req.String()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(str)
}
