# IssueCardAPIResponseBody

Successful Card Issuance


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `CardDetail`                                                                                   | [*operations.IssueCardAPICardDetail](../../../pkg/models/operations/issuecardapicarddetail.md) | :heavy_minus_sign:                                                                             | Card deatils                                                                                   |
| `Code`                                                                                         | **float64*                                                                                     | :heavy_minus_sign:                                                                             | Unique code.                                                                                   |
| `MerchantCustomerID`                                                                           | **string*                                                                                      | :heavy_minus_sign:                                                                             | Unique customer ID                                                                             |
| `Message`                                                                                      | **string*                                                                                      | :heavy_minus_sign:                                                                             | Description of the response code.                                                              |
| `ResponseTime`                                                                                 | **string*                                                                                      | :heavy_minus_sign:                                                                             | Exact TimeStamp of the response in Unix Nanoseconds format.                                    |