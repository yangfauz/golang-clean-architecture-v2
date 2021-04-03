package model

type GetUserResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}

type CreateUserResponse struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}
