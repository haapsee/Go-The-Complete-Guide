package routes

import (
  "strconv"
  "net/http"

  "example.com/event-booking/models"

  "github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
  eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
    return
  }

  event, err := models.GetEventByID(eventID)
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
    return
  }
  context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
  events, err := models.GetAllEvents()
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
    return
  }
  context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
  var event models.Event
  err := context.ShouldBindJSON(&event)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
  }

  userID := context.GetInt64("userid")
  event.UserID = userID

  err = event.Save()

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event."})
    return
  }

  context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
  eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
    return
  }

  event, err := models.GetEventByID(eventID)
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
    return
  }

  userID := context.GetInt64("userid")

  if userID != event.UserID {
    context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
    return
  }

  var updatedEvent models.Event
  err = context.ShouldBindJSON(&updatedEvent)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
  }

  updatedEvent.ID = eventID
  err = updatedEvent.Update()

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
    return
  }

  context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func deleteEvent(context *gin.Context) {
  eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
    return
  }

  event, err := models.GetEventByID(eventID)
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
    return
  }

  userID := context.GetInt64("userid")

  if userID != event.UserID {
    context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
    return
  }

  err = event.Delete()

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
    return
  }

  context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}
