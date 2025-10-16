package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name          string
		a             []int
		b             []int
		expectedBool  bool
		expectedSlice []int
	}{
		{
			name:          "пример из задания",
			a:             []int{65, 3, 58, 678, 64},
			b:             []int{64, 2, 3, 43},
			expectedBool:  true,
			expectedSlice: []int{64, 3},
		},
		{
			name:          "нет пересечений",
			a:             []int{1, 2, 3},
			b:             []int{4, 5, 6},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "полное пересечение",
			a:             []int{10, 20, 30},
			b:             []int{10, 20, 30},
			expectedBool:  true,
			expectedSlice: []int{10, 20, 30},
		},
		{
			name:          "частичное пересечение",
			a:             []int{1, 2, 3, 4, 5},
			b:             []int{3, 4, 5, 6, 7},
			expectedBool:  true,
			expectedSlice: []int{3, 4, 5},
		},
		{
			name:          "дубликаты в исходных слайсах",
			a:             []int{1, 2, 2, 3},
			b:             []int{2, 3, 3, 4},
			expectedBool:  true,
			expectedSlice: []int{2, 3, 3},
		},
		{
			name:          "один пустой слайс - первый",
			a:             []int{},
			b:             []int{1, 2, 3},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "один пустой слайс - второй",
			a:             []int{1, 2, 3},
			b:             []int{},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "оба пустых слайса",
			a:             []int{},
			b:             []int{},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "отрицательные числа",
			a:             []int{-1, -2, -3},
			b:             []int{-2, -3, -4},
			expectedBool:  true,
			expectedSlice: []int{-2, -3},
		},
		{
			name:          "ноль в пересечении",
			a:             []int{-1, 0, 1},
			b:             []int{0, 2, 3},
			expectedBool:  true,
			expectedSlice: []int{0},
		},
		{
			name:          "большие числа",
			a:             []int{1000000, 999999},
			b:             []int{999999, 1000001},
			expectedBool:  true,
			expectedSlice: []int{999999},
		},
		{
			name:          "один элемент в каждом слайсе - есть пересечение",
			a:             []int{42},
			b:             []int{42},
			expectedBool:  true,
			expectedSlice: []int{42},
		},
		{
			name:          "один элемент в каждом слайсе - нет пересечения",
			a:             []int{42},
			b:             []int{43},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "повторяющиеся элементы только в одном слайсе",
			a:             []int{1, 1, 1, 2},
			b:             []int{1, 2, 3},
			expectedBool:  true,
			expectedSlice: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasIntersection, common := Intersection(tt.a, tt.b)

			if hasIntersection != tt.expectedBool {
				t.Errorf("Intersection(%v, %v) bool = %t, expected %t",
					tt.a, tt.b, hasIntersection, tt.expectedBool)
			}

			// Для пустых слайсов проверяем только длину
			if len(common) == 0 && len(tt.expectedSlice) == 0 {
				return
			}

			if !reflect.DeepEqual(common, tt.expectedSlice) {
				t.Errorf("Intersection(%v, %v) slice = %v, expected %v",
					tt.a, tt.b, common, tt.expectedSlice)
			}
		})
	}
}

