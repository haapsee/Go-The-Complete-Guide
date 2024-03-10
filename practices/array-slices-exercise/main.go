package main

import (
  "fmt"
)

type product struct {
  id int
  title string
  price float64
}

func main() {
  // 1)
  hobbies := [3]string{"Football", "Hockey", "Floorball"}
  fmt.Println(hobbies)

  // 2)
  fmt.Println(hobbies[0])
  fmt.Println(hobbies[1:])

  // 3)
  hobbiesSlice := hobbies[0:2]
  fmt.Println(hobbiesSlice)

  // 4)
  hobbiesReSlice := hobbiesSlice[1:3]
  fmt.Println(hobbiesReSlice)

  // 5)
  courseGoals := []string{"Pass the course", "Learn go"}

  // 6)
  courseGoals[1] = "Learn to program with go"
  courseGoals = append(courseGoals, "Earn certificate")

  // 7)
  product0 := product{
    id: 0,
    title: "sausage",
    price: 0.99,
  }
  product1 := product{
    id: 1,
    title: "beef",
    price: 19.99,
  }
  products := []product{product0, product1}

  product2 := product{
    id: 2,
    title: "pork",
    price: 11.99,
  }
  products = append(products, product2)

  fmt.Println(products)
}

