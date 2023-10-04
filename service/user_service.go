package service

import (
	"Aszaychik/go-simple-clean-rest-api/helper"
	"Aszaychik/go-simple-clean-rest-api/model/domain"
	"Aszaychik/go-simple-clean-rest-api/model/web"
	"Aszaychik/go-simple-clean-rest-api/repository"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	CreateUser(ctx echo.Context, request web.UserCreateRequest) (*domain.User, error)
	LoginUser(ctx echo.Context, request web.UserLoginRequest) (*domain.User, error)
	UpdateUser(ctx echo.Context, request web.UserUpdateRequest, id int) (*domain.User, error)
	FindById(ctx echo.Context, id int) (*domain.User, error)
	FindAll(ctx echo.Context) ([]domain.User, error)
	DeleteUser(ctx echo.Context, id int) error
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate  *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate: validate,
	}
}

func (service *UserServiceImpl) CreateUser(ctx echo.Context, request web.UserCreateRequest) (*domain.User, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	// Check if the email already exists
	existingUser, _ := service.UserRepository.FindByEmail(request.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("Email already exists")
	}

	// Convert request to domain
	user := helper.UserCreateRequestToUserDomain(request)
	// Convert password to hash
	user.Password = helper.HashPassword(user.Password)

	result, err := service.UserRepository.Create(user)
	if err != nil {
		return nil, fmt.Errorf("Error when creating user: %s", err.Error())
	}

	return result, nil
}

func (service *UserServiceImpl) LoginUser(ctx echo.Context, request web.UserLoginRequest) (*domain.User, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	// Check if the user exists
	existingUser, err := service.UserRepository.FindByEmail(request.Email)
	if err != nil  {
		return nil, fmt.Errorf("Invalid email or password")
	}
	
	// Convert request to domain
	user := helper.UserLoginRequestToUserDomain(request)

	// Compare password
	err = helper.ComparePassword(existingUser.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("Invalid email or password")
	}

	return existingUser, nil
}

func (service *UserServiceImpl) UpdateUser(ctx echo.Context, request web.UserUpdateRequest, id int) (*domain.User, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	// Check if the user exists
	existingUser, _ := service.UserRepository.FindById(id)
	if existingUser == nil {
		return nil, fmt.Errorf("User not found")
	}

	// Convert request to domain
	user := helper.UserUpdateRequestToUserDomain(request)
	user.Password = helper.HashPassword(user.Password)

	result, err := service.UserRepository.Update(user, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating user: %s", err.Error())
	}

	return result, nil
}

func (service *UserServiceImpl) FindById(ctx echo.Context, id int) (*domain.User, error) {
	// Check if the user exists
	existingUser, _ := service.UserRepository.FindById(id)
	if existingUser == nil {
		return nil, fmt.Errorf("User not found")
	}

	return existingUser, nil
}

func (service *UserServiceImpl) FindAll(ctx echo.Context) ([]domain.User, error) {
	users, err := service.UserRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Users not found")
	}

	return users, nil
}

func (service *UserServiceImpl) DeleteUser(ctx echo.Context, id int) error {
	// Check if the user exists
	existingUser, _ := service.UserRepository.FindById(id)
	if existingUser == nil {
		return fmt.Errorf("User not found")
	}

	err := service.UserRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting user: %s", err)
	}

	return nil
}