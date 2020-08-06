package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jkueh/aramex-tracker/aramex"
)

const apiEndpoint = "https://www.aramex.com.au/tracking-api/?callback"

var debug bool

func init() {
	debug = os.Getenv("DEBUG") == "true"
}

func showUsage() {
	fmt.Println("Usage:", os.Args[0], "<label_number_1> [<label_number_N>]")
}

func main() {

	if debug {
		log.Println("Arg Count:", len(os.Args))
		log.Println("Args:", os.Args)
	}

	if len(os.Args) < 2 {
		showUsage()
		os.Exit(1)
	}

	for _, labelNumber := range os.Args[1:] {
		URL := fmt.Sprintf("%s&LabelNo=%s&dataFormat=json", apiEndpoint, labelNumber)
		resp, err := http.Get(URL)
		if err != nil {
			log.Println("An error occurred while trying to issue a GET request to", apiEndpoint)
			log.Fatalln(err)
		}

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("An error occured while trying to read the response body from", apiEndpoint)
			log.Fatalln(err)
		}

		if debug {
			log.Println("Raw JSON response:")
			fmt.Println(string(content))
		}

		// Attempt to marshal the data into a TrackingResultSummary
		var resultSummary aramex.TrackingResultSummary
		err = json.Unmarshal(content, &resultSummary)
		if err != nil {
			log.Println("An error occurred while trying to unmarshal the JSON response")
			log.Fatalln(err)
		}

		if debug {
			log.Println("Resultant resultSummary struct:")
			fmt.Println(resultSummary)
		}

		if resultSummary.Error != "" {
			log.Printf("Error message from API endpoint for tracking number %s: %s", labelNumber, resultSummary.Error)
			os.Exit(5)
		}

		if debug {
			log.Println("Response Status:", resp.Status)
		}

		if resp.StatusCode != 200 {
			log.Println("Non-200 status code received:", resp.StatusCode)
			log.Println("Status message:", resp.Status)
			os.Exit(10)
		}

		fmt.Println("Label Number:", resultSummary.Result.LabelNumber)
		fmt.Println("Scan Events:")
		for _, event := range resultSummary.Result.Scans {
			fmt.Println()
			fmt.Printf("- [%s] Scan: %s\n", event.Type, event.Date)
			fmt.Printf("  Status: [%s] %s\n", event.Status, event.StatusDescription)
			fmt.Println("  Location: ", event.Name) // I _think_ this is what it's used for.
			fmt.Println("  Description:", event.Description)
			if debug {
				log.Println("Full scan event:")
				fmt.Println(event)
			}
		}
		fmt.Println()
		fmt.Println("Other fun stats™:")
		fmt.Println("  Scan event in last 24 hours:", resultSummary.Result.WasScannedLast24Hours)
		fmt.Println("  Results generated server-side in", resultSummary.GeneratedIn)
	}
}
