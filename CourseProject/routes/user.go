package routes

import (
  "fmt"
  "net/http"

  "example.com/event-booking/models"
  "example.com/event-booking/utils"

  "github.com/gin-gonic/gin"
)

func signin(context *gin.Context) {
  var user models.User
  err := context.ShouldBindJSON(&user)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
    return
  }

  err = user.ValidateCredentials()

  if err != nil {
    fmt.Println(err)
    context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
    return
  }

  token, err := utils.GenerateToken(user.Email, user.ID)

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
    return
  }

  context.JSON(http.StatusOK, gin.H{"message": "Login succesfull!", "token": token})
}

func signup(context *gin.Context) {
  var user models.User
  err := context.ShouldBindJSON(&user)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
    return
  }

  err = user.Save()

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
    return
  }

  context.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})
}
