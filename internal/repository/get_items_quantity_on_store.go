package repository

import (
	"context"
	"errors"

	"github.com/Gamilkarr/store/internal/models"
)

func (r *Repository) GetItemsQuantityOnStore(ctx context.Context, storeID int64) ([]models.Item, error) {
	q := `SELECT item_id, name, reserved_item, item_quantity - reserved_item AS availability 
		  FROM available join items i on i.id = available.item_id 
		  WHERE store_id = $1`

	tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	var isAvailable bool

	err = tx.QueryRow(ctx, isStoreAvailable, storeID).Scan(&isAvailable)
	if err != nil {
		return nil, err
	}

	if !isAvailable {
		return nil, errors.New("store is not available")
	}

	var items []models.Item
	rows, _ := tx.Query(ctx, q, storeID)
	for rows.Next() {
		var i models.Item
		err = rows.Scan(&i.ItemID, &i.Name, &i.ReservedQuantity, &i.AvailableQuantity)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return items, nil
}
