package handlers

type UnreservedRequest struct {
	ItemsIDs []int64 `json:"items_ids"`
}

type UnreservedResponse struct {
	ItemsIDs []int64 `json:"items_ids"`
}

func (s *Store) Unreserved(req UnreservedRequest, res *UnreservedResponse) error {
	res.ItemsIDs = append(res.ItemsIDs, req.ItemsIDs[0])
	return nil
}
