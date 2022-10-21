package kn_map

import "fmt"

func MapStruct() map[string]User {
	t4map := make(map[string]User, 4)
	t4map["user2"] = User{Name: "user2"}
	t4map["user4"] = User{Name: "user4"}
	t4map["user6"] = User{Name: "user6"}
	t4map["user8"] = User{Name: "user8"}
	fmt.Printf("map,%p\r\n", t4map)
	return t4map
}
func MapStructPointer() map[string]*User {
	t4map := make(map[string]*User, 4)
	t4map["user2"] = &User{Name: "user2"}
	t4map["user4"] = &User{Name: "user4"}
	t4map["user6"] = &User{Name: "user6"}
	t4map["user8"] = &User{Name: "user8"}
	fmt.Printf("map,%p\r\n", t4map)
	return t4map
}
