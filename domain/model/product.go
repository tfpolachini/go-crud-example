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

type Product struct {
	id        string
	name      string
	status    Status
	createdAt time.Time
	updatedAt time.Time
}

func UnmarshalProductFromDatabase(id, name string, status Status, createdAt, updatedAt time.Time) (*Product, error) {
	return &Product{
		id:        id,
		name:      name,
		status:    status,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func NewProduct(name string) (*Product, error) {
	if name == "" {
		return nil, errors.New("product name can not be empty")
	}

	product := Product{
		name: name,
	}

	product.id = uuid.NewV4().String()
	product.status = ENABLED
	product.createdAt = time.Now()
	product.updatedAt = time.Now()

	return &product, nil
}

func (p Product) ID() string {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Status() Status {
	return p.status
}

func (p Product) CreatedAt() time.Time {
	return p.createdAt
}

func (p Product) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Product) Edit(name string) (*Product, error) {
	if p.status == DISABLED {
		return p, errors.New("you can't edit a disabled product")
	}

	p.name = name
	p.updatedAt = time.Now()

	return p, nil
}

func (p *Product) Enable() *Product {
	p.status = ENABLED
	p.updatedAt = time.Now()

	return p
}

func (p *Product) Disable() *Product {
	p.status = DISABLED
	p.updatedAt = time.Now()

	return p
}
