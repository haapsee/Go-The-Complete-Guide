package utils

import (
  "fmt"
  "time"
  "errors"

  "github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretkey"

func VerifyToken(token string) (int64, error) {
  parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error){
    _, ok := token.Method.(*jwt.SigningMethodHMAC)

    if !ok {
      return 0, errors.New("Unexpected signing method.")
    }

    return []byte(secretKey), nil
  })

  if err != nil {
    return 0, errors.New("Could not parse token.")
  }

  tokenIsValid := parsedToken.Valid

  if !tokenIsValid {
    return 0, errors.New("Invalid token.")
  }

  claims, ok := parsedToken.Claims.(jwt.MapClaims)

  if !ok {
    return 0, errors.New("Invalid token claims.")
  }

  fmt.Println(claims)

  userID := int64(claims["userID"].(float64))

  return userID, nil
}

func GenerateToken(email string, userID int64) (string, error) {
  token := jwt.NewWithClaims(
    jwt.SigningMethodHS256,
    jwt.MapClaims{
      "email": email,
      "userID": userID,
      "exp": time.Now().Add(time.Hour * 2).Unix(),
    },
  )

  return token.SignedString([]byte(secretKey))
}
