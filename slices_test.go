package slices_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/pypaut/slices"
)

func TestMapAddOne(t *testing.T) {
	numbers := []int{2, 3, 4, 5, 6}
	f := func(i int) (int, error) { return i + 1, nil }
	mapped, _ := slices.Map(numbers, f)

	expected := []int{3, 4, 5, 6, 7}
	if !reflect.DeepEqual(mapped, expected) {
		t.Errorf("Expected %v, got %v", expected, mapped)
	}
}

func TestMapMultiplyByThree(t *testing.T) {
	numbers := []int{2, 3, 4, 5, 6}
	f := func(i int) (int, error) { return i * 3, nil }
	mapped, _ := slices.Map(numbers, f)

	expected := []int{6, 9, 12, 15, 18}
	if !reflect.DeepEqual(mapped, expected) {
		t.Errorf("Expected %v, got %v", expected, numbers)
	}
}

func TestMapConvertToString(t *testing.T) {
	numbers := []int{2, 3, 4, 5, 6, 198}
	f := func(i int) (string, error) { return strconv.Itoa(i), nil }
	mapped, _ := slices.Map(numbers, f)

	expected := []string{"2", "3", "4", "5", "6", "198"}
	if !reflect.DeepEqual(mapped, expected) {
		t.Errorf("Expected %v, got %v", expected, numbers)
	}
}

func TestMapConvertToInt(t *testing.T) {
	strNumbers := []string{"2", "3", "4", "5", "6", "198"}
	f := func(s string) (int, error) { return strconv.Atoi(s) }
	mapped, _ := slices.Map(strNumbers, f)

	expected := []int{2, 3, 4, 5, 6, 198}
	if !reflect.DeepEqual(mapped, expected) {
		t.Errorf("Expected %v, got %v", expected, strNumbers)
	}
}

func TestMapConvertToIntWithError(t *testing.T) {
	strNumbers := []string{"2", "3", "4", "5", "6", "198eijo"}
	f := func(s string) (int, error) { return strconv.Atoi(s) }
	mapped, err := slices.Map(strNumbers, f)

	if mapped != nil {
		t.Error("should be nil")
	}

	if err == nil {
		t.Error("err should not be nil")
	}
}

func TestMapGetFieldFromStruct(t *testing.T) {
	type Person struct {
		Name string
	}

	persons := []*Person{
		{Name: "Priscille"},
		{Name: "Geoffrey"},
	}

	f := func(p *Person) (string, error) { return p.Name, nil }
	mapped, err := slices.Map(persons, f)

	if err != nil {
		t.Errorf("expected err to be nil, got %v", err)
	}

	expected := []string{"Priscille", "Geoffrey"}
	if !reflect.DeepEqual(mapped, expected) {
		t.Errorf("expected %v, got %v", expected, mapped)
	}
}

func TestFilterIntegers(t *testing.T) {
	numbers := []int{2, 3, 4, 5, 6}
	f := func(i int) bool { return i > 2 }
	filtered := slices.Filter(numbers, f)

	expected := []int{3, 4, 5, 6}
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Expected %v, got %v", expected, filtered)
	}
}

func TestFilterStrLen(t *testing.T) {
	names := []string{"Priscille", "Geoffrey", "Nala", "Kira"}
	f := func(s string) bool { return len(s) < 5 }
	filtered := slices.Filter(names, f)

	expected := []string{"Nala", "Kira"}
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Expected %v, got %v", expected, filtered)
	}
}

func TestCheckDuplicatesStringPositive(t *testing.T) {
	names := []string{"Priscille", "Geoffrey", "Nala", "Kira", "Kira"}

	if !slices.CheckDuplicates(names) {
		t.Errorf("Did not find duplicates in %v", names)
	}
}

func TestCheckDuplicatesStringNegative(t *testing.T) {
	names := []string{"Priscille", "Geoffrey", "Nala", "Kira"}

	if slices.CheckDuplicates(names) {
		t.Errorf("Found duplicates in %v", names)
	}
}

