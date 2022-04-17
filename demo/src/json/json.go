package main

import (
  "encoding/json"
  "fmt"
)

func main() {
  fmt.Println("demo run")
  StractEncodeToJson()
  JsonDecodeIntoStruct()
  MapEncodeToJson()
  JsonDecodeIntoMap()
  fmt.Println("demo finish")
}

//struct的标签可以定义json解析时的一些操作
//ArrNum []int //没有标签，表示被解析成的json键名直接就是ArrNum
//ArrNum []int `json:"arr_num"` //这样表示，这个字段解析成的json键名是arr_num
//ArrNum []int `json:"arr_num,omitempty"` //这样表示，如果struct中这个字段是空的就不解析
//ArrNum []int `json:"-"` //这样表示，struct中这个字段不解析

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

//stract to json
func StractEncodeToJson() {
  data := TestStruct{
    Num:    1,
    Str:    "str",
    ArrNum: []int{1, 2, 3},
    ArrStr: []string{"str1", "str2", "str3"},
    ArrObj: []ObjStruct{{Num: 1, Str: "str1"}, {Num: 2, Str: "str2"}},
    Map:    MapStruct{NumKey: 1, StrKey: "str"},
  }
  jsonStr, _ := json.Marshal(data)
  fmt.Println("stract to json")
  fmt.Println(string(jsonStr))
}

//json to stract
func JsonDecodeIntoStruct() {
  var data TestStruct
  jsonStr := "{\"num\":1,\"str\":\"str\",\"arr_num\":[1,2,3],\"arr_str\":[\"str1\",\"str2\",\"str3\"],\"arr_obj\":[{\"num\":1,\"str\":\"str1\"},{\"num\":2,\"str\":\"str2\"}],\"map\":{\"num_key\":1,\"str_key\":\"str\"}}"
  json.Unmarshal([]byte(jsonStr), &data)
  fmt.Println("json to stract")
  fmt.Printf("%+v\r\n", data)
}

//map to json
func MapEncodeToJson() {
  data := make(map[string]interface{})
  data["num"] = 1
  data["str"] = "str"
  data["arr_num"] = []int{1, 2, 3}
  data["arr_str"] = []string{"str1", "str2", "str3"}
  data["arr_obj"] = []ObjStruct{{Num: 1, Str: "str1"}, {Num: 2, Str: "str2"}}
  data["map"] = MapStruct{NumKey: 1, StrKey: "str"}
  jsonStr, _ := json.Marshal(data)
  fmt.Println("map to json")
  fmt.Println(string(jsonStr))
}

//json to map
func JsonDecodeIntoMap() {
  var data map[string]interface{}
  jsonStr := "{\"num\":1,\"str\":\"str\",\"arr_num\":[1,2,3],\"arr_str\":[\"str1\",\"str2\",\"str3\"],\"arr_obj\":[{\"num\":1,\"str\":\"str1\"},{\"num\":2,\"str\":\"str2\"}],\"map\":{\"num_key\":1,\"str_key\":\"str\"}}"
  json.Unmarshal([]byte(jsonStr), &data)
  fmt.Println("json to map")
  fmt.Printf("%+v\r\n", data)
}
