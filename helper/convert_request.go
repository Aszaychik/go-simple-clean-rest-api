package helper

import (
	"Aszaychik/go-simple-clean-rest-api/model/domain"
	"Aszaychik/go-simple-clean-rest-api/model/web"
)

func UserCreateRequestToUserDomain(request web.UserCreateRequest) *domain.User {
	return &domain.User{
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
	}
}

func UserLoginRequestToUserDomain(request web.UserLoginRequest) *domain.User {
	return &domain.User{
		Email: request.Email,
		Password: request.Password,
	}
}

func UserUpdateRequestToUserDomain(request web.UserUpdateRequest) *domain.User {
	return &domain.User{
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
	}
}