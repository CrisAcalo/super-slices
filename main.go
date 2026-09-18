package main

import (
	"strings"

	"github.com/CrisAcalo/super-slices/slices"
	"rsc.io/quote"
)

func main() {
	//Filter
	nums := []int{1, 2, 3, 4, 5}
	slices.Filter(nums, func(n int) bool {
		return n > 3
	})
	// fmt.Println(result)

	strs := []string{"apple", "banana", "gchSerry", quote.Hello()}
	slices.Filter(strs, func(s string) bool {
		return len(s) > 5
	})

	slices.Filter(strs, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "g")
	})

	// fmt.Println(result2)

	//Includes
	// fmt.Println(slices.Includes(nums, 3))
	// fmt.Println(slices.Includes(strs, "apple"))

	slices.Includes(nums, 3)
	slices.Includes(strs, "apple")
}
