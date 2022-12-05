package mapkn

import "fmt"

type s6User struct {
	Name string
}

func f8MapStruct() map[string]s6User {
	m3s6User := make(map[string]s6User, 4)
	m3s6User["user2"] = s6User{Name: "user2"}
	m3s6User["user4"] = s6User{Name: "user4"}
	m3s6User["user6"] = s6User{Name: "user6"}
	m3s6User["user8"] = s6User{Name: "user8"}
	fmt.Printf("map,%p\r\n", m3s6User)
	return m3s6User
}

func f8MapStructPointer() map[string]*s6User {
	m3p7s6User := make(map[string]*s6User, 4)
	m3p7s6User["user2"] = &s6User{Name: "user2"}
	m3p7s6User["user4"] = &s6User{Name: "user4"}
	m3p7s6User["user6"] = &s6User{Name: "user6"}
	m3p7s6User["user8"] = &s6User{Name: "user8"}
	fmt.Printf("map,%p\r\n", m3p7s6User)
	return m3p7s6User
}
