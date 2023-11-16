# github.com/speakeasy-sdks/GDROID-sample-sdk

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/GDROID-sample-sdk.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/GDROID-sample-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README
<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/GDROID-sample-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
### Example

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
	res, err := s.GetCustomerSCardsAPI(ctx, aCorrelationID, aMerchantCode, phone)
	if err != nil {
		log.Fatal(err)
	}

	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [VirtualCardsAPISpecification SDK](docs/sdks/virtualcardsapispecification/README.md)

* [GetCustomerSCardsAPI](docs/sdks/virtualcardsapispecification/README.md#getcustomerscardsapi) - Get cards
* [IssueCardAPI](docs/sdks/virtualcardsapispecification/README.md#issuecardapi) - Issue card
* [ModifySpendlimit](docs/sdks/virtualcardsapispecification/README.md#modifyspendlimit) - Modify card
* [ViewCardAPI](docs/sdks/virtualcardsapispecification/README.md#viewcardapi) - View card
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->

<!-- End Dev Containers -->

<!-- Start Error Handling -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	gdroidsamplesdk "github.com/speakeasy-sdks/GDROID-sample-sdk"
	"github.com/speakeasy-sdks/GDROID-sample-sdk/pkg/models/sdkerrors"
	"log"
)

func main() {
	s := gdroidsamplesdk.New()

	var aCorrelationID string = "a16a1bd9-6411-48a8-aeed-94c383c481ea"

	var aMerchantCode string = "AYOPOP"

	var phone string = "62-9999999999"

	ctx := context.Background()
	res, err := s.GetCustomerSCardsAPI(ctx, aCorrelationID, aMerchantCode, phone)
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->

<!-- Start Server Selection -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://dev.api.of.ayoconnect.id` | None |
| 1 | `https://sandbox.api.of.ayoconnect.id` | None |

#### Example

```go
package main

import (
	"context"
	gdroidsamplesdk "github.com/speakeasy-sdks/GDROID-sample-sdk"
	"log"
)

func main() {
	s := gdroidsamplesdk.New(
		gdroidsamplesdk.WithServerIndex(1),
	)

	var aCorrelationID string = "a16a1bd9-6411-48a8-aeed-94c383c481ea"

	var aMerchantCode string = "AYOPOP"

	var phone string = "62-9999999999"

	ctx := context.Background()
	res, err := s.GetCustomerSCardsAPI(ctx, aCorrelationID, aMerchantCode, phone)
	if err != nil {
		log.Fatal(err)
	}

	if res.Object != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	gdroidsamplesdk "github.com/speakeasy-sdks/GDROID-sample-sdk"
	"log"
)

func main() {
	s := gdroidsamplesdk.New(
		gdroidsamplesdk.WithServerURL("https://dev.api.of.ayoconnect.id"),
	)

	var aCorrelationID string = "a16a1bd9-6411-48a8-aeed-94c383c481ea"

	var aMerchantCode string = "AYOPOP"

	var phone string = "62-9999999999"

	ctx := context.Background()
	res, err := s.GetCustomerSCardsAPI(ctx, aCorrelationID, aMerchantCode, phone)
	if err != nil {
		log.Fatal(err)
	}

	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->

<!-- Start Custom HTTP Client -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->

<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
