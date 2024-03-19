package routes

import (
  "strconv"
  "net/http"

  "example.com/event-booking/models"

  "github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
  userID := context.GetInt64("userid")
  eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
    return
  }

  event, err := models.GetEventByID(eventID)

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
    return
  }

  err = event.Register(userID)

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user"})
    return
  }

  context.JSON(http.StatusCreated, gin.H{"message": "User registered!"})
}

func cancelRegistration(context *gin.Context) {
  userID := context.GetInt64("userid")
  eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
    return
  }

  event, err := models.GetEventByID(eventID)

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
    return
  }

  err = event.CancelRegistration(userID)

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
    return
  }

  context.JSON(http.StatusCreated, gin.H{"message": "Cancelled!"})
}
