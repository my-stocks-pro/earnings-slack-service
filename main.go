package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Data struct {
	Location    string `json:"location"`
	IdI         string `json:"id_i"`
	Description string `json:"description"`
	ImageURL    string `json:"image_source"`
	ImageLink   string `json:"image_link"`
}

func main() {
	router := gin.Default()

	slackService := NewService()

	router.GET("/slack/earnings", func(c *gin.Context) {

		var data Data
		body := c.Request.Body
		b, e := ioutil.ReadAll(body)
		if e != nil {
			fmt.Println(e)
		}

		err := json.Unmarshal(b, &data)
		if err != nil {
			fmt.Println(err)
		}

		c.String(200, "Success")

		slackService.PostMessage(data)

	})

	router.Run(fmt.Sprintf("https://%s:%s", slackService.Port, slackService.Host))
}
