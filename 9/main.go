package main

import (
	"fmt"
)

func main() {
	// Тестовые данные
	testNumbers := []uint8{1, 2, 3, 4, 5, 10}

	fmt.Println("Входные данные (uint8):", testNumbers)

	// Запускаем конвейер
	results := RunPipeline(testNumbers)

	fmt.Println("Результаты (float64, возведенные в куб):")
	for i, num := range testNumbers {
		fmt.Printf("  %d³ = %.2f\n", num, results[i])
	}

	// Демонстрация работы с большими числами
	fmt.Println("\nДемонстрация с большими числами:")
	bigNumbers := []uint8{10, 20, 30, 40, 50}
	bigResults := RunPipeline(bigNumbers)

	for i, num := range bigNumbers {
		fmt.Printf("  %d³ = %.2f\n", num, bigResults[i])
	}

	// Демонстрация с пустым слайсом
	fmt.Println("\nДемонстрация с пустым слайсом:")
	emptyResults := RunPipeline([]uint8{})
	fmt.Println("  Результаты:", emptyResults)
}

// ProcessPipeline создает конвейер обработки чисел
func ProcessPipeline(input <-chan uint8, output chan<- float64) {
	defer close(output)

	for num := range input {
		// Преобразуем uint8 в float64 и возводим в куб
		result := float64(num) * float64(num) * float64(num)
		output <- result
	}
}

// GenerateNumbers генерирует последовательность чисел и отправляет в канал
func GenerateNumbers(numbers []uint8) <-chan uint8 {
	out := make(chan uint8)

	go func() {
		defer close(out)
		for _, num := range numbers {
			out <- num
		}
	}()

	return out
}

// CollectResults собирает результаты из выходного канала в слайс
func CollectResults(output <-chan float64) []float64 {
	var results []float64
	for result := range output {
		results = append(results, result)
	}
	return results
}

// RunPipeline запускает полный конвейер обработки
func RunPipeline(inputNumbers []uint8) []float64 {
	// Создаем каналы
	inputChan := GenerateNumbers(inputNumbers)
	outputChan := make(chan float64)

	// Запускаем обработку
	go ProcessPipeline(inputChan, outputChan)

	// Собираем результаты
	return CollectResults(outputChan)
}
