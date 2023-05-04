package handlers

type ReservedRequest struct {
	ItemsIDs []int64 `json:"items_ids"`
}

type ReservedResponse struct {
	ItemsIDs []int64 `json:"items_ids"`
}

func (s *Store) Reserved(req ReservedRequest, res *ReservedResponse) error {
	res.ItemsIDs = req.ItemsIDs
	return nil
}
