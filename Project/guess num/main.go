package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	maxnum := 100
	r := rand.Intn(maxnum)
	//fmt.Println("rand num is ", r)

	for {
		println("please input your guess")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read input error")
			continue
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		//println(input)

		if err != nil {
			fmt.Println("convert string to int error")
			continue
		}
		println("guess is ", guess)
		if guess > r {
			fmt.Println("guess is bigger")
		} else if guess < r {
			fmt.Println("guess is smaller")
		} else {
			fmt.Println("guess is right")
			break
		}

	}

}
