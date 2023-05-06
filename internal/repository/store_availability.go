package repository

import (
	"context"
	"errors"
	"strconv"
)

func (r *Repository) ItemReserved(storeID int64, items map[int64]int64) error {
	q := `UPDATE store_availability
	  	  SET reserved_item = $3 + reserved_item
	  	  WHERE store_id = $1 AND item_id = $2 AND item_quantity - reserved_item >= $3
	  	  RETURNING item_id`

	tx, err := r.Conn.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	if !r.IsStoreAvailable(storeID) {
		return errors.New("store is not available")
	}
	for itemId, quantity := range items {
		var resItemID int64
		err := r.Conn.QueryRow(context.Background(), q, storeID, itemId, quantity).Scan(&resItemID)
		if err != nil {
			return err
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ItemUnreserved(storeID int64, items map[int64]int64) error {
	q := `UPDATE store_availability
		  SET reserved_item = reserved_item - $3
		  WHERE store_id = $1 AND item_id = $2 AND reserved_item >= $3
		  RETURNING item_id`
	tx, err := r.Conn.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	if !r.IsStoreAvailable(storeID) {
		return errors.New("store is not available")
	}

	for itemId, quantity := range items {
		var resItemID int64
		err := r.Conn.QueryRow(context.Background(), q, storeID, itemId, quantity).Scan(&resItemID)
		if err != nil {
			return err
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

type test struct {
	itemId int64
	name   string
	res    int64
	ave    int64
}

func (r *Repository) GetItemsQuantityOnStore(storeID int64) ([]map[string]string, error) {
	q := `SELECT item_id, name, reserved_item, item_quantity - reserved_item AS availability 
		  FROM store_availability join items i on i.id = store_availability.item_id 
		  WHERE store_id = $1`

	tx, err := r.Conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(context.Background())

	if !r.IsStoreAvailable(storeID) {
		return nil, errors.New("store is not available")
	}

	var items []test
	rows, _ := r.Conn.Query(context.Background(), q, storeID)
	for rows.Next() {
		var t test
		rows.Scan(&t.itemId, &t.name, &t.res, &t.ave)
		items = append(items, t)
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return nil, err
	}
	return structToMap(items), nil
}

func structToMap(tt []test) []map[string]string {
	var result []map[string]string
	for _, t := range tt {
		x := map[string]string{
			"itemID":            strconv.Itoa(int(t.itemId)),
			"name":              t.name,
			"availableQuantity": strconv.Itoa(int(t.ave)),
			"reservedQuantity":  strconv.Itoa(int(t.res)),
		}
		result = append(result, x)
	}
	return result
}
