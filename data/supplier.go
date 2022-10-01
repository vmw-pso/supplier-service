package data

import (
	"context"
	"database/sql"
)

type Supplier struct {
	ID   int    `json:"id"`
	Name string `json:"name:"`
}

type SupplierWithAddress struct {
	Supplier Supplier `json:"supplier"`
	Address  Address  `json:"address"`
}

func (s *Supplier) GetNames(db *sql.DB) ([]*Supplier, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name FROM suppliers`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []*Supplier

	for rows.Next() {
		var supplier Supplier
		err := rows.Scan(
			&supplier.ID,
			&supplier.Name,
		)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, &supplier)
	}

	return suppliers, nil
}

func (s *Supplier) GetAll(db *sql.DB) ([]SupplierWithAddress, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT suppliers.id, suppliers.name, address.id, address.building, 
	          address.unit_floor, address.street_number, address.street_name, 
			  address.city, address.zip_or_postcode, address.state_or_teritory, address.country 
			  from suppliers
			  JOIN supplieraddresses ON (supplieraddresses.supplier_id = suppliers.id)
			  JOIN addresses ON (address.id = supplieraddresses.address_id`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sa []SupplierWithAddress

	for rows.Next() {
		var s SupplierWithAddress
		err := rows.Scan(
			&s.Supplier.ID,
			&s.Supplier.Name,
			&s.Address.ID,
			&s.Address.Building,
			&s.Address.UnitFloor,
			&s.Address.StreetNumber,
			&s.Address.StreetName,
			&s.Address.City,
			&s.Address.ZipOrPostcode,
			&s.Address.StateTerritory,
			&s.Address.Country,
		)
		if err != nil {
			return nil, err
		}
		sa = append(sa, s)
	}
	return sa, nil
}

func (s *Supplier) GetById(db *sql.DB, id int) (*Supplier, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name FROM suppliers WHERE id=$1`

	var supplier Supplier
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&supplier.ID,
		&supplier.Name,
	)
	if err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (s *Supplier) Insert(supplier Supplier, db *sql.DB) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `INSERT INTO suppliers (name) values ($1) RETURNING id`

	err := db.QueryRowContext(ctx, stmt, supplier.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Supplier) Update(supplier Supplier, db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `UPDATE suppliers SET name=$1 WHERE id=$2`

	_, err := db.ExecContext(ctx, stmt, supplier.Name, supplier.ID)
	return err
}

func (s *Supplier) Delete(id int, db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `DELETE FROM suppliers WHERE id=$1`

	_, err := db.ExecContext(ctx, stmt, id)
	return err
}
