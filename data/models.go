package data

type models struct {
	Supplier Supplier
}

func New() *models {
	return &models{
		Supplier: Supplier{},
	}
}
