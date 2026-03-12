package main

import "fmt"

func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    // Space Complexity: O(1) since there's most 26 different characters
    count := make(map[rune]int)
    // Time Complexity: O(n+m) since n as length of s and m as length of t
    for _, c := range s{
        count[c] ++
    }
    for _, c:= range t{
        if count[c]--; count[c] == 0{
            delete(count, c)
        }
    }
    if len(count) == 0{
        return true
    }
    return false
}


func testcase_1() {
	s := "racecar"
	t := "carrace"
	fmt.Println(isAnagram(s, t)) // true
}

func testcase_2() {
	s := "jar"
	t := "jam"
	fmt.Println(isAnagram(s, t)) // false
}

func main() {
	testcase_1()
	testcase_2()
}