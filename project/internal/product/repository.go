package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Bribeltran/ejercicio/project/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("section not found")
)

// Repository encapsulates the storage of a section.
type Repository interface {
	Save(ctx context.Context, s domain.Product) (int, error)
	Update(ctx context.Context, s domain.Product) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Exists(ctx context.Context, price float64) bool {
	query := "SELECT price FROM product WHERE price=?;"
	row := r.db.QueryRow(query, price)
	err := row.Scan(&price)
	return err == nil
}

func (r *repository) Save(ctx context.Context, s domain.Product) (int, error) {
	query := "INSERT INTO product (description, price) VALUES (?, ?);"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&s.Description, &s.Price)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Product) error {
	query := "UPDATE product SET description=?, price=? WHERE id=?;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&s.Description, &s.Price, &s.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}
