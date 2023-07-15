package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := []int{4, 3, 2, 1}

	fmt.Println(doubles(nums))
	fmt.Println(fizzbuzz())

	s := "hello"
	str := "asdf"
	fmt.Println(stringIncrement(s))
	fmt.Println(stringUnique(str))

	target := 6
	nums1 := []int{3, 2, 4}
	fmt.Println(twoSums(nums1, target))
	fmt.Println(OFB())

}

func doubles(n []int) bool {
	// sort the int array
	sort.Ints(n)
	// if the same int exist next to each other, return true
	for i := 0; i < len(n)-1; i++ {
		if n[i] == n[i+1] {
			return true
		}
	}
	return false
}

func fizzbuzz() int {
	sum := 0
	for i := 0; i <= 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func anagram(a, b string) bool {
	// convert to array, sort, convert to string, compare
	return false
}

func stringIncrement(s string) string {
	split := strings.Split(s, "")
	last := len(split) - 1

	// Atoi takes in a string and returns an error if it is not an int
	// var int is the int at the last index
	if int, err := strconv.Atoi(split[last]); err == nil {
		// add 1 and convert last index value back to string
		split[last] = strconv.Itoa(int + 1)
	} else {
		split = append(split, "1")
	}

	join := strings.Join(split, "")

	return join
}

func stringUnique(s string) bool {
	// convert to array
	strArray := strings.Split(s, "")
	// sort
	sort.Strings(strArray)
	// iterate to check if a value is next to itself
	for i := 0; i < len(strArray)-1; i++ {
		if strArray[i] == strArray[i+1] {
			return false
		}
	}
	return true
}
func twoSums(nums []int, target int) []int {
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func OFB() int {
	fizzbuzz := "fIzZbUzz"
	b := len(fizzbuzz) % 2

	for f := 0; f < len(fizzbuzz); f++ {
		b++
	}

	bin := 0
	f := 1
	for b != 0 {
		r := b % 2
		bin = bin + r*f
		f = f * 10
		b = b / 2
	}
	x := 0
	for d := 0; d < bin; d++ {
		if d%3 == 0 || d%5 == 0 {
			x += d
		}
	}
	return x
}
