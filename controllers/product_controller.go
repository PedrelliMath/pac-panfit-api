package controllers

import "github.com/gin-gonic/gin"

type ProductController struct{
  prefix string
  handler func(*gin.Context)
}

