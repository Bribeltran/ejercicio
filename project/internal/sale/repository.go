package sale

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
	Save(ctx context.Context, s domain.Sale) (int, error)
	Update(ctx context.Context, s domain.Sale) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Exists(ctx context.Context, idInvoice float64) bool {
	query := "SELECT id_invoice FROM sale WHERE id_invoice=?;"
	row := r.db.QueryRow(query, idInvoice)
	err := row.Scan(&idInvoice)
	return err == nil
}

func (r *repository) Save(ctx context.Context, s domain.Sale) (int, error) {
	query := "INSERT INTO sale (id_invoice, id_product, quantity) VALUES (?, ?, ?);"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&s.IdInvoice, &s.IdProduct, &s.Quantity)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Sale) error {
	query := "UPDATE sale SET id_invoice=?, id_product=?, Quantity=?  WHERE id=?;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&s.IdInvoice, &s.IdProduct, &s.Quantity, &s.ID)
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