func TestIntersectionUnique(t *testing.T) {
	tests := []struct {
		name          string
		a             []int
		b             []int
		expectedBool  bool
		expectedSlice []int
	}{
		{
			name:          "удаление дубликатов",
			a:             []int{1, 2, 2, 3, 4},
			b:             []int{2, 3, 3, 5},
			expectedBool:  true,
			expectedSlice: []int{2, 3},
		},
		{
			name:          "без дубликатов в результате",
			a:             []int{1, 2, 3},
			b:             []int{2, 3, 4},
			expectedBool:  true,
			expectedSlice: []int{2, 3},
		},
		{
			name:          "нет пересечений",
			a:             []int{1, 2},
			b:             []int{3, 4},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "много дубликатов",
			a:             []int{1, 1, 1, 2, 2},
			b:             []int{1, 1, 2, 2, 2},
			expectedBool:  true,
			expectedSlice: []int{1, 2},
		},
		{
			name:          "пустые слайсы",
			a:             []int{},
			b:             []int{},
			expectedBool:  false,
			expectedSlice: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasIntersection, common := IntersectionUnique(tt.a, tt.b)

			if hasIntersection != tt.expectedBool {
				t.Errorf("IntersectionUnique(%v, %v) bool = %t, expected %t",
					tt.a, tt.b, hasIntersection, tt.expectedBool)
			}

			// Для пустых слайсов проверяем только длину
			if len(common) == 0 && len(tt.expectedSlice) == 0 {
				return
			}

			if !reflect.DeepEqual(common, tt.expectedSlice) {
				t.Errorf("IntersectionUnique(%v, %v) slice = %v, expected %v",
					tt.a, tt.b, common, tt.expectedSlice)
			}

			// Дополнительная проверка: в результате не должно быть дубликатов
			if hasIntersection {
				seen := make(map[int]bool)
				for _, num := range common {
					if seen[num] {
						t.Errorf("В результате есть дубликаты: %v", common)
						break
					}
					seen[num] = true
				}
			}
		})
	}
}

func TestIntersection_OrderIndependence(t *testing.T) {
	// Тест на независимость от порядка
	a := []int{1, 2, 3, 4}
	b := []int{3, 1, 5}

	hasIntersection, common := Intersection(a, b)

	if !hasIntersection {
		t.Error("Должно быть пересечение")
	}

	// Проверяем что все ожидаемые элементы присутствуют
	expectedElements := map[int]bool{1: true, 3: true}
	for _, num := range common {
		if !expectedElements[num] {
			t.Errorf("Неожиданный элемент в результате: %d", num)
		}
		delete(expectedElements, num)
	}

	// Проверяем что все ожидаемые элементы найдены
	if len(expectedElements) > 0 {
		t.Errorf("Не все ожидаемые элементы найдены: %v", expectedElements)
	}
}

func TestIntersection_Commutative(t *testing.T) {
	// Тест на коммутативность: Intersection(a, b) должно быть равно Intersection(b, a)
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}

	hasIntersection1, common1 := Intersection(a, b)
	hasIntersection2, common2 := Intersection(b, a)

	if hasIntersection1 != hasIntersection2 {
		t.Errorf("Коммутативность нарушена: Intersection(a,b) = %t, Intersection(b,a) = %t",
			hasIntersection1, hasIntersection2)
	}

	// Для слайсов порядок может отличаться, поэтому проверяем множества
	common1Set := make(map[int]bool)
	for _, num := range common1 {
		common1Set[num] = true
	}

	common2Set := make(map[int]bool)
	for _, num := range common2 {
		common2Set[num] = true
	}

	if !reflect.DeepEqual(common1Set, common2Set) {
		t.Errorf("Коммутативность нарушена: разные множества %v vs %v", common1Set, common2Set)
	}
}

func TestIntersection_Performance(t *testing.T) {
	// Тест на производительность с большими слайсами
	size := 10000
	a := make([]int, size)
	b := make([]int, size)

	for i := 0; i < size; i++ {
		a[i] = i
		b[i] = i + size/2 // Пересечение size/2 элементов
	}

	hasIntersection, common := Intersection(a, b)

	if !hasIntersection {
		t.Error("Должно быть пересечение")
	}

	if len(common) != size/2 {
		t.Errorf("Ожидалось %d пересекающихся элементов, получено %d", size/2, len(common))
	}
}

