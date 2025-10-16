package main

import "fmt"

func main() {
	// Пример из задания
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := Difference(slice1, slice2)
	fmt.Println("Задание 4: Разность слайсов строк")
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
	fmt.Printf("Результат (slice1 - slice2): %v\n", result)

	// Дополнительные примеры
	fmt.Println("\nДополнительные примеры:")

	// Пример 1: Пустой второй слайс
	emptySlice := []string{}
	result1 := Difference(slice1, emptySlice)
	fmt.Printf("Разность с пустым слайсом: %v\n", result1)

	// Пример 2: Пустой первый слайс
	result2 := Difference(emptySlice, slice2)
	fmt.Printf("Разность пустого слайса: %v\n", result2)

	// Пример 3: Одинаковые слайсы
	result3 := Difference(slice1, slice1)
	fmt.Printf("Разность одинаковых слайсов: %v\n", result3)

	// Пример 4: Нет общих элементов
	slice3 := []string{"x", "y", "z"}
	result4 := Difference(slice1, slice3)
	fmt.Printf("Разность без общих элементов: %v\n", result4)
}

// Difference возвращает элементы, которые есть в slice1, но отсутствуют в slice2
func Difference(slice1, slice2 []string) []string {
	result := make([]string, 0)

	// Если первый слайс пустой, сразу возвращаем пустой результат
	if len(slice1) == 0 {
		return result
	}

	// Создаем множество из второго слайса для быстрого поиска
	set := make(map[string]bool)
	for _, item := range slice2 {
		set[item] = true
	}

	// Собираем элементы из первого слайса, которых нет во множестве
	for _, item := range slice1 {
		if !set[item] {
			result = append(result, item)
		}
	}

	return result
}
