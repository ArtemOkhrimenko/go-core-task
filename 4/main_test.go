package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "пример из задания",
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:     "пустой второй слайс",
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "пустой первый слайс",
			slice1:   []string{},
			slice2:   []string{"a", "b", "c"},
			expected: []string{},
		},
		{
			name:     "оба слайса пустые",
			slice1:   []string{},
			slice2:   []string{},
			expected: []string{},
		},
		{
			name:     "одинаковые слайсы",
			slice1:   []string{"x", "y", "z"},
			slice2:   []string{"x", "y", "z"},
			expected: []string{},
		},
		{
			name:     "нет общих элементов",
			slice1:   []string{"1", "2", "3"},
			slice2:   []string{"a", "b", "c"},
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "все элементы второго слайса есть в первом",
			slice1:   []string{"1", "2", "3", "4"},
			slice2:   []string{"2", "3"},
			expected: []string{"1", "4"},
		},
		{
			name:     "дубликаты в первом слайсе",
			slice1:   []string{"a", "b", "a", "c"},
			slice2:   []string{"b"},
			expected: []string{"a", "a", "c"},
		},
		{
			name:     "специальные символы и числа",
			slice1:   []string{"hello", "123", "!@#", "world"},
			slice2:   []string{"123", "world"},
			expected: []string{"hello", "!@#"},
		},
		{
			name:     "чувствительность к регистру",
			slice1:   []string{"Apple", "banana", "Cherry"},
			slice2:   []string{"apple", "BANANA", "cherry"},
			expected: []string{"Apple", "banana", "Cherry"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Difference(tt.slice1, tt.slice2)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Difference(%v, %v) = %v, expected %v",
					tt.slice1, tt.slice2, result, tt.expected)
			}
		})
	}
}

func TestDifference_OrderIndependence(t *testing.T) {
	// Тест на независимость от порядка элементов
	slice1 := []string{"a", "b", "c", "d"}
	slice2 := []string{"c", "a"}

	result1 := Difference(slice1, slice2)
	expected := []string{"b", "d"}

	if !reflect.DeepEqual(result1, expected) {
		t.Errorf("Difference не сохранил порядок или вернул неверный результат: %v", result1)
	}

	// Меняем порядок во втором слайсе
	slice2Reordered := []string{"a", "c"}
	result2 := Difference(slice1, slice2Reordered)

	if !reflect.DeepEqual(result2, expected) {
		t.Errorf("Difference зависит от порядка во втором слайсе: %v", result2)
	}
}

// Benchmark тесты
func BenchmarkDifference_Small(b *testing.B) {
	slice1 := []string{"a", "b", "c", "d", "e"}
	slice2 := []string{"b", "d"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Difference(slice1, slice2)
	}
}

func BenchmarkDifference_Medium(b *testing.B) {
	slice1 := make([]string, 1000)
	slice2 := make([]string, 500)

	for i := 0; i < 1000; i++ {
		slice1[i] = fmt.Sprintf("item%d", i)
	}
	for i := 0; i < 500; i++ {
		slice2[i] = fmt.Sprintf("item%d", i*2)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Difference(slice1, slice2)
	}
}

func BenchmarkDifference_Large(b *testing.B) {
	slice1 := make([]string, 10000)
	slice2 := make([]string, 5000)

	for i := 0; i < 10000; i++ {
		slice1[i] = fmt.Sprintf("item%d", i)
	}
	for i := 0; i < 5000; i++ {
		slice2[i] = fmt.Sprintf("item%d", i*2)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Difference(slice1, slice2)
	}
}
