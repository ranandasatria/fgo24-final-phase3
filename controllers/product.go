package controllers

import (
	"fmt"
	"net/http"
	"test-fase-3/dto"
	"test-fase-3/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateProduct(c *gin.Context) {
	claims := c.MustGet("user").(jwt.MapClaims)
	if role, ok := claims["role"].(string); !ok || role != "admin" {
		c.JSON(http.StatusForbidden, dto.Response{
			Success: false,
			Message: "Only admin can add products",
		})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	if err := models.CreateProduct(input); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to create product",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Product created",
		Results: input,
	})
}

func GetProducts(c *gin.Context) {
	products, err := models.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to fetch products",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "List of products",
		Results: products,
	})
	fmt.Printf("%+v\n", products)

}
