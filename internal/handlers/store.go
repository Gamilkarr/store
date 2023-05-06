package handlers

import "github.com/Gamilkarr/store/internal/services"

type Store struct {
	service Services
}

type Services interface {
	Remainder(storeID int64) ([]services.Item, error)
	Reserved(storeID int64, items []services.ItemForReserved) error
	Unreserved(storeID int64, items []services.ItemForUnreserved) error
}
