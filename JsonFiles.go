package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//read files iterate over the files and read the file
//It returns the slice of all the objects read from lower to upper in the format below
func startRead(lower int, upper int, location string) []FundRaiser {
	var data []FundRaiser
	//A chan where the data will be sent
	dataChan := make(chan []FundRaiser)
	//A chan to wait for all the routines
	c := make(chan struct{})
	go handleData(&data, dataChan)
	for i := lower; i <= upper; i++ {
		//Format to read the file
		file := fmt.Sprintf("%s%s%d%s", location, "/Fundraiser_", i, ".txt")
		go readFile(file, dataChan, c)
	}
	for i := lower; i <= upper; i++ {
		<-c
	}
	return data
}
func readFile(file string, dataChan chan []FundRaiser, wait chan struct{}) {
	var d []FundRaiser
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	json.Unmarshal(raw, &d)
	// spew.Dump(d)
	dataChan <- d
	wait <- struct{}{} // signal that the routine has completed

}
func handleData(data *[]FundRaiser, dataChan <-chan []FundRaiser) {
	for {
		select {
		case relay := <-dataChan:
			*data = append(relay, *data...)
		}
	}
}
