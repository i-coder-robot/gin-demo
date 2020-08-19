package model

type Category struct {
	CategoryID string `json:"categoryID"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Order int `json:"order"`
	Children []*Category `json:"children"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}
