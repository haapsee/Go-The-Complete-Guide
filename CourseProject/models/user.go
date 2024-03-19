package models

import (
  "errors"
  "example.com/event-booking/db"
  "example.com/event-booking/utils"
)

type User struct {
  ID        int64
  Email     string `binding:"required"`
  Password  string `binding:"required"`
}

func (u *User) ValidateCredentials() error {
  query := "SELECT id, password FROM users WHERE email = ?"

  row := db.DB.QueryRow(query, u.Email)

  var retrievedPassword string
  err := row.Scan(&u.ID, &retrievedPassword)

  if err != nil {
    return errors.New("Invalid credentials")
  }

  if !utils.CheckPasswordHash(u.Password, retrievedPassword) {
    return errors.New("Invalid credentials")
  }

  return nil
}

func (u User) Save() error {
  query := "INSERT INTO users(email, password) VALUES (?, ?)"
  statement, err := db.DB.Prepare(query)

  if err != nil {
    return err
  }

  defer statement.Close()

  hashedPassword, err := utils.HashPassword(u.Password)

  if err != nil {
    return err
  }

  result, err := statement.Exec(u.Email, hashedPassword)

  if err != nil {
    return err
  }

  userID, err := result.LastInsertId()
  u.ID = userID

  return err
}
