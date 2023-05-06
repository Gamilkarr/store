package services

import (
	"errors"
)

type ItemForUnreserved struct {
	ItemID                int64 `json:"item_id"`
	QuantityForUnreserved int64 `json:"quantity_for_reserved"`
}

func (s *StoreService) Unreserved(storeID int64, items []ItemForUnreserved) error {
	if !s.repository.IsStoreAvailable(storeID) {
		return errors.New("store is not available")
	}
	if err := s.repository.ItemUnreserved(storeID, items); err != nil {
		return errors.New("error reserved")
	}
	return nil
}
