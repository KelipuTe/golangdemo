package main

import "fmt"

func main() {
  obj := Constructor(2)
  obj.Put(1, 1)
  obj.Put(2, 2)
  fmt.Println(obj.Get(1))
  obj.Put(3, 3)
  fmt.Println(obj.Get(2))
  obj.Put(4, 4)
  fmt.Println(obj.Get(1))
  fmt.Println(obj.Get(3))
  fmt.Println(obj.Get(4))
}

//运用你所掌握的数据结构，设计和实现一个LRU(最近最少使用)缓存机制。
//实现LRUCache类：
//`LRUCache(int capacity)`以正整数作为容量capacity初始化LRU缓存。
//`int get(int key)`如果关键字key存在于缓存中，则返回关键字的值，否则返回-1。
//`void put(int key, int value)`如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组[关键字-值]。
//当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//进阶：在O(1)时间复杂度内完成这两种操作。

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

//哈希表，双向链表
//哈希表用于实现O(1)快速找到缓存
//双向链表用于实现O(1)删除最近最少使用的缓存
//双向链表增加头尾伪结点，可以方便链表结点的插入和删除操作

//146-LRU缓存机制

type ListNode struct {
  iKey, iValue int       //键，值
  plnP, plnN   *ListNode //前，后结点指针
} //缓存结点

type LRUCache struct {
  iNowSize, iMaxSize int               //缓存已用大小和最大大小
  mapCache           map[int]*ListNode //哈希表，指向每个缓存结点
  plnTou2, plnWei3   *ListNode         //双向链表头尾伪结点，方便插入和删除操作
}

func Constructor(capacity int) LRUCache {
  var lruCacne LRUCache = LRUCache{
    iNowSize: 0,
    iMaxSize: capacity,
    mapCache: map[int]*ListNode{},
    plnTou2:  &ListNode{iKey: 0, iValue: 0},
    plnWei3:  &ListNode{iKey: 0, iValue: 0},
  }
  lruCacne.plnTou2.plnN = lruCacne.plnWei3
  lruCacne.plnWei3.plnP = lruCacne.plnTou2

  return lruCacne
}

func (this *LRUCache) Get(key int) int {
  if _, bExist := this.mapCache[key]; bExist {
    tpln := this.mapCache[key]
    //从链表中移除这个结点
    tpln.plnP.plnN = tpln.plnN
    tpln.plnN.plnP = tpln.plnP
    //移动这个结点到头部
    tpln.plnN = this.plnTou2.plnN
    tpln.plnP = this.plnTou2
    this.plnTou2.plnN.plnP = tpln
    this.plnTou2.plnN = tpln

    return tpln.iValue
  }
  return -1
}

func (this *LRUCache) Put(key int, value int) {
  if _, bExist := this.mapCache[key]; bExist {
    //键存在
    tpln := this.mapCache[key]
    tpln.iValue = value //重新赋值
    //从链表中移除这个结点
    tpln.plnP.plnN = tpln.plnN
    tpln.plnN.plnP = tpln.plnP
    //移动这个结点到头部
    tpln.plnN = this.plnTou2.plnN
    tpln.plnP = this.plnTou2
    this.plnTou2.plnN.plnP = tpln
    this.plnTou2.plnN = tpln
  } else {
    //键不存在
    tpln := &ListNode{iKey: key, iValue: value}
    //移动这个结点到头部
    tpln.plnN = this.plnTou2.plnN
    tpln.plnP = this.plnTou2
    this.plnTou2.plnN.plnP = tpln
    this.plnTou2.plnN = tpln
    this.mapCache[key] = tpln //哈希表添加结点
    this.iNowSize++           //计数增加
    if this.iNowSize > this.iMaxSize {
      //触发淘汰机制
      tplnDel := this.plnWei3.plnP //要淘汰的节点
      //从链表中移除这个结点
      tplnDel.plnP.plnN = tplnDel.plnN
      tplnDel.plnN.plnP = tplnDel.plnP
      delete(this.mapCache, tplnDel.iKey) //哈希表移除结点
      this.iNowSize--                     //计数减少
    }
  }
}
