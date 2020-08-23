package model

type Order struct {
	OrderId string `json:"orderId" gorm:"column:order_id"`
	UserId string `json:"userId" gorm:"column:user_id"`
	TotalPrice int64 `json:"totalPrice" gorm:"column:total_price"`
	PayStatus int `json:"payStatus" gorm:"column:pay_status"`
	PayType int `json:"payType" gorm:"column:pay_type"`
	PayTime string `json:"payTime" gorm:"column:pay_time"`
	OrderStatus int `json:"orderStatus" gorm:"column:order_status"`
	ExtraInfo string `json:"extra_info" gorm:"column:extra_info"`
	UserAddress string `json:"userAddress" gorm:"column:user_address"`
	IsDeleted bool `json:"isDeleted" gorm:"column:is_deleted"`
	CreateAt string `json:"createAt" gorm:"column:create_at"`
	UpdateAt string `json:"updateAt" gorm:"column:update_at"`
}
