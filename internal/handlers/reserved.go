package handlers

import "github.com/Gamilkarr/store/internal/services"

type ReservedRequest struct {
	StoreID int64                      `json:"store_id"`
	Items   []services.ItemForReserved `json:"items_ids"`
}

type ReservedResponse struct {
	Status string `json:"status"`
}

func (s *Store) Reserved(req ReservedRequest, res *ReservedResponse) error {
	if err := s.service.Reserved(req.StoreID, req.Items); err != nil {
		return err
	}
	res.Status = "ok"
	return nil
}
