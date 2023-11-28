package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	apimodel "github.com/scalent-io/orchestration-framework/apimodel"
	"github.com/scalent-io/orchestration-framework/pkg/errors"
)

type InvalidValidationError struct {
	Field string `json:"Field"`
	Msg   string `json:"Msg"`
}

const (
	STATUS_SUCCESS               = "success"
	STATUS_FAILED                = "failed"
	STATUS_INTERNAL_SERVER_ERROR = "internal server error"
)

func DataResponse(c *gin.Context, statusCode int, msg string, payload interface{}) {
	res := apimodel.Response{
		StatusCode: statusCode,
		Status:     STATUS_SUCCESS,
		Message:    msg,
		Data:       payload,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(res)
	c.Writer.Write(response)
	c.Writer.WriteHeader(statusCode)
}

func ErrorResponse(c *gin.Context, err errors.Response, payload interface{}) {

	if err == nil {
		ErrorResponse(c, errors.ResponseInternalServerError(STATUS_INTERNAL_SERVER_ERROR), nil)
		return
	}

	errMsg := err.Error()

	res := apimodel.Response{
		StatusCode: err.StatusCode(),
		Status:     STATUS_FAILED,
		Message:    errMsg,
		Data:       payload,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(res)
	c.Writer.WriteHeader(err.StatusCode())
	c.Writer.Write(response)
}

func ValidationErrorResponse(c *gin.Context, fieldErrs []InvalidValidationError, payload interface{}) {

	if fieldErrs == nil {
		ErrorResponse(c, errors.ResponseInternalServerError(STATUS_INTERNAL_SERVER_ERROR), nil)
		return
	}

	type Response struct {
		StatusCode int                      `json:"code"`
		Status     string                   `json:"status"`
		Message    []InvalidValidationError `json:"message"`
		Data       interface{}              `json:"data,omitempty"`
	}

	temp := Response{
		StatusCode: 400,
		Status:     STATUS_FAILED,
		Message:    fieldErrs,
		Data:       payload,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(temp)
	c.Writer.WriteHeader(400)
	c.Writer.Write(response)
}
