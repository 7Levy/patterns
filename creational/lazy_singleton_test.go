package creational

import (
	"fmt"
	"testing"
)

func TestNewLazySingleton(t *testing.T) {
	s1 := NewLazySingleton()
	s1["this"] = "one"
	s2 := NewLazySingleton()
	fmt.Println("This is", s2["this"])
}
