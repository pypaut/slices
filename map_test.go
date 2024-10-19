package slices

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMapAddOne(t *testing.T) {
	numbers := []int{2, 3, 4, 5, 6}
	f := func(i int) (int, error) { return i + 1, nil }
	mapped, _ := Map(numbers, f)

	expected := []int{3, 4, 5, 6, 7}
	if !reflect.DeepEqual(mapped, expected) {
		t.Errorf("Expected %v, got %v", expected, mapped)
	}
}

func TestMapMultiplyByThree(t *testing.T) {
	numbers := []int{2, 3, 4, 5, 6}
	f := func(i int) (int, error) { return i * 3, nil }
	mapped, _ := Map(numbers, f)

	expected := []int{6, 9, 12, 15, 18}
	if !reflect.DeepEqual(mapped, expected) {
		t.Errorf("Expected %v, got %v", expected, numbers)
	}
}

func TestMapConvertToString(t *testing.T) {
	numbers := []int{2, 3, 4, 5, 6, 198}
	f := func(i int) (string, error) { return strconv.Itoa(i), nil }
	mapped, _ := Map(numbers, f)

	expected := []string{"2", "3", "4", "5", "6", "198"}
	if !reflect.DeepEqual(mapped, expected) {
		t.Errorf("Expected %v, got %v", expected, numbers)
	}
}

func TestMapConvertToInt(t *testing.T) {
	strNumbers := []string{"2", "3", "4", "5", "6", "198"}
	f := func(s string) (int, error) { return strconv.Atoi(s) }
	mapped, _ := Map(strNumbers, f)

	expected := []int{2, 3, 4, 5, 6, 198}
	if !reflect.DeepEqual(mapped, expected) {
		t.Errorf("Expected %v, got %v", expected, strNumbers)
	}
}

func TestMapConvertToIntWithError(t *testing.T) {
	strNumbers := []string{"2", "3", "4", "5", "6", "198eijo"}
	f := func(s string) (int, error) { return strconv.Atoi(s) }
	mapped, err := Map(strNumbers, f)

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
		&Person{Name: "Priscille"},
		&Person{Name: "Geoffrey"},
	}

	f := func(p *Person) (string, error) { return p.Name, nil }
	mapped, err := Map(persons, f)

	if err != nil {
		t.Errorf("expected err to be nil, got %v", err)
	}

	expected := []string{"Priscille", "Geoffrey"}
	if !reflect.DeepEqual(mapped, expected) {
		t.Errorf("expected %v, got %v", expected, mapped)
	}
}
