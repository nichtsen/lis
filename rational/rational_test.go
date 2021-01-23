package rational

import "fmt"

func Example_rational() {
	a := New(1, 2)
	b := New(1, 3)
	fmt.Printf("%v %v %v %v %v %v", a, b, Add(a, b), Subtract(a, b), Multiply(a, b), Divide(a, b))
	//Output:
	//1/2 1/3 5/6 1/6 1/6 3/2
}
