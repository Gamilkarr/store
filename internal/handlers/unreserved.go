package handlers

import "github.com/Gamilkarr/store/internal/services"

type UnreservedRequest struct {
	StoreID int64                        `json:"store_id"`
	Items   []services.ItemForUnreserved `json:"items_ids"`
}

type UnreservedResponse struct {
	Status string `json:"status"`
}

func (s *Store) Unreserved(req UnreservedRequest, res *UnreservedResponse) error {
	if err := s.service.Unreserved(req.StoreID, req.Items); err != nil {
		return err
	}
	res.Status = "ok"
	return nil
}
