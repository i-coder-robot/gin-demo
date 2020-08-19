package model

type User struct {
	UserId string `json:"userId"`
	NickName string `json:"nickName"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password"`
	Address string `json:"address"`
	IsDeleted bool `json:"isDeleted"`
	IsLocked bool `json:"isLocked"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}
