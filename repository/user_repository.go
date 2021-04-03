package repository

import "api_project/entity"

type UserRepository interface {
	FindAll() (users []entity.User)
	FindById(Id int) (user entity.User)
	Insert(user entity.User)
	Update(Id int, user entity.User)
	Delete(Id int) (user bool)
}
