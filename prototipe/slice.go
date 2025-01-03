package prototipe

import (
	"errors"
	"fmt"
)

func (p *Prototype) Append(element interface{}) (*Prototype, error) {
	return p.processSlice(func(slice []interface{}) ([]interface{}, error) {
		return append(slice, element), nil
	})
}

func (p *Prototype) Remove(index int) (*Prototype, error) {
	return p.processSlice(func(slice []interface{}) ([]interface{}, error) {
		if index < 0 || index >= len(slice) {
			return nil, errors.New("index out of bounds")
		}
		return append(slice[:index], slice[index+1:]...), nil
	})
}

func (p *Prototype) PrintSlice() (*Prototype, error) {
	if arr, ok := p.value.([]interface{}); ok {
		result := fmt.Sprintf("%v", arr)
		return &Prototype{value: result}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *Prototype) Unique() (*Prototype, error) {
	return p.processSlice(func(slice []interface{}) ([]interface{}, error) {
		seen := make(map[interface{}]bool)
		var uniqueSlice []interface{}

		for _, elem := range slice {
			if !seen[elem] {
				seen[elem] = true
				uniqueSlice = append(uniqueSlice, elem)
			}
		}
		return uniqueSlice, nil
	})
}

func (p *Prototype) Contains(element interface{}) (*Prototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		for _, v := range slice {
			if v == element {
				return &Prototype{value: true}, nil
			}
		}
		return &Prototype{value: false}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *Prototype) IndexOf(element interface{}) (*Prototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		for i, v := range slice {
			if v == element {
				return &Prototype{value: i}, nil
			}
		}
		return &Prototype{value: -1}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *Prototype) processSlice(operation func([]interface{}) ([]interface{}, error)) (*Prototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		result, err := operation(slice)
		if err != nil {
			return nil, err
		}
		return &Prototype{value: result}, nil
	}
	return nil, errors.New("value is not a slice")
}
