package model

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type ProductRepositoryInterface interface {
	Save(product *Product) (*Product, error)
	CountByName(name string) (int, error)
	FindByName(name string) (*Product, error)
	FindById(id string) (*Product, error)
}

const (
	ENABLED  = "ENABLED"
	DISABLED = "DISABLED"
)

type Product struct {
	ID        string
	Name      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProduct(name string) (*Product, error) {
	product := Product{
		Name: name,
	}

	product.ID = uuid.NewV4().String()
	product.Status = ENABLED
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	return &product, nil
}

func (p *Product) Edit(name string) (*Product, error) {

	if p.Status == DISABLED {
		return p, errors.New("you can't edit a disabled product")
	}

	p.Name = name
	p.UpdatedAt = time.Now()

	return p, nil
}

func (p *Product) Enable() *Product {
	p.Status = ENABLED
	p.UpdatedAt = time.Now()

	return p
}

func (p *Product) Disable() *Product {
	p.Status = DISABLED
	p.UpdatedAt = time.Now()

	return p
}
