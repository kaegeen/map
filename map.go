package main

import (
	"fmt"
)

type Prime struct{}

func main() {
	p := Prime{}
	a := []int{1, 2, 3, 4, 5, 6}
	result := p.Map(p.IsPrime, a)
	fmt.Println(result)
}

func (p Prime) IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func (p Prime) Map(fn func(int) bool, slice []int) []bool {
	result := make([]bool, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}
