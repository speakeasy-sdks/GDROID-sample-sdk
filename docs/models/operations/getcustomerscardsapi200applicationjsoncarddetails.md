# GetCustomerSCardsAPI200ApplicationJSONCardDetails

Card details


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `AvailableBalance`                          | **float64*                                  | :heavy_minus_sign:                          | Available balance on the card               |
| `CardID`                                    | **string*                                   | :heavy_minus_sign:                          | Unique card ID                              |
| `CumulativeLimit`                           | **float64*                                  | :heavy_minus_sign:                          | Initial limit of Card                       |
| `DeactivatedBalance`                        | **float64*                                  | :heavy_minus_sign:                          | Remaining balance after card is deactivated |
| `Expiry`                                    | **string*                                   | :heavy_minus_sign:                          | Expiry of card in YYMM format               |
| `MaskedCard`                                | **string*                                   | :heavy_minus_sign:                          | Masked card number                          |
| `Status`                                    | **string*                                   | :heavy_minus_sign:                          | Status of the card                          |