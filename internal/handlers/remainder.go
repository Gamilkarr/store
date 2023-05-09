package handlers

import (
	"context"
	"time"

	"github.com/Gamilkarr/store/internal/models"
)

type RemainderRequest struct {
	StoreID int64 `json:"store_id"`
}

type RemainderResponse struct {
	Items []models.Item `json:"items"`
}

func (s *Store) Remainder(req RemainderRequest, res *RemainderResponse) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*250)
	defer cancel()
	res.Items, err = s.Repository.GetItemsQuantityOnStore(ctx, req.StoreID)
	return err
}
