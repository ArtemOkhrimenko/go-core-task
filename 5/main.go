package main

import "fmt"

func main() {
	fmt.Println("=== Задание 5: Пересечение слайсов int ===")

	// Пример из задания
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasIntersection, common := Intersection(a, b)
	fmt.Println("Пример из задания:")
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("Есть пересечение: %t\n", hasIntersection)
	fmt.Printf("Пересекающиеся значения: %v\n", common)

	// Версия без дубликатов
	_, commonUnique := IntersectionUnique(a, b)
	fmt.Printf("Без дубликатов: %v\n", commonUnique)

	// Дополнительные примеры
	fmt.Println("\nДополнительные примеры:")

	// Пример 1: Нет пересечений
	c := []int{1, 2, 3}
	d := []int{4, 5, 6}
	hasIntersection1, common1 := Intersection(c, d)
	fmt.Printf("\nПример 1 - Нет пересечений:\n")
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("Есть пересечение: %t\n", hasIntersection1)
	fmt.Printf("Пересекающиеся значения: %v\n", common1)

	// Пример 2: Полное пересечение
	e := []int{10, 20, 30}
	f := []int{10, 20, 30}
	hasIntersection2, common2 := Intersection(e, f)
	fmt.Printf("\nПример 2 - Полное пересечение:\n")
	fmt.Printf("e: %v\n", e)
	fmt.Printf("f: %v\n", f)
	fmt.Printf("Есть пересечение: %t\n", hasIntersection2)
	fmt.Printf("Пересекающиеся значения: %v\n", common2)

	// Пример 3: Частичное пересечение
	g := []int{1, 2, 3, 4, 5}
	h := []int{3, 4, 5, 6, 7}
	hasIntersection3, common3 := Intersection(g, h)
	fmt.Printf("\nПример 3 - Частичное пересечение:\n")
	fmt.Printf("g: %v\n", g)
	fmt.Printf("h: %v\n", h)
	fmt.Printf("Есть пересечение: %t\n", hasIntersection3)
	fmt.Printf("Пересекающиеся значения: %v\n", common3)

	// Пример 4: Дубликаты в исходных слайсах
	i := []int{1, 2, 2, 3, 4}
	j := []int{2, 3, 3, 5}
	hasIntersection4, common4 := Intersection(i, j)
	fmt.Printf("\nПример 4 - С дубликатами:\n")
	fmt.Printf("i: %v\n", i)
	fmt.Printf("j: %v\n", j)
	fmt.Printf("Есть пересечение: %t\n", hasIntersection4)
	fmt.Printf("Пересекающиеся значения: %v\n", common4)

	// Версия без дубликатов
	_, common4Unique := IntersectionUnique(i, j)
	fmt.Printf("Без дубликатов: %v\n", common4Unique)

	// Пример 5: Один пустой слайс
	k := []int{1, 2, 3}
	l := []int{}
	hasIntersection5, common5 := Intersection(k, l)
	fmt.Printf("\nПример 5 - Один пустой слайс:\n")
	fmt.Printf("k: %v\n", k)
	fmt.Printf("l: %v\n", l)
	fmt.Printf("Есть пересечение: %t\n", hasIntersection5)
	fmt.Printf("Пересекающиеся значения: %v\n", common5)

	// Пример 6: Оба пустых слайса
	m := []int{}
	n := []int{}
	hasIntersection6, common6 := Intersection(m, n)
	fmt.Printf("\nПример 6 - Оба пустых слайса:\n")
	fmt.Printf("m: %v\n", m)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("Есть пересечение: %t\n", hasIntersection6)
	fmt.Printf("Пересекающиеся значения: %v\n", common6)

	// Пример 7: Большие числа и отрицательные
	o := []int{-5, -3, 0, 1000, 9999}
	p := []int{-3, 0, 9999, 10000}
	hasIntersection7, common7 := Intersection(o, p)
	fmt.Printf("\nПример 7 - С отрицательными и большими числами:\n")
	fmt.Printf("o: %v\n", o)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("Есть пересечение: %t\n", hasIntersection7)
	fmt.Printf("Пересекающиеся значения: %v\n", common7)

	fmt.Println("\n=== Демонстрация завершена ===")
}

// Intersection проверяет пересечение двух слайсов int и возвращает:
// - есть ли хотя бы одно пересечение
// - слайс с общими элементами
func Intersection(a, b []int) (bool, []int) {
	// Если один из слайсов пустой, сразу возвращаем false и пустой слайс
	if len(a) == 0 || len(b) == 0 {
		return false, []int{}
	}

	// Создаем множество из первого слайса для быстрого поиска
	set := make(map[int]bool)
	for _, num := range a {
		set[num] = true
	}

	// Находим общие элементы
	var common []int
	for _, num := range b {
		if set[num] {
			common = append(common, num)
		}
	}

	hasIntersection := len(common) > 0
	return hasIntersection, common
}

// IntersectionUnique возвращает пересечение без дубликатов
func IntersectionUnique(a, b []int) (bool, []int) {
	hasIntersection, common := Intersection(a, b)
	if !hasIntersection {
		return false, common
	}

	// Удаляем дубликаты из результата
	uniqueSet := make(map[int]bool)
	var uniqueCommon []int
	for _, num := range common {
		if !uniqueSet[num] {
			uniqueSet[num] = true
			uniqueCommon = append(uniqueCommon, num)
		}
	}

	return true, uniqueCommon
}
