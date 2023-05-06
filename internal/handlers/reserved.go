package handlers

type ReservedRequest struct {
	StoreID          int64             `json:"store_id"`
	ItemsForReserved []ItemForReserved `json:"items_for_reserved"`
}

type ReservedResponse struct {
	Status string `json:"status"`
}

type ItemForReserved struct {
	ID       int64 `json:"id"`
	Quantity int64 `json:"quantity"`
}

func (s *Store) Reserved(req ReservedRequest, res *ReservedResponse) error {
	items := make(map[int64]int64, len(req.ItemsForReserved))

	// учти что резерв не суммируется, а перезаписывается
	for _, item := range req.ItemsForReserved {
		items[item.ID] = item.Quantity
	}

	if err := s.Repository.ItemReserved(req.StoreID, items); err != nil {
		return err
	}
	res.Status = "ok"
	return nil
}
