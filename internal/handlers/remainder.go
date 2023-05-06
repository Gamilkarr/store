package handlers

import "github.com/Gamilkarr/store/internal/services"

type RemainderRequest struct {
	StoreID int64 `json:"store_id"`
}

type RemainderResponse struct {
	Items []services.Item `json:"items"`
}

func (s *Store) Remainder(req RemainderRequest, res *RemainderResponse) error {
	var err error
	res.Items, err = s.service.Remainder(req.StoreID)
	return err
}
