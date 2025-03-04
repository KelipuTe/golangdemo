package unsafe

type User struct {
	Name string
	Sex  int
	age  int
}

func (this User) GetName() string {
	return this.Name
}

func (p1this *User) SetSex(sex int) int {
	p1this.Sex = sex
	return sex
}

func (p1this *User) resetAge() {
	p1this.age = 18
}

type UserV2 struct {
	Name    string
	Age     int32
	Alias   []byte
	Address string
}

type UserV4 struct {
	Name    string
	Age     int32
	AgeV2   int32
	Alias   []byte
	Address string
}

type UserV6 struct {
	Name    string
	Age     int32
	Alias   []byte
	Address string
	AgeV2   int32
}

type UserV8 struct {
	Name    string
	Alias   []byte
	Address string
	Age     int32
}
