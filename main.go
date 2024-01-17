package main

import (
	"fmt"
	"strconv"
)

func PrintSliceInts(i []int) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func PrintSliceStrings(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func PrintSlince[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// interface
func f[T fmt.Stringer](xs []T) []string {
	result := []string{}

	for _, x := range xs {
		result = append(result, x.String())
	}
	
	return result
}

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

// 型指定
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Max[T Number] (x, y T) T {
	if x >= y {
		return x
	}
	return y
}

type Vector[T any][]T

type IntVector = Vector[int]

type T[A any, B []C, C *A] struct {
	a A
	b B
	c C
}

func main() {
	
 }