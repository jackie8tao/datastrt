package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jackie8tao/datastrt/dst"
)

func main() {
	bt := dst.NewBPTree()
	for i := 0; i < 15; i++ {
		num := rand.Int() % 1000
		fmt.Println(num)
		bt.Insert(num, time.Now().String())
	}
	bt.Print()
}
