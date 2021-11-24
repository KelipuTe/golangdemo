package main

import "fmt"

//结构体方法
func main() {
  ac := AnimalCategory{species: "dog"}
  fmt.Printf("ac: %s\r\n", ac)
  ac.SetSpecies("cat")
  fmt.Printf("ac: %s\r\n", ac)

  animal := Animal{
    scientificName: "shorthair",
    AnimalCategory: ac,
  }
  fmt.Printf("animal: %s\r\n", animal)
}

//生物分类结构体
type AnimalCategory struct {
  kingdom string //界
  phylum  string //门
  class   string //纲
  order   string //目
  family  string //科
  genus   string //属
  species string //种
}

//这个方法可以自定义该类型的字符串表示形式
//参数是值传递，相当于原数据的一个副本
func (ac AnimalCategory) String() string {
  return fmt.Sprintf("category:%s%s%s%s%s%s%s", ac.kingdom, ac.phylum, ac.class, ac.order, ac.family, ac.genus, ac.species)
}

//这个方法可以修改该类型内部字段（种）的值
//引用传递，通过指针可以直接修改原数据
func (ac *AnimalCategory) SetSpecies(s string) {
  ac.species = s
}

//自定义数据类型的方法集合中仅会包含它的所有值方法，而该类型的指针类型的方法集合中包会包含所有值方法和所有指针方法。
//严格来讲，在值传递时只能调用到它的值方法。但是，go会进行自动地转译，使得在值传递时也能调用到它的指针方法。
func (ac AnimalCategory) SetSpeciesTest(s string) {
  ac.SetSpecies(s) //语句被自动转译为(&ac).SetSpecies(s)
}

//动物结构体
type Animal struct {
  scientificName string //学名
  AnimalCategory        //生物分类
}

//如果不给Animal定义String方法，那么会调用AnimalCategory的String方法
//如果给Animal定义String方法，那么AnimalCategory的String方法会被屏蔽
//结构体多层嵌入时，以嵌入的层级为依据嵌入层级越深的字段或方法越可能被屏蔽
func (a Animal) String() string {
  return fmt.Sprintf("name:%s,category:%s", a.scientificName, a.AnimalCategory.String())
}
