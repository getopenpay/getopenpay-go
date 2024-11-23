# DeleteSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**SubscriptionExternal**](SubscriptionExternal.md) |  | 
**Invoice** | Pointer to [**NullableInvoiceExternal**](InvoiceExternal.md) |  | [optional] 

## Methods

### NewDeleteSubscriptionResponse

`func NewDeleteSubscriptionResponse(subscription SubscriptionExternal, ) *DeleteSubscriptionResponse`

NewDeleteSubscriptionResponse instantiates a new DeleteSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSubscriptionResponseWithDefaults

`func NewDeleteSubscriptionResponseWithDefaults() *DeleteSubscriptionResponse`

NewDeleteSubscriptionResponseWithDefaults instantiates a new DeleteSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *DeleteSubscriptionResponse) GetSubscription() SubscriptionExternal`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *DeleteSubscriptionResponse) GetSubscriptionOk() (*SubscriptionExternal, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *DeleteSubscriptionResponse) SetSubscription(v SubscriptionExternal)`

SetSubscription sets Subscription field to given value.


### GetInvoice

`func (o *DeleteSubscriptionResponse) GetInvoice() InvoiceExternal`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *DeleteSubscriptionResponse) GetInvoiceOk() (*InvoiceExternal, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *DeleteSubscriptionResponse) SetInvoice(v InvoiceExternal)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *DeleteSubscriptionResponse) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### SetInvoiceNil

`func (o *DeleteSubscriptionResponse) SetInvoiceNil(b bool)`

 SetInvoiceNil sets the value for Invoice to be an explicit nil

### UnsetInvoice
`func (o *DeleteSubscriptionResponse) UnsetInvoice()`

UnsetInvoice ensures that no value is present for Invoice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


