package http_handler

import (
	"log"
	"net/http"
	"perqara/lib/helper"
	"perqara/lib/helper/http_response"
	"perqara/lib/helper/timestamp"

	"perqara/domain/request"
	"perqara/domain/response"

	"github.com/gin-gonic/gin"
)

// CreateClient godoc
// @Summary Create Client
// @Description API to create client
// @Tags Client
// @Accept json
// @Produce json
// @Param	client	body	request.CreateClientRequest	true	"Client Data"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} http_response.Response{data=response.SuccessMessageResponse,statusCode=int,message=string}
// @Failure 400 {object} http_response.Response{code=int,message=string}
// @Failure 500 {object} http_response.Response{code=int,message=string}
// @Router /client [post]
func (h *handler) CreateClient(c *gin.Context) {
	select {
	case <-c.Request.Context().Done():
		log.Println(http_response.SendErrorAborted(c))
		return
	default:
	}

	resp := response.SuccessMessageResponse{
		SuccessMessage: "failed",
	}

	in := request.CreateClientRequest{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		log.Println(http_response.SendError(err, c))
		return
	}
	authData, _ := helper.GetToken(c, helper.TokenPayload{
		Secret: h.config.Auth.ApiSecret,
	})
	in.AuthData = authData

	err = h.usecase.CreateClient(c.Request.Context(), &in)
	if err != nil {
		log.Println(http_response.SendError(err, c))
		return
	}

	resp.SuccessMessage = "success"
	response := http_response.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Status:     http_response.STANDARD_200_STATUS,
		Timestamp:  timestamp.GetNow(),
		Data:       resp,
	}

	c.JSON(http.StatusOK, response)
}
