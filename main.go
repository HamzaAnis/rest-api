package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

func main() {
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))
	jsonObjects := startRead(900374, 900375, "files")
	fmt.Println(len(jsonObjects))
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
