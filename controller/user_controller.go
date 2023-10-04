package controller

import (
	"Aszaychik/go-simple-clean-rest-api/helper"
	"Aszaychik/go-simple-clean-rest-api/model/web"
	"Aszaychik/go-simple-clean-rest-api/service"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	RegisterUserController(ctx echo.Context ) error
	LoginUserController(ctx echo.Context ) error
	UpdateUserController(ctx echo.Context) error
	GetUserController(ctx echo.Context) error
	GetUsersController(ctx echo.Context) error
	DeleteUserController(ctx echo.Context) error
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func(c *UserControllerImpl) RegisterUserController(ctx echo.Context) error {
	userCreateRequest := web.UserCreateRequest{}
	err := ctx.Bind(&userCreateRequest) 
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}

	response, err := c.UserService.CreateUser(ctx, userCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "Email already exists") {
			return helper.StatusEmailAlreadyExist(ctx, err)
		}

		return helper.StatusInternalServerError(ctx, err)
	}

	return helper.StatusCreated(ctx, "Success to create user", response)
}

func (c *UserControllerImpl) LoginUserController(ctx echo.Context) error {
	userLoginRequest := web.UserLoginRequest{}
	err := ctx.Bind(&userLoginRequest)
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}
	
	response, err := c.UserService.LoginUser(ctx, userLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}
		
		if strings.Contains(err.Error(), "Invalid email or password") {
			return helper.StatusBadRequest(ctx, err)
		}
		
		return helper.StatusInternalServerError(ctx, err)
	}

	userLoginResponse := helper.UserDomainToUserLoginResponse(response)

	token, err := helper.GenerateToken(&userLoginResponse, response.ID)
	if err != nil {
		return helper.StatusInternalServerError(ctx, err)
	}

	userLoginResponse.Token = token

	return helper.StatusOK(ctx, "Success to login user", userLoginResponse)
}

func (c *UserControllerImpl) UpdateUserController(ctx echo.Context) error {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return helper.StatusInternalServerError(ctx, err)
	}

	userUpdateRequest := web.UserUpdateRequest{}
	err = ctx.Bind(&userUpdateRequest)
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}

	response, err := c.UserService.UpdateUser(ctx, userUpdateRequest, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "User not found") {
			return helper.StatusNotFound(ctx, err)
		}

		return helper.StatusInternalServerError(ctx, err)
	}

	return helper.StatusOK(ctx, "Success to update user", response)
}

func (c *UserControllerImpl) GetUserController(ctx echo.Context) error {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return helper.StatusInternalServerError(ctx, err)
	}
	
	response, err := c.UserService.FindById(ctx, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "User not found") {
			return helper.StatusNotFound(ctx, err)
		}

		return helper.StatusInternalServerError(ctx, err)
	}

	return helper.StatusOK(ctx, "Success to get user", response)
}

func (c *UserControllerImpl) GetUsersController(ctx echo.Context) error {
	response, err := c.UserService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Users not found") {
			return helper.StatusNotFound(ctx, err)
		}

		return helper.StatusInternalServerError(ctx, err)
	}

	return helper.StatusOK(ctx, "Success to get users", response)
}

func (c *UserControllerImpl) DeleteUserController(ctx echo.Context) error {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return helper.StatusInternalServerError(ctx, err)
	}

	err = c.UserService.DeleteUser(ctx, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "User not found") {
			return helper.StatusNotFound(ctx, err)
		}

		return helper.StatusInternalServerError(ctx, err)
	}

	return helper.StatusOK(ctx, "Success to delete user", nil)
}