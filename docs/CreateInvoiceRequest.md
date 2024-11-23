# CreateInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**PaymentMethodId** | Pointer to **NullableString** |  | [optional] 
**CollectionMethod** | Pointer to [**NullableCollectionMethodEnum**](CollectionMethodEnum.md) |  | [optional] 
**Description** | Pointer to **string** | Description for newly created invoice | [optional] [default to "Manual creation of invoice"]
**CouponId** | Pointer to **NullableString** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**SelectedProductPriceQuantity** | Pointer to [**[]SelectedPriceQuantity**](SelectedPriceQuantity.md) |  | [optional] 
**InvoiceType** | Pointer to [**InvoiceType**](InvoiceType.md) |  | [optional] [default to INVOICETYPE_STANDARD]
**CustomerId** | **string** | The external id of the customer. | 
**NetD** | Pointer to **NullableInt32** |  | [optional] 
**EmailInvoiceOnFinalization** | Pointer to **NullableBool** |  | [optional] 
**FinalizeInvoiceImmediately** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCreateInvoiceRequest

`func NewCreateInvoiceRequest(customerId string, ) *CreateInvoiceRequest`

NewCreateInvoiceRequest instantiates a new CreateInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvoiceRequestWithDefaults

`func NewCreateInvoiceRequestWithDefaults() *CreateInvoiceRequest`

NewCreateInvoiceRequestWithDefaults instantiates a new CreateInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *CreateInvoiceRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateInvoiceRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateInvoiceRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateInvoiceRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *CreateInvoiceRequest) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *CreateInvoiceRequest) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetPaymentMethodId

`func (o *CreateInvoiceRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *CreateInvoiceRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *CreateInvoiceRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *CreateInvoiceRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### SetPaymentMethodIdNil

`func (o *CreateInvoiceRequest) SetPaymentMethodIdNil(b bool)`

 SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil

### UnsetPaymentMethodId
`func (o *CreateInvoiceRequest) UnsetPaymentMethodId()`

UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
### GetCollectionMethod

`func (o *CreateInvoiceRequest) GetCollectionMethod() CollectionMethodEnum`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *CreateInvoiceRequest) GetCollectionMethodOk() (*CollectionMethodEnum, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *CreateInvoiceRequest) SetCollectionMethod(v CollectionMethodEnum)`

SetCollectionMethod sets CollectionMethod field to given value.

### HasCollectionMethod

`func (o *CreateInvoiceRequest) HasCollectionMethod() bool`

HasCollectionMethod returns a boolean if a field has been set.

### SetCollectionMethodNil

`func (o *CreateInvoiceRequest) SetCollectionMethodNil(b bool)`

 SetCollectionMethodNil sets the value for CollectionMethod to be an explicit nil

### UnsetCollectionMethod
`func (o *CreateInvoiceRequest) UnsetCollectionMethod()`

UnsetCollectionMethod ensures that no value is present for CollectionMethod, not even an explicit nil
### GetDescription

`func (o *CreateInvoiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInvoiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInvoiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateInvoiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCouponId

`func (o *CreateInvoiceRequest) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *CreateInvoiceRequest) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *CreateInvoiceRequest) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *CreateInvoiceRequest) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponIdNil

`func (o *CreateInvoiceRequest) SetCouponIdNil(b bool)`

 SetCouponIdNil sets the value for CouponId to be an explicit nil

### UnsetCouponId
`func (o *CreateInvoiceRequest) UnsetCouponId()`

UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
### GetCustomFields

