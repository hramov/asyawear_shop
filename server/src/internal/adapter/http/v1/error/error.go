package error_handler

import "asyawear/shared/lib/utils"

type AppError struct {
	Code    utils.HttpCode
	Message string
}

func ParseError(err error) AppError {
	return AppError{
		Code:    utils.StatusBadRequest,
		Message: err.Error(),
	}
}
