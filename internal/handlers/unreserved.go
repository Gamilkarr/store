package handlers

type UnreservedRequest struct {
	StoreID           int64               `json:"store_id"`
	ItemForUnreserved []ItemForUnreserved `json:"items_for_unreserved"`
}

type UnreservedResponse struct {
	Status string `json:"status"`
}

type ItemForUnreserved struct {
	ID       int64 `json:"id"`
	Quantity int64 `json:"quantity"`
}

func (s *Store) Unreserved(req UnreservedRequest, res *UnreservedResponse) error {
	items := make(map[int64]int64, len(req.ItemForUnreserved))

	// учти что резерв не суммируется, а перезаписывается
	for _, item := range req.ItemForUnreserved {
		items[item.ID] = item.Quantity
	}

	if err := s.Repository.ItemUnreserved(req.StoreID, items); err != nil {
		return err
	}
	res.Status = "ok"
	return nil
}
