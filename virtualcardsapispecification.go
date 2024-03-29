// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package gdroidsamplesdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/GDROID-sample-sdk/internal/hooks"
	"github.com/speakeasy-sdks/GDROID-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/GDROID-sample-sdk/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/GDROID-sample-sdk/pkg/utils"
	"io"
	"net/http"
	"net/url"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://dev.api.of.ayoconnect.id",
	"https://sandbox.api.of.ayoconnect.id",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	Client HTTPClient

	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
	Hooks             *hooks.Hooks
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// VirtualCardsAPISpecification - Virtual cards API Specification: <h2>Authentication</h2>
// <p>Ayoconnect API calls require your application to authenticate itself. This is done using some API keys associated with your account. You can get your API keys by signing in the portal (visit the 'Sign in' page). A pair of API keys consists of a 'client id' and a 'client secret'. <br><br>for example:
//
//	</p>
//
// <pre>
// <code>
//
//	{
//	  "client_id": "xxxxxxxxxxx",
//	  "client_secret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
//	}
//	</code>
//	</pre>
//	<p>
//	Note that once the API keys are generated and distributed with your application, you should keep them in a safe storage as they should not be shared with anybody outside your organization.
//
//	All requests to other endpoints are authenticated using HTTP bearer authentication header, with the following format:<br>
//	  <pre>
//	  <code>
//	  Authorization: Bearer {authToken}
//	  </code>
//	  </pre>
//	</p>
//	<p>
//	A missing, incorrect, or revoked token will cause a 401 HTTP error to be returned.
//	</p>
//	<h2>Resource Types</h2>
//	<p>
//	URIs are relative to https://sandbox.api.of.ayoconnect.id unless otherwise noted.
//	</p>
//	    <h4>Access Token</h4>
//	    <ld>
//	      <p>Generated Using <i>Generate Access Token</i> <code>/v1/oauth/client_credential/accesstoken</code> and passed in headers as Bearer Token.
//	          B2B Token is mandatory in all API requests.
//	      </p>
//
//	```Authorization : Bearer <b2b_token>```
//	    </ld>
//
// <br>
// <br>
// <h2>Sequence</h2>
// <p>
// <img src="https://storage.googleapis.com/vcn-static-ui-sandbox/assets/sequence.drawio.png">
// </p>
//
// <h2>Transaction Notification</h2>
// <p>
// For every transaction notification Ayoconnect receives from MasterCard, Ayoconnect will forward the payload via a callback request to your callback endpoint.
//
// Your callback endpoint must accept a POST method and return HTTP Status Code 200 to acknowledge the payload was received and processed successfully.
//
// </p>
// <h3>Security</h3>
// <p>All callback requests from Ayoconnect will originate from following IP Addresses respective to the environments.
//
//	Clients are required to whitelist all of the following IP addresses for successfully receiving the callback notification.
//
// </p>
//
// | Envrironment  | IP / CIDR                                       |
// |---------------|-------------------------------------------------|
// | Sandbox       | <code>34.128.108.0/32</code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<code>34.101.83.121</code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<code>34.101.107.60</code>     |
// | Production    | <code>34.101.108.126/32</code>&nbsp;&nbsp;<code>34.101.184.87</code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<code>34.101.242.235</code>|
// <br>
//
// | Parameter	     |Type    |Description             |Example   |
// |------------------|--------|------------------------|-------------|
// | <code>X-Authorization</code>  | string | Unique Hash generated using client_secret and merchant_code| secret : <code>ABCD</code><br>merchant code : <code>AY</code> <br>input : <code>ABCDAY</code> <br>hash : <code>$2a$12$rBv3GkRaa05etVVDAAIgRO0bwRjtRJ9IC/jNhbJNWuQZJKJHxkXv.</code>|
// | <code>X-Content-Length</code> | number | length of body payload | <code>10</code> |
//
// <br>
// <h3>Payload</h3>
// <p>This is the sample payload that Ayoconnect will send to the callback URL.</p>
//
//	<p><pre><code>{
//	            "transactionYear": "2023",
//	            "transactionDate": "0303",
//	            "transactionTime": "030303",
//	            "approvalCode": "123456",
//	            "transactionID": "GCMNFH8669330303030303",
//	            "customerID": "CUHTCrR4YQOG",
//	            "cardID": "CIjb5aE981V4",
//	            "cumulativeLimit": 1000,
//	            "availableBalance": 800,
//	            "amount": 1000,
//	            "responseCode": "00",
//	            "merchant": {
//	              "ID": "999999999999",
//	              "categoryCode": "5072",
//	              "name": "Warung Kopi",
//	              "stateOrCountryCode": "MO"
//	            }
//	        }
//	  </code></pre></p>
//
// <h3>Response code</h3>
// <table>
//
//	<tr>
//	    <th>
//	        <p>
//	            Code
//	        </p>
//	    </th>
//	    <th>
//	        <p>
//	            Description
//	        </p>
//	    </th>
//	    <th>
//	        <p>
//	          Action
//	        </p>
//	    </th>
//	</tr>
//	<tr>
//	    <td>
//	        00
//	    </td>
//	    <td>
//	        Approved or completed successfully
//	    </td>
//	    <td>
//	        <p>success</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>08</p>
//	    </td>
//	    <td>
//	        <p>
//	            Success - Honor with ID
//	        </p>
//	    </td>
//	    <td>
//	        <p>
//	            success</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            10</p>
//	    </td>
//	    <td>
//	        <p>
//	            Success - Partial Approval</p>
//	    </td>
//	    <td>
//	        <p>
//	            success</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            85</p>
//	    </td>
//	    <td>
//	        <p>
//	            Success - Not declined, Valid for all zero amount transactions.</p>
//	    </td>
//	    <td>
//	        <p>
//	            success</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            87</p>
//	    </td>
//	    <td>
//	        <p>
//	            Success - Purchase Amount Only, No Cash Back Allowed</p>
//	    </td>
//	    <td>
//	        <p>
//	            success</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            01</p>
//	    </td>
//	    <td>
//	        <p>
//	            Issuer Connection lost</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            03</p>
//	    </td>
//	    <td>
//	        <p>
//	            Invalid merchant</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            04</p>
//	    </td>
//	    <td>
//	        <p>
//	            Invalid merchant</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            05</p>
//	    </td>
//	    <td>
//	        <p>
//	            Invalid Card</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            06</p>
//	    </td>
//	    <td>
//	        <p>
//	            Partial Reversal</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            07</p>
//	    </td>
//	    <td>
//	        <p>
//	            Full Reversal</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            12</p>
//	    </td>
//	    <td>
//	        <p>
//	            Invalid transaction</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            13</p>
//	    </td>
//	    <td>
//	        <p>
//	            Invalid amount</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            14</p>
//	    </td>
//	    <td>
//	        <p>
//	            Invalid card number</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            15</p>
//	    </td>
//	    <td>
//	        <p>
//	            Invalid issuer</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            30</p>
//	    </td>
//	    <td>
//	        <p>
//	            card issuer does not recognize the transaction details being entered.</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            41</p>
//	    </td>
//	    <td>
//	        <p>
//	            Lost card</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            43</p>
//	    </td>
//	    <td>
//	        <p>
//	            Stolen card</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            51</p>
//	    </td>
//	    <td>
//	        <p>
//	            Insufficient funds/over credit limit</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            54</p>
//	    </td>
//	    <td>
//	        <p>
//	            Expired card</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            55</p>
//	    </td>
//	    <td>
//	        <p>
//	            Invalid PIN</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            57</p>
//	    </td>
//	    <td>
//	        <p>
//	            Transaction not permitted to issuer/cardholder</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            58</p>
//	    </td>
//	    <td>
//	        <p>
//	            Transaction not permitted to acquirer/terminal</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            61</p>
//	    </td>
//	    <td>
//	        <p>
//	            Exceeds withdrawal amount limit</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            62</p>
//	    </td>
//	    <td>
//	        <p>
//	            Issuer declined transaction as card has some restrictions</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            63</p>
//	    </td>
//	    <td>
//	        <p>
//	            Breach</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            65</p>
//	    </td>
//	    <td>
//	        <p>
//	            Exceeds withdrawal count limit(merchant should resubmit authentication)</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            70</p>
//	    </td>
//	    <td>
//	        <p>
//	            Contact Card Issuer</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            71</p>
//	    </td>
//	    <td>
//	        <p>
//	            PIN Not Changed</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            75</p>
//	    </td>
//	    <td>
//	        <p>
//	            Allowable number of PIN tries exceeded</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            76</p>
//	    </td>
//	    <td>
//	        <p>
//	            Invalid/nonexistent “To Account” specified</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            77</p>
//	    </td>
//	    <td>
//	        <p>
//	            Invalid/nonexistent “From Account” specified</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            78</p>
//	    </td>
//	    <td>
//	        <p>
//	            Invalid/nonexistent account specified (general)</p>
//	    </td>
//	    <td>
//	        <p>fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>81</p>
//	    </td>
//	    <td>
//	        <p>
//	            Domestic Debit Transaction Not Allowed</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	    <td>
//	        <p><br /></p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            84</p>
//	    </td>
//	    <td>
//	        <p>
//	            Connection lost</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            86</p>
//	    </td>
//	    <td>
//	        <p>
//	            PIN Validation Fail</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            88</p>
//	    </td>
//	    <td>
//	        <p>
//	            Connection lost</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            89</p>
//	    </td>
//	    <td>
//	        <p>
//	            Unacceptable PIN - Retry</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            91</p>
//	    </td>
//	    <td>
//	        <p>
//	            Error in MC or Issuer system</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//	<tr>
//	    <td>
//	        <p>
//	            92</p>
//	    </td>
//	    <td>
//	        <p>
//	            Connection lost</p>
//	    </td>
//	    <td>
//	        <p>
//	            fail</p>
//	    </td>
//	</tr>
//
// </table>
type VirtualCardsAPISpecification struct {
	sdkConfiguration sdkConfiguration
}

