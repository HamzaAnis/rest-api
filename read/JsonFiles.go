package read

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"
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

func readFile() {
	var data []FundRaiser
	raw, err := ioutil.ReadFile("files/Fundraiser_900374.txt")
	// fmt.Println(string(raw))
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	json.Unmarshal(raw, &data)
	spew.Dump(data[0])
}
