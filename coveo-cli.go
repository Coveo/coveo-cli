package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
)

type configS struct {
	endpoint        string
	token           string
	fields          string
	numberOfResults int
	showQueryStatus bool
	showHelp        bool
	printJSON       bool
	q               string
	username        string
	password        string
}

var (
	config configS
)

func init() {
	flag.StringVar(&config.endpoint, "e", "https://cloudplatform.coveo.com/rest/search/", "access endpoint")
	flag.StringVar(&config.fields, "f", "systitle,syssource", "fields to show")
	flag.IntVar(&config.numberOfResults, "n", 10, "numbers of results to return")
	flag.BoolVar(&config.showQueryStatus, "s", true, "show query count & duration")
	// TODO: printJSON not enabled
	flag.BoolVar(&config.printJSON, "j", false, "print original json format")
	flag.BoolVar(&config.showHelp, "help", false, "show query count & duration")
	flag.BoolVar(&config.showHelp, "h", false, "show query count & duration")

	flag.StringVar(&config.q, "q", "", "Query \"q\" term")

	// Username & password empty by default, if there is a username we will do a basic auth
	flag.StringVar(&config.username, "u", "", "Username")
	flag.StringVar(&config.password, "p", "", "Password")
	flag.StringVar(&config.token, "t", "52d806a2-0f64-4390-a3f2-e0f41a4a73ec", "access token")

	flag.Parse()
}

func main() {

	// Show help and exit
	if config.showHelp {
		fmt.Println("coveo-cli: usage")
		flag.PrintDefaults()
		return
	}

	q := &Query{}
	q.Q = config.q
	q.NumberOfResults = config.numberOfResults

	marshalledQuery, err := json.Marshal(q)
	if err != nil {
		// handle error
	}

	buf := bytes.NewReader(marshalledQuery)
	req, err := http.NewRequest("POST", config.endpoint, buf)
	req.Header.Add("Authorization", "Bearer "+config.token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accepts", "application/json")

	// If we have a username & a password do a basic auth
	if len(config.username) != 0 {
		req.SetBasicAuth(config.username, config.password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	queryResponse := &QueryResponse{}
	err = json.NewDecoder(resp.Body).Decode(queryResponse)
	if err != nil {
		// handle error
	}
	//pp.Print(queryResponse)

	fmt.Printf("Total: %d, Duration: %d\n", queryResponse.TotalCount, queryResponse.Duration)

	fields := strings.Split(config.fields, ",")

	for _, result := range queryResponse.Results {

		line := make([]string, 0, 0)
		for _, field := range fields {

			f := result.Raw[field]
			if f == nil {
				f = ""
			}

			line = append(line, f.(string))
		}

		fmt.Println(strings.Join(line, "\t|\t"))
		//    fmt.Print(result.raw[])
		//		fmt.Printf("%v", result)
	}

	fmt.Println(config.numberOfResults)

}
