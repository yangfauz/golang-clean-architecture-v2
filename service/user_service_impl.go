package service

import (
	"api_project/entity"
	"api_project/model"
	"api_project/repository"
	"api_project/validation"
)

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{*userRepository}
}

func (service *userServiceImpl) List() (responses []model.GetUserResponse) {
	users := service.UserRepository.FindAll()
	for _, user := range users {
		responses = append(responses, model.GetUserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Hobby: user.Hobby,
		})
	}
	return responses
}

func (service *userServiceImpl) Detail(Id int) (response model.GetUserResponse) {
	user := service.UserRepository.FindById(Id)
	response = model.GetUserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Hobby: user.Hobby,
	}
	return response
}

func (service *userServiceImpl) Create(request model.CreateUserRequest) (response model.CreateUserResponse) {
	validation.Validate(request)

	user := entity.User{
		Name:  request.Name,
		Hobby: request.Hobby,
	}

	service.UserRepository.Insert(user)

	response = model.CreateUserResponse{
		Name:  user.Name,
		Hobby: user.Hobby,
	}

	return response
}

func (service *userServiceImpl) Update(Id int, request model.CreateUserRequest) (response model.GetUserResponse) {
	validation.Validate(request)

	user := entity.User{
		Name:  request.Name,
		Hobby: request.Hobby,
	}

	service.UserRepository.Update(Id, user)

	response = model.GetUserResponse{
		Id:    Id,
		Name:  user.Name,
		Hobby: user.Hobby,
	}

	return response
}

func (service *userServiceImpl) Delete(Id int) (response bool) {
	user := service.UserRepository.Delete(Id)
	return user
}
