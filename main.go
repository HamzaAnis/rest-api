package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/parnurzeal/gorequest"
	// "github.com/davecgh/go-spew/spew"
)

type restAPI struct {
	config []config
	complete chan struct{}
	//A chan to wait for all the routines
}
type config struct {
	ApiEndPoint string `json:"APiEndPoint"`
	ApiKey      string `json:"ApiKey"`
	ContentType string `json:"ContentType"`
	Schema      string `json:"Schema"`
	Username    string `json:"Username"`
	HostName    string `json:"HostName"`
	Port        string `json:"Port"`
	Start       int `json:"StartFile"`
	End 		int `json:"EndFile"`
}

func main() {
	api := restAPI{
		config: readConfig(),
		complete:  make(chan struct{}),
	}

	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))
	
	//A main function
	api.startRead(api.config[0].Start, api.config[0].End)

	for i := api.config[0].Start; i <=api.config[0].End; i++ {
		<-api.complete
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
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
//read files iterate over the files and read the file
//It returns the slice of all the objects read from lower to upper in the format below
func (api restAPI) startRead(lower int, upper int) {
	//A chan where the data will be sent
	dataChan := make(chan string)
	go api.handleData(dataChan)
	for i := lower; i <= upper; i++ {
		//Format to read the file
		file := fmt.Sprintf("%s%d%s","files/Fundraiser_", i, ".txt")
		go api.readFile(file, dataChan)
	}

}
func (api restAPI) readFile(file string, dataChan chan string) {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error)
	}
	dataChan <- string(raw)
	

}
func (api restAPI) handleData(dataChan <-chan string) {
	for {
		select {
		case paylord := <-dataChan:
			go makeRequest(paylord, api.config[0].ApiEndPoint, api.config[0].ApiKey, api.config[0].ContentType,api.complete)
		}
	}
}

func makeRequest(paylord string, endpoint string, apiKey string, contentType string,complete chan struct{}) {
	request := gorequest.New()
	fmt.Println("Calling")
	_, body, err := request.Post(endpoint).Set("content-type", contentType).
		Set("Authorization", apiKey).
		End()
	fmt.Println("Calling2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Response Body: \n%s\n", body)
	complete <- struct{}{} // signal that the routine has completed
}
