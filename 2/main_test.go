package main

import (
	"reflect"
	"testing"
)

// Тесты для createRandomSlice
func TestCreateRandomSlice(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"длина 10", 10},
		{"длина 0", 0},
		{"длина 1", 1},
		{"длина 100", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slice := createRandomSlice(tt.length)

			// Проверяем длину
			if len(slice) != tt.length {
				t.Errorf("createRandomSlice(%d) создал слайс длиной %d, ожидалось %d",
					tt.length, len(slice), tt.length)
			}

			// Проверяем, что все элементы в допустимом диапазоне
			for i, value := range slice {
				if value < 0 || value >= 100 {
					t.Errorf("элемент [%d] = %d вне диапазона [0, 99]", i, value)
				}
			}
		})
	}
}

// Тесты для sliceExample (фильтрация четных чисел)
func TestSliceExample(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "только четные числа",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{2, 4, 6},
		},
		{
			name:     "все нечетные",
			input:    []int{1, 3, 5, 7},
			expected: []int{},
		},
		{
			name:     "все четные",
			input:    []int{2, 4, 6, 8},
			expected: []int{2, 4, 6, 8},
		},
		{
			name:     "пустой слайс",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "с нулями",
			input:    []int{0, 1, 0, 3},
			expected: []int{0, 0},
		},
		{
			name:     "отрицательные числа",
			input:    []int{-2, -1, 0, 1, 2},
			expected: []int{-2, 0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceExample(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sliceExample(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

// Тесты для addElements
func TestAddElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		num      int
		expected []int
	}{
		{
			name:     "добавление к непустому слайсу",
			input:    []int{1, 2, 3},
			num:      4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "добавление к пустому слайсу",
			input:    []int{},
			num:      1,
			expected: []int{1},
		},
		{
			name:     "добавление отрицательного числа",
			input:    []int{1, 2},
			num:      -3,
			expected: []int{1, 2, -3},
		},
		{
			name:     "добавление нуля",
			input:    []int{1, 2, 3},
			num:      0,
			expected: []int{1, 2, 3, 0},
		},
		{
			name:     "добавление к слайсу с одним элементом",
			input:    []int{42},
			num:      100,
			expected: []int{42, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addElements(tt.input, tt.num)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("addElements(%v, %d) = %v, expected %v", tt.input, tt.num, result, tt.expected)
			}
		})
	}
}

// Тесты для copySlice
func TestCopySlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "копирование непустого слайса",
			input: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "копирование пустого слайса",
			input: []int{},
		},
		{
			name:  "копирование слайса с одним элементом",
			input: []int{42},
		},
		{
			name:  "копирование слайса с отрицательными числами",
			input: []int{-5, -3, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Сохраняем оригинальные значения
			original := make([]int, len(tt.input))
			copy(original, tt.input)

			result := copySlice(tt.input)

			// Проверяем, что содержимое одинаковое
			if !reflect.DeepEqual(result, tt.input) {
				t.Errorf("copySlice(%v) = %v, expected %v", tt.input, result, tt.input)
			}

			// Проверяем независимость: изменяем оригинал и проверяем, что копия не изменилась
			if len(tt.input) > 0 {
				originalValue := tt.input[0]
				tt.input[0] = 999
				if result[0] == 999 {
					t.Error("изменения в оригинальном слайсе не должны влиять на копию")
				}
				// Восстанавливаем оригинал
				tt.input[0] = originalValue
			}
		})
	}
}

// Тесты для removeElement
func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		index       int
		expected    []int
		expectError bool
	}{
		{
			name:        "удаление из середины",
			input:       []int{1, 2, 3, 4, 5},
			index:       2,
			expected:    []int{1, 2, 4, 5},
			expectError: false,
		},
		{
			name:        "удаление первого элемента",
			input:       []int{1, 2, 3},
			index:       0,
			expected:    []int{2, 3},
			expectError: false,
		},
		{
			name:        "удаление последнего элемента",
			input:       []int{1, 2, 3},
			index:       2,
			expected:    []int{1, 2},
			expectError: false,
		},
		{
			name:        "отрицательный индекс",
			input:       []int{1, 2, 3},
			index:       -1,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "индекс больше длины",
			input:       []int{1, 2, 3},
			index:       5,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "индекс равен длине",
			input:       []int{1, 2, 3},
			index:       3,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "удаление из слайса с одним элементом",
			input:       []int{42},
			index:       0,
			expected:    []int{},
			expectError: false,
		},
		{
			name:        "удаление из пустого слайса",
			input:       []int{},
			index:       0,
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := removeElement(tt.input, tt.index)

			if tt.expectError {
				if err == nil {
					t.Errorf("removeElement(%v, %d) ожидалась ошибка, но получен nil", tt.input, tt.index)
				}
			} else {
				if err != nil {
					t.Errorf("removeElement(%v, %d) неожиданная ошибка: %v", tt.input, tt.index, err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("removeElement(%v, %d) = %v, expected %v", tt.input, tt.index, result, tt.expected)
				}
			}
		})
	}
}

