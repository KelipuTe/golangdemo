package main

import (
  "testing"
)

//单元测试的方法名以Test开头，以(t *testing.T)作为参数
func Test1(t *testing.T) {
  num := AddOperate(1, 2)
  if num != 3 {
    t.Error("测试失败")
  }
}

//性能测试的方法名以Benchmark开头，以(t *testing.B)做为参数
func Benchmark1(t *testing.B) {
  for i := 0; i < t.N; i++ {
    AddOperate(1, 2)
  }
}
