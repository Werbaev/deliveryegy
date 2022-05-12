package models

type Empty struct{}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//Error ...
type Error struct {
	Error interface{}
}

//ValidationError ...
type ValidationError struct {
	Code        string
	Message     string
	UserMessage string
}

type ResponseOK struct {
	Message interface{}
}

type Response struct {
	ID interface{} `json:"id"`
}

// Find query ...
type FindQueryModel struct {
	Page     int64  `json:"page,string"`
	Search   string `json:"search"`
	Active   bool   `json:"active"`
	Inactive bool   `josn:"inactive"`
	Limit    int64  `json:"limit,string"`
	Sort     string `json:"sort" example:"name|asc"`
	Lang     string `json:"lang"`
}

type GetRequest struct {
	Id   string `json:"id"`
	Slug string `json:"slug"`
	Lang string `json:"lang"`
}

type AuthorizationModel struct {
	Token string `header:"Authorization"`
}

type ResponseID struct {
	ID string `json:"id"`
}

type CreateResponse struct {
	Uuid string `json:"uuid"`
}