// Тесты граничных случаев
func TestIntersection_EdgeCases(t *testing.T) {
	t.Run("nil слайсы", func(t *testing.T) {
		// В Go передача nil в range безопасна
		hasIntersection, common := Intersection(nil, []int{1, 2, 3})
		if hasIntersection {
			t.Error("Intersection(nil, slice) должна возвращать false")
		}
		if common == nil || len(common) != 0 {
			t.Error("Intersection(nil, slice) должна возвращать пустой слайс")
		}

		hasIntersection2, common2 := Intersection([]int{1, 2, 3}, nil)
		if hasIntersection2 {
			t.Error("Intersection(slice, nil) должна возвращать false")
		}
		if common2 == nil || len(common2) != 0 {
			t.Error("Intersection(slice, nil) должна возвращать пустой слайс")
		}

		hasIntersection3, common3 := Intersection(nil, nil)
		if hasIntersection3 {
			t.Error("Intersection(nil, nil) должна возвращать false")
		}
		if common3 == nil || len(common3) != 0 {
			t.Error("Intersection(nil, nil) должна возвращать пустой слайс")
		}
	})

	t.Run("один элемент", func(t *testing.T) {
		hasIntersection, common := Intersection([]int{5}, []int{5})
		if !hasIntersection {
			t.Error("Должно быть пересечение для [5] и [5]")
		}
		if len(common) != 1 || common[0] != 5 {
			t.Errorf("Ожидалось [5], получено %v", common)
		}
	})

	t.Run("все элементы уникальны", func(t *testing.T) {
		a := []int{1, 3, 5, 7, 9}
		b := []int{2, 4, 6, 8, 10}
		hasIntersection, common := Intersection(a, b)
		if hasIntersection {
			t.Error("Не должно быть пересечения")
		}
		if len(common) != 0 {
			t.Errorf("Ожидался пустой слайс, получено %v", common)
		}
	})
}

// Benchmark тесты
func BenchmarkIntersection_Small(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	bSlice := []int{3, 4, 6, 7}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Intersection(a, bSlice)
	}
}

func BenchmarkIntersection_Medium(b *testing.B) {
	a := make([]int, 1000)
	bSlice := make([]int, 1000)

	for i := 0; i < 1000; i++ {
		a[i] = i
		bSlice[i] = i + 500
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Intersection(a, bSlice)
	}
}

func BenchmarkIntersection_Large(b *testing.B) {
	a := make([]int, 10000)
	bSlice := make([]int, 10000)

	for i := 0; i < 10000; i++ {
		a[i] = i
		bSlice[i] = i + 5000
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Intersection(a, bSlice)
	}
}

func BenchmarkIntersectionUnique_Small(b *testing.B) {
	a := []int{1, 2, 2, 3, 4}
	bSlice := []int{2, 3, 3, 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IntersectionUnique(a, bSlice)
	}
}

func BenchmarkIntersectionUnique_Large(b *testing.B) {
	a := make([]int, 10000)
	bSlice := make([]int, 10000)

	for i := 0; i < 10000; i++ {
		a[i] = i % 1000
		bSlice[i] = (i + 500) % 1000
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IntersectionUnique(a, bSlice)
	}
}

// Интеграционные тесты
func TestIntersection_Integration(t *testing.T) {
	t.Run("цепочка операций", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		b := []int{2, 4, 6, 8, 10, 12, 14}
		c := []int{6, 8, 10, 16, 18}

		hasAB, ab := Intersection(a, b)
		if !hasAB {
			t.Fatal("Должно быть пересечение между a и b")
		}

		hasFinal, final := Intersection(ab, c)
		if !hasFinal {
			t.Fatal("Должно быть пересечение между ab и c")
		}

		if len(final) != 3 {
			t.Errorf("Ожидалось 3 элемента, получено %d: %v", len(final), final)
		}
	})

	t.Run("уникальные значения после нескольких операций", func(t *testing.T) {
		a := []int{1, 1, 2, 2, 3, 3}
		b := []int{2, 2, 3, 4}
		c := []int{3, 4, 5}

		hasAB, ab := Intersection(a, b)
		if !hasAB {
			t.Fatal("Должно быть пересечение")
		}

		hasUnique, unique := IntersectionUnique(ab, c)
		if !hasUnique {
			t.Fatal("Должно быть пересечение")
		}

		expected := []int{3}
		if !reflect.DeepEqual(unique, expected) {
			t.Errorf("IntersectionUnique(ab, c) = %v, expected %v", unique, expected)
		}
	})
}
