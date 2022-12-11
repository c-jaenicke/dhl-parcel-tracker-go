package main

import (
	"dhl-parcel-tracker-go/dhl"
	"dhl-parcel-tracker-go/getenvironment"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("No package id provided!")
		os.Exit(1)
	}

	var trackingNumber string
	if args[0] == ".env" {
		trackingNumber = getenvironment.GetEnv("trackingNumber")
	} else {
		trackingNumber = args[0]
	}

	dhl.ShowCurrentStatus(trackingNumber)
}
