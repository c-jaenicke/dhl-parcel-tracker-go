# dhl-parcel-tracker-go

Get current status and past events of parcel send using DHL, utilizing their unified shipping API.

## Usage

Either run `./dhl-parcel-tracker-go <package id>` or `./dhl-parcel-tracker-go .env` to read the package id from a `.env` file in the same directory.

## .env File

Requires a .env file

```text
API-KEY=<your app api key>
trackingNumber=<tracking code>
language=<ISO 639-1 2-character language code for the response>
recipientPostalCode=<postal code of destination address for more accurate info>
```

`recipientPostalCode` is optional, it is needed however to get more accurate tracking and data, much like on the official DHL tracking site.

`language` is optional, it provides a translated status if given.

### API Key

You need a DHL API Key, register as a developer here [https://developer.dhl.com/](https://developer.dhl.com/), create a new application and add the [Shipment Tracking - Unified](https://developer.dhl.com/api-reference/shipment-tracking#get-started-section/) to the application.

Wait until your application gets approved and copy your API Key.
