package main

import (
	"Encoding/json"
	"fmt"
	"io/ioutil"
)

//A struct to hold the json ibject
type FundRaiser struct {
	EpicFundraiserID  string `json:"EpicFundraiserId"`
	FundraiserName    string `json:"FundraiserName"`
	FundraiserTitle   string `json:"FundraiserTitle"`
	FundraiserSummary string `json:"FundraiserSummary"`
	InactiveMessage   string `json:"InactiveMessage"`
	IsActive          string `json:"IsActive"`
	BranchCode        string `json:"BranchCode"`
	Category          string `json:"Category"`
	Goal              string `json:"Goal"`
	Country           string `json:"Country"`
	State             string `json:"State"`
	City              string `json:"City"`
	Postcode          string `json:"Postcode"`
	Address           string `json:"Address"`
	AccountNo         string `json:"AccountNo"`
	OrderEmail        string `json:"OrderEmail"`
	OrderPhone        string `json:"OrderPhone"`
	PriceType         string `json:"PriceType"`
	SellingEdition    string `json:"SellingEdition"`
	GroupPageName     string `json:"GroupPageName"`
	GroupKey          string `json:"GroupKey"`
	GroupPassword     string `json:"GroupPassword"`
	ThankYouTitle     string `json:"ThankYouTitle"`
	ThankYouMessage   string `json:"ThankYouMessage"`
	MerchantAccountNo string `json:"MerchantAccountNo"`
	Theme             string `json:"Theme"`
	acceptDonations   string `json:"acceptDonations"`
	isDeleted         string `json:"isDeleted"`
	videoUrl          string `json:"videoUrl"`
	DeliveryOptions   string `json:"DeliveryOptions"`
	Survey            string `json:"Survey"`
}

func main() {
	plan, _ := ioutil.ReadFile("files/Fundraiser_900374.txt")
	var data interface{}
	json.Unmarshal(plan, &data)
	fmt.Println(data)
}
