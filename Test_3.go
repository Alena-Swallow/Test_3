go
package main

import "fmt"

type Calculator struct{}

func (c *Calculator) calcSum(a, b, c int) {
    aCopy := a
    bCopy := b
    c = 10 // изменяем значение c
}

func main() {
    calc := &Calculator{}
    a, b, c := 1, 2, 3
    fmt.Println("Before:", a, b, c)
    calc.calcSum(a, b, c)
    fmt.Println("After:", a, b, c)
}
