package main

import (
	"strings"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/parnurzeal/gorequest"
	// "github.com/davecgh/go-spew/spew"
)

type restAPI struct {
	config   []config
	complete chan []string
	id       int
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
	Start       int    `json:"StartFile"`
	End         int    `json:"EndFile"`
	Location    string `json:"Location"`
}

var data = [][]string{}

func main() {
	log.Printf("Calls Started\n")
	api := restAPI{
		config:   readConfig(),
		complete: make(chan []string),
		id:       0,
	}

	start := time.Now()

	files, err := ioutil.ReadDir("files")
    if err != nil {
        log.Fatal(err)
	}
	totalFile:=0
	fmt.Println("Total files are:",len(files))
    for _, f := range files {
		if strings.Contains(f.Name(),".txt"){
			totalFile++;
		}
    }
	//A main function
	api.startRead(api.config[0].Start, files)

	for i := 0; i < totalFile; i++ {
		res := <-api.complete
		data = append(data, res)
	}
	writeCsv()
	elapsed := time.Since(start)
	log.Printf("It took %s\n", elapsed)

	fmt.Printf("Press any key to exit...")
	fmt.Scanf("%d")
}

func writeCsv() {
	t := time.Now()
	filename:=t.Format("201-02-21-01-01-52")
	file, err := os.Create(filename+".csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Id", "TimeStamp", "Response"}
	writer.Write(header)
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatal(err)
		}
	}
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
func (api restAPI) startRead(lower int, files []os.FileInfo) {
	dataChan := make(chan string)
	go api.handleData(dataChan)

    for _, f := range files {
		if strings.Contains(f.Name(),".txt"){
			go api.readFile(f.Name(), dataChan)
		}
    }
}
func (api restAPI) readFile(file string, dataChan chan string) {

	fmt.Println("It is : ",api.config[0].Location+file)
	raw, err := ioutil.ReadFile(api.config[0].Location+file)
	if err != nil {
		fmt.Println(err)
	}
	dataChan <- string(raw)

}
func (api restAPI) handleData(dataChan <-chan string) {
	for {
		select {
		case paylord := <-dataChan:
			api.id = api.id + 1
			go makeRequest(paylord, api.config[0].ApiEndPoint, api.config[0].ApiKey, api.config[0].ContentType, api.complete, api.id)
		}
	}
}

func makeRequest(paylord string, endpoint string, apiKey string, contentType string, complete chan []string, id int) {
	request := gorequest.New()
	_, body, err := request.Post(endpoint).Set("content-type", contentType).
		Set("Authorization", apiKey).
		End()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Response Body: \n%s\n", body)
	complete <- []string{strconv.Itoa(id), time.Now().String(), body} // signal that the routine has completed
}
