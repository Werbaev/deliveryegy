package entity

type Couriers struct {
	Guid      string `json:"guid" db:"guid"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Name      string `json:"name" db:"name"`
	Login     string `json:"login" db:"login"`
	Password  string `json:"password" db:"password"`
}

type CreateCourierRequest struct {
	Name     string `json:"name" db:"name"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}

type CreateCourierResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type GetCourierRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type GetCourierResponse struct {
	Guid      string `json:"guid" db:"guid"`
	Name      string `json:"name" db:"name"`
	Login     string `json:"login" db:"login"`
	Password  string `json:"password" db:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type UpdateCourierRequest struct {
	Name     string `json:"name" db:"name"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
	Guid     string `json:"guid" db:"guid"`
}

type UpdateCourierResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type DeleteCourierRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type ListCourierRequest struct {
}

type ListCourierResponse struct {
	Couriers []GetCourierResponse `json:"couriers"`
}

type ListCourierOrdersRequest struct {
	CourierId string `json:"courier_id" db:"courier_id"`
}

type ListCourierOrdersResponse struct {
	Orders []GetOrderResponse `json:"orders"`
}

type LoginCourierRequest struct {
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}

type LoginCourierResponse struct {
	Exist bool   `json:"exist"`
	Guid  string `json:"guid" db:"guid"`
	Name  string `json:"name" db:"name"`
}
