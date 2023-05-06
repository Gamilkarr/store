package handlers

type UnreservedRequest struct {
	StoreID           int64               `json:"store_id"`
	ItemForUnreserved []ItemForUnreserved `json:"items_ids"`
}

type UnreservedResponse struct {
	Status string `json:"status"`
}

type ItemForUnreserved struct {
	ItemID                int64 `json:"item_id"`
	QuantityForUnreserved int64 `json:"quantity_for_reserved"`
}

func (s *Store) Unreserved(req UnreservedRequest, res *UnreservedResponse) error {
	var items map[int64]int64

	// учти что резерв не суммируется, а перезаписывается
	for _, item := range req.ItemForUnreserved {
		items[item.ItemID] = item.QuantityForUnreserved
	}

	if err := s.Service.Unreserved(req.StoreID, items); err != nil {
		return err
	}
	res.Status = "ok"
	return nil
}
