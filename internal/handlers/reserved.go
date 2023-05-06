package handlers

type ReservedRequest struct {
	StoreID          int64             `json:"store_id"`
	ItemsForReserved []ItemForReserved `json:"items_ids"`
}

type ReservedResponse struct {
	Status string `json:"status"`
}

type ItemForReserved struct {
	ItemID              int64 `json:"item_id"`
	QuantityForReserved int64 `json:"quantity_for_reserved"`
}

func (s *Store) Reserved(req ReservedRequest, res *ReservedResponse) error {
	var items map[int64]int64

	// учти что резерв не суммируется, а перезаписывается
	for _, item := range req.ItemsForReserved {
		items[item.ItemID] = item.QuantityForReserved
	}

	if err := s.Service.Reserved(req.StoreID, items); err != nil {
		return err
	}
	res.Status = "ok"
	return nil
}
