package services

import "errors"

type Item struct {
	ItemID            int64 `json:"item_id"`
	AvailableQuantity int64 `json:"available_quantity"`
	ReservedQuantity  int64 `json:"reserved_quantity"`
}

func (s *StoreService) Remainder(storeID int64) ([]map[string]int64, error) {
	if !s.Repository.IsStoreAvailable(storeID) {
		return nil, errors.New("store is not available")
	}
	items, err := s.Repository.GetItemsQuantityOnStore(storeID)
	if err != nil {
		return nil, err
	}
	return items, nil
}
