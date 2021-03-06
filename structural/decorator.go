package structural

import (
	"fmt"
	"time"
)

func myFunc() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

func coolFunc(a func()) {
	fmt.Printf("Starting function execution: %s\n", time.Now())
	a()
	fmt.Printf("End of function execution: %s\n", time.Now())
}

func Decorator() {
	fmt.Printf("Type: %T\n", myFunc)
	coolFunc(myFunc)
}
