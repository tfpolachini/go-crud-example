package repository

import (
	"database/sql"
	"time"

	"github.com/tfpolachini/go-crud-example/domain/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) Save(p *model.Product) (*model.Product, error) {
	stmt, err := r.db.Prepare(`INSERT INTO products(id, name, status, created_at, updated_at) VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.ID(), p.Name(), p.Status().String(), p.CreatedAt(), p.UpdatedAt())
	if err != nil {
		return nil, err
	}

	return r.FindById(p.ID())
}

func (r *ProductRepository) FindByName(n string) (*model.Product, error) {
	stmt, err := r.db.Prepare(`SELECT p.id, p.name, p.status, p.created_at, p.updated_at FROM products p WHERE p.name = ?`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var id, name, statusStr string
	var createdAt, updatedAt time.Time

	err = stmt.QueryRow(n).Scan(&id, &name, &statusStr, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	status, err := model.NewStatusFromString(statusStr)
	if err != nil {
		return nil, err
	}

	p, err := model.UnmarshalProductFromDatabase(id, name, status, createdAt, updatedAt)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *ProductRepository) CountByName(n string) (int, error) {
	stmt, err := r.db.Prepare(`SELECT COUNT(*) FROM products p WHERE p.name = ?`)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	var count int

	err = stmt.QueryRow(n).Scan(&count)
	if err == sql.ErrNoRows {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *ProductRepository) FindById(id string) (*model.Product, error) {
	stmt, err := r.db.Prepare(`SELECT id, name, status, created_at, updated_at FROM products p WHERE p.id = ?`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var uuid, name, statusStr string
	var createdAt, updatedAt time.Time

	err = stmt.QueryRow(id).Scan(&uuid, &name, &statusStr, &createdAt, &updatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	status, err := model.NewStatusFromString(statusStr)
	if err != nil {
		return nil, err
	}

	p, err := model.UnmarshalProductFromDatabase(uuid, name, status, createdAt, updatedAt)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *ProductRepository) CountById(id string) (int, error) {
	stmt, err := r.db.Prepare(`SELECT COUNT(*) FROM products p WHERE p.id = ?`)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	var count int

	err = stmt.QueryRow(id).Scan(&count)
	if err == sql.ErrNoRows {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return count, nil
}
