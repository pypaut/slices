package slices

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMappingAddOne(t *testing.T) {
	numbers := []int{2, 3, 4, 5, 6}
	f := func(i int) (int, error) { return i + 1, nil }
	mapped, _ := SliceMapping(numbers, f)

	expectedResult := []int{3, 4, 5, 6, 7}
	if !reflect.DeepEqual(mapped, expectedResult) {
		t.Errorf("Expected %v, got %v", expectedResult, numbers)
	}
}

func TestMappingMultiplyByThree(t *testing.T) {
	numbers := []int{2, 3, 4, 5, 6}
	f := func(i int) (int, error) { return i * 3, nil }
	mapped, _ := SliceMapping(numbers, f)

	expectedResult := []int{6, 9, 12, 15, 18}
	if !reflect.DeepEqual(mapped, expectedResult) {
		t.Errorf("Expected %v, got %v", expectedResult, numbers)
	}
}

func TestMappingConvertToString(t *testing.T) {
	numbers := []int{2, 3, 4, 5, 6, 198}
	f := func(i int) (string, error) { return strconv.Itoa(i), nil }
	mapped, _ := SliceMapping(numbers, f)

	expectedResult := []string{"2", "3", "4", "5", "6", "198"}
	if !reflect.DeepEqual(mapped, expectedResult) {
		t.Errorf("Expected %v, got %v", expectedResult, numbers)
	}
}

func TestMappingConvertToInt(t *testing.T) {
	strNumbers := []string{"2", "3", "4", "5", "6", "198"}
	f := func(s string) (int, error) { return strconv.Atoi(s) }
	mapped, _ := SliceMapping(strNumbers, f)

	expectedResult := []int{2, 3, 4, 5, 6, 198}
	if !reflect.DeepEqual(mapped, expectedResult) {
		t.Errorf("Expected %v, got %v", expectedResult, strNumbers)
	}
}

func TestMappingConvertToIntWithError(t *testing.T) {
	strNumbers := []string{"2", "3", "4", "5", "6", "198eijo"}
	f := func(s string) (int, error) { return strconv.Atoi(s) }
	mapped, err := SliceMapping(strNumbers, f)

	if mapped != nil {
		t.Error("should be nil")
	}

	if err == nil {
		t.Error("err should not be nil")
	}
}

func TestMappingGetFieldFromStruct(t *testing.T) {
	type Person struct {
		Name string
	}

	persons := []*Person{
		&Person{Name: "Priscille"},
		&Person{Name: "Geoffrey"},
	}

	f := func(p *Person) (string, error) { return p.Name, nil }
	mapped, err := SliceMapping(persons, f)

	if err != nil {
		t.Errorf("expected err to be nil, got %v", err)
	}

	expectedResult := []string{"Priscille", "Geoffrey"}
	if !reflect.DeepEqual(mapped, expectedResult) {
		t.Errorf("expected %v, got %v", expectedResult, mapped)
	}
}
