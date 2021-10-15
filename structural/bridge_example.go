package structural

import "fmt"

func ExampleCommonSMS() {
	fmt.Println("ExampleCommonSMS.")
	m := NewCommonMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
}

func ExampleCommonEmail() {
	fmt.Println("ExampleCommonEmail.")
	m := NewCommonMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
}
