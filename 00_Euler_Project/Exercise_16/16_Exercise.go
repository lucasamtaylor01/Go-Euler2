package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	//Define 2^100
	a := new(big.Int)
	a.Exp(big.NewInt(2), big.NewInt(1000), nil)

	// Convert the big number to string
	str := a.String()

	// Define the sum
	var sum int = 0

	// Loop to sum the digits of 2^100
	for i := 0; i < len(str); i++ {
		n, err := strconv.Atoi(string(str[i]))
		if err != nil {
			fmt.Print(err)
		} else {
			sum += n
		}
	}

	//Show solution
	fmt.Println(sum)
}
