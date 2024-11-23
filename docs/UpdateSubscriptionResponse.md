# UpdateSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriptions** | [**[]SubscriptionExternal**](SubscriptionExternal.md) |  | 
**Invoices** | [**[]InvoiceExternal**](InvoiceExternal.md) |  | 
**RenewalInvoices** | Pointer to [**[]InvoiceExternal**](InvoiceExternal.md) |  | [optional] 

## Methods

### NewUpdateSubscriptionResponse

`func NewUpdateSubscriptionResponse(subscriptions []SubscriptionExternal, invoices []InvoiceExternal, ) *UpdateSubscriptionResponse`

NewUpdateSubscriptionResponse instantiates a new UpdateSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptionResponseWithDefaults

`func NewUpdateSubscriptionResponseWithDefaults() *UpdateSubscriptionResponse`

NewUpdateSubscriptionResponseWithDefaults instantiates a new UpdateSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptions

`func (o *UpdateSubscriptionResponse) GetSubscriptions() []SubscriptionExternal`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *UpdateSubscriptionResponse) GetSubscriptionsOk() (*[]SubscriptionExternal, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *UpdateSubscriptionResponse) SetSubscriptions(v []SubscriptionExternal)`

SetSubscriptions sets Subscriptions field to given value.


### GetInvoices

`func (o *UpdateSubscriptionResponse) GetInvoices() []InvoiceExternal`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *UpdateSubscriptionResponse) GetInvoicesOk() (*[]InvoiceExternal, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *UpdateSubscriptionResponse) SetInvoices(v []InvoiceExternal)`

SetInvoices sets Invoices field to given value.


### GetRenewalInvoices

`func (o *UpdateSubscriptionResponse) GetRenewalInvoices() []InvoiceExternal`

GetRenewalInvoices returns the RenewalInvoices field if non-nil, zero value otherwise.

### GetRenewalInvoicesOk

`func (o *UpdateSubscriptionResponse) GetRenewalInvoicesOk() (*[]InvoiceExternal, bool)`

GetRenewalInvoicesOk returns a tuple with the RenewalInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalInvoices

`func (o *UpdateSubscriptionResponse) SetRenewalInvoices(v []InvoiceExternal)`

SetRenewalInvoices sets RenewalInvoices field to given value.

### HasRenewalInvoices

`func (o *UpdateSubscriptionResponse) HasRenewalInvoices() bool`

HasRenewalInvoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


