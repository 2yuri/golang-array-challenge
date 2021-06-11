package service_test

import (
	"log"
	"testing"

	"github.com/hyperyuri/interview-test/service"
)

func TestMaxValueFromArray(t *testing.T) {
	array := [][]int{
		{3, -1, -1, 4, 3, -1},
		{1, 2, 4, 5, 7},
		{14, -2, 4, -51},
	}

	expectedValues := []int{
		4,
		7,
		14,
	}

	service := service.NewArrayService()

	for i, val := range array {
		result := service.MaxValueFromArray(val)
		if result != expectedValues[i] {
			t.Errorf("expected to be equal %v, but receive a %v", expectedValues[i], result)
		}
	}
}

func TestArrayChallenge(t *testing.T) {
	array := [][]int{
		{3, -1, -1, 4, 3, -1},
		{1, -2, 0, 3},
		{-2, 5, -1, 7, -3},
	}

	expectedValues := []int{
		8,
		3,
		11,
	}

	service := service.NewArrayService()

	for i, val := range array {
		result := service.ArrayChallenge(val)
		if result != expectedValues[i] {
			t.Errorf("expected to be equal %v, but receive a %v", expectedValues[i], result)
		}
	}
}

func TestBigestSumOfArray(t *testing.T) {
	array := [][]int{
		{3, -1, -1, 4, 3, -1},
		{1, 2, 4, 5, 7},
		{14, -2, 4, -51},
	}

	expectedValues := []int{
		8,
		19,
		16,
	}

	service := service.NewArrayService()

	for i := 0; i < len(array); i++ {
		result := service.BigestSumOfArray(array[i])
		if result != expectedValues[i] {
			log.Printf("expected to be equal %v, but receive a %v", expectedValues[i], result)
			t.Errorf("expected to be equal %v, but receive a %v", expectedValues[i], result)
		}
	}
}

func TestRemoveIndexesFromArray(t *testing.T) {
	array := [][]int{
		{3, -1, -1, 4, 3, -1},
		{1, 2, 4, 5, 7},
		{14, -2, 4, -51},
	}

	fakeIndexes := []int{
		2,
		0,
		1,
	}

	ExpectedValues := [][]int{
		{4, 3, -1},
		{2, 4, 5, 7},
		{4, -51},
	}

	service := service.NewArrayService()

	for i, val := range array {
		result := service.RemoveIndexesFromArray(val, fakeIndexes[i])
		if !equalArrays(result, ExpectedValues[i]) {
			t.Errorf("expected to be equal %v, but receive a %v in the index %v", ExpectedValues[i], result, fakeIndexes[i])
		}
	}
}

func equalArrays(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
