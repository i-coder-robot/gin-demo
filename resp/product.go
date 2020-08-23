package resp

type Product struct {
	ProductId string `json:"productId"`
	ProductName string `json:"productName"`;
	ProductIntro string `json:"productIntro"`
	CategoryId string `json:"categoryId"`
	ProductCoverImg string `json:"productCoverImg"`
	ProductBanner string `json:"productBanner"`
	OriginalPrice int `json:"originalPrice"`;
	SellingPrice int `json:"sellingPrice"`
	StockNum int `json:"stockNum"`
	Tag string `json:"tag"`
	SellStatus bool `json:"sellStatus"`
	ProductDetailContent string `json:"productDetailContent"`
	IsDeleted bool `json:"isDeleted"`
}
