package creational

import "testing"

func TestNewFactory(t *testing.T) {
	NewFactory("+")
	NewFactory("-")
	NewFactory("*")
	NewFactory("/")
}
