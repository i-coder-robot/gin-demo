package model

type Banner struct {
	BannerID    string    `json:"bannerID"`
	Url         string `json:"url"`
	RedirectUrl string `json:"redirectUrl"`
	OrderBy     int    `json:"order"`
	CreateUser  string `json:"createUser"`
	UpdateUser  string `json:"updateUser"`

	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`

}