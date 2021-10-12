package creational

import (
	"fmt"
	"testing"
)

func TestUnsafeLazySingleton(t *testing.T) {
	s1 := UnsafeLazySingleton()
	s1["this"] = "one"
	s2 := UnsafeLazySingleton()
	fmt.Println("This is", s2["this"])
}

func TestSafeLazySingleton(t *testing.T) {
	s1 := SafeLazySingleton()
	s1["this"] = "one"
	s2 := SafeLazySingleton()
	fmt.Println("This is", s2["this"])
}
