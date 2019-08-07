package main



import (
	"fmt"
	// "sync"
	
)

const N = 10 

func test(a []int) {
	b := make([]int, len(a))
	copy(b, a)
	for i := range b {
		b[i] += 100
	}

}

func main() {

	a := []int {1, 2, 3}

	test(a)

	fmt.Println(a)
	// m := make(map[int]int) 
	// wg := &sync.WaitGroup{}
	// mu := &sync.Mutex{}

	// wg.Add(N)
	// for i := 0; i < N; i ++ {
	// 	go func(value int) {
	// 		defer wg.Done()
	// 		mu.Lock()
	// 		m[value] = value
	// 		mu.Unlock()
	// 	}(i)
	// }
	// wg.Wait()

	// fmt.Println("len map = ", len(m))

	// for k, v := range m {
	// 	fmt.Println("key = ", k)
	// 	fmt.Println("v = ", v)
	// }
}