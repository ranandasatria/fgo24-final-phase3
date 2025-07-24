package controllers

import (
	"fmt"
	"net/http"
	"test-fase-3/dto"
	"test-fase-3/models"
	"test-fase-3/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	var input dto.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	hashed, err := utils.HashString(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to hash password",
		})
		return
	}

	user := models.User{
		Nama: input.Nama,
		Email: input.Email,
		Password: hashed,
		Role: "user",
	}

	if err := models.CreateUser(user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "User created",
		Results: gin.H{
			"nama": user.Nama, 
			"email": user.Email,
			"role": user.Role}, 
	})
}

func Login(c *gin.Context){
	var input dto.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	user, err := models.FindUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "Wrong email or password",
		})
		return
	}

	if err := utils.CompareHash(user.Password, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "Wrong email or password",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Login success",
		Results: gin.H{
			"token": token,
		},
	})
}
