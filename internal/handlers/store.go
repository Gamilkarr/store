package handlers

import (
	"context"

	"github.com/Gamilkarr/store/internal/models"
)

type Store struct {
	Repository Repository
}

//go:generate mockgen -source=./store.go -destination=./store_mock.go -package=handlers
type Repository interface {
	ItemReserved(ctx context.Context, storeID int64, items []models.Reserved) error
	ItemUnreserved(ctx context.Context, storeID int64, items []models.Reserved) error
	GetItemsQuantityOnStore(ctx context.Context, storeID int64) ([]models.Item, error)
}
