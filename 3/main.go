package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Демонстрация StringIntMap ===")

	// Создаем новую карту
	myMap := NewStringIntMap()

	// 1. Добавление элементов
	fmt.Println("1. Добавление элементов:")
	myMap.Add("apple", 10)
	myMap.Add("banana", 20)
	myMap.Add("orange", 30)
	myMap.Add("grape", 40)

	// 2. Проверка наличия ключей
	fmt.Println("\n2. Проверка наличия ключей:")
	keysToCheck := []string{"apple", "banana", "cherry"}
	for _, key := range keysToCheck {
		exists := myMap.Exists(key)
		fmt.Printf("   Ключ '%s' существует: %t\n", key, exists)
	}

	// 3. Получение значений
	fmt.Println("\n3. Получение значений:")
	keysToGet := []string{"orange", "kiwi"}
	for _, key := range keysToGet {
		value, exists := myMap.Get(key)
		if exists {
			fmt.Printf("   Ключ '%s' имеет значение: %d\n", key, value)
		} else {
			fmt.Printf("   Ключ '%s' не найден\n", key)
		}
	}

	// 4. Копирование карты
	fmt.Println("\n4. Копирование карты:")
	copiedMap := myMap.Copy()
	fmt.Printf("   Оригинальная карта: %s\n", myMap)
	fmt.Printf("   Скопированная карта: %v\n", copiedMap)

	// Демонстрируем независимость копии
	myMap.Add("pineapple", 50)
	fmt.Printf("   После добавления 'pineapple' в оригинал:\n")
	fmt.Printf("   Оригинальная карта: %s\n", myMap)
	fmt.Printf("   Скопированная карта: %v\n", copiedMap)

	// 5. Удаление элементов
	fmt.Println("\n5. Удаление элементов:")
	fmt.Printf("   До удаления: %s\n", myMap)
	myMap.Remove("banana")
	myMap.Remove("cherry") // Попытка удалить несуществующий ключ
	fmt.Printf("   После удаления 'banana' и 'cherry': %s\n", myMap)

	fmt.Println("\n=== Демонстрация завершена ===")
}

// StringIntMap структура для хранения пар "строка - целое число"
type StringIntMap struct {
	data map[string]int
}

// NewStringIntMap создает новый экземпляр StringIntMap
func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

// Add добавляет новую пару "ключ-значение" в карту
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

// Remove удаляет элемент по ключу из карты
func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

// Copy возвращает новую карту, содержащую все элементы текущей карты
func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int, len(m.data))
	for key, value := range m.data {
		newMap[key] = value
	}
	return newMap
}

// Exists проверяет, существует ли ключ в карте
func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}

// Get возвращает значение по ключу и булевый флаг, указывающий на успешность операции
func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}
