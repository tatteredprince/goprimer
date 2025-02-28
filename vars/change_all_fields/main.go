package main

import "fmt"

type Character struct {
	Name string
	Age  int
	Sex  string
}

func setName(chr *Character, name string) {
	chr.Name = name
}

func setAge(chr *Character, age int) {
	chr.Age = age
}

func setSex(chr *Character, sex string) {
	chr.Sex = sex
}

func changePerson(chr *Character, name string, age int, sex string) {
	// fmt.Printf("dbg person = %p\n", chr)
	chr = &Character{Name: chr.Name, Age: chr.Age, Sex: chr.Sex}
	// fmt.Printf("dbg person = %p\n", chr)
}

func changePerson2(chr *Character, name string, age int, sex string) {
	chr.Name, chr.Age, chr.Sex = name, age, sex
}

func main() {
	char := Character{Name: "Mary", Age: 31, Sex: "Female"}

	fmt.Println(char) // {Mary 31 Female}

	setName(&char, "Alexander")
	setAge(&char, 25)
	setSex(&char, "Male")

	fmt.Println(char) // {Alexander 25 Male}

	changePerson(&char, "Alice", 54, "Female")

	fmt.Println(char) // {Alexander 25 Male}

	changePerson2(&char, "Alice", 54, "Female")

	fmt.Println(char) // {Alice 54 Female}
}
