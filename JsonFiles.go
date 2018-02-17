package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//A struct to hold the json object
type FundRaiser struct {
	FundraiserName    string `json:FundraiserName`
	FundraiserTitle   string `json:FundraiserTitle`
	ID                int64  `json:"EpicFundraiserId"`
	FundraiserSummary string `json:FundraiserSummary`
	InactiveMessage   string `json:InactiveMessage`
	IsActive          bool   `json:IsActive`
	BranchCode        string `json:BranchCode`
	Category          string `json:Category`
	Goal              string `json:Goal`
	Country           string `json:Country`
	State             string `json:State`
	City              string `json:City`
	Postcode          string `json:Postcode`
	Address           string `json:Address`
	AccountNo         string `json:AccountNo`
	OrderEmail        string `json:OrderEmail`
	OrderPhone        string `json:OrderPhone`
	PriceType         string `json:PriceType`
	SellingEdition    string `json:SellingEdition`
	GroupPageName     string `json:GroupPageName`
	GroupKey          string `json:GroupKey`
	GroupPassword     string `json:GroupPassword`
	ThankYouTitle     string `json:ThankYouTitle`
	ThankYouMessage   string `json:ThankYouMessage`
	MerchantAccountNo string `json:MerchantAccountNo`
	Theme             string `json:Theme`
	acceptDonations   bool   `json:acceptDonations`
	isDeleted         bool   `json:isDeleted`
	videoUrl          string `json:videoUrl`
	DeliveryOptions   string `json:DeliveryOptions`
	Survey            string `json:Survey`
}

//read files iterate over the files and read the file
//It returns the slice of all the objects read from lower to upper in the format below
func startRead(lower int, upper int, location string) []FundRaiser {
	var data []FundRaiser
	//A chan where the data will be sent
	dataChan := make(chan []FundRaiser)
	//A chan to wait for all the routines
	c := make(chan struct{})
	go handleData(&data, dataChan)
	for i := 1; i <= 77; i++ {
		//Format to read the file
		file := fmt.Sprintf("%s%s%d%s", location, "/Fundraiser_900375.", i, ".txt")
		go readFile(file, dataChan, c)
	}
	for i := 1; i < 77; i++ {
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
			fmt.Println(len(*data))
			// spew.Dump(relay)
			*data = append(relay, *data...)
		}
	}
}
