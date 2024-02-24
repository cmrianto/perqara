package http_handler

import (
	"log"
	"net/http"
	"perqara/lib/helper/http_response"
	"perqara/lib/helper/timestamp"

	"perqara/domain/request"
	"perqara/domain/response"

	"github.com/gin-gonic/gin"
)

// UpdateRole godoc
// @Summary Update Role
// @Description API to update role
// @Tags Auth
// @Accept json
// @Produce json
// @Param	role	body	request.UpdateRoleRequest	true	"Role Data"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} http_response.Response{data=response.SuccessMessageResponse,statusCode=int,message=string}
// @Failure 400 {object} http_response.Response{code=int,message=string}
// @Failure 500 {object} http_response.Response{code=int,message=string}
// @Router /role [put]
func (h *handler) UpdateRole(c *gin.Context) {
	select {
	case <-c.Request.Context().Done():
		log.Println(http_response.SendErrorAborted(c))
		return
	default:
	}

	resp := response.SuccessMessageResponse{
		SuccessMessage: "failed",
	}

	in := request.UpdateRoleRequest{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		log.Println(http_response.SendError(err, c))
		return
	}

	err = h.usecase.UpdateRole(c.Request.Context(), &in)
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
