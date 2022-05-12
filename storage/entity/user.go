package entity

type Users struct {
	Guid        string `json:"guid" db:"guid"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	Name        string `json:"name" db:"name"`
	Login       string `json:"login" db:"login"`
	Password    string `json:"password" db:"password"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}

type CreateUserRequest struct {
	Name        string `json:"name" db:"name"`
	Login       string `json:"login" db:"login"`
	Password    string `json:"password" db:"password"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}

type CreateUserResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type GetUserRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type GetUserResponse struct {
	Guid        string `json:"guid" db:"guid"`
	Name        string `json:"name" db:"name"`
	Login       string `json:"login" db:"login"`
	Password    string `json:"password" db:"password"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}

type UpdateUserRequest struct {
	Name        string `json:"name" db:"name"`
	Login       string `json:"login" db:"login"`
	Password    string `json:"password" db:"password"`
	Guid        string `json:"guid" db:"guid"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}

type UpdateUserResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type DeleteUserRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type ListUserRequest struct {
}

type ListUserResponse struct {
	Users []GetUserResponse `json:"users"`
}

type ResponseOK struct {
	Message interface{}
}

type LoginUserRequest struct {
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}

type LoginUserResponse struct {
	Exist       bool   `json:"exist"`
	Guid        string `json:"guid" db:"guid"`
	Name        string `json:"name" db:"name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}

type GetUserOrdersRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type GetUserOrdersResponse struct {
	Orders []GetOrderResponse `json:"orders"`
}
