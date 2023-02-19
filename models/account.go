package models

type Account struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type SignUpInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type SignUpOutput struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type AccountGetByIdInput struct {
	Id int `form:"accountId" binding:"required,numeric"`
}

type AccountGetByIdOutput struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type AccountSearchInput struct {
	FirstName string `form:"firstName"`
	LastName  string `form:"lastName"`
	Email     string `form:"email"`
	Offset    int    `form:"from"`
	Limit     int    `form:"size"`
}

type AccountSearchOutput struct {
	Id        int    `form:"id"`
	FirstName string `form:"firstName"`
	LastName  string `form:"lastName"`
	Email     string `form:"email"`
}

type AccountUpdateInput struct {
	Id        int    `form:"accountId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AccountUpdateOutput struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type AccountRemoveInput struct {
	Id int `form:"accountId"`
}
