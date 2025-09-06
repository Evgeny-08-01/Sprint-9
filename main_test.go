package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name        string
		lenReal     int64
		lenExpected int64
	}{
		{"Slice length matching test=1", 1, 1}, {"Slice length matching test=7", 7, 7},
		{"Slice length matching test=8", 8, 8}, {"Slice length matching test=9", 9, 9},
		{"Slice length matching test =100", 100, 100}, {"Slice length matching test=300", 300, 300},
		{"Slice length matching test=500", 500, 500}, {"Slice length matching test=1000", 1000, 1000},
		{"Slice length matching test=5000", 5000, 5000}, {"Slice length matching test=20000", 20000, 20000},
		{"Slice length matching test=100000000", 100000000, 100000000}}
	for _, tt := range tests {
		resultReal := generateRandomElements(tt.lenReal)
		t.Run(tt.name, func(t *testing.T) {
			assert.Len(t, resultReal, int(tt.lenExpected), fmt.Sprintf("Expected length: %d, Length: %d", tt.lenExpected, len(resultReal)))
		})
	}
}
func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		in       []int64
		expected int64
	}{
		{"Test with positive numbers", []int64{0, 1, 2, 3, 4, 5, 6, 7, 8,
			9, 10, 100, 200, 300, 500, 1000, 10000, 100000, 100000000}, 100000000},
		{"Test with the same numbers", []int64{20, 20, 20, 20, 20, 20, 20}, 20},
		{"Test with an empty slice", []int64{}, 0},
		{"Single-number test", []int64{200}, 200},
	}
	for _, tt := range tests {
		t.Run((tt.name), func(t *testing.T) {
			assert.Equal(t, tt.expected, maximum(tt.in), tt.name+".  Test failed")
		})
	}
}
func TestComparisonMax(t *testing.T) {
	testData := generateRandomElements(100000000)
	max1 := maximum(testData)
	max2 := maxChunks(testData)
	assert.Equal(t, max1, max2, "Maximum search results are different. Test failed")
}
