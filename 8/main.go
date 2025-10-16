package main

import (
	"fmt"
	"time"
)

func main() {
	// Демонстрация базовой кастомной WaitGroup
	demonstrateCustomWaitGroup()

	fmt.Println("\n=== Тестирование edge cases ===")

	// Тест 1: Wait без Add
	fmt.Println("Тест 1: Wait без Add")
	wg1 := NewCustomWaitGroup()
	wg1.Wait()
	fmt.Println("✓ Успешно завершено")

	// Тест 2: Несколько Add/Done
	fmt.Println("\nТест 2: Несколько Add/Done")
	wg2 := NewCustomWaitGroup()
	wg2.Add(3)
	wg2.Done()
	wg2.Done()
	wg2.Done()
	wg2.Wait()
	fmt.Println("✓ Успешно завершено")

	// Тест 3: Конкурентный доступ
	fmt.Println("\nТест 3: Конкурентный доступ")
	wg3 := NewCustomWaitGroup()
	const numGoroutines = 5

	for i := 0; i < numGoroutines; i++ {
		wg3.Add(1)
		go func(id int) {
			defer wg3.Done()
			// Быстрая операция без задержки
		}(i)
	}

	wg3.Wait()
	fmt.Println("✓ Все конкурентные операции завершены")

	fmt.Println("=== Все тесты пройдены ===")
}

// CustomWaitGroup - кастомная реализация WaitGroup на основе семафора
type CustomWaitGroup struct {
	count     int
	semaphore chan struct{}
}

// NewCustomWaitGroup создает новую кастомную WaitGroup
func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		count:     0,
		semaphore: make(chan struct{}, 1), // буферизированный канал как мьютекс
	}
}

// Add добавляет delta к счетчику
func (wg *CustomWaitGroup) Add(delta int) {
	// Блокируем семафор для атомарного доступа
	wg.semaphore <- struct{}{}
	defer func() {
		<-wg.semaphore // освобождаем семафор
	}()

	wg.count += delta

	if wg.count < 0 {
		panic("negative WaitGroup counter")
	}
}

// Done уменьшает счетчик на 1
func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

// Wait блокируется until счетчик станет 0
func (wg *CustomWaitGroup) Wait() {
	for {
		// Блокируем для чтения счетчика
		wg.semaphore <- struct{}{}

		if wg.count == 0 {
			<-wg.semaphore
			return
		}

		// Освобождаем и ждем немного перед следующей проверкой
		<-wg.semaphore
		time.Sleep(1 * time.Millisecond) // уменьшаем задержку
	}
}

// Демонстрация работы кастомной WaitGroup
func demonstrateCustomWaitGroup() {
	fmt.Println("=== Демонстрация кастомной WaitGroup ===")

	wg := NewCustomWaitGroup()

	fmt.Println("Запускаем 3 горутины...")
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Горутина %d начала работу\n", id)
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			fmt.Printf("Горутина %d завершила работу", id)
		}(i)
	}

	fmt.Println("Ожидание завершения всех горутин...")
	wg.Wait()
	fmt.Println("Все горутины завершены!")
}
