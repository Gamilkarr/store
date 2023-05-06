package handlers

type RemainderRequest struct {
	StoreID int64 `json:"store_id"`
}

type RemainderResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	ItemID            string `json:"item_id"`
	Name              string `json:"name"`
	AvailableQuantity string `json:"available_quantity"`
	ReservedQuantity  string `json:"reserved_quantity"`
}

func (s *Store) Remainder(req RemainderRequest, res *RemainderResponse) error {
	items, err := s.Repository.GetItemsQuantityOnStore(req.StoreID)
	for _, item := range items {
		res.Items = append(res.Items, Item{
			ItemID:            item["itemID"],
			Name:              item["name"],
			AvailableQuantity: item["availableQuantity"],
			ReservedQuantity:  item["reservedQuantity"],
		})
	}
	return err
}
