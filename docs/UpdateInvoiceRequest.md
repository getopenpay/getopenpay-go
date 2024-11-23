# UpdateInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** | Description for newly created invoice | [optional] [default to "Manual creation of invoice"]
**CouponId** | Pointer to **NullableString** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**EmailInvoiceOnFinalization** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateInvoiceRequest

`func NewUpdateInvoiceRequest() *UpdateInvoiceRequest`

NewUpdateInvoiceRequest instantiates a new UpdateInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInvoiceRequestWithDefaults

`func NewUpdateInvoiceRequestWithDefaults() *UpdateInvoiceRequest`

NewUpdateInvoiceRequestWithDefaults instantiates a new UpdateInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodId

`func (o *UpdateInvoiceRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *UpdateInvoiceRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *UpdateInvoiceRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *UpdateInvoiceRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### SetPaymentMethodIdNil

`func (o *UpdateInvoiceRequest) SetPaymentMethodIdNil(b bool)`

 SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil

### UnsetPaymentMethodId
`func (o *UpdateInvoiceRequest) UnsetPaymentMethodId()`

UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
### GetDescription

`func (o *UpdateInvoiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateInvoiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateInvoiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateInvoiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCouponId

`func (o *UpdateInvoiceRequest) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *UpdateInvoiceRequest) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *UpdateInvoiceRequest) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *UpdateInvoiceRequest) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponIdNil

`func (o *UpdateInvoiceRequest) SetCouponIdNil(b bool)`

 SetCouponIdNil sets the value for CouponId to be an explicit nil

### UnsetCouponId
`func (o *UpdateInvoiceRequest) UnsetCouponId()`

UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
### GetCustomFields

`func (o *UpdateInvoiceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *UpdateInvoiceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *UpdateInvoiceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *UpdateInvoiceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *UpdateInvoiceRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *UpdateInvoiceRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetEmailInvoiceOnFinalization

`func (o *UpdateInvoiceRequest) GetEmailInvoiceOnFinalization() bool`

GetEmailInvoiceOnFinalization returns the EmailInvoiceOnFinalization field if non-nil, zero value otherwise.

### GetEmailInvoiceOnFinalizationOk

`func (o *UpdateInvoiceRequest) GetEmailInvoiceOnFinalizationOk() (*bool, bool)`

GetEmailInvoiceOnFinalizationOk returns a tuple with the EmailInvoiceOnFinalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailInvoiceOnFinalization

`func (o *UpdateInvoiceRequest) SetEmailInvoiceOnFinalization(v bool)`

SetEmailInvoiceOnFinalization sets EmailInvoiceOnFinalization field to given value.

### HasEmailInvoiceOnFinalization

`func (o *UpdateInvoiceRequest) HasEmailInvoiceOnFinalization() bool`

HasEmailInvoiceOnFinalization returns a boolean if a field has been set.

### SetEmailInvoiceOnFinalizationNil

`func (o *UpdateInvoiceRequest) SetEmailInvoiceOnFinalizationNil(b bool)`

 SetEmailInvoiceOnFinalizationNil sets the value for EmailInvoiceOnFinalization to be an explicit nil

### UnsetEmailInvoiceOnFinalization
`func (o *UpdateInvoiceRequest) UnsetEmailInvoiceOnFinalization()`

UnsetEmailInvoiceOnFinalization ensures that no value is present for EmailInvoiceOnFinalization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


