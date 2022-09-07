package challenge90

import (
	"reflect"
	"testing"
)

func TestSplit2DigNum(t *testing.T) {
	expectedA1 := 5
	expectedB1 := 8
	expectedA2 := 0
	expectedB2 := 8
	expectedA3 := 51
	expectedB3 := 8
	resultA1, resultB1 := Split2DigNum(58)
	if !(resultA1 == expectedA1 && resultB1 == expectedB1) {
		t.Errorf("resultA1: %d, resultB1: %d not equal to expectedA1: %d, expetedB1: %d\n", resultA1, resultB1, expectedA1, expectedB1)
	}
	resultA2, resultB2 := Split2DigNum(8)
	if !(resultA2 == expectedA2 && resultB2 == expectedB2) {
		t.Errorf("resultA2: %d, resultB2: %d not equal to expectedA2: %d, expectedB2: %d\n", resultA2, resultB2, expectedA2, expectedB2)
	}
	resultA3, resultB3 := Split2DigNum(518)
	if !(resultA3 == expectedA3 && resultB3 == expectedB3) {
		t.Errorf("resultA3: %d, resultB3: %d not equal to expectedA1: %d, expetedB3: %d\n", resultA3, resultB3, expectedA3, expectedB3)
	}
}

func TestContainsAll(t *testing.T) {
	sa1 := []int{1, 5, 6, 8, 0, 7}
	sb1 := []int{2, 6, 4, 1, 3, 8}
	sa2 := []int{3, 6, 1, 4, 8, 7}
	sb2 := []int{4, 8, 3, 2, 0, 5}
	sa3 := []int{1, 2, 5, 6, 3, 0}
	sb3 := []int{2, 8, 7, 1, 4, 6}

	expected := []bool{true, false, true}
	result := []bool{ContainsAll(sa1, sb1), ContainsAll(sa2, sb2), ContainsAll(sa3, sb3)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result: %v but got %v\n", expected, result)
	}
}

func TestContains01(t *testing.T) {
	sa1 := []int{1, 5, 6, 8, 0}
	sb1 := []int{2, 8, 7, 0, 4}
	sa2 := []int{2, 5, 6, 8, 0}
	sb2 := []int{2, 8, 7, 5, 4}
	sa3 := []int{2, 1, 6, 8, 0}
	sb3 := []int{2, 8, 6, 2, 0}

	expected := []bool{true, false, true}
	result := []bool{Contains01(sa1, sb1), Contains01(sa2, sb2), Contains01(sa3, sb3)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result: %v but got %v\n", expected, result)
	}
}

func TestContains06(t *testing.T) {
	sa1 := []int{1, 5, 6, 8, 0}
	sb1 := []int{2, 8, 7, 0, 4}
	sa2 := []int{2, 5, 6, 8, 0}
	sb2 := []int{2, 8, 7, 5, 4}
	sa3 := []int{2, 5, 6, 8, 0}
	sb3 := []int{2, 8, 6, 2, 1}

	expected := []bool{true, false, true}
	result := []bool{Contains06(sa1, sb1), Contains06(sa2, sb2), Contains06(sa3, sb3)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result: %v but got %v\n", expected, result)
	}
}

func TestContains16(t *testing.T) {
	sa1 := []int{1, 5, 6, 8, 0}
	sb1 := []int{2, 8, 7, 1, 4}
	sa2 := []int{2, 5, 6, 8, 0}
	sb2 := []int{2, 8, 7, 5, 4}
	sa3 := []int{2, 5, 6, 8, 0}
	sb3 := []int{2, 8, 7, 2, 1}

	expected := []bool{true, false, true}
	result := []bool{Contains16(sa1, sb1), Contains16(sa2, sb2), Contains16(sa3, sb3)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result: %v but got %v\n", expected, result)
	}
}

func TestContains25(t *testing.T) {
	sa1 := []int{1, 5, 6, 8, 0}
	sb1 := []int{2, 8, 7, 1, 4}
	sa2 := []int{2, 5, 6, 8, 0}
	sb2 := []int{1, 8, 7, 7, 4}
	sa3 := []int{2, 7, 6, 8, 0}
	sb3 := []int{2, 5, 7, 2, 1}

	expected := []bool{true, false, true}
	result := []bool{Contains25(sa1, sb1), Contains25(sa2, sb2), Contains25(sa3, sb3)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result: %v but got %v\n", expected, result)
	}
}

