package handlers

type Store struct {
	service Services
}

type Services interface {
	Remainder(storeID int64) ([]map[string]int64, error)
	Reserved(storeID int64, items map[int64]int64) error
	Unreserved(storeID int64, items map[int64]int64) error
}
