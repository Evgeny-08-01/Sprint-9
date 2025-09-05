package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		lenReal     int
		lenExpected int
	}{
		{1, 1}, {7, 7}, {8, 8}, {9, 9}, {100, 100}, {300, 300}, {500, 500}, {1000, 1000}, {5000, 5000},
		{20000, 20000}, {100000000, 100000000}}

	for _, tt := range tests {
		resultReal := generateRandomElements(tt.lenReal)
		t.Run(fmt.Sprintf("Тест на соответствие длины среза %d", tt.lenReal), func(t *testing.T) {
			require.Len(t, resultReal, tt.lenExpected, fmt.Sprintf("Ожидаемая длина: %d, фактическая: %d", tt.lenExpected, len(resultReal)))
		})
		// расчет среднего значения
		sum := 0
		for _, value := range resultReal {
			sum += value
		}
		average := float64(sum) / float64(len(resultReal))
		if len(resultReal) > 500 {
			minValue := average * 0.9
			maxValue := average * 1.1
			if average < minValue || average > maxValue {
				t.Errorf("Ожидаемое значение вне пределов ожидаемого распределения от  %.2f до  %.2f, получено  %.2f", minValue, maxValue, average)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		in       []int
		expected int
	}{
		{
			name:     "тест с положительными числами",
			in:       []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 200, 300, 500, 1000, 10000, 100000, 100000000},
			expected: 100000000},
		{
			name:     "тест с одинаковыми числами",
			in:       []int{20, 20, 20, 20, 20, 20, 20},
			expected: 20},
		{
			name:     "тест с одним числом",
			in:       []int{200},
			expected: 200},
		{
			name:     "тест с пустым срезом",
			in:       []int{},
			expected: 0},
	}
	for _, tt := range tests {
		t.Run((tt.name), func(t *testing.T) {
			require.Equal(t, tt.expected, maximum(tt.in), tt.name+"тест не пройден")
		})
	}
}
func TestComparisonMax(t *testing.T) {
	// Генерируем тестовый массив
	testData := generateRandomElements(100000000)
	max1 := maximum(testData)
	max2 := maxChunks(testData)
	require.Equal(t, max1, max2, "Результаты поиска максимума отличаются. Ошибка")
}
