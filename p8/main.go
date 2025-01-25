package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := readFile("main.txt") // fmt.Println(genPrimes())
	fmt.Println(compute13th(data))
}

func compute13th(data string) int {
	largest := 0

	for i := 0; i < len(data)-13; i++ {
		size := i + 13
		if size > len(data) {
			break
		}
		tmp := prod(data[i:size])
		if tmp > largest {
			largest = tmp
		}
	}

	return largest
}

func prod(s string) int {
	res := 1
	for _, c := range s {
		res *= int(c - '0')
	}
	return res
}

func readFile(filename string) string {
	fd, _ := os.Open(filename)

	scanner := bufio.NewScanner(fd)

	res := ""
	for scanner.Scan() {
		res += strings.TrimSpace(scanner.Text())
	}
	return res
}
