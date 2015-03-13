package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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
	groups          string
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
	flag.BoolVar(&config.printJSON, "json", false, "print original json format")

	// Username & password empty by default, if there is a username we will do a basic auth
	flag.StringVar(&config.username, "u", "", "Username")
	flag.StringVar(&config.password, "p", "", "Password")
	flag.StringVar(&config.token, "t", "52d806a2-0f64-4390-a3f2-e0f41a4a73ec", "access token")

	// Query Parameters
	flag.StringVar(&config.q, "q", "", "Query \"q\" term")
	flag.StringVar(&config.groups, "g", "", "Facets to query, if you query facets you cant query normal results")
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

	q := buildQuery()
	req, err := buildRequest(q)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	queryResponse := &QueryResponse{}
	if config.printJSON {
		d, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s\n", d)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(queryResponse)
	if err != nil {
		log.Fatal(err)
	}

	if config.showQueryStatus {
		queryStatusString := queryStatusStringFormatter(queryResponse)
		fmt.Println(queryStatusString)
	}

	// Prepare result printing
	groupByString := groupByStringFormatter(queryResponse)
	fmt.Println(groupByString)

	resultString := resultStringFormatter(queryResponse)
	fmt.Println(resultString)
}

func buildQuery() *Query {
	q := &Query{}
	q.Q = config.q
	q.FirstResult = config.skip
	q.NumberOfResults = config.numberOfResults

	// If you get facets you wont get normal results
	if len(config.groups) > 0 {
		q.NumberOfResults = 0

		for _, group := range strings.Split(config.groups, ",") {
			q.AddGroupByRequest(&GroupByRequest{
				Field: "@" + group,
				MaximumNumberOfValues: config.numberOfResults,
				SortCriteria:          "chiSquare",
				InjectionDepth:        1000,
			})
		}
	}
	return q
}
func buildRequest(q *Query) (*http.Request, error) {
	marshalledQuery, err := json.Marshal(q)
	if err != nil {
		return nil, err
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
	return req, nil
}

func queryStatusStringFormatter(qr *QueryResponse) string {
	n := config.numberOfResults
	if qr.TotalCount < n {
		n = qr.TotalCount
	}
	return fmt.Sprintf("Result: [%d-%d]/%d, Duration: %dms\n", config.skip, config.skip+config.numberOfResults-1, qr.TotalCount, qr.Duration)
}

func groupByStringFormatter(qr *QueryResponse) string {
	s := ""

	for _, groupByResult := range qr.GroupByResults {
		flen := len(groupByResult.Field)

		s = fmt.Sprintf("%s%s:\n", s, groupByResult.Field)
		for _, value := range groupByResult.Values {
			s = fmt.Sprintf("%s%s %s : %d\n", s, strings.Repeat(" ", flen), value.Value, value.NumberOfResults)
		}
	}

	return s
}

func resultStringFormatter(qr *QueryResponse) string {
	s := ""

	fields := strings.Split(config.fields, ",")

	for _, result := range qr.Results {

		s += fmt.Sprintln(result.Title)
		s += fmt.Sprintln(result.URI)

		for _, field := range fields {
			f := result.Raw[field]
			if f != nil {
				s += fmt.Sprintf("\t%s: %s\n", field, f)
			}
		}
	}

	return s
}
