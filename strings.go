package main

import (
  "fmt"
  "encoding/base64"
)

func main() {

  StringToEncode := "ABSJKFJABSKFJBAKSDJFBASLJKDHAI5465435438/43849874685"

  Encoding := base64.StdEncoding.EncodeToString([]byte(StringToEncode))
  fmt.Println(Encoding)
}
