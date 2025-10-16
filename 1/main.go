package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== Программа для работы с различными типами данных ===")

	// Выполняем весь процесс
	types, combinedString, runes, hash := processVariables()

	// Выводим результаты
	fmt.Println("1. Типы переменных:")
	for i, t := range types {
		fmt.Printf("   Переменная %d: %s\n", i+1, t)
	}

	fmt.Println("\n2. Объединенная строка с солью:")
	fmt.Printf("   %s\n", combinedString)

	fmt.Println("\n3. Срез рун (первые 20 элементов):")
	fmt.Printf("   %v\n", runes[:min(20, len(runes))])
	fmt.Printf("   Всего рун: %d\n", len(runes))

	fmt.Println("\n4. SHA256 хеш:")
	fmt.Printf("   %s\n", hash)

	fmt.Println("\n=== Тестирование функций ===")

	// Демонстрация отдельных функций
	numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum := createVariables()

	fmt.Printf("\nСозданные переменные:\n")
	fmt.Printf("Decimal: %d, Octal: %d, Hexadecimal: %d\n", numDecimal, numOctal, numHexadecimal)
	fmt.Printf("Float: %.2f, String: %s, Bool: %t, Complex: %v\n", pi, name, isActive, complexNum)

	// Тестирование преобразования в строки
	stringVars := variablesToStrings(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Printf("\nПеременные как строки: %v\n", stringVars)

	// Тестирование хеширования
	testString := "test string"
	testRunes := stringToRunes(testString)
	testHash := hashRunes(testRunes)
	fmt.Printf("\nТестовое хеширование: '%s' -> %s\n", testString, testHash)
}

// Создание переменных различных типов
func createVariables() (int, int, int, float64, string, bool, complex64) {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	return numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum
}

// Определение типов переменных
func getVariableTypes(vars ...interface{}) []string {
	types := make([]string, len(vars))
	for i, v := range vars {
		types[i] = fmt.Sprintf("%T", v)
	}
	return types
}

// Преобразование переменных в строки
func variablesToStrings(numDecimal, numOctal, numHexadecimal int, pi float64, name string, isActive bool, complexNum complex64) []string {
	return []string{
		strconv.Itoa(numDecimal),
		strconv.Itoa(numOctal),
		strconv.Itoa(numHexadecimal),
		strconv.FormatFloat(pi, 'f', -1, 64),
		name,
		strconv.FormatBool(isActive),
		fmt.Sprintf("%v", complexNum),
	}
}

// Объединение строк в одну с добавлением соли
func concatenateWithSalt(strings []string, salt string) string {
	result := ""
	for i, s := range strings {
		result += s
		// Добавляем соль после половины элементов (округление вниз)
		if i == (len(strings)-1)/2 {
			result += salt
		}
	}
	return result
}

// Преобразование строки в срез рун
func stringToRunes(s string) []rune {
	return []rune(s)
}

// Хеширование среза рун с SHA256
func hashRunes(runes []rune) string {
	// Преобразуем срез рун обратно в строку для хеширования
	str := string(runes)
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

// Основная функция для выполнения всего процесса
func processVariables() ([]string, string, []rune, string) {
	// 1. Создание переменных
	numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum := createVariables()

	// 2. Определение типов
	types := getVariableTypes(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	// 3. Преобразование в строки и объединение
	stringVars := variablesToStrings(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	combinedString := concatenateWithSalt(stringVars, "go-2024")

	// 4. Преобразование в срез рун
	runes := stringToRunes(combinedString)

	// 5. Хеширование
	hash := hashRunes(runes)

	return types, combinedString, runes, hash
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
