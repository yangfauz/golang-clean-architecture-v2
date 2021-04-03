package service

import "api_project/model"

type UserService interface {
	List() (responses []model.GetUserResponse)
	Detail(Id int) (response model.GetUserResponse)
	Create(request model.CreateUserRequest) (response model.CreateUserResponse)
	Update(Id int, request model.CreateUserRequest) (response model.GetUserResponse)
	Delete(Id int) (response bool)
}
