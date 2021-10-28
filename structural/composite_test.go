package structural

import (
	"fmt"
	"testing"
)

func TestEmployee(t *testing.T) {
	ceo := GetEmployee("John", "CEO", 30000)

	headSales := GetEmployee("Robert", "Head Sales", 20000)
	headMarketing := GetEmployee("Michel", "Head Marketing", 20000)

	ceo.Add(headSales)
	ceo.Add(headMarketing)

	salesExecutive1 := GetEmployee("Richard", "Sales", 10000)
	salesExecutive2 := GetEmployee("Rob", "Sales", 10000)

	headSales.Add(salesExecutive1)
	headSales.Add(salesExecutive2)

	clerk1 := GetEmployee("Laura", "Marketing", 10000)
	clerk2 := GetEmployee("Bob", "Marketing", 10000)

	headMarketing.Add(clerk1)
	headMarketing.Add(clerk2)

	ceo.PrintSubordinates()

	ceo.Remove(headSales)
	//ceo.Remove(headMarketing)
	fmt.Println("")
	ceo.PrintSubordinates()

}
