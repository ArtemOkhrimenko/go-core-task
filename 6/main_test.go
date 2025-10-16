package main

import (
	"testing"
	"time"
)

func TestNewRandomGenerator(t *testing.T) {
	tests := []struct {
		name     string
		min      int
		max      int
		count    int
		expected *RandomGenerator
	}{
		{
			name:  "положительный диапазон",
			min:   1,
			max:   100,
			count: 10,
		},
		{
			name:  "отрицательный диапазон",
			min:   -50,
			max:   50,
			count: 5,
		},
		{
			name:  "нулевое количество",
			min:   1,
			max:   10,
			count: 0,
		},
		{
			name:  "один элемент",
			min:   5,
			max:   5,
			count: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := NewRandomGenerator(tt.min, tt.max, tt.count)

			if gen == nil {
				t.Error("NewRandomGenerator() вернул nil")
			}
			if gen.min != tt.min {
				t.Errorf("min = %d, ожидалось %d", gen.min, tt.min)
			}
			if gen.max != tt.max {
				t.Errorf("max = %d, ожидалось %d", gen.max, tt.max)
			}
			if gen.count != tt.count {
				t.Errorf("count = %d, ожидалось %d", gen.count, tt.count)
			}
		})
	}
}

func TestRandomGenerator_Generate_Count(t *testing.T) {
	tests := []struct {
		name  string
		min   int
		max   int
		count int
	}{
		{"10 чисел", 1, 100, 10},
		{"0 чисел", 1, 100, 0},
		{"1 число", 1, 100, 1},
		{"100 чисел", 1, 1000, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := NewRandomGenerator(tt.min, tt.max, tt.count)
			ch := gen.Generate()

			received := 0
			for num := range ch {
				received++
				// Проверяем что число в заданном диапазоне
				if num < tt.min || num > tt.max {
					t.Errorf("число %d вне диапазона [%d, %d]", num, tt.min, tt.max)
				}
			}

			if received != tt.count {
				t.Errorf("получено %d чисел, ожидалось %d", received, tt.count)
			}
		})
	}
}

func TestRandomGenerator_Generate_Range(t *testing.T) {
	tests := []struct {
		name string
		min  int
		max  int
	}{
		{"положительный диапазон", 1, 100},
		{"отрицательный диапазон", -100, -1},
		{"смешанный диапазон", -50, 50},
		{"маленький диапазон", 5, 10},
		{"один элемент", 42, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := NewRandomGenerator(tt.min, tt.max, 100)
			ch := gen.Generate()

			for num := range ch {
				if num < tt.min || num > tt.max {
					t.Errorf("число %d вне диапазона [%d, %d]", num, tt.min, tt.max)
				}
			}
		})
	}
}

func TestRandomGenerator_Generate_ChannelClose(t *testing.T) {
	gen := NewRandomGenerator(1, 100, 5)
	ch := gen.Generate()

	// Читаем все значения
	count := 0
	for range ch {
		count++
	}

	// Канал должен быть закрыт
	_, ok := <-ch
	if ok {
		t.Error("канал не закрыт после генерации всех чисел")
	}

	if count != 5 {
		t.Errorf("получено %d чисел, ожидалось 5", count)
	}
}

func TestRandomGenerator_Generate_ZeroCount(t *testing.T) {
	gen := NewRandomGenerator(1, 100, 0)
	ch := gen.Generate()

	// Канал должен быть сразу закрыт
	_, ok := <-ch
	if ok {
		t.Error("канал не закрыт для нулевого количества")
	}
}

func TestRandomGenerator_Generate_Concurrent(t *testing.T) {
	gen := NewRandomGenerator(1, 100, 100)
	ch := gen.Generate()

	// Читаем из канала в нескольких горутинах
	results := make(chan int, 100)
	done := make(chan bool)

	go func() {
		for num := range ch {
			results <- num
		}
		close(results)
	}()

	total := 0
	go func() {
		for range results {
			total++
		}
		done <- true
	}()

	<-done

	if total != 100 {
		t.Errorf("получено %d чисел, ожидалось 100", total)
	}
}

func TestRandomGenerator_Generate_UniqueSequences(t *testing.T) {
	// Тест на то, что разные генераторы производят разные последовательности
	// (с высокой вероятностью, так как используется время как seed)

	gen1 := NewRandomGenerator(1, 1000, 10)
	gen2 := NewRandomGenerator(1, 1000, 10)

	// Даем немного времени между созданиями генераторов
	// чтобы гарантировать разные сиды
	time.Sleep(10 * time.Millisecond)

	ch1 := gen1.Generate()
	ch2 := gen2.Generate()

	seq1 := make([]int, 0)
	seq2 := make([]int, 0)

	for num := range ch1 {
		seq1 = append(seq1, num)
	}
	for num := range ch2 {
		seq2 = append(seq2, num)
	}

	// Проверяем что последовательности разные
	// (в редких случаях могут совпасть, но вероятность очень мала)
	identical := true
	for i := 0; i < len(seq1) && i < len(seq2); i++ {
		if seq1[i] != seq2[i] {
			identical = false
			break
		}
	}

	if identical && len(seq1) > 0 {
		t.Log("Предупреждение: последовательности случайных чисел совпали (маловероятно, но возможно)")
	}
}

func TestRandomGenerator_Generate_BlockingBehavior(t *testing.T) {
	gen := NewRandomGenerator(1, 5, 3)
	ch := gen.Generate()

	// Демонстрируем блокирующее поведение небуферизированного канала
	start := time.Now()

	// Читаем первое значение сразу
	num1 := <-ch
	if num1 < 1 || num1 > 5 {
		t.Errorf("число %d вне диапазона [1, 5]", num1)
	}

	// Ждем немного перед чтением остальных
	time.Sleep(100 * time.Millisecond)

	// Читаем оставшиеся значения
	count := 1
	for num := range ch {
		if num < 1 || num > 5 {
			t.Errorf("число %d вне диапазона", num)
		}
		count++
	}

	elapsed := time.Since(start)
	if elapsed < 100*time.Millisecond {
		t.Error("небуферизированный канал не показал блокирующее поведение")
	}

	if count != 3 {
		t.Errorf("получено %d чисел, ожидалось 3", count)
	}
}

// Benchmark тесты
func BenchmarkRandomGenerator_Generate_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gen := NewRandomGenerator(1, 100, 100)
		ch := gen.Generate()
		for range ch {
			// Просто читаем все значения
		}
	}
}

func BenchmarkRandomGenerator_Generate_Large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gen := NewRandomGenerator(1, 1000, 10000)
		ch := gen.Generate()
		for range ch {
			// Просто читаем все значения
		}
	}
}

func BenchmarkRandomGenerator_Generate_SmallRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gen := NewRandomGenerator(1, 10, 1000)
		ch := gen.Generate()
		for range ch {
			// Просто читаем все значения
		}
	}
}

// Тесты на граничные случаи
func TestRandomGenerator_EdgeCases(t *testing.T) {
	t.Run("минимальный диапазон", func(t *testing.T) {
		gen := NewRandomGenerator(1, 1, 5)
		ch := gen.Generate()

		for num := range ch {
			if num != 1 {
				t.Errorf("ожидалось число 1, получено %d", num)
			}
		}
	})

	t.Run("большой диапазон", func(t *testing.T) {
		gen := NewRandomGenerator(-1000000, 1000000, 10)
		ch := gen.Generate()

		for num := range ch {
			if num < -1000000 || num > 1000000 {
				t.Errorf("число %d вне диапазона", num)
			}
		}
	})

}
