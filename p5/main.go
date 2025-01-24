package main

import "fmt"

func main() {
        s := 2520
        for {
                divisibleByAll := true 
                for i := 1; i <= 20; i++ {
                        if s % i != 0 {
                                divisibleByAll = false
                                break
                        }
                }
                if divisibleByAll {
                        fmt.Println(s)
                        break
                }
                s++
        }
}
