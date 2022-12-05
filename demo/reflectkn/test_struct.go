package reflectkn

type User struct {
	Name string
	Sex  int
	age  int
}

func (this User) GetName() string {
	return this.Name
}

func (p7this *User) SetSex(sex int) int {
	p7this.Sex = sex
	return sex
}

func (p7this *User) resetAge() {
	p7this.age = 18
}
