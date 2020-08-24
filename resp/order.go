package resp

type Order struct {
	OrderId string `json:"orderId"`
	UserId string `json:"userId"`
	TotalPrice int64 `json:"totalPrice"`
	PayStatus int `json:"payStatus"`
	PayType int `json:"payType"`
	PayTime string `json:"payTime"`
	OrderStatus int `json:"orderStatus"`
	ExtraInfo string `json:"extraInfo"`
	UserAddress string `json:"userAddress"`
	IsDeleted bool `json:"isDeleted"`
	//TODO后加的字段，需要后面加上，数据库也要加上这个字段
	NickName string `json:"nickName"`
	Mobile string `json:"mobile"`
}

