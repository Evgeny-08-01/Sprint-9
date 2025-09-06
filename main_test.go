package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name        string
		lenReal     int64
		lenExpected int64
	}{
		{"Тест на соответствие длины среза=1", 1, 1}, {"Тест на соответствие длины среза=7", 7, 7},
		{"Тест на соответствие длины среза=8", 8, 8}, {"Тест на соответствие длины среза=9", 9, 9},
		{"Тест на соответствие длины среза=100", 100, 100}, {"Тест на соответствие длины среза=300", 300, 300},
		{"Тест на соответствие длины среза=500", 500, 500}, {"Тест на соответствие длины среза=1000", 1000, 1000},
		{"Тест на соответствие длины среза=5000", 5000, 5000}, {"Тест на соответствие длины среза=20000", 20000, 20000},
		{"Тест на соответствие длины среза=100000000", 100000000, 100000000}}

	for _, tt := range tests {
		resultReal := generateRandomElements(tt.lenReal)
		t.Run(tt.name, func(t *testing.T) {
			assert.Len(t, resultReal, int(tt.lenExpected), fmt.Sprintf("Ожидаемая длина: %d, фактическая: %d", tt.lenExpected, len(resultReal)))
		})

	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		in       []int64
		expected int64
	}{
		{
			name:     "тест с положительными числами",
			in:       []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 200, 300, 500, 1000, 10000, 100000, 100000000},
			expected: 100000000},
		{
			name:     "тест с одинаковыми числами",
			in:       []int64{20, 20, 20, 20, 20, 20, 20},
			expected: 20},
		{
			name:     "тест с одним числом",
			in:       []int64{200},
			expected: 200},
		{
			name:     "тест с пустым срезом",
			in:       []int64{},
			expected: 0},
	}
	for _, tt := range tests {
		t.Run((tt.name), func(t *testing.T) {
			assert.Equal(t, tt.expected, maximum(tt.in), tt.name+"тест не пройден")
		})
	}
}

func TestComparisonMax(t *testing.T) {
	// Генерируем тестовый массив
	testData := generateRandomElements(100000000)
	max1 := maximum(testData)
	max2 := maxChunks(testData)
	assert.Equal(t, max1, max2, "Результаты поиска максимума отличаются. Ошибка")
}

func Test_maxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int64
		want int64
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if true {
				t.Errorf("maxChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
