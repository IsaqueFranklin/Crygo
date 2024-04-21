package main

import (
  "fmt"
  "encoding/json"
  "io/ioutil"
  "log"
  "bytes"
  "net/http"
  "os"
  "strings"
  "sync"
)

func main(){
  fmt.Println("Hello.")
}

//This function is gonna read the passwords in the file passwords.txt and return a slice of strings.
func readInPasswords(passwordFile string) []string {
  b, err := ioutil.ReadFile(passwordFile) //Passing the file name.
  if err != nil {
    fmt.Println(err)
  }

  str := string(b) //Convert content b to a string

  return string.Split(str, "\n")
}

//postToURL
//Succeeded = true if res.status == 200
func postToURL(url string, username string, password string) (succeeded bool) {
  fmt.Println("password: ", password)
  values := map[string]string{"username": username, "email": username, "password": password}

  jsonValue, _ := json.Marshal(values)

  resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
  if err != nil {
    fmt.Println("Error occured: ", err)
    return false
  }
  defer resp.Body.Close()
  fmt.Println("Response status: ", resp.Status)
  var body interface{}
  json.NewDecoder(resp.Body).Decode(&body)
  fmt.Println(body)
  return resp.Status == "200"
}
