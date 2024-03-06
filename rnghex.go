package main

import (
	"math/rand"
	"os"
	"strconv"
)

const HEX_CHARS = "0123456789abcdef"

func main() {
	if len(os.Args) != 2 {
		println("rnghex requires exactly 1 argument `length`")

		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println("`length` must be an integer value")

		return
	}

	for i := 0; i < n; i++ {
		print(string(HEX_CHARS[rand.Intn(len(HEX_CHARS))]))
	}

	println()
}
