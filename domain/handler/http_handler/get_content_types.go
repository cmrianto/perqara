package http_handler

import (
	"log"
	"net/http"
	"perqara/lib/helper/http_response"
	"perqara/lib/helper/timestamp"

	"perqara/domain/request"

	"github.com/gin-gonic/gin"
)

// GetContentTypes godoc
// @Summary Get Content Type List
// @Description API to get content types
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param	q	query	request.GetContentTypesRequest	true	"Query Params"
// @Success 200 {object} http_response.Response{data=response.GetContentTypesResponse,statusCode=int,message=string}
// @Failure 400 {object} http_response.Response{code=int,message=string}
// @Failure 500 {object} http_response.Response{code=int,message=string}
// @Router /content-type [get]
func (h *handler) GetContentTypes(c *gin.Context) {
	select {
	case <-c.Request.Context().Done():
		log.Println(http_response.SendErrorAborted(c))
		return
	default:
	}

	in := request.GetContentTypesRequest{
		Search:        c.Query("search"),
		Limit:         c.Query("size"),
		Offset:        c.Query("page"),
		LastCreatedAt: c.Query("last_created_at"),
	}

	result, err := h.usecase.GetContentTypes(c.Request.Context(), in.ToPayload())
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
