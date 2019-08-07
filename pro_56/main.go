package main

import (
	"fmt"
)

// 合并区间



type Interval struct {
	Start int
	End   int
}
 
func merge(intervals []Interval) []Interval {
	
	res := []Interval {}
	flag := 0;
	interval := Interval {}
	for i := 0; i < len(intervals); i ++ {
		if flag == 0 {
			//新开始
			interval.Start = intervals[i].Start 
			interval.End = intervals[i].End
			flag = 1
			if i == len(intervals) - 1 {
				res = append(res, interval)
			}
		} else {
			if interval.End > intervals[i].Start {
				//证明可以合并
				interval.End = intervals[i].End
				flag = 1
			} else {
				//不可合并
				i--
				flag = 0
				res = append(res, interval)
			}
		}
	}

	return res
    
}

func main() {

	intervals := []Interval {Interval{Start:1, End: 3}, Interval{Start: 2, End: 6}, Interval{Start: 8, End: 10}, Interval {Start: 15, End: 18}}
	res := merge(intervals)
	fmt.Println(res)
	// fmt.Println("hello world");
}