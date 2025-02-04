package tests

import (
	"github.com/genov8/goprototype/v2/prototype"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	num := prototype.NewPrototype(10)
	newNum, err := num.Add(5)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if newNum.Value() != 15 {
		t.Errorf("Expected 15, got %v", newNum.Value())
	}
}

func TestMultiply(t *testing.T) {
	num := prototype.NewPrototype(10)
	newNum, err := num.Multiply(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if newNum.Value() != 30 {
		t.Errorf("Expected 30, got %v", newNum.Value())
	}
}

func TestNumber(t *testing.T) {
	numProto := prototype.NewPrototype(10)
	resultProto, err := numProto.Divide(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 5.0 {
		t.Errorf("Expected 5, got %v", resultProto.Value())
	}

	floatProto := prototype.NewPrototype(15.5)
	resultProto, err = floatProto.Divide(3.1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 15.5/3.1 {
		t.Errorf("Expected %v, got %v", 15.5/3.1, resultProto.Value())
	}

	_, err = numProto.Divide(0)
	if err == nil {
		t.Error("Expected division by zero error, got nil")
	}

	strProto := prototype.NewPrototype("hello")
	_, err = strProto.Divide(2)
	if err == nil {
		t.Error("Expected an error for non-number value, got nil")
	}
}

func TestIsEven(t *testing.T) {
	intProto := prototype.NewPrototype(10)
	resultProto, err := intProto.IsEven()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != true {
		t.Errorf("Expected true, got %v", resultProto.Value())
	}

	oddProto := prototype.NewPrototype(7)
	resultProto, err = oddProto.IsEven()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != false {
		t.Errorf("Expected false, got %v", resultProto.Value())
	}

	floatProto := prototype.NewPrototype(10.5)
	_, err = floatProto.IsEven()
	if err == nil {
		t.Error("Expected an error for float value, got nil")
	}

	strProto := prototype.NewPrototype("hello")
	_, err = strProto.IsEven()
	if err == nil {
		t.Error("Expected an error for non-number value, got nil")
	}
}

func TestIsOdd(t *testing.T) {
	intProto := prototype.NewPrototype(7)
	resultProto, err := intProto.IsOdd()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != true {
		t.Errorf("Expected true, got %v", resultProto.Value())
	}

	evenProto := prototype.NewPrototype(10)
	resultProto, err = evenProto.IsOdd()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != false {
		t.Errorf("Expected false, got %v", resultProto.Value())
	}

	floatProto := prototype.NewPrototype(10.5)
	_, err = floatProto.IsOdd()
	if err == nil {
		t.Error("Expected an error for float value, got nil")
	}

	strProto := prototype.NewPrototype("hello")
	_, err = strProto.IsOdd()
	if err == nil {
		t.Error("Expected an error for non-number value, got nil")
	}
}

func TestModulo(t *testing.T) {
	intProto := prototype.NewPrototype(10)
	resultProto, err := intProto.Modulo(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 1 {
		t.Errorf("Expected 1, got %v", resultProto.Value())
	}

	_, err = intProto.Modulo(0)
	if err == nil {
		t.Error("Expected division by zero error, got nil")
	}

	floatProto := prototype.NewPrototype(10.5)
	_, err = floatProto.Modulo(3)
	if err == nil {
		t.Error("Expected error for float number, got nil")
	}

	strProto := prototype.NewPrototype("hello")
	_, err = strProto.Modulo(3)
	if err == nil {
		t.Error("Expected error for non-number value, got nil")
	}
}

func TestPower(t *testing.T) {
	intProto := prototype.NewPrototype(2)
	resultProto, err := intProto.Power(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 8.0 {
		t.Errorf("Expected 8, got %v", resultProto.Value())
	}

	floatProto := prototype.NewPrototype(2.5)
	resultProto, err = floatProto.Power(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 6.25 {
		t.Errorf("Expected 6.25, got %v", resultProto.Value())
	}

	strProto := prototype.NewPrototype("hello")
	_, err = strProto.Power(2)
	if err == nil {
		t.Error("Expected an error for non-number value, got nil")
	}
}

func TestSubtract(t *testing.T) {
	intProto := prototype.NewPrototype(10)
	resultProto, err := intProto.Subtract(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 7 {
		t.Errorf("Expected 7, got %v", resultProto.Value())
	}

	floatProto := prototype.NewPrototype(15.5)
	resultProto, err = floatProto.Subtract(2.5)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 13.0 {
		t.Errorf("Expected 13.0, got %v", resultProto.Value())
	}

	strProto := prototype.NewPrototype("hello")
	_, err = strProto.Subtract(3)
	if err == nil {
		t.Error("Expected an error for non-number value, got nil")
	}
}

func TestAbs(t *testing.T) {
	intPrototype := prototype.NewPrototype(-10)
	intResult, err := intPrototype.Abs()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if intResult.Value() != 10 {
		t.Errorf("Expected 10, got %v", intResult.Value())
	}

	floatPrototype := prototype.NewPrototype(-10.5)
	floatResult, err := floatPrototype.Abs()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if floatResult.Value() != 10.5 {
		t.Errorf("Expected 10.5, got %v", floatResult.Value())
	}

	posPrototype := prototype.NewPrototype(10)
	posResult, err := posPrototype.Abs()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if posResult.Value() != 10 {
		t.Errorf("Expected 10, got %v", posResult.Value())
	}

	stringPrototype := prototype.NewPrototype("string")
	_, err = stringPrototype.Abs()
	if err == nil {
		t.Errorf("Expected an error for non-numeric value, got none")
	}
}

func TestRound(t *testing.T) {
	numPrototype := prototype.NewPrototype(10.4)
	result, err := numPrototype.Round()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != 10.0 {
		t.Errorf("Expected 10, got %v", result.Value())
	}

	numPrototype = prototype.NewPrototype(10.6)
	result, err = numPrototype.Round()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != 11.0 {
		t.Errorf("Expected 11, got %v", result.Value())
	}

	numPrototype = prototype.NewPrototype(10)
	result, err = numPrototype.Round()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != 10 {
		t.Errorf("Expected 10, got %v", result.Value())
	}

	strPrototype := prototype.NewPrototype("string")
	_, err = strPrototype.Round()
	if err == nil {
		t.Errorf("Expected an error for non-numeric value, got none")
	}
}

func TestSqrt(t *testing.T) {
	intPrototype := prototype.NewPrototype(16)
	result, err := intPrototype.Sqrt()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != 4.0 {
		t.Errorf("Expected 4, got %v", result.Value())
	}

	floatPrototype := prototype.NewPrototype(20.25)
	result, err = floatPrototype.Sqrt()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != 4.5 {
		t.Errorf("Expected 4.5, got %v", result.Value())
	}

	negativePrototype := prototype.NewPrototype(-16)
	_, err = negativePrototype.Sqrt()
	if err == nil {
		t.Errorf("Expected an error for negative number, got none")
	}

	strPrototype := prototype.NewPrototype("string")
	_, err = strPrototype.Sqrt()
	if err == nil {
		t.Errorf("Expected an error for non-numeric value, got none")
	}
}

func TestFactorial(t *testing.T) {
	num := prototype.NewPrototype(5)
	result, err := num.Factorial()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := 120
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	num = prototype.NewPrototype(0)
	result, err = num.Factorial()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = 1
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	num = prototype.NewPrototype(-3)
	_, err = num.Factorial()
	if err == nil {
		t.Errorf("Expected an error for negative numbers, got none")
	}
}

func TestIsPrime(t *testing.T) {
	num := prototype.NewPrototype(7)
	result, err := num.IsPrime()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != true {
		t.Errorf("Expected true, got %v", result.Value())
	}

	num = prototype.NewPrototype(10)
	result, err = num.IsPrime()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != false {
		t.Errorf("Expected false, got %v", result.Value())
	}

	num = prototype.NewPrototype(1)
	result, err = num.IsPrime()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != false {
		t.Errorf("Expected false, got %v", result.Value())
	}

	num = prototype.NewPrototype("not a number")
	_, err = num.IsPrime()
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestClamp(t *testing.T) {
	num := prototype.NewPrototype(15.0)
	result, err := num.Clamp(10, 20)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := 15.0
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	num = prototype.NewPrototype(5.0)
	result, err = num.Clamp(10, 20)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = 10.0
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	num = prototype.NewPrototype(25.0)
	result, err = num.Clamp(10, 20)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = 20.0
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	str := prototype.NewPrototype("string")
	_, err = str.Clamp(10, 20)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	num = prototype.NewPrototype(15.0)
	_, err = num.Clamp(20, 10)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestFibonacci(t *testing.T) {
	num := prototype.NewPrototype(5)
	result, err := num.Fibonacci()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := []int{0, 1, 1, 2, 3}
	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	num = prototype.NewPrototype(0)
	result, err = num.Fibonacci()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = []int{0}
	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	num = prototype.NewPrototype(1)
	result, err = num.Fibonacci()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = []int{0, 1}
	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	str := prototype.NewPrototype("string")
	_, err = str.Fibonacci()
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	num = prototype.NewPrototype(-5)
	_, err = num.Fibonacci()
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	num := prototype.NewPrototype(8)
	result, err := num.IsPowerOfTwo()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != true {
		t.Errorf("Expected true, got %v", result.Value())
	}

	num = prototype.NewPrototype(10)
	result, err = num.IsPowerOfTwo()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != false {
		t.Errorf("Expected false, got %v", result.Value())
	}

	num = prototype.NewPrototype(1)
	result, err = num.IsPowerOfTwo()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != true {
		t.Errorf("Expected true, got %v", result.Value())
	}

	notNum := prototype.NewPrototype("string")
	_, err = notNum.IsPowerOfTwo()
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	num = prototype.NewPrototype(-4)
	_, err = num.IsPowerOfTwo()
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}
