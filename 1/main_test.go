package main

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"testing"
)

func TestCreateVariables(t *testing.T) {
	numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum := createVariables()

	// Проверяем значения
	if numDecimal != 42 {
		t.Errorf("Expected numDecimal to be 42, got %d", numDecimal)
	}
	if numOctal != 42 {
		t.Errorf("Expected numOctal to be 42, got %d", numOctal)
	}
	if numHexadecimal != 42 {
		t.Errorf("Expected numHexadecimal to be 42, got %d", numHexadecimal)
	}
	if pi != 3.14 {
		t.Errorf("Expected pi to be 3.14, got %f", pi)
	}
	if name != "Golang" {
		t.Errorf("Expected name to be 'Golang', got '%s'", name)
	}
	if isActive != true {
		t.Errorf("Expected isActive to be true, got %t", isActive)
	}
	if complexNum != 1+2i {
		t.Errorf("Expected complexNum to be 1+2i, got %v", complexNum)
	}
}

func TestGetVariableTypes(t *testing.T) {
	numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum := createVariables()
	types := getVariableTypes(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	expectedTypes := []string{"int", "int", "int", "float64", "string", "bool", "complex64"}

	if !reflect.DeepEqual(types, expectedTypes) {
		t.Errorf("Expected types %v, got %v", expectedTypes, types)
	}
}

func TestVariablesToStrings(t *testing.T) {
	numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum := createVariables()
	stringVars := variablesToStrings(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	expected := []string{
		"42", "42", "42", "3.14", "Golang", "true", "(1+2i)",
	}

	if !reflect.DeepEqual(stringVars, expected) {
		t.Errorf("Expected string vars %v, got %v", expected, stringVars)
	}
}

func TestConcatenateWithSalt(t *testing.T) {
	testStrings := []string{"a", "b", "c", "d"}
	salt := "salt"
	result := concatenateWithSalt(testStrings, salt)

	// Соль должна быть добавлена после второго элемента (индекс 1)
	expected := "absaltcd"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestStringToRunes(t *testing.T) {
	testString := "Hello, 世界"
	runes := stringToRunes(testString)

	expected := []rune{'H', 'e', 'l', 'l', 'o', ',', ' ', '世', '界'}

	if !reflect.DeepEqual(runes, expected) {
		t.Errorf("Expected runes %v, got %v", expected, runes)
	}

	// Проверяем длину
	if len(runes) != 9 {
		t.Errorf("Expected 9 runes, got %d", len(runes))
	}
}

func TestHashRunes(t *testing.T) {
	testRunes := []rune("test string")
	hash := hashRunes(testRunes)

	// Вычисляем ожидаемый хеш для проверки
	expectedHash := sha256.Sum256([]byte("test string"))
	expected := hex.EncodeToString(expectedHash[:])

	if hash != expected {
		t.Errorf("Expected hash '%s', got '%s'", expected, hash)
	}
}

func TestProcessVariables(t *testing.T) {
	types, combinedString, runes, hash := processVariables()

	// Проверяем, что все возвращаемые значения не пустые
	if len(types) == 0 {
		t.Error("Types slice should not be empty")
	}
	if combinedString == "" {
		t.Error("Combined string should not be empty")
	}
	if len(runes) == 0 {
		t.Error("Runes slice should not be empty")
	}
	if hash == "" {
		t.Error("Hash should not be empty")
	}

	// Проверяем, что хеш имеет правильную длину (64 символа для SHA256 в hex)
	if len(hash) != 64 {
		t.Errorf("Hash should be 64 characters long, got %d", len(hash))
	}

	// Проверяем, что соль присутствует в объединенной строке
	if !contains(combinedString, "go-2024") {
		t.Error("Combined string should contain salt 'go-2024'")
	}
}

// Вспомогательная функция для проверки наличия подстроки
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Benchmark тесты
func BenchmarkProcessVariables(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processVariables()
	}
}

func BenchmarkHashRunes(b *testing.B) {
	testRunes := []rune("benchmark test string")
	for i := 0; i < b.N; i++ {
		hashRunes(testRunes)
	}
}
