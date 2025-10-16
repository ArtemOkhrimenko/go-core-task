package main

import (
	"testing"
	"time"
)

func TestCustomWaitGroup_Basic(t *testing.T) {
	wg := NewCustomWaitGroup()

	wg.Add(2)

	// Запускаем горутины которые сразу завершаются
	go func() {
		defer wg.Done()
		// Немедленное завершение
	}()

	go func() {
		defer wg.Done()
		// Немедленное завершение
	}()

	// Wait должен завершиться быстро
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		// Успех - Wait завершился
	case <-time.After(100 * time.Millisecond):
		t.Error("Wait заблокировался навсегда")
	}
}

func TestCustomWaitGroup_ZeroWait(t *testing.T) {
	wg := NewCustomWaitGroup()

	// Wait должен завершиться сразу для нулевого счетчика
	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)

	if elapsed > 10*time.Millisecond {
		t.Error("Wait должен завершаться мгновенно для нулевого счетчика")
	}
}

func TestCustomWaitGroup_SingleGoroutine(t *testing.T) {
	wg := NewCustomWaitGroup()

	wg.Add(1)

	go func() {
		defer wg.Done()
		// Быстрая операция
	}()

	// Wait не должно блокироваться
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		// Успех
	case <-time.After(100 * time.Millisecond):
		t.Error("Wait заблокировался для одной горутины")
	}
}

func TestCustomWaitGroup_MultipleOperations(t *testing.T) {
	wg := NewCustomWaitGroup()

	// Тестируем несколько операций Add/Done
	wg.Add(3)
	wg.Done()
	wg.Add(-1) // эквивалентно Done()
	wg.Done()

	// Счетчик должен быть 0
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		// Успех
	case <-time.After(100 * time.Millisecond):
		t.Error("Wait заблокировался после нескольких операций")
	}
}

func TestCustomWaitGroup_Concurrent(t *testing.T) {
	wg := NewCustomWaitGroup()
	const numGoroutines = 10

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			// Быстрая операция без задержки
		}(i)
	}

	// Wait не должно блокироваться
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		// Успех
	case <-time.After(100 * time.Millisecond):
		t.Error("Wait заблокировался при конкурентном доступе")
	}
}

func TestCustomWaitGroup_NegativeCounter(t *testing.T) {
	wg := NewCustomWaitGroup()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Ожидалась паника при отрицательном счетчике")
		}
	}()

	wg.Add(-1) // Должно вызвать панику
}

func TestCustomWaitGroup_WithDelayedGoroutines(t *testing.T) {
	wg := NewCustomWaitGroup()

	wg.Add(2)

	// Горутины с небольшой задержкой
	go func() {
		time.Sleep(5 * time.Millisecond)
		defer wg.Done()
	}()

	go func() {
		time.Sleep(10 * time.Millisecond)
		defer wg.Done()
	}()

	// Wait должен дождаться завершения
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		// Успех
	case <-time.After(50 * time.Millisecond):
		t.Error("Wait не дождался завершения горутин с задержкой")
	}
}

func TestCustomWaitGroup_SequentialUsage(t *testing.T) {
	wg := NewCustomWaitGroup()

	// Последовательное использование
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()
	wg.Wait()

	// Повторное использование
	wg.Add(2)
	go func() {
		defer wg.Done()
	}()
	go func() {
		defer wg.Done()
	}()
	wg.Wait()

}
