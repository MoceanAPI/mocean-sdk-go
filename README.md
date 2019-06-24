MoceanAPI Client Library for [Go](https://golang.org/) 
============================

This is the [Go](https://golang.org/) library for use Mocean's API. To use this, you'll need a Mocean account. Sign up [for free at 
moceanapi.com][signup].

 * [Installation](#installation)
 * [Usage](#usage)
 * [Example](#example)

## Installation

To install the client

```bash
go get "gopkg.in/MoceanAPI/mocean-sdk-go.v1"
```

## Usage

Create a client with your API key and secret

```go
import "gopkg.in/MoceanAPI/mocean-sdk-go.v1"

mocean := moceansdk.NewMoceanClient("apiKey", "apiSecret")
```

## Available API
```go
mocean.Message().Send()             //Send SMS
mocean.Message().GetMessageStatus() //Get Message Status
mocean.Account().GetBalance()       //Get Account Balance
mocean.Account().GetPricing()       //Get Account Pricing
mocean.Verify().SendCode()          //Send Verify Code
mocean.Verify().VerifyCode()        //Check Verify Code
mocean.NumberLookup().Inquiry()     //Number Lookup
```

## Example

To use [Mocean's SMS API][doc_sms] to send an SMS message, call the `mocean.Message.Send()` method.

The API can be called directly, using a simple array of parameters, the keys match the [parameters of the API][doc_sms].

```go
res, err := mocean.Message().Send(url.Values{
    "mocean-to": {"60123456789"},
    "mocean-from": {"MOCEAN"},
    "mocean-text": {"Hello World"}
})

if err != nil {
    fmt.Println(err)
} else {
    fmt.Printf("res: %v", res)
}
```

## Responses

For your convenient, the API response has been parsed to specific struct

```go
fmt.Printf("res: %v", res)        // show full response string
fmt.Printf("res: %s", res.Status) // show response status, "0" in this case
```

## Documentation

Kindly visit [MoceanApi Docs][doc_main] for more usage

## License

This library is released under the [MIT License][license]

[signup]: https://dashboard.moceanapi.com/register?medium=github&campaign=sdk-go
[doc_main]: https://moceanapi.com/docs/?go
[doc_sms]: https://moceanapi.com/docs/?go#send-sms
[license]: LICENSE