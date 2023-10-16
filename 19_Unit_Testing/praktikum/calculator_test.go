package calculator

import (
    "testing"
)

func TestAddition(t *testing.T) {
    result := Addition(3, 2)
    if result != 5 {
        t.Errorf("Expected 5, but got %v", result)
    }
}

func TestSubtraction(t *testing.T) {
    result := Subtraction(5, 2)
    if result != 3 {
        t.Errorf("Expected 3, but got %v", result)
    }
}

func TestDivision(t *testing.T) {
    result, err := Division(10, 2)
    if err != nil {
        t.Errorf("Expected no error, but got an error: %v", err)
    }
    if result != 5 {
        t.Errorf("Expected 5, but got %v", result)
    }

    _, err = Division(5, 0)
    if err == nil {
        t.Errorf("Expected an error, but got none.")
    }
}

func TestMultiplication(t *testing.T) {
    result := Multiplication(4, 3)
    if result != 12 {
        t.Errorf("Expected 12, but got %v", result)
    }
}