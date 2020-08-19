package model

type Banner struct {
	BannerID int `json:"bannerID"`
	Url string `json:"url"`
	RedirectUrl string `json:"redirectUrl"`
	Order int `json:"order"`
	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`

	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`

}