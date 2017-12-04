package main

import (
	"fmt"
	"go/types"
	"math"
	"strings"
)
var g int = 0;

func main() {
	//var age int;
	//fmt.Println( age)
	//var p int;
	//p = 41; //static type declaration.
	//p1 := 23 //dynamic type declaration.
	//fmt.Println(p1);

	const length  = 2;
	var p3 int;
	p3 = length * 2;
	fmt.Println(p3)
	if(p3 == 4) {
		fmt.Println(p3+1)
	}else{
		fmt.Println(p3+2)
	}

	var x interface{}
	x = 5
	switch x.(type) {
	case types.Nil:
		print("Hello")
	case string:
		fmt.Println("It's String")
	default:
		fmt.Println("don't know the type")
	}

	fmt.Println(add(1,12))

	//Declare function variable.
	getSqrt := func(no float64) float64 {
		return math.Sqrt(no);
	}

	fmt.Println(getSqrt(1.2))

	//Call anonymous function.
	var i = getSquence();
	//call the function i()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())

	var greeting  = "Hello";
	fmt.Printf("%X\n" ,greeting[0]);

	fmt.Println(len(greeting))
	fmt.Println(strings.Join([]string{"Iniyan","Paramasivam"}," "))


	//Usage of Array.

	var a[10] int;
	var j int;

	for j=0; j< 10; j++ {
		a[j] = j+1
	}

	for j=0;j<10;j++ {
		fmt.Println(a[j])
	}

	//var l int  = 10;
	var l1 *int;
	//l1 = &l;
	fmt.Println( l1);

	//Use make to create slice.

	//last argument in capacity. first is length.
	//I think slice is similar to list. (Variable array length.)
	var numbers []int = make([]int,5,51)
	fmt.Println(cap(numbers))

}


func getSquence() func() int {
	i := 0
	//Anonymous function.
	return func() int {
		i = i + 1
		return i
	}
}

//write return type. you can even return multiple types.
func add(num1 int, num2 int) (int,int,int) {
	return 	num1 + num2, 1, 2
}
