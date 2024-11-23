# DeleteSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancellationDetails** | Pointer to [**NullableSubscriptionCancellationDetails**](SubscriptionCancellationDetails.md) |  | [optional] 
**CancelUnpaidInvoices** | Pointer to **bool** | Mark unpaid invoices as uncollectible | [optional] [default to false]
**Prorate** | Pointer to **bool** | Will generate a proration invoice_item that credits remaining unused time until the subscription period end, also creates invoice_item for un-invoiced metered usage.Setting this to false will not invoice for un-invoiced metered usage. | [optional] [default to true]
**FullRefund** | Pointer to **bool** | Flag to decide whether full refund should be given or not. | [optional] [default to false]
**IsPreview** | Pointer to **bool** | Whether the request is in preview mode (subscriptions won&#39;t actually be deleted) | [optional] [default to false]

## Methods

### NewDeleteSubscriptionRequest

`func NewDeleteSubscriptionRequest() *DeleteSubscriptionRequest`

NewDeleteSubscriptionRequest instantiates a new DeleteSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSubscriptionRequestWithDefaults

`func NewDeleteSubscriptionRequestWithDefaults() *DeleteSubscriptionRequest`

NewDeleteSubscriptionRequestWithDefaults instantiates a new DeleteSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancellationDetails

`func (o *DeleteSubscriptionRequest) GetCancellationDetails() SubscriptionCancellationDetails`

GetCancellationDetails returns the CancellationDetails field if non-nil, zero value otherwise.

### GetCancellationDetailsOk

`func (o *DeleteSubscriptionRequest) GetCancellationDetailsOk() (*SubscriptionCancellationDetails, bool)`

GetCancellationDetailsOk returns a tuple with the CancellationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationDetails

`func (o *DeleteSubscriptionRequest) SetCancellationDetails(v SubscriptionCancellationDetails)`

SetCancellationDetails sets CancellationDetails field to given value.

### HasCancellationDetails

`func (o *DeleteSubscriptionRequest) HasCancellationDetails() bool`

HasCancellationDetails returns a boolean if a field has been set.

### SetCancellationDetailsNil

`func (o *DeleteSubscriptionRequest) SetCancellationDetailsNil(b bool)`

 SetCancellationDetailsNil sets the value for CancellationDetails to be an explicit nil

### UnsetCancellationDetails
`func (o *DeleteSubscriptionRequest) UnsetCancellationDetails()`

UnsetCancellationDetails ensures that no value is present for CancellationDetails, not even an explicit nil
### GetCancelUnpaidInvoices

`func (o *DeleteSubscriptionRequest) GetCancelUnpaidInvoices() bool`

GetCancelUnpaidInvoices returns the CancelUnpaidInvoices field if non-nil, zero value otherwise.

### GetCancelUnpaidInvoicesOk

`func (o *DeleteSubscriptionRequest) GetCancelUnpaidInvoicesOk() (*bool, bool)`

GetCancelUnpaidInvoicesOk returns a tuple with the CancelUnpaidInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUnpaidInvoices

`func (o *DeleteSubscriptionRequest) SetCancelUnpaidInvoices(v bool)`

SetCancelUnpaidInvoices sets CancelUnpaidInvoices field to given value.

### HasCancelUnpaidInvoices

`func (o *DeleteSubscriptionRequest) HasCancelUnpaidInvoices() bool`

HasCancelUnpaidInvoices returns a boolean if a field has been set.

### GetProrate

`func (o *DeleteSubscriptionRequest) GetProrate() bool`

GetProrate returns the Prorate field if non-nil, zero value otherwise.

### GetProrateOk

`func (o *DeleteSubscriptionRequest) GetProrateOk() (*bool, bool)`

GetProrateOk returns a tuple with the Prorate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrate

`func (o *DeleteSubscriptionRequest) SetProrate(v bool)`

SetProrate sets Prorate field to given value.

### HasProrate

`func (o *DeleteSubscriptionRequest) HasProrate() bool`

HasProrate returns a boolean if a field has been set.

### GetFullRefund

`func (o *DeleteSubscriptionRequest) GetFullRefund() bool`

GetFullRefund returns the FullRefund field if non-nil, zero value otherwise.

### GetFullRefundOk

`func (o *DeleteSubscriptionRequest) GetFullRefundOk() (*bool, bool)`

GetFullRefundOk returns a tuple with the FullRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRefund

`func (o *DeleteSubscriptionRequest) SetFullRefund(v bool)`

SetFullRefund sets FullRefund field to given value.

### HasFullRefund

`func (o *DeleteSubscriptionRequest) HasFullRefund() bool`

HasFullRefund returns a boolean if a field has been set.

### GetIsPreview

`func (o *DeleteSubscriptionRequest) GetIsPreview() bool`

GetIsPreview returns the IsPreview field if non-nil, zero value otherwise.

### GetIsPreviewOk

`func (o *DeleteSubscriptionRequest) GetIsPreviewOk() (*bool, bool)`

GetIsPreviewOk returns a tuple with the IsPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreview

`func (o *DeleteSubscriptionRequest) SetIsPreview(v bool)`

SetIsPreview sets IsPreview field to given value.

### HasIsPreview

`func (o *DeleteSubscriptionRequest) HasIsPreview() bool`

HasIsPreview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


