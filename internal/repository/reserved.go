package repository

import (
	"context"
	"errors"
	"github.com/Gamilkarr/store/internal/models"
)

func (r *Repository) ItemReserved(ctx context.Context, storeID int64, items []models.Reserved) error {
	q := `UPDATE available
	  	  SET reserved_item = $3 + reserved_item
	  	  WHERE store_id = $1 AND item_id = $2 AND item_quantity - reserved_item >= $3
		  RETURNING item_id`

	tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	var isAvailable bool

	err = tx.QueryRow(ctx, isStoreAvailable, storeID).Scan(&isAvailable)
	if err != nil {
		return err
	}

	if !isAvailable {
		return errors.New("store is not available")
	}

	for _, item := range items {
		var resItemID int64
		err := tx.QueryRow(ctx, q, storeID, item.ItemID, item.Quantity).Scan(&resItemID)
		if err != nil {
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
