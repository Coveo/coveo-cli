package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type configS struct {
	endpoint string
	token    string
	username string
	password string

	showQueryStatus bool
	showHelp        bool
	printJSON       bool

	fields          string
	facets          string
	q               string
	numberOfResults int
	skip            int
}

var (
	config configS
)

func init() {
	flag.BoolVar(&config.showHelp, "help", false, "show query count & duration")
	flag.BoolVar(&config.showHelp, "h", false, "show query count & duration")

	// Endpoint
	flag.StringVar(&config.endpoint, "e", "https://cloudplatform.coveo.com/rest/search/", "access endpoint")

	// Debug Params
	flag.BoolVar(&config.showQueryStatus, "s", true, "show query count & duration")
	flag.BoolVar(&config.printJSON, "j", false, "print original json format")

	// Username & password empty by default, if there is a username we will do a basic auth
	flag.StringVar(&config.username, "u", "", "Username")
	flag.StringVar(&config.password, "p", "", "Password")
	flag.StringVar(&config.token, "t", "52d806a2-0f64-4390-a3f2-e0f41a4a73ec", "access token")

	// Query Parameters
	flag.StringVar(&config.q, "q", "", "Query \"q\" term")
	flag.StringVar(&config.facets, "facets", "", "Facets to query, if you query facets you cant query normal results")
	flag.StringVar(&config.fields, "f", "systitle,syssource", "fields to show")
	flag.IntVar(&config.numberOfResults, "n", 10, "numbers of results to return")
	flag.IntVar(&config.skip, "skip", 0, "number of results to skip")

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
	q.FirstResult = config.skip
	q.NumberOfResults = config.numberOfResults

	marshalledQuery, err := json.Marshal(q)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	defer resp.Body.Close()

	queryResponse := &QueryResponse{}
	err = json.NewDecoder(resp.Body).Decode(queryResponse)
	if err != nil {
		log.Fatal(err)
	}
	if config.printJSON {
		b, err := json.MarshalIndent(queryResponse, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
		return
	}

	fmt.Printf("Results: %d, Skipped: %d,Total: %d, Duration: %dms\n", config.numberOfResults, config.skip, queryResponse.TotalCount, queryResponse.Duration)

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
}
