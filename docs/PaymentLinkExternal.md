# PaymentLinkExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique Identifier of the account. | 
**AccountName** | **string** | Name of the account. | 
**Active** | **bool** | Whether a payment link is active or not. | 
**AlwaysCreateNewCustomer** | Pointer to **bool** | If this flag is set to True and there is no customer attached to the payment link, the payment link will always create a new customer, rather than the default behavior of de-duping by email. | [optional] [default to false]
**CheckoutPreferences** | [**NullableCheckoutPreferencesOutput**](CheckoutPreferencesOutput.md) |  | 
**CouponId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**Currency** | **NullableString** |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomerEmail** | **NullableString** |  | 
**CustomerId** | **NullableString** |  | 
**Id** | **string** | Unique Identifier of the checkout session. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**LineItems** | [**[]PaymentLinkLineItemExternal**](PaymentLinkLineItemExternal.md) | The line items purchased by the customers. | 
**Mode** | [**CheckoutMode**](CheckoutMode.md) | The mode of the checkout sessions created by this payment link. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**SecureToken** | **string** | The random secure token associated with the payment link. | 
**SuccessUrl** | **NullableString** |  | 
**TrialEnd** | **NullableTime** |  | 
**TrialFromPrice** | **NullableBool** |  | 
**TrialPeriodDays** | **NullableInt32** |  | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**Url** | **string** | The main URL for this payment link. | 

## Methods

### NewPaymentLinkExternal

`func NewPaymentLinkExternal(accountId string, accountName string, active bool, checkoutPreferences NullableCheckoutPreferencesOutput, createdAt time.Time, currency NullableString, customerEmail NullableString, customerId NullableString, id string, lineItems []PaymentLinkLineItemExternal, mode CheckoutMode, secureToken string, successUrl NullableString, trialEnd NullableTime, trialFromPrice NullableBool, trialPeriodDays NullableInt32, updatedAt time.Time, url string, ) *PaymentLinkExternal`

NewPaymentLinkExternal instantiates a new PaymentLinkExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentLinkExternalWithDefaults

`func NewPaymentLinkExternalWithDefaults() *PaymentLinkExternal`

NewPaymentLinkExternalWithDefaults instantiates a new PaymentLinkExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PaymentLinkExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PaymentLinkExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PaymentLinkExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccountName

`func (o *PaymentLinkExternal) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PaymentLinkExternal) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PaymentLinkExternal) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetActive

`func (o *PaymentLinkExternal) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PaymentLinkExternal) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PaymentLinkExternal) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAlwaysCreateNewCustomer

`func (o *PaymentLinkExternal) GetAlwaysCreateNewCustomer() bool`

GetAlwaysCreateNewCustomer returns the AlwaysCreateNewCustomer field if non-nil, zero value otherwise.

### GetAlwaysCreateNewCustomerOk

`func (o *PaymentLinkExternal) GetAlwaysCreateNewCustomerOk() (*bool, bool)`

GetAlwaysCreateNewCustomerOk returns a tuple with the AlwaysCreateNewCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysCreateNewCustomer

`func (o *PaymentLinkExternal) SetAlwaysCreateNewCustomer(v bool)`

SetAlwaysCreateNewCustomer sets AlwaysCreateNewCustomer field to given value.

### HasAlwaysCreateNewCustomer

`func (o *PaymentLinkExternal) HasAlwaysCreateNewCustomer() bool`

HasAlwaysCreateNewCustomer returns a boolean if a field has been set.

### GetCheckoutPreferences

`func (o *PaymentLinkExternal) GetCheckoutPreferences() CheckoutPreferencesOutput`

GetCheckoutPreferences returns the CheckoutPreferences field if non-nil, zero value otherwise.

### GetCheckoutPreferencesOk

`func (o *PaymentLinkExternal) GetCheckoutPreferencesOk() (*CheckoutPreferencesOutput, bool)`

GetCheckoutPreferencesOk returns a tuple with the CheckoutPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutPreferences

`func (o *PaymentLinkExternal) SetCheckoutPreferences(v CheckoutPreferencesOutput)`

SetCheckoutPreferences sets CheckoutPreferences field to given value.


### SetCheckoutPreferencesNil

`func (o *PaymentLinkExternal) SetCheckoutPreferencesNil(b bool)`

 SetCheckoutPreferencesNil sets the value for CheckoutPreferences to be an explicit nil

### UnsetCheckoutPreferences
`func (o *PaymentLinkExternal) UnsetCheckoutPreferences()`

UnsetCheckoutPreferences ensures that no value is present for CheckoutPreferences, not even an explicit nil
### GetCouponId

`func (o *PaymentLinkExternal) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *PaymentLinkExternal) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *PaymentLinkExternal) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *PaymentLinkExternal) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponIdNil

`func (o *PaymentLinkExternal) SetCouponIdNil(b bool)`

 SetCouponIdNil sets the value for CouponId to be an explicit nil

### UnsetCouponId
`func (o *PaymentLinkExternal) UnsetCouponId()`

UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
### GetCreatedAt

`func (o *PaymentLinkExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentLinkExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentLinkExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *PaymentLinkExternal) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentLinkExternal) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentLinkExternal) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *PaymentLinkExternal) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PaymentLinkExternal) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCustomFields

