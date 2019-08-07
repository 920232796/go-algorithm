package main 

import (
	"fmt"
)


func canCompleteCircuit(gas []int, cost []int) int {
    for i := 0; i < len(gas); i ++ {
		startCost := 0
		flag := 0
		for j := i; j < i + len(gas); j ++ {
			k := j % len(gas)//k就是现在位于的加油站
			startCost += gas[k]//ji加上油了
			if startCost - cost[k] < 0 {
				//z证明不能到达
				flag = 1
				if k > i {
                    i = k
                }
				break
			}
			startCost -= cost[k]
		}
		if flag == 0 {
			return i
		}
	}
	return -1
}


func main() {

	// res := canCompleteCircuit([]int {1, 2, 3, 4, 5}, []int {3, 4, 5, 1, 2})

	// fmt.Println(res)

	for i := 0; i < 10; i ++ {
		fmt.Println(i)
		if i == 1 {
			i = 5
		}
		fmt.Println(i)
	}
	
}