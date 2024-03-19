package middlewares

import (
  "net/http"

  "example.com/event-booking/utils"
  "github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
  token := context.Request.Header.Get("Authorization")

  if token == "" {
    context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
    return
  }

  userID, err := utils.VerifyToken(token)

  if err != nil {
    context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
    return
  }

  context.Set("userid", userID)
  context.Next()
}