`func (o *PaymentLinkExternal) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PaymentLinkExternal) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PaymentLinkExternal) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PaymentLinkExternal) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *PaymentLinkExternal) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *PaymentLinkExternal) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetCustomerEmail

`func (o *PaymentLinkExternal) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *PaymentLinkExternal) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *PaymentLinkExternal) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.


### SetCustomerEmailNil

`func (o *PaymentLinkExternal) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *PaymentLinkExternal) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetCustomerId

`func (o *PaymentLinkExternal) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PaymentLinkExternal) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PaymentLinkExternal) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *PaymentLinkExternal) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *PaymentLinkExternal) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetId

`func (o *PaymentLinkExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentLinkExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentLinkExternal) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *PaymentLinkExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PaymentLinkExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PaymentLinkExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PaymentLinkExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLineItems

`func (o *PaymentLinkExternal) GetLineItems() []PaymentLinkLineItemExternal`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *PaymentLinkExternal) GetLineItemsOk() (*[]PaymentLinkLineItemExternal, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *PaymentLinkExternal) SetLineItems(v []PaymentLinkLineItemExternal)`

SetLineItems sets LineItems field to given value.


### GetMode

`func (o *PaymentLinkExternal) GetMode() CheckoutMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PaymentLinkExternal) GetModeOk() (*CheckoutMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PaymentLinkExternal) SetMode(v CheckoutMode)`

SetMode sets Mode field to given value.


### GetObject

`func (o *PaymentLinkExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentLinkExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentLinkExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *PaymentLinkExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetSecureToken

`func (o *PaymentLinkExternal) GetSecureToken() string`

GetSecureToken returns the SecureToken field if non-nil, zero value otherwise.

### GetSecureTokenOk

`func (o *PaymentLinkExternal) GetSecureTokenOk() (*string, bool)`

GetSecureTokenOk returns a tuple with the SecureToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureToken

`func (o *PaymentLinkExternal) SetSecureToken(v string)`

SetSecureToken sets SecureToken field to given value.


### GetSuccessUrl

`func (o *PaymentLinkExternal) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *PaymentLinkExternal) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *PaymentLinkExternal) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.


### SetSuccessUrlNil

`func (o *PaymentLinkExternal) SetSuccessUrlNil(b bool)`

 SetSuccessUrlNil sets the value for SuccessUrl to be an explicit nil

### UnsetSuccessUrl
`func (o *PaymentLinkExternal) UnsetSuccessUrl()`

UnsetSuccessUrl ensures that no value is present for SuccessUrl, not even an explicit nil
### GetTrialEnd

`func (o *PaymentLinkExternal) GetTrialEnd() time.Time`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *PaymentLinkExternal) GetTrialEndOk() (*time.Time, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *PaymentLinkExternal) SetTrialEnd(v time.Time)`

SetTrialEnd sets TrialEnd field to given value.


### SetTrialEndNil

`func (o *PaymentLinkExternal) SetTrialEndNil(b bool)`

 SetTrialEndNil sets the value for TrialEnd to be an explicit nil

### UnsetTrialEnd
`func (o *PaymentLinkExternal) UnsetTrialEnd()`

UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
### GetTrialFromPrice

`func (o *PaymentLinkExternal) GetTrialFromPrice() bool`

GetTrialFromPrice returns the TrialFromPrice field if non-nil, zero value otherwise.

### GetTrialFromPriceOk

`func (o *PaymentLinkExternal) GetTrialFromPriceOk() (*bool, bool)`

GetTrialFromPriceOk returns a tuple with the TrialFromPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialFromPrice

`func (o *PaymentLinkExternal) SetTrialFromPrice(v bool)`

SetTrialFromPrice sets TrialFromPrice field to given value.


### SetTrialFromPriceNil

`func (o *PaymentLinkExternal) SetTrialFromPriceNil(b bool)`

 SetTrialFromPriceNil sets the value for TrialFromPrice to be an explicit nil

### UnsetTrialFromPrice
`func (o *PaymentLinkExternal) UnsetTrialFromPrice()`

UnsetTrialFromPrice ensures that no value is present for TrialFromPrice, not even an explicit nil
### GetTrialPeriodDays

`func (o *PaymentLinkExternal) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *PaymentLinkExternal) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *PaymentLinkExternal) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.


### SetTrialPeriodDaysNil

`func (o *PaymentLinkExternal) SetTrialPeriodDaysNil(b bool)`

 SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil

### UnsetTrialPeriodDays
`func (o *PaymentLinkExternal) UnsetTrialPeriodDays()`

UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil
### GetUpdatedAt

`func (o *PaymentLinkExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentLinkExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentLinkExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *PaymentLinkExternal) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PaymentLinkExternal) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PaymentLinkExternal) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


