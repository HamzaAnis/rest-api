package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
	"io/ioutil"
	"os"
)

func main() {
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))
	startRead(900374, 900375, "files")
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

//read files iterate over the files and read the file
//It returns the slice of all the objects read from lower to upper in the format below
func startRead(lower int, upper int, location string) {
	//A chan where the data will be sent
	dataChan := make(chan string)
	//A chan to wait for all the routines
	c := make(chan struct{})
	go handleData(dataChan)
	for i := 1; i <= 77; i++ {
		//Format to read the file
		file := fmt.Sprintf("%s%s%d%s", location, "/Fundraiser_900375.", i, ".txt")
		go readFile(file, dataChan, c)
	}
	for i := 1; i <= 77; i++ {
		<-c
	}
}
func readFile(file string, dataChan chan string, wait chan struct{}) {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	// spew.Dump(d)
	dataChan <- string(raw)
	wait <- struct{}{} // signal that the routine has completed

}
func handleData(dataChan <-chan string) {
	for {
		select {
		case relay := <-dataChan:
			fmt.Println(relay)
			fmt.Println()
			// *data = append(relay, *data...)
		}
	}
}