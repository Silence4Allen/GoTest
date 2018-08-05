package demo

type mac struct {
	name string
	age  int
}

func NewMac(name string, age int) (a *mac) {
	a = new(mac)
	a.name = name
	a.age = age
	return a
}
