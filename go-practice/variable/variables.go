package variable

import "fmt"

func Variables() {
	var student1 = "john"
	fmt.Println(student1)

	student2 := "Jane"
	fmt.Println(student2)

	var student3 string
	student3 = "test"
	fmt.Println(student3)

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var d, e, f, g int = 1, 3, 5, 7

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	h, i := "hello ", "world"
	fmt.Print(h)
	fmt.Print(i)
}
