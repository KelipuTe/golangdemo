package main

import "fmt"

func main() {
	ac := AnimalCategory{
		species: "dog",
	}
	fmt.Printf("ac: %s\r\n", ac)

	animal := Animal{
		scientificName: "shorthair",
		AnimalCategory: ac,
	}
	fmt.Printf("animal: %s\r\n", animal)
}

// 生物分类结构体
type AnimalCategory struct {
	kingdom string // 界
	phylum  string // 门
	class   string // 纲
	order   string // 目
	family  string // 科
	genus   string // 属
	species string // 种
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("category:%s%s%s%s%s%s%s", ac.kingdom, ac.phylum, ac.class, ac.order, ac.family, ac.genus, ac.species)
}

// 动物结构体
type Animal struct {
	scientificName string // 学名
	AnimalCategory        // 生物分类
}

// 如果不给 Animal 定义 String 方法，那么会调用 AnimalCategory 的 String 方法
// 如果给 Animal 定义 String 方法，那么 AnimalCategory 的 String 方法会被屏蔽
// 结构体多层嵌入时，以嵌入的层级为依据嵌入层级越深的字段或方法越可能被屏蔽
func (a Animal) String() string {
	return fmt.Sprintf("name:%s,category:%s", a.scientificName, a.AnimalCategory.String())
}
