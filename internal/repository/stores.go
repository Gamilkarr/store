package repository

import (
	"context"
	"fmt"
)

func (r *Repository) IsStoreAvailable(storeID int64) bool {
	q := `SELECT availability FROM stores WHERE id=$1`

	var ave bool

	err := r.Conn.QueryRow(context.Background(), q, storeID).Scan(&ave)
	if err != nil {
		fmt.Print(err)
	}

	return ave
}
