package main

import (
	"fmt"
	"sync"
)

func main() {
	// Пример 1: Слияние трех каналов
	fmt.Println("=== Пример 1: Слияние трех каналов ===")

	// Создаем три канала
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Запускаем горутины для отправки данных в каналы
	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()

	go func() {
		defer close(ch2)
		ch2 <- 10
		ch2 <- 20
	}()

	go func() {
		defer close(ch3)
		ch3 <- 100
		ch3 <- 200
		ch3 <- 300
	}()

	// Сливаем каналы в один
	merged := MergeChannels(ch1, ch2, ch3)

	// Читаем все значения из объединенного канала
	fmt.Println("Объединенные значения:")
	for value := range merged {
		fmt.Printf("Получено: %d\n", value)
	}

	fmt.Println("\n=== Программа завершена ===")
}

// MergeChannels сливает несколько каналов в один
func MergeChannels(channels ...<-chan int) <-chan int {
	// Создаем объединенный канал
	merged := make(chan int)

	// Используем WaitGroup для отслеживания завершения всех горутин
	var wg sync.WaitGroup

	// Функция для перенаправления данных из одного канала в объединенный
	forward := func(ch <-chan int) {
		// Уменьшаем счетчик WaitGroup при завершении горутины
		defer wg.Done()

		// Читаем все значения из канала и отправляем в объединенный
		for value := range ch {
			merged <- value
		}
	}

	// Устанавливаем счетчик WaitGroup равным количеству каналов
	wg.Add(len(channels))

	// Запускаем горутину для каждого входного канала
	for _, ch := range channels {
		go forward(ch)
	}

	// Запускаем горутину для закрытия объединенного канала
	// после завершения всех рабочих горутин
	go func() {
		// Ждем завершения всех горутин
		wg.Wait()
		// Закрываем объединенный канал
		close(merged)
	}()

	return merged
}