type SDKOption func(*VirtualCardsAPISpecification)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *VirtualCardsAPISpecification) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *VirtualCardsAPISpecification) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *VirtualCardsAPISpecification) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *VirtualCardsAPISpecification) {
		sdk.sdkConfiguration.Client = client
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *VirtualCardsAPISpecification) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *VirtualCardsAPISpecification {
	sdk := &VirtualCardsAPISpecification{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.6.7",
			SDKVersion:        "0.6.2",
			GenVersion:        "2.291.0",
			UserAgent:         "speakeasy-sdk/go 0.6.2 2.291.0 1.6.7 github.com/speakeasy-sdks/GDROID-sample-sdk",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	return sdk
}

// GetCustomerSCardsAPI - Get cards
// This API will fetch all cards that belong to the specified customer
func (s *VirtualCardsAPISpecification) GetCustomerSCardsAPI(ctx context.Context, aCorrelationID string, aMerchantCode string, phone string) (*operations.GetCustomerSCardsAPIResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "getCustomerSCardsApi",
		SecuritySource: nil,
	}

	request := operations.GetCustomerSCardsAPIRequest{
		ACorrelationID: aCorrelationID,
		AMerchantCode:  aMerchantCode,
		Phone:          phone,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/api/v1/cards/virtual/customer/{phone}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request, nil)

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.GetCustomerSCardsAPIResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out operations.GetCustomerSCardsAPIResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// IssueCardAPI - Issue card
// This API is used to issue new cards to your customers
func (s *VirtualCardsAPISpecification) IssueCardAPI(ctx context.Context, aCorrelationID string, aMerchantCode string, requestBody *operations.IssueCardAPIRequestBody) (*operations.IssueCardAPIResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "issueCardApi",
		SecuritySource: nil,
	}

	request := operations.IssueCardAPIRequest{
		ACorrelationID: aCorrelationID,
		AMerchantCode:  aMerchantCode,
		RequestBody:    requestBody,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/api/v1/cards/virtual/issue")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "RequestBody", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request, nil)

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.IssueCardAPIResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out operations.IssueCardAPIResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// ModifySpendlimit - Modify card
// This API is used to make modifications on an exisitng virtual card
func (s *VirtualCardsAPISpecification) ModifySpendlimit(ctx context.Context, aCorrelationID string, aMerchantCode string, cardID string, requestBody *operations.ModifySpendlimitRequestBody) (*operations.ModifySpendlimitResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "modifySpendlimit",
		SecuritySource: nil,
	}

	request := operations.ModifySpendlimitRequest{
		ACorrelationID: aCorrelationID,
		AMerchantCode:  aMerchantCode,
		CardID:         cardID,
		RequestBody:    requestBody,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/api/v1/cards/virtual/{cardID}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "RequestBody", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request, nil)

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.ModifySpendlimitResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out operations.ModifySpendlimitResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// ViewCardAPI - View card
// This API is used to get card information and card webview URL
func (s *VirtualCardsAPISpecification) ViewCardAPI(ctx context.Context, aCorrelationID string, aMerchantCode string, cardID string, merchantCustomerID string) (*operations.ViewCardAPIResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "viewCardApi",
		SecuritySource: nil,
	}

	request := operations.ViewCardAPIRequest{
		ACorrelationID:     aCorrelationID,
		AMerchantCode:      aMerchantCode,
		CardID:             cardID,
		MerchantCustomerID: merchantCustomerID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/api/v1/cards/virtual/{cardID}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request, nil)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.ViewCardAPIResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out operations.ViewCardAPIResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
