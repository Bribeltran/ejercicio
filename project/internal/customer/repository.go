package customer

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
	Save(ctx context.Context, s domain.Customer) (int, error)
	Update(ctx context.Context, s domain.Customer) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Exists(ctx context.Context, last_name string) bool {
	query := "SELECT last_name FROM customer WHERE last_name=?;"
	row := r.db.QueryRow(query, last_name)
	err := row.Scan(&last_name)
	return err == nil
}

func (r *repository) Save(ctx context.Context, s domain.Customer) (int, error) {
	query := "INSERT INTO customer (last_name, first_name, condition) VALUES (?, ?, ?);"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&s.LastName, &s.FirstName, &s.Condition)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
func (r *repository) Update(ctx context.Context, s domain.Customer) error {
	query := "UPDATE customer SET last_name=?, first_name=?, condition=? WHERE id=?;"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&s.LastName, &s.FirstName, &s.Condition, &s.ID)
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
