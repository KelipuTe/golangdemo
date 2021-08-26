package main

import (
	"fmt"
	"sort"
)

func main() {
  fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
  fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
}

//第i个人的体重为people[i]，每艘船可以承载的最大重量为limit。
//每艘船最多可同时载两人，但条件是这些人的重量之和最多为limit。
//返回载到每一个人所需的最小船数。(保证每个人都能被船载)。
//1<=people.length<=500001<=people[i]<=limit<=30000;

//排序，双指针
//因为，每个船最多装两人。所以，要么是装一个大体重，要么是装一个大体重加一个小体重。
//如果一个大体重加一个小体重小于船最大重量，就能装两个，否则只能装一个。
//这样的问题在有序序列中更容易求解，所以先排序。
//然后，使用两个指针，一个指向未处理的最小体重，一个指向未处理的最大体重。
//如果大体重加小体重小于船最大重量，属于装一个大体重加一个小体重的情况，两个指针各自移动一位。
//如果大体重加小体重大于船最大重量，属于装一个大体重的情况，大体重指针动一位。

//881-救生艇
func numRescueBoats(people []int, limit int) int {
  var countRes int = 0
  var indexSmall, indexBig int

  sort.Ints(people)
  indexSmall, indexBig = 0, len(people)-1
  for indexSmall <= indexBig {
    if people[indexSmall]+people[indexBig] <= limit {
      //装一个大体重加一个小体重的情况
      countRes++
      indexSmall++
      indexBig--
    } else {
      //装一个大体重的情况
      countRes++
      indexBig--
    }
  }

  return countRes
}
