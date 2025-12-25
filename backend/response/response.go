package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Status uint        `json:"status"`
}

func SuccessWithData(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Msg:    msg,
		Data:   data,
		Status: 200,
	})
}

func SuccessWithNoData(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Msg:    msg,
		Data:   nil,
		Status: 200,
	})
}

func FailBadRequest(c *gin.Context, msg string, err error) {
	c.JSON(http.StatusBadRequest, Response{
		Data:   nil,
		Msg:    msg,
		Status: 400,
	})
}
