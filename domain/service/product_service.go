package service

import (
	"fmt"

	"github.com/tfpolachini/go-crud-example/domain/model"
)

type ProductServiceInterface interface {
	CreateProduct(inputDto CreateProductInputDto) (*CreateProductOutputDto, error)
}

type ProductService struct {
	repository model.ProductRepositoryInterface
}

func NewProductService(repository model.ProductRepositoryInterface) *ProductService {
	return &ProductService{repository}
}

func (svc *ProductService) CreateProduct(inputDto CreateProductInputDto) (*CreateProductOutputDto, error) {
	count, err := svc.repository.CountByName(inputDto.Name)
	if err != nil {
		return nil, err
	}

	if count != 0 {
		return nil, fmt.Errorf("product with name %s already exists: %w", inputDto.Name, ErrResourceAlreadyExists)
	}

	product, err := model.NewProduct(inputDto.Name)
	if err != nil {
		return nil, err
	}

	product, err = svc.repository.Save(product)
	if err != nil {
		return nil, err
	}

	outputDto := NewCreateProductOutputDto(product)

	return outputDto, nil
}

func (svc *ProductService) UpdateProduct(inputDto *UpdateProductInputDto) (*UpdateProductOutputDto, error) {
	product, err := svc.repository.FindById(inputDto.ID)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, fmt.Errorf("product with id %s not found: %w", inputDto.ID, ErrResourceNotFound)
	}

	p, err := product.Edit(inputDto.NewName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err.Error(), ErrInvalidAction)
	}

	p, err = svc.repository.Save(p)
	if err != nil {
		return nil, err
	}

	outputDto := NewUpdateProductOutputDto(p)

	return outputDto, nil
}
