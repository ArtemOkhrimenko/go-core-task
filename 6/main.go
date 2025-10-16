package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	demonstrateGenerator()
}

// RandomGenerator представляет генератор случайных чисел
type RandomGenerator struct {
	min   int
	max   int
	count int
}

// NewRandomGenerator создает новый генератор случайных чисел
func NewRandomGenerator(min, max, count int) *RandomGenerator {
	return &RandomGenerator{
		min:   min,
		max:   max,
		count: count,
	}
}

// Generate генерирует случайные числа и отправляет их в небуферизированный канал
func (rg *RandomGenerator) Generate() <-chan int {
	// Создаем небуферизированный канал
	ch := make(chan int)

	go func() {
		defer close(ch)

		// Создаем локальный генератор случайных чисел
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		for i := 0; i < rg.count; i++ {
			// Генерируем случайное число в заданном диапазоне
			num := r.Intn(rg.max-rg.min+1) + rg.min
			// Отправляем в канал (блокируется, пока значение не будет прочитано)
			ch <- num
		}
	}()

	return ch
}

// Функция-помощник для демонстрации
func demonstrateGenerator() {
	fmt.Println("=== Демонстрация генератора случайных чисел ===")

	// Пример 1: Базовый генератор
	fmt.Println("1. Базовый генератор (10 чисел от 1 до 100):")
	gen1 := NewRandomGenerator(1, 100, 10)
	ch1 := gen1.Generate()

	count := 0
	for num := range ch1 {
		fmt.Printf("Получено: %d\n", num)
		count++
	}
	fmt.Printf("Всего получено чисел: %d\n\n", count)
}
