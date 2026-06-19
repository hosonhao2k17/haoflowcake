package dto

import (
	"fmt"
	"haoflowcake/common/constant"
	"time"
)

type ErrorResDto struct {
	Message    string             `json:"message"`
	StatusCode int                `json:"status_code"`
	Timestamp  time.Time          `json:"timestamp"`
	ErrorCode  constant.ErrorCode `json:"error_code"`
	Details    []ErrorDetail      `json:"details,omitempty"`
}

type ErrorDetail struct {
	Property string `json:"property"`
	Code     string `json:"code"`
	Message  string `json:"message"`
}

func (err *ErrorResDto) Error() string {
	return err.Message
}

func NewErrorResponse(errCode constant.ErrorCode) *ErrorResDto {
	fmt.Println(errCode)
	currentTime := time.Now()
	return &ErrorResDto{
		ErrorCode:  errCode,
		Message:    constant.MapErrorCode[errCode].Message,
		StatusCode: constant.MapErrorCode[errCode].StatusCode,
		Timestamp:  currentTime,
	}
}

func NewFullErrorResponse(errCode constant.ErrorCode, details []ErrorDetail) *ErrorResDto {

	currentTime := time.Now()
	return &ErrorResDto{
		ErrorCode:  errCode,
		Message:    constant.MapErrorCode[errCode].Message,
		StatusCode: constant.MapErrorCode[errCode].StatusCode,
		Timestamp:  currentTime,
		Details:    details,
	}
}
