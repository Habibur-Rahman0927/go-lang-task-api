package functions

import (
	"fmt"
)

func Functions() {
	fmt.Println("I just got executed!")
	familyName("Rock")
	fmt.Println(myFunction(1, 2))

}

func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func myFunction(x int, y int) int {
	return x + y
}
