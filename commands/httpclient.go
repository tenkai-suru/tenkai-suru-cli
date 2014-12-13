package commands

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func httpGet(path string) {
  resp, _ := http.Get("https://tenkai-suru-api.herokuapp.com/" + path)

  contents, _ := ioutil.ReadAll(resp.Body)

  fmt.Printf("%s\n", string(contents))
  fmt.Println(resp.Status)
}

func httpPost(path string) {
  resp, _ := http.Post("https://tenkai-suru-api.herokuapp.com/" + path, "application/text", nil)

  contents, _ := ioutil.ReadAll(resp.Body)

  fmt.Printf("%s\n", string(contents))
  fmt.Println(resp.Status)
}
