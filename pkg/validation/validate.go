package validation

import (
	"fmt"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/scalent-io/orchestration-framework/pkg/errors"
	"github.com/scalent-io/orchestration-framework/pkg/utils"
)

func DecodeAndValidate(r io.Reader, requestInstance interface{}, c *gin.Context) ([]utils.InvalidValidationError, errors.Response) {
	// decode the request
	InvalidValidationErrors := []utils.InvalidValidationError{}
	err := c.ShouldBindJSON(requestInstance)
	fmt.Println("---errr--------", err)
	if err != nil {
		fields, customValidationErrs := ValidationError(err)
		for i := 0; i < len(fields); i++ {
			log.Error().Str("validation  ", "Error").Msg(customValidationErrs[i])
			InvalidValidationErrors = append(InvalidValidationErrors, utils.InvalidValidationError{
				Field: fields[i],
				Msg:   customValidationErrs[i],
			})
		}

		return InvalidValidationErrors, errors.ResponseInternalServerError(err.Error())
	}

	return nil, nil
}

func DecodeAndValidateForQueryParams(c *gin.Context, requestInstance interface{}) errors.Response {
	// decode the request
	err := c.ShouldBindQuery(requestInstance)
	if err != nil {
		if numErr, ok := err.(*strconv.NumError); ok {
			msg1 := fmt.Sprintf("Invalid numeric value is " + numErr.Num)
			log.Error().Str("validation  ", "Error").Msg(err.Error())
			return errors.ResponseBadRequestError(msg1)
		}

		customValidationErrs := ValidationQueryParamsError(err)
		log.Error().Str("validation  ", "Error").Msg(err.Error())
		return errors.ResponseBadRequestError(customValidationErrs)
	}

	return nil
}
