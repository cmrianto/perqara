package http_response

import (
	"context"
	"net/http"
	"perqara/lib/helper"
	"perqara/lib/helper/timestamp"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

func SendError(err error, c *gin.Context) string {
	// set error codes, only if previous error doesnt have any
	httpStatus := http.StatusInternalServerError
	errMsg := err.Error()
	var response Response
	response.Data = Empty{}
	if st, ok := status.FromError(err); ok {
		httpStatus = STATUS_MAP[st.Code()]
		errMsg = st.Message()
		errMsg = helper.GetMessageFromFormattedErr(errMsg)
		dStatus, dMessage, dData, isFilled := responseFromDetail(st)
		if isFilled {
			if dStatus != 0 {
				httpStatus = dStatus
			}

			if dMessage != "" {
				errMsg = dMessage
			}

			if dData != nil {
				response.Data = map[string]interface{}{
					"errorMessage": dData,
					"isSuccess":    false,
				}
			}
		}
	}

	response.StatusCode = httpStatus
	response.Message = errMsg
	response.Status = STATUS_DESC[httpStatus]
	response.Timestamp = timestamp.GetNow()

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Authorization, X-SKIP-AUTH")
	c.Header("Access-Control-Allow-Methods", "DELETE, POST, HEAD, PATCH, OPTIONS, GET, PUT")

	c.JSON(httpStatus, response)
	return errMsg
}

func SendErrorAborted(c *gin.Context) string {
	switch c.Request.Context().Err() {
	case context.Canceled:
		return SendError(status.Error(codes.Canceled, "Request Canceled"), c)
	case context.DeadlineExceeded:
		return SendError(status.Error(codes.DeadlineExceeded, "Deadline Exceeded"), c)
	default:
		return ""
	}
}

func SendAbortedUnauthenticate(c *gin.Context) string {
	return SendError(status.Error(codes.Unauthenticated, "Request Unauthenticate"), c)
}

func SendAbortedUnauthorized(c *gin.Context) string {
	return SendError(status.Error(codes.Unauthenticated, "Request Unauthorized"), c)
}

func AbortError(err error, c *gin.Context) {
	httpStatus := http.StatusInternalServerError
	errMsg := err.Error()
	if st, ok := status.FromError(err); ok {
		httpStatus = STATUS_MAP[st.Code()]
		errMsg = st.Message()
	}

	response := Response{
		StatusCode: httpStatus,
		Message:    errMsg,
		Status:     STATUS_DESC[httpStatus],
		Timestamp:  timestamp.GetNow(),
		Data:       Empty{},
	}
	c.AbortWithStatusJSON(httpStatus, response)
}

func AbortUnauthenticate(c *gin.Context) {
	AbortError(status.Error(codes.Unauthenticated, "Request Unauthenticate"), c)
}

func AbortUnauthorized(c *gin.Context) {
	AbortError(status.Error(codes.Unauthenticated, "Request Unauthorized"), c)
}

func responseFromDetail(st *status.Status) (int, string, map[string]interface{}, bool) {
	var status int
	var errorMesasge string
	data := map[string]interface{}{}
	dt := st.Details()
	if len(dt) > 0 {
		for idx := range dt {
			if a, ok := dt[idx].(*anypb.Any); ok {
				var t structpb.Struct
				if err := anypb.UnmarshalTo(a, &t, proto.UnmarshalOptions{}); err == nil {
					for k, v := range t.GetFields() {
						switch k {
						case "statusCode":
							sc := v.GetNumberValue()
							status = int(sc)
						case "errorMessage":
							msg := v.GetStringValue()
							errorMesasge = msg
						case "errorDetail":
							src := v.GetStructValue()
							for kc, vc := range src.GetFields() {
								data[kc] = vc.AsInterface()
							}
						}
					}
				}
			}
		}
		return status, errorMesasge, data, true
	}
	return status, errorMesasge, data, false
}
