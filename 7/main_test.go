package main

import (
	"sync"
	"testing"
)

func TestMergeChannels_Basic(t *testing.T) {
	// Создаем три тестовых канала
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Сливаем каналы
	merged := MergeChannels(ch1, ch2, ch3)

	// Запускаем отправку данных в каналы
	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()

	go func() {
		defer close(ch3)
		ch3 <- 5
		ch3 <- 6
	}()

	// Собираем все значения из объединенного канала
	results := make([]int, 0)
	for value := range merged {
		results = append(results, value)
	}

	// Проверяем количество полученных значений
	if len(results) != 6 {
		t.Errorf("Получено %d значений, ожидалось 6", len(results))
	}
}

func TestMergeChannels_SingleChannel(t *testing.T) {
	ch := make(chan int)
	merged := MergeChannels(ch)

	go func() {
		defer close(ch)
		ch <- 42
		ch <- 24
	}()

	results := make([]int, 0)
	for value := range merged {
		results = append(results, value)
	}

	if len(results) != 2 {
		t.Errorf("Получено %d значений, ожидалось 2", len(results))
	}
}

func TestMergeChannels_Empty(t *testing.T) {
	// Тестируем слияние пустого списка каналов
	merged := MergeChannels()

	// Канал должен быть сразу закрыт
	_, ok := <-merged
	if ok {
		t.Error("Объединенный канал должен быть закрыт для пустого списка входных каналов")
	}
}

func TestMergeChannels_Concurrent(t *testing.T) {
	const numChannels = 5
	const valuesPerChannel = 10

	channels := make([]<-chan int, numChannels)
	var wg sync.WaitGroup

	// Создаем каналы и запускаем горутины для отправки данных
	for i := 0; i < numChannels; i++ {
		ch := make(chan int)
		channels[i] = ch

		wg.Add(1)
		go func(ch chan int, start int) {
			defer wg.Done()
			defer close(ch)
			for j := 0; j < valuesPerChannel; j++ {
				ch <- start + j
			}
		}(ch, i*100)
	}

	// Сливаем все каналы
	merged := MergeChannels(channels...)

	// Читаем все значения
	totalValues := 0
	for range merged {
		totalValues++
	}

	expectedTotal := numChannels * valuesPerChannel
	if totalValues != expectedTotal {
		t.Errorf("Получено %d значений, ожидалось %d", totalValues, expectedTotal)
	}

	wg.Wait()
}

func TestMergeChannels_OrderIndependence(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	merged := MergeChannels(ch1, ch2)

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

	// Просто проверяем, что получаем все значения
	count := 0
	for range merged {
		count++
	}

	if count != 5 {
		t.Errorf("Получено %d значений, ожидалось 5", count)
	}
}

// Benchmark тест
func BenchmarkMergeChannels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch1 := make(chan int)
		ch2 := make(chan int)
		ch3 := make(chan int)

		merged := MergeChannels(ch1, ch2, ch3)

		go func() {
			defer close(ch1)
			for j := 0; j < 100; j++ {
				ch1 <- j
			}
		}()

		go func() {
			defer close(ch2)
			for j := 0; j < 100; j++ {
				ch2 <- j + 100
			}
		}()

		go func() {
			defer close(ch3)
			for j := 0; j < 100; j++ {
				ch3 <- j + 200
			}
		}()

		// Читаем все значения
		for range merged {
		}
	}
}
