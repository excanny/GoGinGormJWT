package controllers

import (
	"net/http"
  	"github.com/gin-gonic/gin"
	"ginjwt2/models"
	"ginjwt2/utils/token"
)

func CurrentUser(c *gin.Context){

	user_id, err := token.ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	u,err := models.GetUserByID(user_id)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","data":u})
}

type RegisterInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func Register(ctx *gin.Context){
	var input RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}

	user.Email = input.Email
	user.Password = input.Password

	_,err := user.SaveUser()

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message" : "Registration successful"})     
}

type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context){
	var input LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}

	user.Email = input.Email
	user.Password = input.Password

	token, err := models.LoginCheck(user.Email, user.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token":token})
}