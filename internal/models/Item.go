package models

type Item struct {
	ItemID            int64  `json:"item_id"`
	Name              string `json:"name"`
	AvailableQuantity int64  `json:"available_quantity"`
	ReservedQuantity  int64  `json:"reserved_quantity"`
}

type Reserved struct {
	ItemID   int64 `json:"id"`
	Quantity int64 `json:"quantity"`
}
