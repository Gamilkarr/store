package handlers

import (
	"context"
	"time"

	"github.com/Gamilkarr/store/internal/models"
)

type UnreservedRequest struct {
	StoreID           int64             `json:"store_id"`
	ItemForUnreserved []models.Reserved `json:"items_for_unreserved"`
}

type UnreservedResponse struct {
	Status string `json:"status"`
}

func (s *Store) Unreserved(req UnreservedRequest, res *UnreservedResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*250)
	defer cancel()
	if err := s.Repository.ItemUnreserved(ctx, req.StoreID, req.ItemForUnreserved); err != nil {
		return err
	}
	res.Status = "ok"
	return nil
}
