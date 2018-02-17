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

func readFile(lower int, upper int) []FundRaiser {
	var data []FundRaiser
	for i := lower; i <= upper; i++ {
		var d []FundRaiser
		file := fmt.Sprintf("%s%d%s", "files/Fundraiser_", i, ".txt")
		raw, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err.Error)
			os.Exit(1)
		}
		json.Unmarshal(raw, &d)
		data = append(d, data...)
	}
	return data
}
