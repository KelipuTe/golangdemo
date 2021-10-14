package main

import (
  "encoding/json"
  "fmt"
)

func main() {
  jsonEncode()
  jsonDecodeIntoMap()
  jsonDecodeIntoStruct()
}

type ObjStruct struct {
  Num int    `json:"num"`
  Str string `json:"str"`
}

type MapStruct struct {
  NumKey int    `json:"num_key"`
  StrKey string `json:"str_key"`
}

type TestStruct struct {
  Num    int         `json:"num"`
  Str    string      `json:"str"`
  ArrNum []int       `json:"arr_num"`
  ArrStr []string    `json:"arr_str"`
  ArrObj []ObjStruct `json:"arr_obj"`
  Map    MapStruct   `json:"map"`
}

func jsonEncode() {
  data := TestStruct{
    Num:    1,
    Str:    "str",
    ArrNum: []int{1, 2, 3},
    ArrStr: []string{"str1", "str2", "str3"},
    ArrObj: []ObjStruct{{Num: 1, Str: "str1"}, {Num: 2, Str: "str2"}},
    Map:    MapStruct{NumKey: 1, StrKey: "str"},
  }
  jsonStr, _ := json.Marshal(data)
  fmt.Println(string(jsonStr))
}

func jsonDecodeIntoMap() {
  var data map[string]interface{}
  jsonStr := "{\"num\":1,\"str\":\"str\",\"arr_num\":[1,2,3],\"arr_str\":[\"str1\",\"str2\",\"str3\"],\"arr_obj\":[{\"num\":1,\"str\":\"str1\"},{\"num\":2,\"str\":\"str2\"}],\"map\":{\"num_key\":1,\"str_key\":\"str\"}}"
  json.Unmarshal([]byte(jsonStr), &data)
  fmt.Println(data)
}

func jsonDecodeIntoStruct() {
  var data TestStruct
  jsonStr := "{\"num\":1,\"str\":\"str\",\"arr_num\":[1,2,3],\"arr_str\":[\"str1\",\"str2\",\"str3\"],\"arr_obj\":[{\"num\":1,\"str\":\"str1\"},{\"num\":2,\"str\":\"str2\"}],\"map\":{\"num_key\":1,\"str_key\":\"str\"}}"
  json.Unmarshal([]byte(jsonStr), &data)
  fmt.Println(data)
}
