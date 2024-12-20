package model

type CustomerCreateInput struct {
	CustomerName string `json:"customer_name"`
	Password     string `json:"password"`
	Mobile       string `json:"mobile"`
	Email        string `json:"email"`
}

type CustomerUpdateInput struct {
	Id           int64  `json:"id"`
	CustomerName string `json:"customer_name"`
	Password     string `json:"password"`
	Mobile       string `json:"mobile"`
	Email        string `json:"email"`
}

type CustomerGetListInput struct {
	CustomerName string `json:"customer_name"`
}

type CustomerLoginInput struct {
	CustomerName string `json:"customer_name"`
	Password     string `json:"password"`
}

type CustomerResetPasswordInput struct {
	Id       int64  `v:"required" dc:"customer id"`
	Password string `json:"password"`
}
