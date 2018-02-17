package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))
	jsonObjects := startRead(900374, 900375, "files")
	spew.Dump(jsonObjects)
	fmt.Println(len(jsonObjects))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
