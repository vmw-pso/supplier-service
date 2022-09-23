package data

import (
	"context"
	"database/sql"
	"time"
)

const dbTimeout = 3 * time.Second // move to config

type Supplier struct {
	ID   int    `json:"id"`
	Name string `json:"name:"`
}

func (s *Supplier) GetAll(db *sql.DB) ([]*Supplier, error) {
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

var id int

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
