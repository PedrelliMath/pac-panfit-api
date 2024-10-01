package interfaces

import "panfit_api/models"

type Product = models.Product

type ProductRepository interface{
  Save_product(p Product)Product
  Get_product_by_id(id int)(Product, err error)
  Get_all_products()*[]Product
}