func TestCheckDuplicatesIntPositive(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 7}

	if !slices.CheckDuplicates(numbers) {
		t.Errorf("Did not find duplicates in %v", numbers)
	}
}

func TestCheckDuplicatesIntNegative(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	if slices.CheckDuplicates(numbers) {
		t.Errorf("Found duplicates in %v", numbers)
	}
}

func TestCheckDuplicatesSlicesNegative(t *testing.T) {
	sliceOfSlices := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	if slices.CheckDuplicates(sliceOfSlices) {
		t.Errorf("Found duplicates in %v", sliceOfSlices)
	}
}

func TestCheckDuplicatesSlicesPositive(t *testing.T) {
	sliceOfSlices := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8},
		{1, 2, 3},
	}

	if !slices.CheckDuplicates(sliceOfSlices) {
		t.Errorf("Did not find duplicates in %v", sliceOfSlices)
	}
}

func TestCheckDuplicatesEmptySlicesPositive(t *testing.T) {
	sliceOfSlices := [][]int{
		nil,
		nil,
		{1, 2, 3},
		{4, 5, 6},
		{7, 8},
		{1, 2, 3},
	}

	if !slices.CheckDuplicates(sliceOfSlices) {
		t.Errorf("Did not find duplicates in %v", sliceOfSlices)
	}
}

func TestCheckDuplicatesEmptySliceNegative(t *testing.T) {
	numbers := []int{}

	if slices.CheckDuplicates(numbers) {
		t.Errorf("Found duplicates in %v", numbers)
	}
}

func TestCheckDuplicatesCustomStructNegative(t *testing.T) {
	type Cat struct {
		Name  string
		Age   int
		Color string
	}

  cats := []Cat{
    {Name: "Nala", Age: 6, Color: "BROWN"},
    {Name: "Kira", Age: 2, Color: "ORANGE"},
  }

	if slices.CheckDuplicates(cats) {
		t.Errorf("Found duplicates in %v", cats)
	}
}

func TestCheckDuplicatesCustomStructPointersNegative(t *testing.T) {
	type Cat struct {
		Name  string
		Age   int
		Color string
	}

  cats := []*Cat{
    {Name: "Nala", Age: 6, Color: "BROWN"},
    {Name: "Kira", Age: 2, Color: "ORANGE"},
  }

	if slices.CheckDuplicates(cats) {
		t.Errorf("Found duplicates in %v", cats)
	}
}

func TestCheckDuplicatesCustomStructPositive(t *testing.T) {
	type Cat struct {
		Name  string
		Age   int
		Color string
	}

  cats := []Cat{
    {Name: "Nala", Age: 6, Color: "BROWN"},
    {Name: "Kira", Age: 2, Color: "ORANGE"},
    {Name: "Kira", Age: 2, Color: "ORANGE"},
  }

	if !slices.CheckDuplicates(cats) {
		t.Errorf("Did not find duplicates in %v", cats)
	}
}

func TestCheckDuplicatesCustomStructPointersPositive(t *testing.T) {
	type Cat struct {
		Name  string
		Age   int
		Color string
	}

  cats := []*Cat{
    {Name: "Nala", Age: 6, Color: "BROWN"},
    {Name: "Kira", Age: 2, Color: "ORANGE"},
    {Name: "Kira", Age: 2, Color: "ORANGE"},
  }

	if !slices.CheckDuplicates(cats) {
		t.Errorf("Did not find duplicates in %v", cats)
	}
}

func TestCheckDuplicatesCustomStructNegative2(t *testing.T) {
	type Cat struct {
		Name  string
		Age   int
		Color string
	}

  cats := []Cat{
    {Name: "Nala", Age: 6, Color: "BROWN"},
    {Name: "Kira", Age: 2, Color: "ORANGE"},
    {Name: "Kira", Age: 2, Color: "ORANG"},
  }

  if slices.CheckDuplicates(cats) {
    t.Errorf("Found duplicates in %v", cats)
  }
}
