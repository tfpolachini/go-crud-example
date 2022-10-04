package service

import (
	"time"

	"github.com/tfpolachini/go-crud-example/domain/model"
)

type productOutputDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateProductInputDto struct {
	Name string `json:"name"`
}

type CreateProductOutputDto struct {
	productOutputDto
}

func NewCreateProductOutputDto(product *model.Product) *CreateProductOutputDto {
	dto := &CreateProductOutputDto{}

	dto.ID = product.ID
	dto.Name = product.Name
	dto.Status = product.Status
	dto.CreateAt = product.CreatedAt
	dto.UpdatedAt = product.UpdatedAt

	return dto
}

type UpdateProductInputDto struct {
	ID      string `json:"id"`
	NewName string `json:"new_name"`
}

type UpdateProductOutputDto struct {
	productOutputDto
}

func NewUpdateProductOutputDto(product *model.Product) *UpdateProductOutputDto {
	dto := &UpdateProductOutputDto{}

	dto.ID = product.ID
	dto.Name = product.Name
	dto.Status = product.Status
	dto.CreateAt = product.CreatedAt
	dto.UpdatedAt = product.UpdatedAt

	return dto
}
