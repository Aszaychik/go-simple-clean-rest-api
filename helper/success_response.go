package helper

import (
	"Aszaychik/go-simple-clean-rest-api/model/web"
	"net/http"

	"github.com/labstack/echo/v4"
)

func successResponse(ctx echo.Context, code int, message string, data any) error {
	return ctx.JSON(code, web.WebResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func StatusCreated(ctx echo.Context, message string, data any) error {
	return successResponse(ctx, http.StatusCreated, message, data)
}

func StatusOK(ctx echo.Context, message string, data any) error {
	return successResponse(ctx, http.StatusOK, message, data)
}