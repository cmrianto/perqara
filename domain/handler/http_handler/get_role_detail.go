package http_handler

import (
	"log"
	"net/http"
	"perqara/lib/helper"
	"perqara/lib/helper/http_response"
	"perqara/lib/helper/timestamp"

	"perqara/domain/request"

	"github.com/gin-gonic/gin"
)

// GetRoleDetail godoc
// @Summary Get Role Detail
// @Description API to get role detail
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param	id	path	string	true	"Role ID"
// @Success 200 {object} http_response.Response{data=response.Role,statusCode=int,message=string}
// @Failure 400 {object} http_response.Response{code=int,message=string}
// @Failure 500 {object} http_response.Response{code=int,message=string}
// @Router /role/{id} [get]
func (h *handler) GetRoleDetail(c *gin.Context) {
	select {
	case <-c.Request.Context().Done():
		log.Println(http_response.SendErrorAborted(c))
		return
	default:
	}

	in := request.GetSingleIdRequest{
		Id: c.Param("id"),
	}
	authData, _ := helper.GetToken(c, helper.TokenPayload{
		Secret: h.config.Auth.ApiSecret,
	})
	in.AuthData = authData

	result, err := h.usecase.GetRoleDetail(c.Request.Context(), in.ToPayload())
	if err != nil {
		log.Println(http_response.SendError(err, c))
		return
	}

	response := http_response.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Status:     http_response.STANDARD_200_STATUS,
		Timestamp:  timestamp.GetNow(),
		Data:       result,
	}

	c.JSON(http.StatusOK, response)
}
