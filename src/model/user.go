package model

type User struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status int `json:"status"`
}

type UserListResponse struct {
	Status int64 `json:"status"`
	Message string `json:"message"`
	Result []User `json:"result"`
}

type UserResponse struct {
	Status int64 `json:"status"`
	Message string `json:"message"`
	Result User `json:"result"`
}

type UserResponseError struct {
	Status int64 `json:"status"`
	Message string `json:"message"`
	Result string `json:"result"`
}


