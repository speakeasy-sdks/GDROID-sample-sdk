// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ViewCardAPIRequest struct {
	// Unique, hyphen-separated alphanumeric string. We recommend the UUID algorithm, an industry standard that guarantees a high degree of uniqueness length 14 to 36 characters.
	ACorrelationID string `header:"style=simple,explode=false,name=A-Correlation-ID"`
	// Unique Merchant Code provided by Ayoconnect.
	AMerchantCode string `header:"style=simple,explode=false,name=A-Merchant-Code"`
	// Unique Card ID returned in Issue card API
	CardID string `pathParam:"style=simple,explode=false,name=cardID"`
	// Unique Customer ID returned in Issue card API
	MerchantCustomerID string `queryParam:"style=form,explode=true,name=merchantCustomerID"`
}

func (o *ViewCardAPIRequest) GetACorrelationID() string {
	if o == nil {
		return ""
	}
	return o.ACorrelationID
}

func (o *ViewCardAPIRequest) GetAMerchantCode() string {
	if o == nil {
		return ""
	}
	return o.AMerchantCode
}

func (o *ViewCardAPIRequest) GetCardID() string {
	if o == nil {
		return ""
	}
	return o.CardID
}

func (o *ViewCardAPIRequest) GetMerchantCustomerID() string {
	if o == nil {
		return ""
	}
	return o.MerchantCustomerID
}

// ViewCardAPICardDetail - Card deatils
type ViewCardAPICardDetail struct {
	// Available balance on the card
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	// Unique card ID
	CardID *string `json:"cardID,omitempty"`
	// Initial limit of Card
	CumulativeLimit *float64 `json:"cumulativeLimit,omitempty"`
	// Remaining balance after card is deactivated
	DeactivatedBalance *float64 `json:"deactivatedBalance,omitempty"`
	// Expiry of card in YYMM format
	Expiry *string `json:"expiry,omitempty"`
	// Masked card number
	MaskedCard *string `json:"maskedCard,omitempty"`
	// Status of the card
	Status *string `json:"status,omitempty"`
}

func (o *ViewCardAPICardDetail) GetAvailableBalance() *float64 {
	if o == nil {
		return nil
	}
	return o.AvailableBalance
}

func (o *ViewCardAPICardDetail) GetCardID() *string {
	if o == nil {
		return nil
	}
	return o.CardID
}

func (o *ViewCardAPICardDetail) GetCumulativeLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.CumulativeLimit
}

func (o *ViewCardAPICardDetail) GetDeactivatedBalance() *float64 {
	if o == nil {
		return nil
	}
	return o.DeactivatedBalance
}

func (o *ViewCardAPICardDetail) GetExpiry() *string {
	if o == nil {
		return nil
	}
	return o.Expiry
}

func (o *ViewCardAPICardDetail) GetMaskedCard() *string {
	if o == nil {
		return nil
	}
	return o.MaskedCard
}

func (o *ViewCardAPICardDetail) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

// Webview - Card deatils
type Webview struct {
	// Unique card webview URI
	URI *string `json:"URI,omitempty"`
}

func (o *Webview) GetURI() *string {
	if o == nil {
		return nil
	}
	return o.URI
}

// ViewCardAPIResponseBody - Successful Card Issuance
type ViewCardAPIResponseBody struct {
	// Card deatils
	CardDetail *ViewCardAPICardDetail `json:"cardDetail,omitempty"`
	// Unique code.
	Code *float64 `json:"code,omitempty"`
	// Unique customer ID
	MerchantCustomerID *string `json:"merchantCustomerID,omitempty"`
	// Description of the response code.
	Message *string `json:"message,omitempty"`
	// Exact TimeStamp of the response in Unix Nanoseconds format.
	ResponseTime *string `json:"responseTime,omitempty"`
	// Card deatils
	Webview *Webview `json:"webview,omitempty"`
}

func (o *ViewCardAPIResponseBody) GetCardDetail() *ViewCardAPICardDetail {
	if o == nil {
		return nil
	}
	return o.CardDetail
}

func (o *ViewCardAPIResponseBody) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ViewCardAPIResponseBody) GetMerchantCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantCustomerID
}

func (o *ViewCardAPIResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ViewCardAPIResponseBody) GetResponseTime() *string {
	if o == nil {
		return nil
	}
	return o.ResponseTime
}

func (o *ViewCardAPIResponseBody) GetWebview() *Webview {
	if o == nil {
		return nil
	}
	return o.Webview
}

type ViewCardAPIResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful Card Issuance
	Object *ViewCardAPIResponseBody
}

func (o *ViewCardAPIResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ViewCardAPIResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ViewCardAPIResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ViewCardAPIResponse) GetObject() *ViewCardAPIResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}