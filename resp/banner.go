package resp

type Banner struct {
	BannerID    string `json:"bannerID"`
	Url         string `json:"url"`
	RedirectUrl string `json:"redirectUrl"`
	OrderBy     int    `json:"order"`
}
