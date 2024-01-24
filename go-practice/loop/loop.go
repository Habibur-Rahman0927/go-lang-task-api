package loop

import (
	"fmt"
)

func Loop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("=========")

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
