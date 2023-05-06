package handlers

type RemainderRequest struct {
	StoreID int64 `json:"store_id"`
}

type RemainderResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	ItemID            int64 `json:"item_id"`
	AvailableQuantity int64 `json:"available_quantity"`
	ReservedQuantity  int64 `json:"reserved_quantity"`
}

func (s *Store) Remainder(req RemainderRequest, res *RemainderResponse) error {
	items, err := s.Service.Remainder(req.StoreID)
	for _, item := range items {
		res.Items = append(res.Items, Item{
			ItemID:            item["itemID"],
			AvailableQuantity: item["availableQuantity"],
			ReservedQuantity:  item["reservedQuantity"],
		})
	}
	return err
}
