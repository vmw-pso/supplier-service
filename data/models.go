package data

import "time"

const dbTimeout = 3 * time.Second // move to config

var id int

type Models struct {
	Supplier Supplier
	Address  Address
}

func New() *Models {
	return &Models{
		Supplier: Supplier{},
		Address:  Address{},
	}
}
