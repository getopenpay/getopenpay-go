# CreateCheckoutSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutPreferences** | Pointer to [**NullableCheckoutPreferences**](CheckoutPreferences.md) |  | [optional] 
**ClientReferenceId** | Pointer to **NullableString** |  | [optional] 
**CouponId** | Pointer to **NullableString** |  | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomerEmail** | Pointer to **NullableString** |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**Language** | Pointer to [**NullableCustomerLanguage**](CustomerLanguage.md) |  | [optional] 
**LineItems** | Pointer to [**[]CreateCheckoutLineItem**](CreateCheckoutLineItem.md) | The line items purchased by the customer. | [optional] 
**Mode** | [**CheckoutMode**](CheckoutMode.md) | The mode of the Checkout Session. Possible values: payment (one-time payments), setup (collect payment info), subscription (recurring payments). | 
**ReturnUrl** | Pointer to **NullableString** |  | [optional] 
**SuccessUrl** | Pointer to **NullableString** |  | [optional] 
**TrialEnd** | Pointer to **NullableTime** |  | [optional] 
**TrialFromPrice** | Pointer to **NullableBool** |  | [optional] 
**TrialPeriodDays** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateCheckoutSessionRequest

`func NewCreateCheckoutSessionRequest(mode CheckoutMode, ) *CreateCheckoutSessionRequest`

NewCreateCheckoutSessionRequest instantiates a new CreateCheckoutSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCheckoutSessionRequestWithDefaults

`func NewCreateCheckoutSessionRequestWithDefaults() *CreateCheckoutSessionRequest`

NewCreateCheckoutSessionRequestWithDefaults instantiates a new CreateCheckoutSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutPreferences

`func (o *CreateCheckoutSessionRequest) GetCheckoutPreferences() CheckoutPreferences`

GetCheckoutPreferences returns the CheckoutPreferences field if non-nil, zero value otherwise.

### GetCheckoutPreferencesOk

`func (o *CreateCheckoutSessionRequest) GetCheckoutPreferencesOk() (*CheckoutPreferences, bool)`

GetCheckoutPreferencesOk returns a tuple with the CheckoutPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutPreferences

`func (o *CreateCheckoutSessionRequest) SetCheckoutPreferences(v CheckoutPreferences)`

SetCheckoutPreferences sets CheckoutPreferences field to given value.

### HasCheckoutPreferences

`func (o *CreateCheckoutSessionRequest) HasCheckoutPreferences() bool`

HasCheckoutPreferences returns a boolean if a field has been set.

### SetCheckoutPreferencesNil

`func (o *CreateCheckoutSessionRequest) SetCheckoutPreferencesNil(b bool)`

 SetCheckoutPreferencesNil sets the value for CheckoutPreferences to be an explicit nil

### UnsetCheckoutPreferences
`func (o *CreateCheckoutSessionRequest) UnsetCheckoutPreferences()`

UnsetCheckoutPreferences ensures that no value is present for CheckoutPreferences, not even an explicit nil
### GetClientReferenceId

`func (o *CreateCheckoutSessionRequest) GetClientReferenceId() string`

GetClientReferenceId returns the ClientReferenceId field if non-nil, zero value otherwise.

### GetClientReferenceIdOk

`func (o *CreateCheckoutSessionRequest) GetClientReferenceIdOk() (*string, bool)`

GetClientReferenceIdOk returns a tuple with the ClientReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReferenceId

`func (o *CreateCheckoutSessionRequest) SetClientReferenceId(v string)`

SetClientReferenceId sets ClientReferenceId field to given value.

### HasClientReferenceId

`func (o *CreateCheckoutSessionRequest) HasClientReferenceId() bool`

HasClientReferenceId returns a boolean if a field has been set.

### SetClientReferenceIdNil

`func (o *CreateCheckoutSessionRequest) SetClientReferenceIdNil(b bool)`

 SetClientReferenceIdNil sets the value for ClientReferenceId to be an explicit nil

### UnsetClientReferenceId
`func (o *CreateCheckoutSessionRequest) UnsetClientReferenceId()`

UnsetClientReferenceId ensures that no value is present for ClientReferenceId, not even an explicit nil
### GetCouponId

`func (o *CreateCheckoutSessionRequest) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *CreateCheckoutSessionRequest) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *CreateCheckoutSessionRequest) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *CreateCheckoutSessionRequest) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponIdNil

`func (o *CreateCheckoutSessionRequest) SetCouponIdNil(b bool)`

 SetCouponIdNil sets the value for CouponId to be an explicit nil

### UnsetCouponId
`func (o *CreateCheckoutSessionRequest) UnsetCouponId()`

UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
### GetCurrency

`func (o *CreateCheckoutSessionRequest) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCheckoutSessionRequest) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCheckoutSessionRequest) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateCheckoutSessionRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CreateCheckoutSessionRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CreateCheckoutSessionRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCustomFields

