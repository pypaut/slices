package slices_test

import (
	"fmt"

	"github.com/pypaut/slices"
)

func ExampleMap() {
	type Person struct {
		Name string
		Age  int
	}

	persons := []Person{
		{Name: "Alice", Age: 31},
		{Name: "Bob", Age: 23},
		{Name: "Vera", Age: 48},
	}

	names, _ := slices.Map(
		persons, func(p Person) (string, error) { return p.Name, nil },
	)
	ages, _ := slices.Map(
		persons, func(p Person) (int, error) { return p.Age, nil },
	)
	fmt.Println("Names:", names)
	fmt.Println("Ages:", ages)
	// Output:
	// Names: [Alice Bob Vera]
	// Ages: [31 23 48]
}
