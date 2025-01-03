package prototipe

import "errors"

func (p *Prototype) Length() (int, error) {
	switch v := p.value.(type) {
	case string:
		return len(v), nil
	case []interface{}:
		return len(v), nil
	default:
		return 0, errors.New("value is neither a string nor an array")
	}
}

func (p *Prototype) Reverse() (*Prototype, error) {
	switch value := p.value.(type) {
	case string:
		return &Prototype{value: reverseString(value)}, nil
	case []interface{}:
		return &Prototype{value: reverseSlice(value)}, nil
	default:
		return nil, errors.New("value is not a string or slice")
	}
}

func (p *Prototype) Concat(other *Prototype) (*Prototype, error) {
	switch value1 := p.value.(type) {
	case string:
		if value2, ok := other.value.(string); ok {
			return &Prototype{value: value1 + value2}, nil
		}
		return nil, errors.New("second value is not a string")
	case []interface{}:
		if value2, ok := other.value.([]interface{}); ok {
			return &Prototype{value: append(value1, value2...)}, nil
		}
		return nil, errors.New("second value is not a slice")
	default:
		return nil, errors.New("value is not a string or slice")
	}
}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseSlice(input []interface{}) []interface{} {
	length := len(input)
	reversed := make([]interface{}, length)
	for i, v := range input {
		reversed[length-1-i] = v
	}
	return reversed
}
