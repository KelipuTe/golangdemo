package kn_reflect

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
