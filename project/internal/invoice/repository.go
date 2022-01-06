package invoice

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
	Save(ctx context.Context, s domain.Invoice) (int, error)
	Update(ctx context.Context, s domain.Invoice) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Exists(ctx context.Context, idCustomer int) bool {
	query := "SELECT id_customer FROM invoice WHERE id_customer=?;"
	row := r.db.QueryRow(query, idCustomer)
	err := row.Scan(&idCustomer)
	return err == nil
}

func (r *repository) Save(ctx context.Context, s domain.Invoice) (int, error) {
	query := "INSERT INTO invoice (data_time, id_customer, total) VALUES (?, ?, ?);"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&s.DateTime, &s.IdCustomer, &s.Total)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Invoice) error {
	query := "UPDATE invoice SET data_time=?, id_customer=?, total=? WHERE id=?;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&s.DateTime, &s.IdCustomer, &s.Total, &s.ID)
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
