package handlers

type RemainderRequest struct {
	StoreID int64 `json:"store_id"`
}

type RemainderResponse struct {
	StoreID int64 `json:"store_id"`
}

func (s *Store) Remainder(req RemainderRequest, res *RemainderResponse) error {
	res.StoreID = req.StoreID
	return nil
}
