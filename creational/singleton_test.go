package creational

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s1 := newSingleton()
	s1["this"] = "one"
	s2 := newSingleton()
	fmt.Println("This is", s2["this"])
}
