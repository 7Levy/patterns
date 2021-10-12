package creational

import (
	"fmt"
	"testing"
)

func TestNewSingleton(t *testing.T) {
	s1 := NewSingleton()
	s1["this"] = "one"
	s2 := NewSingleton()
	fmt.Println("This is", s2["this"])
}
