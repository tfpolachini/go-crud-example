package repository

import (
	"database/sql"

	"github.com/tfpolachini/go-crud-example/domain/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) Save(product *model.Product) (*model.Product, error) {
	stmt, err := r.db.Prepare(`INSERT INTO products(id, name, status, created_at, updated_at) VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Status, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return r.FindById(product.ID)
}

func (r *ProductRepository) FindByName(name string) (*model.Product, error) {
	stmt, err := r.db.Prepare(`SELECT p.id, p.name, p.status, p.created_at, p.updated_at FROM products p WHERE p.name = ?`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var product model.Product

	err = stmt.QueryRow(name).Scan(&product.ID, &product.Name, &product.Status, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) CountByName(name string) (int, error) {
	stmt, err := r.db.Prepare(`SELECT COUNT(*) FROM products WHERE name = ?`)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	var count int

	err = stmt.QueryRow(name).Scan(&count)
	if err == sql.ErrNoRows {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *ProductRepository) FindById(id string) (*model.Product, error) {
	var product model.Product

	stmt, err := r.db.Prepare(`SELECT id, name, status, created_at, updated_at FROM products WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Status, &product.CreatedAt, &product.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) CountById(id string) (int, error) {
	stmt, err := r.db.Prepare(`SELECT COUNT(*) FROM products WHERE id = ?`)
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
