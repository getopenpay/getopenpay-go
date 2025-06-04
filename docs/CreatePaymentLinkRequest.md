# CreatePaymentLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysCreateNewCustomer** | Pointer to **bool** | If this flag is set to True and there is no customer attached to the payment link, the payment link will always create a new customer, rather than the default behavior of de-duping by email. | [optional] [default to false]
**CheckoutPreferences** | Pointer to [**NullableCheckoutPreferences**](CheckoutPreferences.md) |  | [optional] 
**CouponId** | Pointer to **NullableString** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**LineItems** | Pointer to [**[]CreateCheckoutLineItem**](CreateCheckoutLineItem.md) | The line items to be purchased by the customer. If empty, the checkout sessions produced by the payment link will be in \&quot;add payment method\&quot; mode. | [optional] 
**Mode** | [**CheckoutMode**](CheckoutMode.md) | The mode of the checkout sessions created by this payment link. Possible values: payment (one-time payments), setup (not supported yet), subscription (recurring payments). | 
**SuccessUrl** | **NullableString** |  | 
**TrialEnd** | Pointer to **NullableTime** |  | [optional] 
**TrialFromPrice** | Pointer to **NullableBool** |  | [optional] 
**TrialPeriodDays** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreatePaymentLinkRequest

`func NewCreatePaymentLinkRequest(mode CheckoutMode, successUrl NullableString, ) *CreatePaymentLinkRequest`

NewCreatePaymentLinkRequest instantiates a new CreatePaymentLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentLinkRequestWithDefaults

`func NewCreatePaymentLinkRequestWithDefaults() *CreatePaymentLinkRequest`

NewCreatePaymentLinkRequestWithDefaults instantiates a new CreatePaymentLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysCreateNewCustomer

`func (o *CreatePaymentLinkRequest) GetAlwaysCreateNewCustomer() bool`

GetAlwaysCreateNewCustomer returns the AlwaysCreateNewCustomer field if non-nil, zero value otherwise.

### GetAlwaysCreateNewCustomerOk

`func (o *CreatePaymentLinkRequest) GetAlwaysCreateNewCustomerOk() (*bool, bool)`

GetAlwaysCreateNewCustomerOk returns a tuple with the AlwaysCreateNewCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysCreateNewCustomer

`func (o *CreatePaymentLinkRequest) SetAlwaysCreateNewCustomer(v bool)`

SetAlwaysCreateNewCustomer sets AlwaysCreateNewCustomer field to given value.

### HasAlwaysCreateNewCustomer

`func (o *CreatePaymentLinkRequest) HasAlwaysCreateNewCustomer() bool`

HasAlwaysCreateNewCustomer returns a boolean if a field has been set.

### GetCheckoutPreferences

`func (o *CreatePaymentLinkRequest) GetCheckoutPreferences() CheckoutPreferences`

GetCheckoutPreferences returns the CheckoutPreferences field if non-nil, zero value otherwise.

### GetCheckoutPreferencesOk

`func (o *CreatePaymentLinkRequest) GetCheckoutPreferencesOk() (*CheckoutPreferences, bool)`

GetCheckoutPreferencesOk returns a tuple with the CheckoutPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutPreferences

`func (o *CreatePaymentLinkRequest) SetCheckoutPreferences(v CheckoutPreferences)`

SetCheckoutPreferences sets CheckoutPreferences field to given value.

### HasCheckoutPreferences

`func (o *CreatePaymentLinkRequest) HasCheckoutPreferences() bool`

HasCheckoutPreferences returns a boolean if a field has been set.

### SetCheckoutPreferencesNil

`func (o *CreatePaymentLinkRequest) SetCheckoutPreferencesNil(b bool)`

 SetCheckoutPreferencesNil sets the value for CheckoutPreferences to be an explicit nil

### UnsetCheckoutPreferences
`func (o *CreatePaymentLinkRequest) UnsetCheckoutPreferences()`

UnsetCheckoutPreferences ensures that no value is present for CheckoutPreferences, not even an explicit nil
### GetCouponId

