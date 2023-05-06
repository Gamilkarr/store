package repository

func (r *Repository) ItemReserved(storeID int64, items map[int64]int64) error {
	return nil
}

func (r *Repository) ItemUnreserved(storeID int64, items map[int64]int64) error {
	return nil
}

func (r *Repository) GetItemsQuantityOnStore(storeID int64) ([]map[string]int64, error) {
	return nil, nil
}
