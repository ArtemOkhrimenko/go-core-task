package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== Работа со слайсами в Go ===")

	// 1. Создаем оригинальный слайс
	originalSlice := createRandomSlice(10)
	fmt.Println("1. Оригинальный слайс (случайные числа):")
	printSlice(originalSlice, "originalSlice")

	// 2. Тестируем функцию sliceExample
	evenSlice := sliceExample(originalSlice)
	fmt.Println("\n2. Четные числа из оригинального слайса:")
	printSlice(evenSlice, "evenSlice")

	// 3. Тестируем функцию addElements
	numberToAdd := 99
	extendedSlice := addElements(originalSlice, numberToAdd)
	fmt.Printf("\n3. Добавление числа %d в конец слайса:\n", numberToAdd)
	printSlice(originalSlice, "originalSlice")
	printSlice(extendedSlice, "extendedSlice")

	// 4. Тестируем функцию copySlice
	copiedSlice := copySlice(originalSlice)
	fmt.Println("\n4. Копирование слайса:")
	printSlice(originalSlice, "originalSlice")
	printSlice(copiedSlice, "copiedSlice")

	// Демонстрируем, что копия независима
	fmt.Println("\n   Изменяем оригинальный слайс...")
	originalSlice[0] = 999
	fmt.Println("   После изменения originalSlice[0] = 999:")
	printSlice(originalSlice, "originalSlice")
	printSlice(copiedSlice, "copiedSlice")

	// 5. Тестируем функцию removeElement
	indexToRemove := 3
	sliceWithoutElement, err := removeElement(originalSlice, indexToRemove)
	if err != nil {
		fmt.Printf("\n5. Ошибка при удалении элемента: %v\n", err)
	} else {
		fmt.Printf("\n5. Удаление элемента по индексу %d:\n", indexToRemove)
		printSlice(originalSlice, "originalSlice")
		printSlice(sliceWithoutElement, "sliceWithoutElement")
		fmt.Printf("   Удаленный элемент: %d\n", originalSlice[indexToRemove])
	}

	// Тестируем обработку ошибок
	fmt.Println("\n6. Тестирование обработки ошибок:")

	// Некорректный индекс (отрицательный)
	_, err1 := removeElement(originalSlice, -1)
	if err1 != nil {
		fmt.Printf("   Ошибка при index=-1: %v\n", err1)
	}

	// Некорректный индекс (больше длины)
	_, err2 := removeElement(originalSlice, len(originalSlice))
	if err2 != nil {
		fmt.Printf("   Ошибка при index=%d: %v\n", len(originalSlice), err2)
	}

	// Демонстрация разных случайных слайсов
	fmt.Println("\n=== Демонстрация разных случайных слайсов ===")
	fmt.Println("Первый вызов:")
	slice1 := createRandomSlice(5)
	printSlice(slice1, "slice1")

	fmt.Println("Второй вызов:")
	slice2 := createRandomSlice(5)
	printSlice(slice2, "slice2")

	// Дополнительные демонстрации
	fmt.Println("\n=== Дополнительные примеры ===")

	// Цепочка операций
	fmt.Println("\nЦепочка операций:")
	chainSlice := []int{10, 20, 30, 40, 50}
	fmt.Println("Исходный слайс:", chainSlice)

	// Добавляем -> удаляем -> фильтруем четные
	chainSlice = addElements(chainSlice, 60)
	chainSlice, _ = removeElement(chainSlice, 1) // удаляем элемент с индексом 1 (20)
	chainSlice = sliceExample(chainSlice)
	fmt.Println("После добавления 60, удаления элемента [1], фильтрации четных:", chainSlice)

	fmt.Println("\n=== Программа завершена ===")
}

// 1. Создание слайса с случайными числами
func createRandomSlice(length int) []int {
	// Новый способ инициализации генератора случайных чисел
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = r.Intn(100) // случайные числа от 0 до 99
	}
	return slice
}

// 2. Функция возвращает новый слайс с четными числами
func sliceExample(slice []int) []int {
	result := make([]int, 0)
	for _, value := range slice {
		if value%2 == 0 {
			result = append(result, value)
		}
	}
	return result
}

// 3. Функция добавляет число в конец слайса
func addElements(slice []int, num int) []int {
	// Создаем новый слайс с увеличенной длиной
	newSlice := make([]int, len(slice)+1)
	copy(newSlice, slice)
	newSlice[len(slice)] = num
	return newSlice
}

// 4. Функция создает глубокую копию слайса
func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// 5. Функция удаляет элемент по индексу
func removeElement(slice []int, index int) ([]int, error) {
	if index < 0 || index >= len(slice) {
		return nil, fmt.Errorf("индекс %d вне диапазона [0, %d]", index, len(slice)-1)
	}

	// Создаем новый слайс без элемента по указанному индексу
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice, nil
}

// Вспомогательная функция для красивого вывода слайса
func printSlice(slice []int, name string) {
	fmt.Printf("%s: %v (длина: %d, емкость: %d)\n", name, slice, len(slice), cap(slice))
}
