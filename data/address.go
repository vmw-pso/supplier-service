package data

import (
	"context"
	"database/sql"
	"time"
)

const dbTimeout = 3 * time.Second // move to config

type Address struct {
	ID             int    `json:"id"`
	Building       string `json:"building,omitempty"`
	UnitFloor      string `json:"unitFloor,omitempty"`
	StreetNumber   string `json:"streetNumber"`
	StreetName     string `json:"streetName"`
	City           string `json:"city"`
	ZipOrPostcode  string `json:"zipOrPostcode"`
	StateTerritory string `json:"stateTerritory"`
	Country        string `json:"country"`
}

func (a *Address) GetAll(db *sql.DB) ([]*Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, building, unit_floor, street_number, street_name, city, 
	          zip_or_postcode, state_territory, country FROM addresses`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []*Address

	for rows.Next() {
		var address Address
		err := rows.Scan(
			&address.ID,
			&address.Building,
			&address.UnitFloor,
			&address.StreetNumber,
			&address.StreetName,
			&address.City,
			&address.ZipOrPostcode,
			&address.StateTerritory,
			&address.Country,
		)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, &address)
	}

	return addresses, nil
}

func (a *Address) GetById(db *sql.DB, id int) (*Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, building, unit_floor, street_number, street_name, city, 
	          zip_or_postcode, state_territory, country 
			  FROM addresses WHERE id=$1`

	var address Address
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&address.ID,
		&address.Building,
		&address.UnitFloor,
		&address.StreetNumber,
		&address.StreetName,
		&address.City,
		&address.ZipOrPostcode,
		&address.StateTerritory,
		&address.Country,
	)
	if err != nil {
		return nil, err
	}

	return &address, nil
}

var id int

func (s *Address) Insert(address Address, db *sql.DB) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `INSERT INTO addresses 
	         (building, unit_floor, street_number, street_name, city, zip_or_postcode, state_territory, country) 
	         values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err := db.QueryRowContext(
		ctx,
		stmt,
		address.Building,
		address.UnitFloor,
		address.StreetNumber,
		address.StreetName,
		address.City,
		address.ZipOrPostcode,
		address.Country).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (a *Address) Update(address Address, db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `UPDATE addresses SET 
	         building=$1, 
			 unit_floor = $2
			 street_number=$3, 
			 street_name=$4, 
			 city=$5, 
			 zip_or_postcode=$6, 
			 state_territory=$7, 
			 country=$8 
			 WHERE id=$9`

	_, err := db.ExecContext(
		ctx,
		stmt,
		address.Building,
		address.UnitFloor,
		address.StreetNumber,
		address.StreetName,
		address.City,
		address.ZipOrPostcode,
		address.StateTerritory,
		address.Country,
		address.ID)
	return err
}

func (a *Address) Delete(id int, db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `DELETE FROM addresses WHERE id=$1`

	_, err := db.ExecContext(ctx, stmt, id)
	return err
}
