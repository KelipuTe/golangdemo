package main

import "fmt"

func main() {
  // fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
  // fmt.Println(wordBreak("aaaaaab", []string{"a", "aa", "aaa", "b"}))
  fmt.Println(wordBreak("aaaaaab", []string{"a", "aa", "aaa"}))
}

//给定一个非空字符串s和一个包含非空单词的列表wordDict。
//判定s是否可以被空格拆分为一个或多个在字典中出现的单词。
//拆分时可以重复使用字典中的单词。可以假设字典中没有重复的单词。

//递归，哈希表
//例如，s="leetcode"，wordDict=["leet", "code"]
//判断leetcode能不能拆分，可以分解成，判断l能不能拆分和判断eetcode能不能拆分，两个子问题
//依次向后遍历，当找到第一个可以拆分的部分leet时，问题变成了leet能拆分和判断code能不能拆分
//判断code能不能拆分的问题，和上面的问题是一样的，所以递归解决
//递归过程中，存在大量重复子问题
//例如，s="aaaaaab"，wordDict=["a", "aa", "aaa"]
//[a,aaaaab],[a,a,aaaab],[a,a,a,aaab],[a,a,a,a,aab],[a,a,a,a,a,ab],[a,a,a,a,a,a,b]
//回溯至[a,a,a,a,aab]，递归[a,a,a,a,aa,b]
//回溯至[a,a,a,aaab]，递归[a,a,a,aa,ab],[a,a,a,aaa,b]
//回溯至[a,a,aaaab]，递归[a,a,aa,aab],[a,a,aaa,ab]
//到这里已经重复出现以ab和aab结尾的子问题了，它们会重复派生子问题
//如果递归继续下去，会重复产生大量这样的子问题，所以需要缓存这些子问题的结果

//139-单词拆分
func wordBreak(s string, wordDict []string) bool {
  iSLen := len(s)
  var mapWord map[string]bool = map[string]bool{} //单词哈希表
  var mapResCache map[int]bool = map[int]bool{}   //从int下标开始往后的部分的字符串的判断结果

  for _, sStr := range wordDict {
    mapWord[sStr] = true
  }

  return di4gui1(s, iSLen, 0, mapWord, mapResCache)
}

func di4gui1(s string, iSLen int, iStartIndex int, mapWord map[string]bool, mapResCache map[int]bool) bool {
  if iStartIndex >= iSLen {
    return true //如果开始下标到头了，则找到了
  }
  if _, bExist := mapResCache[iStartIndex]; bExist {
    return mapResCache[iStartIndex] //从缓存中获取重复子问题的结果
  }
  for ii := iStartIndex; ii < iSLen; ii++ {
    ts := s[iStartIndex : ii+1]
    if _, bExist := mapWord[ts]; bExist {
      if di4gui1(s, iSLen, ii+1, mapWord, mapResCache) {
        mapResCache[iStartIndex] = true
        return true
      }
    }
  }
  mapResCache[iStartIndex] = false
  return false
}
