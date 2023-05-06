package services

import "errors"

type Item struct {
	ItemID            int64 `json:"item_id"`
	AvailableQuantity int64 `json:"available_quantity"`
	ReservedQuantity  int64 `json:"reserved_quantity"`
}

func (s *StoreService) Remainder(storeID int64) ([]Item, error) {
	if !s.repository.IsStoreAvailable(storeID) {
		return nil, errors.New("store is not available")
	}
	var err error
	items, err := s.repository.GetItemsQuantityOnStore(storeID)
	if err != nil {
		return nil, err
	}
	return items, nil
}
