package main

import "fmt"

func hasDuplicate(nums []int) bool {
    visit := make(map[int]struct{})
    for _, e := range nums {
        if _, ok := visit[e]; ok{
            return true
        }
        visit[e] = struct{}{}
    }
    return false
}

func testcase_1() {
    nums := []int{1, 2, 3, 3}
    fmt.Println(hasDuplicate(nums)) // true
}

func testcase_2() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(hasDuplicate(nums)) // false
}


func main() {
    testcase_1()
    testcase_2()
}