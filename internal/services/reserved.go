package services

import (
	"errors"
)

type ItemForReserved struct {
	ItemID              int64 `json:"item_id"`
	QuantityForReserved int64 `json:"quantity_for_reserved"`
}

func (s *StoreService) Reserved(storeID int64, items []ItemForReserved) error {
	if !s.repository.IsStoreAvailable(storeID) {
		return errors.New("store is not available")
	}
	if err := s.repository.ItemReserved(storeID, items); err != nil {
		return errors.New("error reserved")
	}
	return nil
}