`func (o *CreateInvoiceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateInvoiceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateInvoiceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateInvoiceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CreateInvoiceRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CreateInvoiceRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetSelectedProductPriceQuantity

`func (o *CreateInvoiceRequest) GetSelectedProductPriceQuantity() []SelectedPriceQuantity`

GetSelectedProductPriceQuantity returns the SelectedProductPriceQuantity field if non-nil, zero value otherwise.

### GetSelectedProductPriceQuantityOk

`func (o *CreateInvoiceRequest) GetSelectedProductPriceQuantityOk() (*[]SelectedPriceQuantity, bool)`

GetSelectedProductPriceQuantityOk returns a tuple with the SelectedProductPriceQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedProductPriceQuantity

`func (o *CreateInvoiceRequest) SetSelectedProductPriceQuantity(v []SelectedPriceQuantity)`

SetSelectedProductPriceQuantity sets SelectedProductPriceQuantity field to given value.

### HasSelectedProductPriceQuantity

`func (o *CreateInvoiceRequest) HasSelectedProductPriceQuantity() bool`

HasSelectedProductPriceQuantity returns a boolean if a field has been set.

### GetInvoiceType

`func (o *CreateInvoiceRequest) GetInvoiceType() InvoiceType`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *CreateInvoiceRequest) GetInvoiceTypeOk() (*InvoiceType, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *CreateInvoiceRequest) SetInvoiceType(v InvoiceType)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *CreateInvoiceRequest) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetCustomerId

`func (o *CreateInvoiceRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateInvoiceRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateInvoiceRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetNetD

`func (o *CreateInvoiceRequest) GetNetD() int32`

GetNetD returns the NetD field if non-nil, zero value otherwise.

### GetNetDOk

`func (o *CreateInvoiceRequest) GetNetDOk() (*int32, bool)`

GetNetDOk returns a tuple with the NetD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetD

`func (o *CreateInvoiceRequest) SetNetD(v int32)`

SetNetD sets NetD field to given value.

### HasNetD

`func (o *CreateInvoiceRequest) HasNetD() bool`

HasNetD returns a boolean if a field has been set.

### SetNetDNil

`func (o *CreateInvoiceRequest) SetNetDNil(b bool)`

 SetNetDNil sets the value for NetD to be an explicit nil

### UnsetNetD
`func (o *CreateInvoiceRequest) UnsetNetD()`

UnsetNetD ensures that no value is present for NetD, not even an explicit nil
### GetEmailInvoiceOnFinalization

`func (o *CreateInvoiceRequest) GetEmailInvoiceOnFinalization() bool`

GetEmailInvoiceOnFinalization returns the EmailInvoiceOnFinalization field if non-nil, zero value otherwise.

### GetEmailInvoiceOnFinalizationOk

`func (o *CreateInvoiceRequest) GetEmailInvoiceOnFinalizationOk() (*bool, bool)`

GetEmailInvoiceOnFinalizationOk returns a tuple with the EmailInvoiceOnFinalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailInvoiceOnFinalization

`func (o *CreateInvoiceRequest) SetEmailInvoiceOnFinalization(v bool)`

SetEmailInvoiceOnFinalization sets EmailInvoiceOnFinalization field to given value.

### HasEmailInvoiceOnFinalization

`func (o *CreateInvoiceRequest) HasEmailInvoiceOnFinalization() bool`

HasEmailInvoiceOnFinalization returns a boolean if a field has been set.

### SetEmailInvoiceOnFinalizationNil

`func (o *CreateInvoiceRequest) SetEmailInvoiceOnFinalizationNil(b bool)`

 SetEmailInvoiceOnFinalizationNil sets the value for EmailInvoiceOnFinalization to be an explicit nil

### UnsetEmailInvoiceOnFinalization
`func (o *CreateInvoiceRequest) UnsetEmailInvoiceOnFinalization()`

UnsetEmailInvoiceOnFinalization ensures that no value is present for EmailInvoiceOnFinalization, not even an explicit nil
### GetFinalizeInvoiceImmediately

`func (o *CreateInvoiceRequest) GetFinalizeInvoiceImmediately() bool`

GetFinalizeInvoiceImmediately returns the FinalizeInvoiceImmediately field if non-nil, zero value otherwise.

### GetFinalizeInvoiceImmediatelyOk

`func (o *CreateInvoiceRequest) GetFinalizeInvoiceImmediatelyOk() (*bool, bool)`

GetFinalizeInvoiceImmediatelyOk returns a tuple with the FinalizeInvoiceImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizeInvoiceImmediately

`func (o *CreateInvoiceRequest) SetFinalizeInvoiceImmediately(v bool)`

SetFinalizeInvoiceImmediately sets FinalizeInvoiceImmediately field to given value.

### HasFinalizeInvoiceImmediately

`func (o *CreateInvoiceRequest) HasFinalizeInvoiceImmediately() bool`

HasFinalizeInvoiceImmediately returns a boolean if a field has been set.

### SetFinalizeInvoiceImmediatelyNil

`func (o *CreateInvoiceRequest) SetFinalizeInvoiceImmediatelyNil(b bool)`

 SetFinalizeInvoiceImmediatelyNil sets the value for FinalizeInvoiceImmediately to be an explicit nil

### UnsetFinalizeInvoiceImmediately
`func (o *CreateInvoiceRequest) UnsetFinalizeInvoiceImmediately()`

UnsetFinalizeInvoiceImmediately ensures that no value is present for FinalizeInvoiceImmediately, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


