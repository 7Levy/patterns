package creational

import (
	"fmt"
	"testing"
)

func TestSafeLazySingleton1(t *testing.T) {
	s1 := SafeLazySingleton1()
	s1["this"] = "one"
	s2 := SafeLazySingleton1()
	fmt.Println("This is", s2["this"])
}

func TestSafeLazySingleton2(t *testing.T) {
	s1 := SafeLazySingleton2()
	s1["this"] = "one"
	s2 := SafeLazySingleton2()
	fmt.Println("This is", s2["this"])
}

func TestUnsafeLazySingleton(t *testing.T) {
	s1 := UnsafeLazySingleton()
	s1["this"] = "one"
	s2 := UnsafeLazySingleton()
	fmt.Println("This is", s2["this"])
}
