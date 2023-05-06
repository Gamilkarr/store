package services

import (
	"errors"
)

type ItemForUnreserved struct {
	ItemID                int64 `json:"item_id"`
	QuantityForUnreserved int64 `json:"quantity_for_reserved"`
}

func (s *StoreService) Unreserved(storeID int64, items map[int64]int64) error {
	if !s.Repository.IsStoreAvailable(storeID) {
		return errors.New("store is not available")
	}
	if err := s.Repository.ItemUnreserved(storeID, items); err != nil {
		return errors.New("error reserved")
	}
	return nil
}
