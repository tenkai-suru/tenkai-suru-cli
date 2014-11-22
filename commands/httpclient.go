package commands

import (
  "net/http"
  "fmt"
)

func httpGet(path string) {
  resp, _ := http.Get("https://tenkai-suru-api.herokuapp.com/" + path)
  fmt.Println("Response body ", resp.Body)
  fmt.Println(resp.Status)
}

func httpPost(path string) {
  resp, _ := http.Post("https://tenkai-suru-api.herokuapp.com/" + path, "application/text", nil)
  fmt.Println("Response body ", resp.Body)
  fmt.Println(resp.Status)
}
