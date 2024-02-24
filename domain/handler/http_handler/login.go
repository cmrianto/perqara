package http_handler

import (
	"log"
	"net/http"
	"perqara/lib/helper/http_response"
	"perqara/lib/helper/timestamp"

	"perqara/domain/request"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login
// @Description API to login
// @Tags Auth
// @Accept json
// @Produce json
// @Param	login	body	request.LoginRequest	true	"Login Data"
// @Success 200 {object} http_response.Response{data=response.LoginResponse,statusCode=int,message=string}
// @Failure 400 {object} http_response.Response{code=int,message=string}
// @Failure 500 {object} http_response.Response{code=int,message=string}
// @Router /login [post]
func (h *handler) Login(c *gin.Context) {
	select {
	case <-c.Request.Context().Done():
		log.Println(http_response.SendErrorAborted(c))
		return
	default:
	}

	in := request.LoginRequest{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		log.Println(http_response.SendError(err, c))
		return
	}

	result, err := h.usecase.Login(c.Request.Context(), in.ToPayload())
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