// Тест на порядок элементов после удаления
func TestRemoveElement_OrderPreservation(t *testing.T) {
	input := []int{10, 20, 30, 40, 50}
	result, err := removeElement(input, 2) // Удаляем 30

	if err != nil {
		t.Fatalf("removeElement вернул ошибку: %v", err)
	}

	expected := []int{10, 20, 40, 50}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeElement нарушил порядок элементов: %v, expected %v", result, expected)
	}
}

// Тесты на граничные случаи
func TestEdgeCases(t *testing.T) {
	t.Run("createRandomSlice с длиной 0", func(t *testing.T) {
		result := createRandomSlice(0)
		if result == nil {
			t.Error("createRandomSlice(0) не должен возвращать nil")
		}
		if len(result) != 0 {
			t.Errorf("createRandomSlice(0) должен возвращать пустой слайс")
		}
	})

	t.Run("sliceExample с nil слайсом", func(t *testing.T) {
		// В Go передача nil в range безопасна
		result := sliceExample(nil)
		if result == nil {
			t.Error("sliceExample(nil) не должен возвращать nil")
		}
		if len(result) != 0 {
			t.Errorf("sliceExample(nil) должен возвращать пустой слайс")
		}
	})

	t.Run("addElements с nil слайсом", func(t *testing.T) {
		result := addElements(nil, 42)
		if len(result) != 1 || result[0] != 42 {
			t.Errorf("addElements(nil, 42) = %v, expected [42]", result)
		}
	})

	t.Run("copySlice с nil слайсом", func(t *testing.T) {
		result := copySlice(nil)
		if result == nil {
			t.Error("copySlice(nil) не должен возвращать nil")
		}
		if len(result) != 0 {
			t.Errorf("copySlice(nil) должен возвращать пустой слайс")
		}
	})

	t.Run("removeElement с nil слайсом", func(t *testing.T) {
		result, err := removeElement(nil, 0)
		if err == nil {
			t.Error("removeElement(nil, 0) должен возвращать ошибку")
		}
		if result != nil {
			t.Error("removeElement(nil, 0) должен возвращать nil")
		}
	})
}

// Benchmark тесты
func BenchmarkSliceExample(b *testing.B) {
	testSlice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		testSlice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceExample(testSlice)
	}
}

func BenchmarkAddElements(b *testing.B) {
	testSlice := []int{1, 2, 3, 4, 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		addElements(testSlice, i)
	}
}

func BenchmarkCopySlice(b *testing.B) {
	testSlice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		testSlice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copySlice(testSlice)
	}
}

func BenchmarkRemoveElement(b *testing.B) {
	testSlice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		testSlice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		removeElement(testSlice, 500)
	}
}

func BenchmarkCreateRandomSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createRandomSlice(100)
	}
}

// Интеграционные тесты
func TestIntegration(t *testing.T) {
	t.Run("цепочка операций", func(t *testing.T) {
		// Создаем случайный слайс
		original := createRandomSlice(20)
		if len(original) != 20 {
			t.Fatalf("Не удалось создать слайс длиной 20")
		}

		// Фильтруем четные числа
		even := sliceExample(original)

		// Добавляем новый элемент
		extended := addElements(even, 999)

		// Копируем результат
		copied := copySlice(extended)

		// Удаляем первый элемент (если он есть)
		if len(copied) > 0 {
			final, err := removeElement(copied, 0)
			if err != nil {
				t.Fatalf("Ошибка при удалении элемента: %v", err)
			}

			// Проверяем, что после удаления длина уменьшилась на 1
			if len(final) != len(copied)-1 {
				t.Errorf("После удаления длина должна быть %d, но получили %d",
					len(copied)-1, len(final))
			}
		}
	})

	t.Run("все функции с фиксированными данными", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		// Фильтруем четные
		even := sliceExample(input)
		expectedEven := []int{2, 4, 6, 8, 10}
		if !reflect.DeepEqual(even, expectedEven) {
			t.Errorf("sliceExample(%v) = %v, expected %v", input, even, expectedEven)
		}

		// Добавляем число
		extended := addElements(even, 99)
		expectedExtended := []int{2, 4, 6, 8, 10, 99}
		if !reflect.DeepEqual(extended, expectedExtended) {
			t.Errorf("addElements(%v, 99) = %v, expected %v", even, extended, expectedExtended)
		}

		// Копируем
		copied := copySlice(extended)
		if !reflect.DeepEqual(copied, expectedExtended) {
			t.Errorf("copySlice(%v) = %v, expected %v", extended, copied, expectedExtended)
		}

		// Удаляем элемент
		final, err := removeElement(copied, 2) // Удаляем 6
		if err != nil {
			t.Fatalf("removeElement вернул ошибку: %v", err)
		}
		expectedFinal := []int{2, 4, 8, 10, 99}
		if !reflect.DeepEqual(final, expectedFinal) {
			t.Errorf("removeElement(%v, 2) = %v, expected %v", copied, final, expectedFinal)
		}
	})
}