func TestContains36(t *testing.T) {
	sa1 := []int{1, 5, 6, 8, 0}
	sb1 := []int{2, 3, 7, 1, 4}
	sa2 := []int{2, 5, 6, 8, 0}
	sb2 := []int{1, 8, 7, 7, 4}
	sa3 := []int{2, 7, 6, 3, 0}
	sb3 := []int{2, 3, 7, 6, 1}

	expected := []bool{true, false, true}
	result := []bool{Contains36(sa1, sb1), Contains36(sa2, sb2), Contains36(sa3, sb3)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result: %v but got %v\n", expected, result)
	}
}

func TestContains46(t *testing.T) {
	sa1 := []int{1, 5, 6, 8, 0}
	sb1 := []int{2, 3, 7, 1, 4}
	sa2 := []int{2, 5, 7, 8, 0}
	sb2 := []int{1, 8, 7, 7, 4}
	sa3 := []int{2, 4, 6, 3, 0}
	sb3 := []int{2, 3, 7, 6, 1}

	expected := []bool{true, false, true}
	result := []bool{Contains46(sa1, sb1), Contains46(sa2, sb2), Contains46(sa3, sb3)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result: %v but got %v\n", expected, result)
	}
}

func TestContains64(t *testing.T) {
	sa1 := []int{1, 5, 6, 8, 0}
	sb1 := []int{2, 3, 7, 1, 4}
	sa2 := []int{2, 5, 7, 8, 0}
	sb2 := []int{1, 8, 7, 7, 4}
	sa3 := []int{2, 4, 6, 3, 0}
	sb3 := []int{2, 3, 7, 6, 1}

	expected := []bool{true, false, true}
	result := []bool{Contains64(sa1, sb1), Contains64(sa2, sb2), Contains64(sa3, sb3)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result: %v but got %v\n", expected, result)
	}
}

func TestContains81(t *testing.T) {
	sa1 := []int{1, 5, 6, 8, 0}
	sb1 := []int{2, 3, 7, 1, 4}
	sa2 := []int{2, 5, 7, 8, 0}
	sb2 := []int{2, 8, 7, 7, 4}
	sa3 := []int{2, 4, 6, 3, 8}
	sb3 := []int{2, 3, 7, 6, 1}

	expected := []bool{true, false, true}
	result := []bool{Contains81(sa1, sb1), Contains81(sa2, sb2), Contains81(sa3, sb3)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result: %v but got %v\n", expected, result)
	}
}

func TestContainsN(t *testing.T) {
	sa := []int{1, 5, 6, 8, 0}
	sb := []int{2, 8, 7, 1, 4}
	val1 := 5
	val2 := 7
	val3 := 3
	expected := []int{1, 2, 0}
	result := []int{
		ContainsN(sa, sb, val1), ContainsN(sa, sb, val2), ContainsN(sa, sb, val3),
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result: %v but got %v\n", expected, result)
	}
}
func TestIntInSlice(t *testing.T) {
	s := []int{1, 5, 6, 8, 9, 0}
	val := 6
	if IntInSlice(s, val) != true {
		t.Errorf("IntInSlice() expected val %d found in %v to be true but got false\n", val, s)
	}
}
func TestSubset(t *testing.T) {
	var result, expected [][]int
	result = Subset([]int{0, 1, 2, 3}, 3)
	expected = [][]int{
		{0, 1, 2},
		{0, 1, 3},
		{0, 2, 3},
		{1, 2, 3},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result: %v\n expected: %v", result, expected)
	}
	result = Subset([]int{0, 1, 2, 3, 4}, 3)
	expected = [][]int{
		{0, 1, 2},
		{0, 1, 3},
		{0, 1, 4},
		{0, 2, 3},
		{0, 2, 4},
		{0, 3, 4},
		{1, 2, 3},
		{1, 2, 4},
		{1, 3, 4},
		{2, 3, 4},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result: %v\n expected: %v", result, expected)
	}
	result = Subset([]int{0, 1, 2}, 2)
	expected = [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result: %v\n expected: %v", result, expected)
	}
}
