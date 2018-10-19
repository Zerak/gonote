package main

type PixelItem struct {
	Color        int    `json:"color"`
	PriceCounter int    `json:"priceCounter"`
	Owner        string `json:"owner"`
}

type PixelRowItem struct {
	Row    int         `json:"row"`
	Pixels []PixelItem `json:"pixels"`
}

type Response struct {
	Rows PixelRowItem `json:"rows"`
}
