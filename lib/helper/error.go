package helper

import (
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

type StatusErrorDetail struct {
	StatusCode   int
	ErrorCode    string
	ErrorMessage string
	ErrorDetail  map[string]interface{}
}

func Error(errCode codes.Code, label string, err error) error {
	// make sure that we didnt have nested error codes
	if st, ok := status.FromError(err); !ok {
		var errMsg string

		if err, ok := err.(*pq.Error); ok {
			if err.Code == "23505" { // Duplicate data violation
				errMsg = "Duplicate data violation"

				if err.Table != "" {
					errMsg += " when trying to insert or update to : " + err.Table
				}
				return status.Error(errCode, errMsg+". Error details : "+label+err.Error())
			}
		}

		if err != nil {
			return status.Error(errCode, label+err.Error())
		} else {
			return status.Error(errCode, label)
		}
	} else {
		return status.Error(errCode, st.Message())
	}
}

func ErrorWithDetail(err error, errorDetail StatusErrorDetail) error {
	detail := make(map[string]interface{})
	if errorDetail.StatusCode != 0 {
		detail["statusCode"] = errorDetail.StatusCode
	}

	if errorDetail.ErrorCode != "" {
		detail["errorCode"] = errorDetail.ErrorCode
	}

	if errorDetail.ErrorMessage != "" {
		detail["errorMessage"] = errorDetail.ErrorMessage
	}

	if len(errorDetail.ErrorDetail) > 0 {
		detail["errorDetail"] = errorDetail.ErrorDetail
	}

	f, _ := structpb.NewStruct(detail)
	t, _ := anypb.New(f)

	if s, ok := status.FromError(err); ok {
		if ns, err := s.WithDetails(t); err == nil {
			return ns.Err()
		}
		return err
	}
	return err
}
