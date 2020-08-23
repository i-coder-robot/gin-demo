package model

type Category struct {
	CategoryID string `json:"categoryID"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Order int `json:"order"`
	Children []*Category2 `json:"children"`
	IsDeleted bool `json:"isDeleted" gorm:"column:is_deleted"`
	CreateAt string `json:"createAt" gorm:"column:create_at"`
	UpdateAt string `json:"updateAt" gorm:"column:update_at"`
}

type Category2 struct {
	CategoryID string `json:"categoryID"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Order int `json:"order"`
	Children []*Category3 `json:"children"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

type Category3 struct {
	CategoryID string `json:"categoryID"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Order int `json:"order"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

type CategoryResult struct {

	C1CategoryID string `gorm:"c1_category_id"`
	C1Name string `gorm:"column:c1_name"`
	C1Desc string `gorm:"column:c1_desc"`
	C1Order int  `gorm:"column:c1_order"`
	C1ParentId string `gorm:"column:c1_parent_id"`

	C2CategoryID string `gorm:"c2_category_id"`
	C2Name string `gorm:"column:c2_name"`
	C2Order int  `gorm:"column:c2_order"`
	C2ParentId string `gorm:"column:c2_parent_id"`

	C3CategoryID string `gorm:"c3_category_id"`
	C3Name string `gorm:"column:c3_name"`
	C3Order int  `gorm:"column:c3_order"`
	C3ParentId string `gorm:"column:c3_parent_id"`
}
