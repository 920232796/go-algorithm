package main



import (
	"fmt"
	"strconv"
)

func main() {

	res := getPermutation1(4, 9)

	fmt.Println(res)
}

	


func getPermutation1(n int, k int) string {
	
	fac := []int {1, 1, 2, 6, 24, 120, 720, 5040, 40320}

	marked := [10]int {}

	for i,_  := range (marked) {
		marked[i] = 0
	}

	
	res := ""
	k = k - 1
	var res_1 int
	var r int
	r = k
	for i := n - 1; i > 0; i -- {
		
		res_1 = r / fac[i]
		r = r % fac[i]
		count := 0

		for j := 1; j <= n; j ++ {
			if res_1 == count && marked[j] == 0 {
				res += strconv.Itoa(j)
				marked[j] = 1
				break
			}

			if marked[j] == 0 {
				//说明没出现过
				count += 1
				
			}

			


		}
		
	}

	for i := 1; i <= n; i ++ {
		if marked[i] == 0 {
			res += strconv.Itoa(i)
		}
	}

	return res

}
