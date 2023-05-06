package services

type StoreService struct {
	Repository Repository
}

type Repository interface {
	IsStoreAvailable(storeID int64) bool
	ItemReserved(storeID int64, items map[int64]int64) error
	ItemUnreserved(storeID int64, items map[int64]int64) error
	GetItemsQuantityOnStore(storeID int64) ([]map[string]int64, error)
}
