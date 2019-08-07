package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello world;")
	res := match("SADFASDGRGMAMAMMIASDASDWF", "MAMAMMIA")
	fmt.Println(res)
}



func match(t string, p string) int {
	next := getNext(p)

	fmt.Println(next)

	i := 0
	j := 0

	for i < len(t) && j < len(p) {

		if 0 > j || t[i] == p[j] {
			i ++
			j ++
		} else {
			j = next[j]
		}
	}

	if j == len(p) {
		return i - j
	}

	return -1

}

func getNext(p string) []int {
	next := make([]int, len(p))

	next[0] = -1
	k := next[0]
	j := 0
	
	for j < len(p) - 1 {

		if 0 > k || p[j] == p[k] {
			next[j + 1] = k + 1
			// fmt.Println(k)
			j = j + 1
			k = k + 1
		} else if p[j] != p[k] {
			k = next[k]
		}
	}
	// fmt.Println(next)
	return next
}