package model

type MerchantCreateInput struct {
	Name    string `json:"merchant_name"`
	Address string `json:"address"`
	Mobile  string `json:"mobile"`
}

type MerchantUpdateInput struct {
	Id      int64  `json:"id"`
	Name    string `json:"merchant_name"`
	Address string `json:"address"`
	Mobile  string `json:"mobile"`
}

type MerchantGetListInput struct {
	Name string `json:"merchant_name"`
}
