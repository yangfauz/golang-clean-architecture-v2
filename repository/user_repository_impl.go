package repository

import (
	"api_project/entity"
	"api_project/exception"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type userRepositoryImpl struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepositoryImpl{database}
}

func (repository *userRepositoryImpl) FindAll() (users []entity.User) {
	rows, err := repository.database.Query("SELECT id, name, hobby FROM users order by id")
	exception.PanicIfNeeded(err)
	defer rows.Close()

	var customers []entity.User

	for rows.Next() {
		var customer entity.User

		err := rows.Scan(&customer.Id, &customer.Name, &customer.Hobby)

		exception.PanicIfNeeded(err)

		customers = append(customers, customer)
	}

	return customers
}

func (repository *userRepositoryImpl) FindById(Id int) (user entity.User) {
	var customers entity.User
	err := repository.database.QueryRow("SELECT id, name, hobby FROM users WHERE id = ? ", Id).Scan(&customers.Id, &customers.Name, &customers.Hobby)

	if err != nil && err != sql.ErrNoRows {
		exception.PanicIfNeeded(err)
	}

	return customers
}

func (repository *userRepositoryImpl) Insert(user entity.User) {
	_, err := repository.database.Query("INSERT INTO users (name, hobby) VALUES (?, ?)", user.Name, user.Hobby)

	exception.PanicIfNeeded(err)
}

func (repository *userRepositoryImpl) Update(Id int, user entity.User) {
	_, err := repository.database.Query("UPDATE users SET name = ? ,hobby = ? WHERE id = ?", user.Name, user.Hobby, Id)

	exception.PanicIfNeeded(err)
}

func (repository *userRepositoryImpl) Delete(Id int) (user bool) {
	_, err := repository.database.Query("DELETE FROM users WHERE id = ? ", Id)
	exception.PanicIfNeeded(err)

	return true
}
