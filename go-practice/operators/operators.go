package operators

import "fmt"

func Operator() {
	var (
		sum1 = 100 + 50    // 150 (100 + 50)
		sum2 = sum1 + 250  // 400 (150 + 250)
		sum3 = sum2 + sum2 // 800 (400 + 400)
	)
	fmt.Println(sum3)

	var x = 5
	var y = 3
	fmt.Println(x > y)
}
