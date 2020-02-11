package model

type student struct {
	Name string
	Id   string
	Age  int
}

func NewStudent(name string, id string, age int) student {
	return student{
		Name: name,
		Id:   id,
		Age:  age,
	}
}
