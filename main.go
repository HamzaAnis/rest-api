package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"
	// "github.com/parnurzeal/gorequest"
)

type restAPI struct {
	config []config
}
type config struct {
	ApiEndPoint string `json:"APiEndPoint"`
	ApiKey      string `json:"ApiKey"`
	ContentType string `json:"ContentType"`
	Schema      string `json:"Schema"`
	Username    string `json:"Username"`
	HostName    string `json:"HostName"`
	Port        string `json:"Port"`
}

func readConfig() []config {
	var d []config
	raw, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	json.Unmarshal(raw, &d)
	// spew.Dump(d)
	return d
}
func main() {

	api := restAPI{
		config: readConfig(),
	}

	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	//A main function
	api.startRead(900374, 900375, "files")

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

//read files iterate over the files and read the file
//It returns the slice of all the objects read from lower to upper in the format below
func (api restAPI) startRead(lower int, upper int, location string) {
	//A chan where the data will be sent
	dataChan := make(chan string)
	//A chan to wait for all the routines
	c := make(chan struct{})

	go handleData(dataChan)
	for i := 1; i <= 2; i++ {
		//Format to read the file
		file := fmt.Sprintf("%s%s%d%s", location, "/Fundraiser_900375.", i, ".txt")
		go api.readFile(file, dataChan, c)
	}
	for i := 1; i <= 2; i++ {
		<-c
	}
}
func (api restAPI) readFile(file string, dataChan chan string, wait chan struct{}) {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	fmt.Println("\t\tYESY", api.config[0].ApiEndPoint)
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

func makeRequest() {
	// request:=gorequest.New()
}
