package helper

import (
	"unit_testing/model/domain"
	"unit_testing/model/web"
)

func UserDomainToUserLoginResponse(user *domain.User) web.UserLoginResponse {
	return web.UserLoginResponse{
		Name: user.Name,
		Email: user.Email,
	}
}