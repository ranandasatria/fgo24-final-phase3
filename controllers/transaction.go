package controllers

import (
	"fmt"
	"net/http"
	"test-fase-3/dto"
	"test-fase-3/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetTransactions(c *gin.Context) {
	transactions, err := models.GetAllTransactions()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to fetch transactions",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "List of transactions",
		Results: transactions,
	})
}

func CreateTransaction(c *gin.Context) {
	claims := c.MustGet("user").(jwt.MapClaims)
	if role, ok := claims["role"].(string); !ok || role != "admin" {
		c.JSON(http.StatusForbidden, dto.Response{
			Success: false,
			Message: "Only admin can add products",
		})
		return
	}

	var input dto.CreateTransactionDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	userIDFloat := claims["user_id"].(float64)
	userID := int(userIDFloat)

	tx := models.Transaction{
		ProdukID: input.ProdukID,
		UserID:   userID,
		Tipe:     input.Tipe,
		Jumlah:   input.Jumlah,
	}

	if err := models.CreateTransaction(tx); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Transaction created",
		Results: input,
	})
}
