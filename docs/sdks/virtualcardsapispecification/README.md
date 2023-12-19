# VirtualCardsAPISpecification SDK


## Overview

Virtual cards API Specification: <h2>Authentication</h2>
<p>Ayoconnect API calls require your application to authenticate itself. This is done using some API keys associated with your account. You can get your API keys by signing in the portal (visit the 'Sign in' page). A pair of API keys consists of a 'client id' and a 'client secret'. <br><br>for example:
  </p>
<pre>
<code>
  {
    "client_id": "xxxxxxxxxxx",
    "client_secret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  }
  </code>
  </pre>
  <p>
  Note that once the API keys are generated and distributed with your application, you should keep them in a safe storage as they should not be shared with anybody outside your organization.

  All requests to other endpoints are authenticated using HTTP bearer authentication header, with the following format:<br>
    <pre>
    <code>
    Authorization: Bearer {authToken}
    </code>
    </pre>
  </p>
  <p>
  A missing, incorrect, or revoked token will cause a 401 HTTP error to be returned.
  </p>
  <h2>Resource Types</h2>
  <p>
  URIs are relative to https://sandbox.api.of.ayoconnect.id unless otherwise noted.
  </p>
      <h4>Access Token</h4>
      <ld>
        <p>Generated Using <i>Generate Access Token</i> <code>/v1/oauth/client_credential/accesstoken</code> and passed in headers as Bearer Token. 
            B2B Token is mandatory in all API requests.
        </p>

  ```Authorization : Bearer <b2b_token>```
      </ld>
<br>
<br>
<h2>Sequence</h2>  
<p>
<img src="https://storage.googleapis.com/vcn-static-ui-sandbox/assets/sequence.drawio.png">
</p>

<h2>Transaction Notification</h2>
<p>
For every transaction notification Ayoconnect receives from MasterCard, Ayoconnect will forward the payload via a callback request to your callback endpoint.

Your callback endpoint must accept a POST method and return HTTP Status Code 200 to acknowledge the payload was received and processed successfully.

</p>
<h3>Security</h3>
<p>All callback requests from Ayoconnect will originate from following IP Addresses respective to the environments.

  Clients are required to whitelist all of the following IP addresses for successfully receiving the callback notification.
</p>


| Envrironment  | IP / CIDR                                       |
|---------------|-------------------------------------------------|
| Sandbox       | <code>34.128.108.0/32</code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<code>34.101.83.121</code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<code>34.101.107.60</code>     |
| Production    | <code>34.101.108.126/32</code>&nbsp;&nbsp;<code>34.101.184.87</code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<code>34.101.242.235</code>|
<br>
    
| Parameter	     |Type    |Description             |Example   |
|------------------|--------|------------------------|-------------|
| <code>X-Authorization</code>  | string | Unique Hash generated using client_secret and merchant_code| secret : <code>ABCD</code><br>merchant code : <code>AY</code> <br>input : <code>ABCDAY</code> <br>hash : <code>$2a$12$rBv3GkRaa05etVVDAAIgRO0bwRjtRJ9IC/jNhbJNWuQZJKJHxkXv.</code>|
| <code>X-Content-Length</code> | number | length of body payload | <code>10</code> |
    
<br>
<h3>Payload</h3>
<p>This is the sample payload that Ayoconnect will send to the callback URL.</p>

<p><pre><code>{
            "transactionYear": "2023",
            "transactionDate": "0303",
            "transactionTime": "030303",
            "approvalCode": "123456",
            "transactionID": "GCMNFH8669330303030303",
            "customerID": "CUHTCrR4YQOG",
            "cardID": "CIjb5aE981V4",
            "cumulativeLimit": 1000,
            "availableBalance": 800,
            "amount": 1000,
            "responseCode": "00",
            "merchant": {
              "ID": "999999999999",
              "categoryCode": "5072",
              "name": "Warung Kopi",
              "stateOrCountryCode": "MO"
            }
        }
  </code></pre></p>
  
<h3>Response code</h3>
<table>
    <tr>
        <th>
            <p>
                Code
            </p>
        </th>
        <th>
            <p>
                Description
            </p>
        </th>
        <th>
            <p>
              Action
            </p>
        </th>
    </tr>
    <tr>
        <td>
            00
        </td>
        <td>
            Approved or completed successfully
        </td>
        <td>
            <p>success</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>08</p>
        </td>
        <td>
            <p>
                Success - Honor with ID
            </p>
        </td>
        <td>
            <p>
                success</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                10</p>
        </td>
        <td>
            <p>
                Success - Partial Approval</p>
        </td>
        <td>
            <p>
                success</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                85</p>
        </td>
        <td>
            <p>
                Success - Not declined, Valid for all zero amount transactions.</p>
        </td>
        <td>
            <p>
                success</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                87</p>
        </td>
        <td>
            <p>
                Success - Purchase Amount Only, No Cash Back Allowed</p>
        </td>
        <td>
            <p>
                success</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                01</p>
        </td>
        <td>
            <p>
                Issuer Connection lost</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                03</p>
        </td>
        <td>
            <p>
                Invalid merchant</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                04</p>
        </td>
        <td>
            <p>
                Invalid merchant</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                05</p>
        </td>
        <td>
            <p>
                Invalid Card</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                06</p>
        </td>
        <td>
            <p>
                Partial Reversal</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                07</p>
        </td>
        <td>
            <p>
                Full Reversal</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                12</p>
        </td>
        <td>
            <p>
                Invalid transaction</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                13</p>
        </td>
        <td>
            <p>
                Invalid amount</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                14</p>
        </td>
        <td>
            <p>
                Invalid card number</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                15</p>
        </td>
        <td>
            <p>
                Invalid issuer</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                30</p>
        </td>
        <td>
            <p>
                card issuer does not recognize the transaction details being entered.</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                41</p>
        </td>
        <td>
            <p>
                Lost card</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                43</p>
        </td>
        <td>
            <p>
                Stolen card</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                51</p>
        </td>
        <td>
            <p>
                Insufficient funds/over credit limit</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                54</p>
        </td>
        <td>
            <p>
                Expired card</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                55</p>
        </td>
        <td>
            <p>
                Invalid PIN</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                57</p>
        </td>
        <td>
            <p>
                Transaction not permitted to issuer/cardholder</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                58</p>
        </td>
        <td>
            <p>
                Transaction not permitted to acquirer/terminal</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                61</p>
        </td>
        <td>
            <p>
                Exceeds withdrawal amount limit</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                62</p>
        </td>
        <td>
            <p>
                Issuer declined transaction as card has some restrictions</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                63</p>
        </td>
        <td>
            <p>
                Breach</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                65</p>
        </td>
        <td>
            <p>
                Exceeds withdrawal count limit(merchant should resubmit authentication)</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                70</p>
        </td>
        <td>
            <p>
                Contact Card Issuer</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                71</p>
        </td>
        <td>
            <p>
                PIN Not Changed</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                75</p>
        </td>
        <td>
            <p>
                Allowable number of PIN tries exceeded</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                76</p>
        </td>
        <td>
            <p>
                Invalid/nonexistent “To Account” specified</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                77</p>
        </td>
        <td>
            <p>
                Invalid/nonexistent “From Account” specified</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                78</p>
        </td>
        <td>
            <p>
                Invalid/nonexistent account specified (general)</p>
        </td>
        <td>
            <p>fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>81</p>
        </td>
        <td>
            <p>
                Domestic Debit Transaction Not Allowed</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
        <td>
            <p><br /></p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                84</p>
        </td>
        <td>
            <p>
                Connection lost</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                86</p>
        </td>
        <td>
            <p>
                PIN Validation Fail</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                88</p>
        </td>
        <td>
            <p>
                Connection lost</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                89</p>
        </td>
        <td>
            <p>
                Unacceptable PIN - Retry</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                91</p>
        </td>
        <td>
            <p>
                Error in MC or Issuer system</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
    <tr>
        <td>
            <p>
                92</p>
        </td>
        <td>
            <p>
                Connection lost</p>
        </td>
        <td>
            <p>
                fail</p>
        </td>
    </tr>
</table>


### Available Operations

* [GetCustomerSCardsAPI](#getcustomerscardsapi) - Get cards
* [IssueCardAPI](#issuecardapi) - Issue card
* [ModifySpendlimit](#modifyspendlimit) - Modify card
* [ViewCardAPI](#viewcardapi) - View card

## GetCustomerSCardsAPI

This API will fetch all cards that belong to the specified customer

### Example Usage

```go
package main

import(
	gdroidsamplesdk "github.com/speakeasy-sdks/GDROID-sample-sdk"
	"context"
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

### Parameters

| Parameter                                                                                                                                                                   | Type                                                                                                                                                                        | Required                                                                                                                                                                    | Description                                                                                                                                                                 | Example                                                                                                                                                                     |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                       | :heavy_check_mark:                                                                                                                                                          | The context to use for the request.                                                                                                                                         |                                                                                                                                                                             |
| `aCorrelationID`                                                                                                                                                            | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | Unique, hyphen-separated alphanumeric string. We recommend the UUID algorithm, an industry standard that guarantees a high degree of uniqueness length 14 to 36 characters. | a16a1bd9-6411-48a8-aeed-94c383c481ea                                                                                                                                        |
| `aMerchantCode`                                                                                                                                                             | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | Unique Merchant Code provided by Ayoconnect.                                                                                                                                | AYOPOP                                                                                                                                                                      |
| `phone`                                                                                                                                                                     | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | Customer phone number used to issue a Card.                                                                                                                                 | 62-9999999999                                                                                                                                                               |


### Response

**[*operations.GetCustomerSCardsAPIResponse](../../pkg/models/operations/getcustomerscardsapiresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## IssueCardAPI

This API is used to issue new cards to your customers

### Example Usage

```go
package main

import(
	gdroidsamplesdk "github.com/speakeasy-sdks/GDROID-sample-sdk"
	"github.com/speakeasy-sdks/GDROID-sample-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := gdroidsamplesdk.New()


    var aCorrelationID string = "a16a1bd9-6411-48a8-aeed-94c383c481ea"

    var aMerchantCode string = "AYOPOP"

    requestBody := &operations.IssueCardAPIRequestBody{
        CardDetail: operations.CardDetail{
            SpendLimit: gdroidsamplesdk.Float64(1000),
            Validity: gdroidsamplesdk.Float64(1),
        },
        Email: "test@example.com",
        FullName: "john",
        GovernmentID: "320822267382a7fd6",
        Phone: "62-999999999",
    }

    ctx := context.Background()
    res, err := s.IssueCardAPI(ctx, aCorrelationID, aMerchantCode, requestBody)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                   | Type                                                                                                                                                                        | Required                                                                                                                                                                    | Description                                                                                                                                                                 | Example                                                                                                                                                                     |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                       | :heavy_check_mark:                                                                                                                                                          | The context to use for the request.                                                                                                                                         |                                                                                                                                                                             |
| `aCorrelationID`                                                                                                                                                            | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | Unique, hyphen-separated alphanumeric string. We recommend the UUID algorithm, an industry standard that guarantees a high degree of uniqueness length 14 to 36 characters. | a16a1bd9-6411-48a8-aeed-94c383c481ea                                                                                                                                        |
| `aMerchantCode`                                                                                                                                                             | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | Unique Merchant Code provided by Ayoconnect.                                                                                                                                | AYOPOP                                                                                                                                                                      |
| `requestBody`                                                                                                                                                               | [*operations.IssueCardAPIRequestBody](../../pkg/models/operations/issuecardapirequestbody.md)                                                                               | :heavy_minus_sign:                                                                                                                                                          | N/A                                                                                                                                                                         |                                                                                                                                                                             |


### Response

**[*operations.IssueCardAPIResponse](../../pkg/models/operations/issuecardapiresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ModifySpendlimit

This API is used to make modifications on an exisitng virtual card

### Example Usage

```go
package main

import(
	gdroidsamplesdk "github.com/speakeasy-sdks/GDROID-sample-sdk"
	"github.com/speakeasy-sdks/GDROID-sample-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := gdroidsamplesdk.New()


    var aCorrelationID string = "a16a1bd9-6411-48a8-aeed-94c383c481ea"

    var aMerchantCode string = "AYOPOP"

    var cardID string = "CIXXXXXXXX"

    requestBody := &operations.ModifySpendlimitRequestBody{
        CardDetail: &operations.ModifySpendlimitCardDetail{
            SpendLimit: gdroidsamplesdk.Float64(5000),
            Status: gdroidsamplesdk.String("ACTIVE/INACTIVE/DEACTIVATE"),
            Validity: gdroidsamplesdk.Float64(24),
        },
        MerchantCustomerID: gdroidsamplesdk.String("CIXXXXXXXXXX"),
    }

    ctx := context.Background()
    res, err := s.ModifySpendlimit(ctx, aCorrelationID, aMerchantCode, cardID, requestBody)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                   | Type                                                                                                                                                                        | Required                                                                                                                                                                    | Description                                                                                                                                                                 | Example                                                                                                                                                                     |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                       | :heavy_check_mark:                                                                                                                                                          | The context to use for the request.                                                                                                                                         |                                                                                                                                                                             |
| `aCorrelationID`                                                                                                                                                            | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | Unique, hyphen-separated alphanumeric string. We recommend the UUID algorithm, an industry standard that guarantees a high degree of uniqueness length 14 to 36 characters. | a16a1bd9-6411-48a8-aeed-94c383c481ea                                                                                                                                        |
| `aMerchantCode`                                                                                                                                                             | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | Unique Merchant Code provided by Ayoconnect.                                                                                                                                | AYOPOP                                                                                                                                                                      |
| `cardID`                                                                                                                                                                    | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | Unique Card ID returned in Issue card API                                                                                                                                   | CIXXXXXXXX                                                                                                                                                                  |
| `requestBody`                                                                                                                                                               | [*operations.ModifySpendlimitRequestBody](../../pkg/models/operations/modifyspendlimitrequestbody.md)                                                                       | :heavy_minus_sign:                                                                                                                                                          | N/A                                                                                                                                                                         |                                                                                                                                                                             |


### Response

**[*operations.ModifySpendlimitResponse](../../pkg/models/operations/modifyspendlimitresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ViewCardAPI

This API is used to get card information and card webview URL

### Example Usage

```go
package main

import(
	gdroidsamplesdk "github.com/speakeasy-sdks/GDROID-sample-sdk"
	"context"
	"log"
)

func main() {
    s := gdroidsamplesdk.New()


    var aCorrelationID string = "a16a1bd9-6411-48a8-aeed-94c383c481ea"

    var aMerchantCode string = "AYOPOP"

    var cardID string = "CIXXXXXXXX"

    var merchantCustomerID string = "CUXXXXXXXX"

    ctx := context.Background()
    res, err := s.ViewCardAPI(ctx, aCorrelationID, aMerchantCode, cardID, merchantCustomerID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                   | Type                                                                                                                                                                        | Required                                                                                                                                                                    | Description                                                                                                                                                                 | Example                                                                                                                                                                     |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                       | :heavy_check_mark:                                                                                                                                                          | The context to use for the request.                                                                                                                                         |                                                                                                                                                                             |
| `aCorrelationID`                                                                                                                                                            | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | Unique, hyphen-separated alphanumeric string. We recommend the UUID algorithm, an industry standard that guarantees a high degree of uniqueness length 14 to 36 characters. | a16a1bd9-6411-48a8-aeed-94c383c481ea                                                                                                                                        |
| `aMerchantCode`                                                                                                                                                             | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | Unique Merchant Code provided by Ayoconnect.                                                                                                                                | AYOPOP                                                                                                                                                                      |
| `cardID`                                                                                                                                                                    | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | Unique Card ID returned in Issue card API                                                                                                                                   | CIXXXXXXXX                                                                                                                                                                  |
| `merchantCustomerID`                                                                                                                                                        | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | Unique Customer ID returned in Issue card API                                                                                                                               | CUXXXXXXXX                                                                                                                                                                  |


### Response

**[*operations.ViewCardAPIResponse](../../pkg/models/operations/viewcardapiresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
