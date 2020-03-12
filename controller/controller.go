package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ProxeusApp/node-proof-of-existence/service"
	"github.com/labstack/echo"
)

type Controller struct {
	twitterService service.TwitterService
}

func NewController(twitterService service.TwitterService) *Controller {
	return &Controller{twitterService: twitterService}
}

func (me *Controller) GetTweetByURLOrID(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	tweetUrlOrID, ok := response["tweetUrlOrID"].(string)
	if !ok {
		return c.String(http.StatusBadRequest, "parameter not a string")
	}

	tweet, err := me.twitterService.GetTweetByUrlOrId(tweetUrlOrID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, tweet)
}
