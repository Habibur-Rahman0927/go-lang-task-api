package datatypes

import "fmt"

func DataTypes() {
	// boolean types
	booleanType()

	// integer types
	integerType()

	// float types
	floatType()

	// string types
	stringType()
}

func customPrintf(n interface{}) {
	fmt.Printf("Type: %T, value: %v\n", n, n)
}

func booleanType() {
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)

	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
}

func integerType() {
	var x int = 500
	var y int = -500
	// fmt.Printf("Type: %T, value: %v\n", x, x)
	// fmt.Printf("Type: %T, value: %v\n", y, y)
	customPrintf(x)
	customPrintf(y)
}

func floatType() {
	var x float32 = 124.34
	var y float32 = 3.4e+38
	customPrintf(x)
	customPrintf(y)
}

func stringType() {
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "world 1"

	customPrintf(txt1)
	customPrintf(txt2)
	customPrintf(txt3)
}
