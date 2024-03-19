package models

import (
  "time"
  "example.com/event-booking/db"
)

type Event struct {
  ID          int64
  Name        string `binding:"required"`
  Description string `binding:"required"`
  Location    string `binding:"required"`
  DateTime    time.Time `binding:"required"`
  UserID      int64
}

func (event *Event) Save() error {
  query := `
    INSERT INTO events(name, description, location, datetime, user_id)
    VALUES (?, ?, ?, ?, ?)
  `
  statement, err := db.DB.Prepare(query)
  if err != nil {
    return err
  }
  defer statement.Close()

  result, err := statement.Exec(
    event.Name,
    event.Description,
    event.Location,
    event.DateTime,
    event.UserID,
  )
  if err != nil {
    return err
  }
  id, err := result.LastInsertId()
  event.ID = id
  return err
}

func GetAllEvents() ([]Event, error) {
  query := `SELECT * FROM events`
  results, err := db.DB.Query(query)

  if err != nil {
    return nil, err
  }
  defer results.Close()

  var events []Event

  for results.Next() {
    var event Event
    err := results.Scan(
      &event.ID,
      &event.Name,
      &event.Description,
      &event.Location,
      &event.DateTime,
      &event.UserID,
    )
    if err != nil {
      return nil, err
    }
    events = append(events, event)
  }

  return events, nil
}

func GetEventByID(id int64) (*Event, error){
  query := "SELECT * FROM EVENTS where ID = ?"
  row := db.DB.QueryRow(query, id)

  var event Event
  err := row.Scan(
    &event.ID,
    &event.Name,
    &event.Description,
    &event.Location,
    &event.DateTime,
    &event.UserID,
  )

  if err != nil {
    return nil, err
  }
  return &event, nil
}

func (event Event) Update() error {
  query := `
    UPDATE events
    SET name = ?, description = ?, location = ?, datetime = ?
    WHERE id = ?
  `

  statement, err := db.DB.Prepare(query)

  if err != nil {
    return err
  }

  defer statement.Close()

  _, err = statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

  return err
}

func (event Event) Delete() error {
  query := `
    DELETE FROM events WHERE id = ?
  `

  statement, err := db.DB.Prepare(query)

  if err != nil {
    return err
  }

  defer statement.Close()

  _, err = statement.Exec(event.ID)

  return err
}

func (event Event) Register(userID int64) error {
  query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"

  statement, err := db.DB.Prepare(query)

  if err != nil {
    return err
  }

  defer statement.Close()

  _, err = statement.Exec(event.ID, userID)
  return err
}

func (event Event) CancelRegistration(userID int64) error {
  query := "DELETE FROM registrations WHERE user_id = ? AND event_id = ?"

  statement, err := db.DB.Prepare(query)

  if err != nil {
    return err
  }

  defer statement.Close()

  _, err = statement.Exec(event.ID, userID)
  return err
}
