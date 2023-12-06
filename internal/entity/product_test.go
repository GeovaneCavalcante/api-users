package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Coca-Cola", 5)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Coca-Cola", product.Name)
	assert.Equal(t, 5.0, product.Price)
	assert.NotEmpty(t, product.CreatedAt)
}

func TestProductWhenNameIsRequerid(t *testing.T) {
	product, err := NewProduct("", 5.0)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}
func TestProductWhenPriceIsRequerid(t *testing.T) {
	product, err := NewProduct("Coca-Cola", 0)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Coca-Cola", -2)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidePrice, err)
}
