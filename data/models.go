package data

type models struct {
	Supplier Supplier
	Address  Address
}

func New() *models {
	return &models{
		Supplier: Supplier{},
		Address:  Address{},
	}
}
