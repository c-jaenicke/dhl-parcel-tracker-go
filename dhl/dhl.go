package dhl

import (
	"dhl-parcel-tracker-go/dhl/request"
	"encoding/json"
	"fmt"
)

// print current status of package as id and status
func ShowCurrentStatus(trackingNumber string) {
	packageData := getPackageData(trackingNumber)

	status := packageData.Shipments[0].Status.Status

	response := fmt.Sprintf("%s: %s", trackingNumber, status)
	fmt.Println(response)
}

// TODO get full event list function here

// make request to get package info, marshal json and return json responseObject
func getPackageData(trackingNumber string) successfulResponse {
	responseData := request.Request(trackingNumber)
	var responseObject successfulResponse
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}

// successful api response with data of package
type successfulResponse struct {
	Shipments []struct {
		ServiceURL string `json:"serviceUrl"`
		ID         string `json:"id"`
		Service    string `json:"service"`
		Origin     struct {
			Address struct {
				CountryCode string `json:"countryCode"`
			} `json:"address"`
		} `json:"origin"`
		Destination struct {
			Address struct {
				CountryCode string `json:"countryCode"`
			} `json:"address"`
		} `json:"destination"`
		Status struct {
			Timestamp string `json:"timestamp"`
			Location  struct {
				Address struct {
					AddressLocality string `json:"addressLocality"`
				} `json:"address"`
			} `json:"location"`
			StatusCode  string `json:"statusCode"`
			Status      string `json:"status"`
			Description string `json:"description"`
		} `json:"status"`
		Details struct {
			Product struct {
				ProductName string `json:"productName"`
			} `json:"product"`
			ProofOfDeliverySignedAvailable bool     `json:"proofOfDeliverySignedAvailable"`
			TotalNumberOfPieces            int      `json:"totalNumberOfPieces"`
			PieceIds                       []string `json:"pieceIds"`
			Weight                         struct {
				Value    float64 `json:"value"`
				UnitText string  `json:"unitText"`
			} `json:"weight"`
		} `json:"details"`
		Events []struct {
			Timestamp string `json:"timestamp"`
			Location  struct {
				Address struct {
					AddressLocality string `json:"addressLocality"`
				} `json:"address"`
			} `json:"location,omitempty"`
			StatusCode  string `json:"statusCode"`
			Status      string `json:"status"`
			Description string `json:"description"`
		} `json:"events"`
	} `json:"shipments"`
}
