package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func main() {
	time := time.Now().UnixNano()
	xn := big.NewInt(time)
	m := new(big.Int).SetInt64((int64(math.Pow(2, 31) - 1))) // Ideally prime
	a := big.NewInt(67280421310721) // Biggest prime by Thomas Clausen
	c := 2531011 // Wiki Parameters

	for i := 0;; i++ {
		// xn = (a * xn + c) % m (Never do big integer arithmetic like this)
		xn = new(big.Int).Mul(a, xn)
		xn = new(big.Int).Add(xn, big.NewInt(int64(c)))
		xn = new(big.Int).Mod(xn, m)
		fmt.Printf("%d   ->  %d\n",i,xn)
	}
}
