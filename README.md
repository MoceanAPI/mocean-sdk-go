MoceanAPI Client Library for [Go](https://golang.org/) 
============================

This is the [Go](https://golang.org/) library for use Mocean's API. To use this, you'll need a Mocean account. Sign up [for free at 
moceanapi.com](https://dashboard.moceanapi.com/register?medium=github&campaign=sdk-go).

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

mocean := moceango.NewMoceanClient("apiKey", "apiSecret")
```

## Available API
```go
mocean.Sms().Send()             //Send SMS
mocean.Sms().getMsgStatus()     //Get Message Status
mocean.Account().getBalance()   //Get Account Balance
mocean.Account().getPricing()   //Get Account Pricing
mocean.Verify().sendCode()      //Send Verify Code
mocean.Verify().verifyCode()    //Check Verify Code
```

## Example

```go
import "gopkg.in/MoceanAPI/mocean-sdk-go.v1"

mocean := moceango.NewMoceanClient(testParams["apiKey"], testParams["apiSecret"])

message := &moceango.Message{
	From: "Mocean",
	To:   "60123456789",
	Text: "testing",
}

res, err := mocean.Sms().Send(message)
if err != nil {
	fmt.println(err)
} else {
	fmt.println("Msgid: " + res.Messages[0].Msgid)
}
```

## License

This library is released under the [MIT License](LICENSE)