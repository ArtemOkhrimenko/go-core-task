package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewStringIntMap(t *testing.T) {
	m := NewStringIntMap()

	if m == nil {
		t.Error("NewStringIntMap() вернул nil")
	}

	if m.data == nil {
		t.Error("Внутренняя карта не инициализирована")
	}

	if m.Size() != 0 {
		t.Errorf("Новая карта должна быть пустой, но размер: %d", m.Size())
	}
}

func TestStringIntMap_Add(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    int
		expected int
	}{
		{"добавление нового ключа", "test", 42, 42},
		{"добавление с отрицательным значением", "negative", -10, -10},
		{"добавление с нулевым значением", "zero", 0, 0},
		{"перезапись существующего ключа", "test", 100, 100},
	}

	m := NewStringIntMap()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.Add(tt.key, tt.value)

			// Проверяем, что значение правильно установлено
			if actualValue, exists := m.Get(tt.key); !exists || actualValue != tt.expected {
				t.Errorf("Add(%s, %d): ожидалось значение %d, получено %d, exists=%t",
					tt.key, tt.value, tt.expected, actualValue, exists)
			}
		})
	}
}

func TestStringIntMap_Remove(t *testing.T) {
	m := NewStringIntMap()

	// Добавляем тестовые данные
	m.Add("key1", 1)
	m.Add("key2", 2)
	m.Add("key3", 3)

	initialSize := m.Size()

	// Удаляем существующий ключ
	t.Run("удаление существующего ключа", func(t *testing.T) {
		m.Remove("key2")

		if m.Exists("key2") {
			t.Error("Ключ 'key2' должен был быть удален")
		}

		if m.Size() != initialSize-1 {
			t.Errorf("Размер карты должен быть %d, но получен %d", initialSize-1, m.Size())
		}
	})

	// Удаляем несуществующий ключ
	t.Run("удаление несуществующего ключа", func(t *testing.T) {
		initialSize := m.Size()
		m.Remove("nonexistent")

		// Размер не должен измениться
		if m.Size() != initialSize {
			t.Errorf("Размер карты не должен измениться при удалении несуществующего ключа")
		}
	})

	// Проверяем, что другие ключи остались нетронутыми
	if !m.Exists("key1") || !m.Exists("key3") {
		t.Error("Другие ключи не должны быть затронуты при удалении")
	}
}

func TestStringIntMap_Copy(t *testing.T) {
	m := NewStringIntMap()

	// Добавляем тестовые данные
	testData := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k, v := range testData {
		m.Add(k, v)
	}

	// Копируем карту
	copied := m.Copy()

	// Проверяем, что содержимое одинаковое
	if !reflect.DeepEqual(copied, testData) {
		t.Errorf("Скопированная карта %v не совпадает с ожидаемой %v", copied, testData)
	}

	// Проверяем независимость копии
	m.Add("d", 4)
	if _, exists := copied["d"]; exists {
		t.Error("Изменения в оригинальной карте не должны влиять на копию")
	}

	// Изменяем копию и проверяем, что оригинал не изменился
	copied["e"] = 5
	if m.Exists("e") {
		t.Error("Изменения в копии не должны влиять на оригинал")
	}
}

func TestStringIntMap_Exists(t *testing.T) {
	m := NewStringIntMap()

	m.Add("exists", 42)
	m.Add("zero", 0)

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{"существующий ключ", "exists", true},
		{"ключ с нулевым значением", "zero", true},
		{"несуществующий ключ", "nonexistent", false},
		{"пустая строка", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := m.Exists(tt.key)
			if result != tt.expected {
				t.Errorf("Exists(%s) = %t, ожидалось %t", tt.key, result, tt.expected)
			}
		})
	}
}

func TestStringIntMap_Get(t *testing.T) {
	m := NewStringIntMap()

	m.Add("key1", 100)
	m.Add("key2", 0)
	m.Add("key3", -50)

	tests := []struct {
		name          string
		key           string
		expectedValue int
		expectedFound bool
	}{
		{"существующий ключ с положительным значением", "key1", 100, true},
		{"существующий ключ с нулевым значением", "key2", 0, true},
		{"существующий ключ с отрицательным значением", "key3", -50, true},
		{"несуществующий ключ", "key4", 0, false},
		{"пустая строка", "", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := m.Get(tt.key)

			if value != tt.expectedValue {
				t.Errorf("Get(%s) вернул значение %d, ожидалось %d", tt.key, value, tt.expectedValue)
			}

			if found != tt.expectedFound {
				t.Errorf("Get(%s) вернул found=%t, ожидалось %t", tt.key, found, tt.expectedFound)
			}
		})
	}
}

