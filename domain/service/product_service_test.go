package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tfpolachini/go-crud-example/domain/model"
	"github.com/tfpolachini/go-crud-example/domain/model/mocks"
)

func TestProductShouldNotBeCreatedBecauseCountByNameMethodReturnAnError(t *testing.T) {

	// SETUP
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockProductRepositoryInterface(ctrl)

	// GIVEN
	inputDto := CreateProductInputDto{"nome"}
	svc := NewProductService(m)
	m.EXPECT().CountByName(gomock.Any()).Return(0, errors.New("database error"))

	// WHEN
	_, err := svc.CreateProduct(inputDto)

	// THEN
	assert.EqualError(t, err, "database error")
}

func TestProductShouldNotBeCreatedBecauseAlreadyExists(t *testing.T) {

	// SETUP
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockProductRepositoryInterface(ctrl)

	// GIVEN
	name := "Product Name"
	inputDto := CreateProductInputDto{name}
	svc := NewProductService(m)
	m.EXPECT().CountByName(gomock.Any()).Return(1, nil)

	// WHEN
	product, err := svc.CreateProduct(inputDto)

	// THEN
	assert.ErrorIs(t, err, ErrResourceAlreadyExists)
	assert.Nil(t, product)
}

func TestProductShouldBeCreatedAndSaved(t *testing.T) {

	// SETUP
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockProductRepositoryInterface(ctrl)

	// GIVEN
	name := "Product Name"
	product, _ := model.NewProduct(name)
	inputDto := CreateProductInputDto{name}
	svc := NewProductService(m)
	m.EXPECT().CountByName(gomock.Any()).Return(0, nil)
	m.EXPECT().Save(gomock.Any()).Return(product, nil)

	// WHEN
	p, err := svc.CreateProduct(inputDto)

	// THEN
	assert.Nil(t, err)
	assert.Equal(t, product.ID, p.ID)
}

func TestProductShouldNotBeCreatedBecauseSaveMethodReturnAnError(t *testing.T) {

	// SETUP
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockProductRepositoryInterface(ctrl)

	// GIVEN
	inputDto := CreateProductInputDto{"nome"}
	svc := NewProductService(m)
	m.EXPECT().CountByName(gomock.Any()).Return(0, nil)
	m.EXPECT().Save(gomock.Any()).Return(nil, errors.New("database error"))

	// WHEN
	_, err := svc.CreateProduct(inputDto)

	// THEN
	assert.EqualError(t, err, "database error")
}

func TestProductShouldNotBeEditedBecauseNotFound(t *testing.T) {

	// SETUP
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockProductRepositoryInterface(ctrl)

	// GIVEN
	id := "Product id"
	name := "Product name"
	inputDto := &UpdateProductInputDto{id, name}
	svc := NewProductService(m)
	m.EXPECT().FindById(inputDto.ID).Return(nil, nil)

	// WHEN
	product, err := svc.UpdateProduct(inputDto)

	// THEN
	assert.Nil(t, product)
	assert.True(t, errors.Is(err, ErrResourceNotFound))
}