`func (o *CreatePaymentLinkRequest) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *CreatePaymentLinkRequest) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *CreatePaymentLinkRequest) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *CreatePaymentLinkRequest) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponIdNil

`func (o *CreatePaymentLinkRequest) SetCouponIdNil(b bool)`

 SetCouponIdNil sets the value for CouponId to be an explicit nil

### UnsetCouponId
`func (o *CreatePaymentLinkRequest) UnsetCouponId()`

UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
### GetCustomFields

`func (o *CreatePaymentLinkRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreatePaymentLinkRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreatePaymentLinkRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreatePaymentLinkRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CreatePaymentLinkRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CreatePaymentLinkRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetCustomerId

`func (o *CreatePaymentLinkRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreatePaymentLinkRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreatePaymentLinkRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CreatePaymentLinkRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *CreatePaymentLinkRequest) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *CreatePaymentLinkRequest) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetLineItems

`func (o *CreatePaymentLinkRequest) GetLineItems() []CreateCheckoutLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CreatePaymentLinkRequest) GetLineItemsOk() (*[]CreateCheckoutLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CreatePaymentLinkRequest) SetLineItems(v []CreateCheckoutLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *CreatePaymentLinkRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMode

`func (o *CreatePaymentLinkRequest) GetMode() CheckoutMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreatePaymentLinkRequest) GetModeOk() (*CheckoutMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreatePaymentLinkRequest) SetMode(v CheckoutMode)`

SetMode sets Mode field to given value.


### GetSuccessUrl

`func (o *CreatePaymentLinkRequest) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CreatePaymentLinkRequest) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CreatePaymentLinkRequest) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.


### SetSuccessUrlNil

`func (o *CreatePaymentLinkRequest) SetSuccessUrlNil(b bool)`

 SetSuccessUrlNil sets the value for SuccessUrl to be an explicit nil

### UnsetSuccessUrl
`func (o *CreatePaymentLinkRequest) UnsetSuccessUrl()`

UnsetSuccessUrl ensures that no value is present for SuccessUrl, not even an explicit nil
### GetTrialEnd

`func (o *CreatePaymentLinkRequest) GetTrialEnd() time.Time`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *CreatePaymentLinkRequest) GetTrialEndOk() (*time.Time, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *CreatePaymentLinkRequest) SetTrialEnd(v time.Time)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *CreatePaymentLinkRequest) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### SetTrialEndNil

`func (o *CreatePaymentLinkRequest) SetTrialEndNil(b bool)`

 SetTrialEndNil sets the value for TrialEnd to be an explicit nil

### UnsetTrialEnd
`func (o *CreatePaymentLinkRequest) UnsetTrialEnd()`

UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
### GetTrialFromPrice

`func (o *CreatePaymentLinkRequest) GetTrialFromPrice() bool`

GetTrialFromPrice returns the TrialFromPrice field if non-nil, zero value otherwise.

### GetTrialFromPriceOk

`func (o *CreatePaymentLinkRequest) GetTrialFromPriceOk() (*bool, bool)`

GetTrialFromPriceOk returns a tuple with the TrialFromPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialFromPrice

`func (o *CreatePaymentLinkRequest) SetTrialFromPrice(v bool)`

SetTrialFromPrice sets TrialFromPrice field to given value.

### HasTrialFromPrice

`func (o *CreatePaymentLinkRequest) HasTrialFromPrice() bool`

HasTrialFromPrice returns a boolean if a field has been set.

### SetTrialFromPriceNil

`func (o *CreatePaymentLinkRequest) SetTrialFromPriceNil(b bool)`

 SetTrialFromPriceNil sets the value for TrialFromPrice to be an explicit nil

### UnsetTrialFromPrice
`func (o *CreatePaymentLinkRequest) UnsetTrialFromPrice()`

UnsetTrialFromPrice ensures that no value is present for TrialFromPrice, not even an explicit nil
### GetTrialPeriodDays

`func (o *CreatePaymentLinkRequest) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *CreatePaymentLinkRequest) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *CreatePaymentLinkRequest) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *CreatePaymentLinkRequest) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### SetTrialPeriodDaysNil

`func (o *CreatePaymentLinkRequest) SetTrialPeriodDaysNil(b bool)`

 SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil

### UnsetTrialPeriodDays
`func (o *CreatePaymentLinkRequest) UnsetTrialPeriodDays()`

UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


