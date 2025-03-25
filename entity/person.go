package entity

type Person struct {
	Name  string
	Age   int
	Class Class
}

type Class struct {
	ClassNo int
}
