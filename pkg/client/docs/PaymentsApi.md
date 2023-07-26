# \PaymentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentsPost**](PaymentsApi.md#PaymentsPost) | **Post** /payments | Create a new payment
[**WebhooksSubscriptionsPost**](PaymentsApi.md#WebhooksSubscriptionsPost) | **Post** /webhooks/subscriptions | Subscribe to events



## PaymentsPost

> Payment PaymentsPost(ctx).Payment(payment).Execute()

Create a new payment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    payment := *openapiclient.NewPayment(float32(123), *openapiclient.NewAccount("AccountHolder_example", "AccountNumber_example"), *openapiclient.NewAccount("AccountHolder_example", "AccountNumber_example")) // Payment | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentsApi.PaymentsPost(context.Background()).Payment(payment).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentsPost`: Payment
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PaymentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payment** | [**Payment**](Payment.md) |  | 

### Return type

[**Payment**](payment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksSubscriptionsPost

> WebhooksSubscriptionsPost(ctx).InlineObject(inlineObject).Execute()

Subscribe to events

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject := *openapiclient.NewInlineObject("https://myserver.com/send/callback/here", "EventType_example") // InlineObject | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentsApi.WebhooksSubscriptionsPost(context.Background()).InlineObject(inlineObject).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.WebhooksSubscriptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

