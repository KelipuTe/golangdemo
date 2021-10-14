package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "strconv"
  "strings"
)

func main() {
  curlGetJson()
  curlGetXml()
  curlGetWithParam()
  curlPostWithForm()
  curlPostWithJson()
}

func errThrow(title string, err error) {
  if err != nil {
    fmt.Printf("%s,err:%v\n", title, err)
    panic(err)
  }
}

func curlGetJson() {
  apiUrl := "http://127.0.0.1:8000/api/test/get_json"

  fmt.Println(apiUrl)
  p0httpRes, err := http.Get(apiUrl)
  errThrow("p0httpRes", err)
  defer p0httpRes.Body.Close()

  resio, err := ioutil.ReadAll(p0httpRes.Body)
  errThrow("resio", err)
  fmt.Println(string(resio))
}

func curlGetXml() {
  apiUrl := "http://127.0.0.1:8000/api/test/get_xml"

  fmt.Println(apiUrl)
  p0httpRes, err := http.Get(apiUrl)
  errThrow("p0httpRes", err)
  defer p0httpRes.Body.Close()

  resio, err := ioutil.ReadAll(p0httpRes.Body)
  errThrow("resio", err)
  fmt.Println(string(resio))
}

func curlGetWithParam() {
  apiUrl := "http://127.0.0.1:8000/api/test/get_json_with_param"
  param := url.Values{}
  param.Set("num", strconv.Itoa(1))
  param.Set("str", "str")

  p0url, err := url.ParseRequestURI(apiUrl)
  errThrow("p0url", err)
  p0url.RawQuery = param.Encode()

  fmt.Println(p0url.String())
  p0httpRes, err := http.Get(p0url.String())
  errThrow("p0httpRes", err)
  defer p0httpRes.Body.Close()

  resio, err := ioutil.ReadAll(p0httpRes.Body)
  errThrow("resio", err)
  fmt.Println(string(resio))
}

func curlPostWithForm() {
  apiUrl := "http://127.0.0.1:8000/api/test/post_json_with_form"
  data := "num=1&str=str"

  fmt.Println(apiUrl)
  p0httpRes, err := http.Post(apiUrl, "application/x-www-form-urlencoded", strings.NewReader(data))
  errThrow("p0httpRes", err)
  defer p0httpRes.Body.Close()

  resio, err := ioutil.ReadAll(p0httpRes.Body)
  errThrow("resio", err)
  fmt.Println(string(resio))
}

func curlPostWithJson() {
  apiUrl := "http://127.0.0.1:8000/api/test/post_json_with_json"
  data := "{\"num\":\"1\",\"str\":\"str\"}"

  fmt.Println(apiUrl)
  p0httpRes, err := http.Post(apiUrl, "application/json", strings.NewReader(data))
  errThrow("p0httpRes", err)
  defer p0httpRes.Body.Close()

  resio, err := ioutil.ReadAll(p0httpRes.Body)
  errThrow("resio", err)
  fmt.Println(string(resio))
}
