package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if l := len(args); l == 0 || l > 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}

	var (
		sum   float64
		nums  [5]float64
		total float64
	)
	fmt.Println(sum)
	fmt.Println(nums)
	fmt.Println(total)

	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}

		total++
		fmt.Println("totol", total)

		nums[i] = n
		fmt.Println("Nums", nums)

		sum += n
	}

	fmt.Println("Your numbers:", nums)
	fmt.Println("Average:", sum/total)
}
