package request

import (
	"dhl-parcel-tracker-go/getenvironment"
	"fmt"
	"io"
	"net/http"
	"os"
)

var apiKey string = getenvironment.GetEnv("API-KEY")
var language string = getenvironment.GetEnv("language")
var recipientPostalCode string = getenvironment.GetEnv("recipientPostalCode")

// Make a request to the given URL and return a []byte with the response body
// Logs fatal if response body is empty
// Prints error if status code is not 200
// Parameters: trackingNumber string
// Returns: responseData []byte
func Request(trackingNumber string) []byte {
	url := fmt.Sprintf("https://api-eu.dhl.com/track/shipments?trackingNumber=%s&language=%s&recipientPostalCode=%s", trackingNumber, language, recipientPostalCode)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Add("DHL-API-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while making request: "+err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Response body was empty. ")
		os.Exit(1)
	}

	if res.StatusCode == 400 {
		fmt.Printf("Response 400: Package could not be found")
		os.Exit(1)
	} else if res.StatusCode != 200 {
		fmt.Printf("Response " + res.Status)
		os.Exit(1)
	}

	//fmt.Println(string(responseData))

	return responseData
}
