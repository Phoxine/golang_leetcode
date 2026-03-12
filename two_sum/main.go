package main

import "fmt"

func twoSum(nums []int, target int) []int {
    // (target - num) as key, and num index as value
    need_value := make(map[int]int)

    for i, v := range nums{
        if need_index, ok :=need_value[v]; ok{
            return []int{need_index, i}
        }
        need_value[target-v] = i
    }
    return []int{}
}

func testcase_1() {
	nums := []int{3, 4, 5, 6}
	target := 7
	fmt.Println(twoSum(nums, target)) // [0, 1]
}

func testcase_2() {
	nums := []int{4, 5, 6}
	target := 10
	fmt.Println(twoSum(nums, target)) // [0, 2]
}

func testcase_3() {
	nums := []int{5, 5}
	target := 10
	fmt.Println(twoSum(nums, target)) // [0, 1]
}

func main() {
	testcase_1()
	testcase_2()
	testcase_3()
}