func TestStringIntMap_Size(t *testing.T) {
	m := NewStringIntMap()

	// Проверяем начальный размер
	if m.Size() != 0 {
		t.Errorf("Начальный размер должен быть 0, но получен %d", m.Size())
	}

	// Добавляем элементы и проверяем размер
	expectedSize := 0
	for i := 1; i <= 5; i++ {
		key := fmt.Sprintf("key%d", i)
		m.Add(key, i)
		expectedSize++

		if m.Size() != expectedSize {
			t.Errorf("После добавления %d элементов размер должен быть %d, но получен %d",
				i, expectedSize, m.Size())
		}
	}

	// Удаляем элементы и проверяем размер
	m.Remove("key2")
	expectedSize--
	if m.Size() != expectedSize {
		t.Errorf("После удаления элемента размер должен быть %d, но получен %d",
			expectedSize, m.Size())
	}

	// Удаляем несуществующий элемент
	m.Remove("nonexistent")
	if m.Size() != expectedSize {
		t.Errorf("Размер не должен измениться при удалении несуществующего элемента")
	}
}

func TestStringIntMap_Clear(t *testing.T) {
	m := NewStringIntMap()

	// Добавляем данные
	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)

	if m.Size() == 0 {
		t.Error("Карта должна содержать элементы перед очисткой")
	}

	// Очищаем карту
	m.Clear()

	if m.Size() != 0 {
		t.Errorf("После очистки размер должен быть 0, но получен %d", m.Size())
	}

	// Проверяем, что все ключи удалены
	keys := []string{"a", "b", "c"}
	for _, key := range keys {
		if m.Exists(key) {
			t.Errorf("Ключ '%s' должен быть удален после очистки", key)
		}
	}
}

func TestStringIntMap_Keys(t *testing.T) {
	m := NewStringIntMap()

	// Тест с пустой картой
	emptyKeys := m.Keys()
	if len(emptyKeys) != 0 {
		t.Errorf("Ключи пустой карты должны быть пустым срезом, но получено: %v", emptyKeys)
	}

	// Тест с данными
	testData := map[string]int{
		"x": 1,
		"y": 2,
		"z": 3,
	}

	for k, v := range testData {
		m.Add(k, v)
	}

	keys := m.Keys()

	// Проверяем количество ключей
	if len(keys) != len(testData) {
		t.Errorf("Ожидалось %d ключей, но получено %d", len(testData), len(keys))
	}

	// Проверяем, что все ожидаемые ключи присутствуют
	keySet := make(map[string]bool)
	for _, key := range keys {
		keySet[key] = true
	}

	for expectedKey := range testData {
		if !keySet[expectedKey] {
			t.Errorf("Ожидаемый ключ '%s' отсутствует в результате", expectedKey)
		}
	}
}

func TestStringIntMap_Values(t *testing.T) {
	m := NewStringIntMap()

	// Тест с пустой картой
	emptyValues := m.Values()
	if len(emptyValues) != 0 {
		t.Errorf("Значения пустой карты должны быть пустым срезом, но получено: %v", emptyValues)
	}

	// Тест с данными
	testData := map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
	}

	for k, v := range testData {
		m.Add(k, v)
	}

	values := m.Values()

	// Проверяем количество значений
	if len(values) != len(testData) {
		t.Errorf("Ожидалось %d значений, но получено %d", len(testData), len(values))
	}

	// Проверяем, что все ожидаемые значения присутствуют
	valueSet := make(map[int]bool)
	for _, value := range values {
		valueSet[value] = true
	}

	for _, expectedValue := range testData {
		if !valueSet[expectedValue] {
			t.Errorf("Ожидаемое значение %d отсутствует в результате", expectedValue)
		}
	}
}

func TestStringIntMap_String(t *testing.T) {
	m := NewStringIntMap()

	// Тест с пустой картой
	emptyString := m.String()
	expectedEmpty := "StringIntMap{}"
	if emptyString != expectedEmpty {
		t.Errorf("String() для пустой карты: получено '%s', ожидалось '%s'", emptyString, expectedEmpty)
	}

	// Тест с одним элементом
	m.Add("single", 42)
	singleString := m.String()
	// Проверяем, что строка содержит ожидаемые данные
	if !contains(singleString, "single: 42") {
		t.Errorf("String() должен содержать 'single: 42', но получено: %s", singleString)
	}

	// Тест с несколькими элементами
	m.Add("another", 100)
	multiString := m.String()

	// Проверяем, что оба элемента присутствуют
	if !contains(multiString, "single: 42") || !contains(multiString, "another: 100") {
		t.Errorf("String() должен содержать оба элемента, но получено: %s", multiString)
	}
}

// Вспомогательная функция для проверки наличия подстроки
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		(len(s) > len(substr) && (s[:len(substr)] == substr ||
			contains(s[1:], substr))))
}

// Benchmark тесты
func BenchmarkStringIntMap_Add(b *testing.B) {
	m := NewStringIntMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i)
		m.Add(key, i)
	}
}

func BenchmarkStringIntMap_Get(b *testing.B) {
	m := NewStringIntMap()

	// Подготавливаем данные
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key%d", i)
		m.Add(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i%1000)
		m.Get(key)
	}
}

func BenchmarkStringIntMap_Exists(b *testing.B) {
	m := NewStringIntMap()

	// Подготавливаем данные
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key%d", i)
		m.Add(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i%1000)
		m.Exists(key)
	}
}

func BenchmarkStringIntMap_Copy(b *testing.B) {
	m := NewStringIntMap()

	// Подготавливаем данные
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key%d", i)
		m.Add(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Copy()
	}
}
