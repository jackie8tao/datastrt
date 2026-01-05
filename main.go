package main

import (
	"gohelper/dst"
	"time"
)

func main() {
	bt := dst.NewBPTree()
	for i := 0; i < 30; i++ {
		bt.Insert(i+1, time.Now().String())
	}
	bt.Print()
}
