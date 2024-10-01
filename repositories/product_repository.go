package repositories

import (
	"errors"
	"panfit_api/models"
)

type Product = models.Product

type ProductRepositoryMemory struct{
  ProductList []Product
} 

func(p *ProductRepositoryMemory) Save_product(produto Product)Product{
  (*p).ProductList = append((*p).ProductList, produto)
  return produto
}

func(p *ProductRepositoryMemory) Get_product_by_id(product_id int)(Product, error){
  for _, produto := range (*p).ProductList{
    if produto.ID == product_id{
      return produto, nil
    }
  }
  return Product{}, errors.New("Product not found")
}

func(p *ProductRepositoryMemory) Get_all_products()[]Product{
  product_list := make([]Product, len((*p).ProductList))
  return product_list
}