`func (o *CreateCheckoutSessionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateCheckoutSessionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateCheckoutSessionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateCheckoutSessionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CreateCheckoutSessionRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CreateCheckoutSessionRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetCustomerEmail

`func (o *CreateCheckoutSessionRequest) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *CreateCheckoutSessionRequest) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *CreateCheckoutSessionRequest) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *CreateCheckoutSessionRequest) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### SetCustomerEmailNil

`func (o *CreateCheckoutSessionRequest) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *CreateCheckoutSessionRequest) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetCustomerId

`func (o *CreateCheckoutSessionRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateCheckoutSessionRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateCheckoutSessionRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CreateCheckoutSessionRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *CreateCheckoutSessionRequest) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *CreateCheckoutSessionRequest) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetLanguage

`func (o *CreateCheckoutSessionRequest) GetLanguage() CustomerLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CreateCheckoutSessionRequest) GetLanguageOk() (*CustomerLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CreateCheckoutSessionRequest) SetLanguage(v CustomerLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CreateCheckoutSessionRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *CreateCheckoutSessionRequest) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *CreateCheckoutSessionRequest) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetLineItems

`func (o *CreateCheckoutSessionRequest) GetLineItems() []CreateCheckoutLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CreateCheckoutSessionRequest) GetLineItemsOk() (*[]CreateCheckoutLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CreateCheckoutSessionRequest) SetLineItems(v []CreateCheckoutLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *CreateCheckoutSessionRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMode

`func (o *CreateCheckoutSessionRequest) GetMode() CheckoutMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateCheckoutSessionRequest) GetModeOk() (*CheckoutMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateCheckoutSessionRequest) SetMode(v CheckoutMode)`

SetMode sets Mode field to given value.


### GetReturnUrl

`func (o *CreateCheckoutSessionRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CreateCheckoutSessionRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CreateCheckoutSessionRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *CreateCheckoutSessionRequest) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *CreateCheckoutSessionRequest) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *CreateCheckoutSessionRequest) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetSuccessUrl

`func (o *CreateCheckoutSessionRequest) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CreateCheckoutSessionRequest) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CreateCheckoutSessionRequest) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *CreateCheckoutSessionRequest) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### SetSuccessUrlNil

`func (o *CreateCheckoutSessionRequest) SetSuccessUrlNil(b bool)`

 SetSuccessUrlNil sets the value for SuccessUrl to be an explicit nil

### UnsetSuccessUrl
`func (o *CreateCheckoutSessionRequest) UnsetSuccessUrl()`

UnsetSuccessUrl ensures that no value is present for SuccessUrl, not even an explicit nil
### GetTrialEnd

`func (o *CreateCheckoutSessionRequest) GetTrialEnd() time.Time`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *CreateCheckoutSessionRequest) GetTrialEndOk() (*time.Time, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *CreateCheckoutSessionRequest) SetTrialEnd(v time.Time)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *CreateCheckoutSessionRequest) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### SetTrialEndNil

`func (o *CreateCheckoutSessionRequest) SetTrialEndNil(b bool)`

 SetTrialEndNil sets the value for TrialEnd to be an explicit nil

### UnsetTrialEnd
`func (o *CreateCheckoutSessionRequest) UnsetTrialEnd()`

UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
### GetTrialFromPrice

`func (o *CreateCheckoutSessionRequest) GetTrialFromPrice() bool`

GetTrialFromPrice returns the TrialFromPrice field if non-nil, zero value otherwise.

### GetTrialFromPriceOk

`func (o *CreateCheckoutSessionRequest) GetTrialFromPriceOk() (*bool, bool)`

GetTrialFromPriceOk returns a tuple with the TrialFromPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialFromPrice

`func (o *CreateCheckoutSessionRequest) SetTrialFromPrice(v bool)`

SetTrialFromPrice sets TrialFromPrice field to given value.

### HasTrialFromPrice

`func (o *CreateCheckoutSessionRequest) HasTrialFromPrice() bool`

HasTrialFromPrice returns a boolean if a field has been set.

### SetTrialFromPriceNil

`func (o *CreateCheckoutSessionRequest) SetTrialFromPriceNil(b bool)`

 SetTrialFromPriceNil sets the value for TrialFromPrice to be an explicit nil

### UnsetTrialFromPrice
`func (o *CreateCheckoutSessionRequest) UnsetTrialFromPrice()`

UnsetTrialFromPrice ensures that no value is present for TrialFromPrice, not even an explicit nil
### GetTrialPeriodDays

`func (o *CreateCheckoutSessionRequest) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *CreateCheckoutSessionRequest) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *CreateCheckoutSessionRequest) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *CreateCheckoutSessionRequest) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### SetTrialPeriodDaysNil

`func (o *CreateCheckoutSessionRequest) SetTrialPeriodDaysNil(b bool)`

 SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil

### UnsetTrialPeriodDays
`func (o *CreateCheckoutSessionRequest) UnsetTrialPeriodDays()`

UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


