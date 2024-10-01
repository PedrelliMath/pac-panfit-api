package services

import (
	"panfit_api/models"
	"panfit_api/repositories"

	"github.com/gin-gonic/gin"
)

type Product = models.Product
type ProductRepository = repositories.ProductRepositoryMemory

type ProductService struct{
  repostiory ProductRepository
}

func(p *ProductService) Create_product(product []byte)Product{
}
