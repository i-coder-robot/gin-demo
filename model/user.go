package model

type User struct {

	UserId string `json:"userId"`

	NickName string `json:"nickName"`

	Password string `json:"password"`

	Address string `json:"address"`

	IsDeleted int `json:"isDeleted"`

	IsLocked int `json:"isLocked"`

	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}
