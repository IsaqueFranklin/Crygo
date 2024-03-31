//The most important thing for encryption is to generate random numbers, whithout it the encrypted data would be very predictable.

package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  //The above line gives us the current time to the second and the seed parameter for rand changes each time.

  fmt.Println(rand.Intn(100))
}
