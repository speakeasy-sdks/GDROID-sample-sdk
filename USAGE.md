<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	gdroidsamplesdk "github.com/speakeasy-sdks/GDROID-sample-sdk"
	"log"
)

func main() {
	s := gdroidsamplesdk.New()

	var aCorrelationID string = "a16a1bd9-6411-48a8-aeed-94c383c481ea"

	var aMerchantCode string = "AYOPOP"

	var phone string = "62-9999999999"

	ctx := context.Background()
	res, err := s.VirtualCardsAPISpecification.GetCustomerSCardsAPI(ctx, aCorrelationID, aMerchantCode, phone)
	if err != nil {
		log.Fatal(err)
	}

	if res.GetCustomerSCardsAPI200ApplicationJSONObject != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->