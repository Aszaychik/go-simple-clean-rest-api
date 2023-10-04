package helper

import (
	"Aszaychik/go-simple-clean-rest-api/model/web"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func errorResponse(ctx echo.Context, code int, message string) error {
	return ctx.JSON(code, web.WebResponse{
		Code: code,
		Message: message,
		Data: nil,
	})
}

func StatusNotFound(ctx echo.Context, err error) error {
	return errorResponse(ctx, http.StatusNotFound, err.Error())
}

func StatusInternalServerError(ctx echo.Context, err error) error {
	logrus.Error(err.Error())
	return errorResponse(ctx, http.StatusInternalServerError, err.Error())
}

func StatusBadRequest(ctx echo.Context, err error) error {
	return errorResponse(ctx, http.StatusBadRequest, err.Error())
}

func StatusEmailAlreadyExist(ctx echo.Context, err error) error {
	return errorResponse(ctx, http.StatusConflict, err.Error())
}