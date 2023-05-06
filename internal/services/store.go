package services

type StoreService struct {
	repository Repository
}

type Repository interface {
	IsStoreAvailable(storeID int64) bool
	ItemReserved(storeID int64, items []ItemForReserved) error
	ItemUnreserved(storeID int64, items []ItemForUnreserved) error
	GetItemsQuantityOnStore(storeID int64) ([]Item, error)
}
