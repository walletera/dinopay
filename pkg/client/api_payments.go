/*
 * DinoPay API
 *
 * This API allows sending and receiving payment through the DinoPay platform.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// PaymentsApiService PaymentsApi service
type PaymentsApiService service

type ApiPaymentsPostRequest struct {
	ctx _context.Context
	ApiService *PaymentsApiService
	payment *Payment
}

func (r ApiPaymentsPostRequest) Payment(payment Payment) ApiPaymentsPostRequest {
	r.payment = &payment
	return r
}

func (r ApiPaymentsPostRequest) Execute() (Payment, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.PaymentsPostExecute(r)
}

/*
 * PaymentsPost Create a new payment
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPaymentsPostRequest
 */
func (a *PaymentsApiService) PaymentsPost(ctx _context.Context) ApiPaymentsPostRequest {
	return ApiPaymentsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return Payment
 */
func (a *PaymentsApiService) PaymentsPostExecute(r ApiPaymentsPostRequest) (Payment, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  Payment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsApiService.PaymentsPost")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.payment == nil {
		executionError.error = "payment is required and must be specified"
		return localVarReturnValue, nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.payment
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiWebhooksSubscriptionsPostRequest struct {
	ctx _context.Context
	ApiService *PaymentsApiService
	inlineObject *InlineObject
}

func (r ApiWebhooksSubscriptionsPostRequest) InlineObject(inlineObject InlineObject) ApiWebhooksSubscriptionsPostRequest {
	r.inlineObject = &inlineObject
	return r
}

func (r ApiWebhooksSubscriptionsPostRequest) Execute() (*_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.WebhooksSubscriptionsPostExecute(r)
}

/*
 * WebhooksSubscriptionsPost Subscribe to events
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiWebhooksSubscriptionsPostRequest
 */
func (a *PaymentsApiService) WebhooksSubscriptionsPost(ctx _context.Context) ApiWebhooksSubscriptionsPostRequest {
	return ApiWebhooksSubscriptionsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *PaymentsApiService) WebhooksSubscriptionsPostExecute(r ApiWebhooksSubscriptionsPostRequest) (*_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsApiService.WebhooksSubscriptionsPost")
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarPath := localBasePath + "/webhooks/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject == nil {
		executionError.error = "inlineObject is required and must be specified"
		return nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.inlineObject
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, executionError
}