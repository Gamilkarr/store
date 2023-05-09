package handlers

import (
	"context"
	"time"

	"github.com/Gamilkarr/store/internal/models"
)

type ReservedRequest struct {
	StoreID          int64             `json:"store_id"`
	ItemsForReserved []models.Reserved `json:"items_for_reserved"`
}

type ReservedResponse struct {
	Status string `json:"status"`
}

func (s *Store) Reserved(req ReservedRequest, res *ReservedResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*250)
	defer cancel()
	if err := s.Repository.ItemReserved(ctx, req.StoreID, req.ItemsForReserved); err != nil {
		return err
	}
	res.Status = "ok"
	return nil
}
