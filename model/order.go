package model

type Order struct {

	OrderId string `json:"orderId"`

	UserId string `json:"userId"`

	TotalPrice int64 `json:"totalPrice"`

	PayStatus int `json:"payStatus"`

	PayType int `json:"payType"`

	PayTime string `json:"payTime"`

	 OrderStatus int `json:"orderStatus"`;

	ExtraInfo string `json:"extra_info"`;

	UserAddress string `json:"userAddress"`

	IsDeleted bool `json:"isDeleted"`

	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}
