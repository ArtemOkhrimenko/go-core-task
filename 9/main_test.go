package main

import (
	"sync"
	"testing"
)

// TestProcessPipeline тестирует обработку конвейера
func TestProcessPipeline(t *testing.T) {
	tests := []struct {
		name     string
		input    []uint8
		expected []float64
	}{
		{
			name:     "basic numbers",
			input:    []uint8{1, 2, 3},
			expected: []float64{1, 8, 27},
		},
		{
			name:     "zero and one",
			input:    []uint8{0, 1},
			expected: []float64{0, 1},
		},
		{
			name:     "empty input",
			input:    []uint8{},
			expected: []float64{},
		},
		{
			name:     "larger numbers",
			input:    []uint8{5, 10},
			expected: []float64{125, 1000},
		},
		{
			name:     "maximum uint8",
			input:    []uint8{255},
			expected: []float64{16581375}, // 255 * 255 * 255
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем каналы
			inputChan := make(chan uint8)
			outputChan := make(chan float64)

			// Запускаем обработку в отдельной горутине
			go ProcessPipeline(inputChan, outputChan)

			// Отправляем данные в отдельной горутине
			go func() {
				for _, num := range tt.input {
					inputChan <- num
				}
				close(inputChan)
			}()

			// Собираем результаты
			var results []float64
			for result := range outputChan {
				results = append(results, result)
			}

			// Проверяем результаты
			if len(results) != len(tt.expected) {
				t.Errorf("Expected %d results, got %d", len(tt.expected), len(results))
				return
			}

			for i, expected := range tt.expected {
				if results[i] != expected {
					t.Errorf("At index %d: expected %.2f, got %.2f", i, expected, results[i])
				}
			}
		})
	}
}

// TestGenerateNumbers тестирует генератор чисел
func TestGenerateNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    []uint8
		expected []uint8
	}{
		{
			name:     "normal case",
			input:    []uint8{1, 2, 3, 4, 5},
			expected: []uint8{1, 2, 3, 4, 5},
		},
		{
			name:     "empty case",
			input:    []uint8{},
			expected: []uint8{},
		},
		{
			name:     "single element",
			input:    []uint8{42},
			expected: []uint8{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := GenerateNumbers(tt.input)

			var results []uint8
			for num := range output {
				results = append(results, num)
			}

			if len(results) != len(tt.expected) {
				t.Errorf("Expected %d numbers, got %d", len(tt.expected), len(results))
				return
			}

			for i, expected := range tt.expected {
				if results[i] != expected {
					t.Errorf("At index %d: expected %d, got %d", i, expected, results[i])
				}
			}
		})
	}
}

// TestCollectResults тестирует сбор результатов
func TestCollectResults(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected []float64
	}{
		{
			name:     "normal case",
			input:    []float64{1.0, 8.0, 27.0},
			expected: []float64{1.0, 8.0, 27.0},
		},
		{
			name:     "empty case",
			input:    []float64{},
			expected: []float64{},
		},
		{
			name:     "fractional results",
			input:    []float64{1.5, 2.5, 3.5},
			expected: []float64{1.5, 2.5, 3.5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputChan := make(chan float64)

			go func() {
				for _, num := range tt.input {
					inputChan <- num
				}
				close(inputChan)
			}()

			results := CollectResults(inputChan)

			if len(results) != len(tt.expected) {
				t.Errorf("Expected %d results, got %d", len(tt.expected), len(results))
				return
			}

			for i, expected := range tt.expected {
				if results[i] != expected {
					t.Errorf("At index %d: expected %.2f, got %.2f", i, expected, results[i])
				}
			}
		})
	}
}

// TestRunPipeline тестирует полный конвейер
func TestRunPipeline(t *testing.T) {
	tests := []struct {
		name     string
		input    []uint8
		expected []float64
	}{
		{
			name:     "basic pipeline",
			input:    []uint8{1, 2, 3, 4, 5},
			expected: []float64{1, 8, 27, 64, 125},
		},
		{
			name:     "empty pipeline",
			input:    []uint8{},
			expected: []float64{},
		},
		{
			name:     "single element",
			input:    []uint8{10},
			expected: []float64{1000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := RunPipeline(tt.input)

			if len(results) != len(tt.expected) {
				t.Errorf("Expected %d results, got %d", len(tt.expected), len(results))
				return
			}

			for i, expected := range tt.expected {
				if results[i] != expected {
					t.Errorf("At index %d: expected %.2f, got %.2f", i, expected, results[i])
				}
			}
		})
	}
}

// TestProcessPipelineConcurrent тестирует конкурентную обработку
func TestProcessPipelineConcurrent(t *testing.T) {
	inputChan := make(chan uint8)
	outputChan := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(1)

	// Запускаем обработку
	go ProcessPipeline(inputChan, outputChan)

	// Запускаем сбор результатов
	var results []float64
	go func() {
		for result := range outputChan {
			results = append(results, result)
		}
		wg.Done()
	}()

	// Отправляем данные
	testData := []uint8{1, 2, 3, 4, 5}
	expected := []float64{1, 8, 27, 64, 125}

	for _, num := range testData {
		inputChan <- num
	}
	close(inputChan)

	// Ждем завершения
	wg.Wait()

	// Проверяем результаты
	if len(results) != len(expected) {
		t.Errorf("Expected %d results, got %d", len(expected), len(results))
		return
	}

	for i, expectedVal := range expected {
		if results[i] != expectedVal {
			t.Errorf("At index %d: expected %.2f, got %.2f", i, expectedVal, results[i])
		}
	}
}

// BenchmarkPipeline бенчмарк для измерения производительности конвейера
func BenchmarkPipeline(b *testing.B) {
	// Создаем тестовые данные
	testData := make([]uint8, 1000)
	for i := range testData {
		testData[i] = uint8(i % 256)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RunPipeline(testData)
	}
}
