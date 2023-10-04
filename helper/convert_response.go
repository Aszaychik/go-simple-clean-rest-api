package helper

import (
	"Aszaychik/go-simple-clean-rest-api/model/domain"
	"Aszaychik/go-simple-clean-rest-api/model/web"
)

func UserDomainToUserLoginResponse(user *domain.User) web.UserLoginResponse {
	return web.UserLoginResponse{
		Name: user.Name,
		Email: user.Email,
	}
}