# CheckoutSessionExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique Identifier of the account. | 
**AccountName** | **string** | Name of the account. | 
**AmountSubtotalAtom** | **int32** | The integer amount representing the subtotal amount for the line items. | 
**AmountTotalAtom** | **int32** | The integer amount representing the total amount for the line items, after discounts and taxes. | 
**ClientReferenceId** | Pointer to **NullableString** |  | [optional] 
**CouponId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**Currency** | **NullableString** |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomerEmail** | **NullableString** |  | 
**CustomerId** | **NullableString** |  | 
**Id** | **string** | Unique Identifier of the checkout session. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Language** | Pointer to [**NullableCustomerLanguage**](CustomerLanguage.md) |  | [optional] 
**LineItems** | [**[]CheckoutSessionLineItemExternal**](CheckoutSessionLineItemExternal.md) | The line items purchased by the customers. | 
**Mode** | [**CheckoutMode**](CheckoutMode.md) | The mode of the Checkout Session. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**ReturnUrl** | **NullableString** |  | 
**SecureToken** | **string** | The random secure token associated with the checkout session. | 
**SetupIntent** | [**NullableSetupIntentExternal**](SetupIntentExternal.md) |  | 
**Status** | [**CheckoutSessionStatus**](CheckoutSessionStatus.md) | The current status of the checkout session. | 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionIds** | **[]string** |  | 
**SuccessUrl** | **NullableString** |  | 
**TaxAmountAtom** | **int32** | The integer amount representing the tax amount for the line items. | 
**TrialEnd** | **NullableTime** |  | 
**TrialFromPrice** | **NullableBool** |  | 
**TrialPeriodDays** | **NullableInt32** |  | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**Url** | **string** | The main URL for this checkout session. | 

## Methods

### NewCheckoutSessionExternal

`func NewCheckoutSessionExternal(accountId string, accountName string, amountSubtotalAtom int32, amountTotalAtom int32, createdAt time.Time, currency NullableString, customerEmail NullableString, customerId NullableString, id string, lineItems []CheckoutSessionLineItemExternal, mode CheckoutMode, returnUrl NullableString, secureToken string, setupIntent NullableSetupIntentExternal, status CheckoutSessionStatus, subscriptionIds []string, successUrl NullableString, taxAmountAtom int32, trialEnd NullableTime, trialFromPrice NullableBool, trialPeriodDays NullableInt32, updatedAt time.Time, url string, ) *CheckoutSessionExternal`

NewCheckoutSessionExternal instantiates a new CheckoutSessionExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionExternalWithDefaults

`func NewCheckoutSessionExternalWithDefaults() *CheckoutSessionExternal`

NewCheckoutSessionExternalWithDefaults instantiates a new CheckoutSessionExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CheckoutSessionExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CheckoutSessionExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CheckoutSessionExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccountName

`func (o *CheckoutSessionExternal) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *CheckoutSessionExternal) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *CheckoutSessionExternal) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetAmountSubtotalAtom

`func (o *CheckoutSessionExternal) GetAmountSubtotalAtom() int32`

GetAmountSubtotalAtom returns the AmountSubtotalAtom field if non-nil, zero value otherwise.

### GetAmountSubtotalAtomOk

`func (o *CheckoutSessionExternal) GetAmountSubtotalAtomOk() (*int32, bool)`

GetAmountSubtotalAtomOk returns a tuple with the AmountSubtotalAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSubtotalAtom

`func (o *CheckoutSessionExternal) SetAmountSubtotalAtom(v int32)`

SetAmountSubtotalAtom sets AmountSubtotalAtom field to given value.


### GetAmountTotalAtom

`func (o *CheckoutSessionExternal) GetAmountTotalAtom() int32`

GetAmountTotalAtom returns the AmountTotalAtom field if non-nil, zero value otherwise.

### GetAmountTotalAtomOk

`func (o *CheckoutSessionExternal) GetAmountTotalAtomOk() (*int32, bool)`

GetAmountTotalAtomOk returns a tuple with the AmountTotalAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountTotalAtom

`func (o *CheckoutSessionExternal) SetAmountTotalAtom(v int32)`

SetAmountTotalAtom sets AmountTotalAtom field to given value.


### GetClientReferenceId

`func (o *CheckoutSessionExternal) GetClientReferenceId() string`

GetClientReferenceId returns the ClientReferenceId field if non-nil, zero value otherwise.

### GetClientReferenceIdOk

`func (o *CheckoutSessionExternal) GetClientReferenceIdOk() (*string, bool)`

GetClientReferenceIdOk returns a tuple with the ClientReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReferenceId

`func (o *CheckoutSessionExternal) SetClientReferenceId(v string)`

SetClientReferenceId sets ClientReferenceId field to given value.

### HasClientReferenceId

`func (o *CheckoutSessionExternal) HasClientReferenceId() bool`

HasClientReferenceId returns a boolean if a field has been set.

### SetClientReferenceIdNil

`func (o *CheckoutSessionExternal) SetClientReferenceIdNil(b bool)`

 SetClientReferenceIdNil sets the value for ClientReferenceId to be an explicit nil

### UnsetClientReferenceId
`func (o *CheckoutSessionExternal) UnsetClientReferenceId()`

UnsetClientReferenceId ensures that no value is present for ClientReferenceId, not even an explicit nil
### GetCouponId

`func (o *CheckoutSessionExternal) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *CheckoutSessionExternal) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *CheckoutSessionExternal) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *CheckoutSessionExternal) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponIdNil

`func (o *CheckoutSessionExternal) SetCouponIdNil(b bool)`

 SetCouponIdNil sets the value for CouponId to be an explicit nil

### UnsetCouponId
`func (o *CheckoutSessionExternal) UnsetCouponId()`

UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
### GetCreatedAt

`func (o *CheckoutSessionExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckoutSessionExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckoutSessionExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *CheckoutSessionExternal) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CheckoutSessionExternal) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CheckoutSessionExternal) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *CheckoutSessionExternal) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CheckoutSessionExternal) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCustomFields

`func (o *CheckoutSessionExternal) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CheckoutSessionExternal) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CheckoutSessionExternal) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CheckoutSessionExternal) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CheckoutSessionExternal) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CheckoutSessionExternal) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetCustomerEmail

`func (o *CheckoutSessionExternal) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *CheckoutSessionExternal) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *CheckoutSessionExternal) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.


### SetCustomerEmailNil

`func (o *CheckoutSessionExternal) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *CheckoutSessionExternal) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetCustomerId

`func (o *CheckoutSessionExternal) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CheckoutSessionExternal) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CheckoutSessionExternal) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *CheckoutSessionExternal) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *CheckoutSessionExternal) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetId

`func (o *CheckoutSessionExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutSessionExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutSessionExternal) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *CheckoutSessionExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CheckoutSessionExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CheckoutSessionExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CheckoutSessionExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLanguage

`func (o *CheckoutSessionExternal) GetLanguage() CustomerLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CheckoutSessionExternal) GetLanguageOk() (*CustomerLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CheckoutSessionExternal) SetLanguage(v CustomerLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CheckoutSessionExternal) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *CheckoutSessionExternal) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *CheckoutSessionExternal) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetLineItems

`func (o *CheckoutSessionExternal) GetLineItems() []CheckoutSessionLineItemExternal`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CheckoutSessionExternal) GetLineItemsOk() (*[]CheckoutSessionLineItemExternal, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CheckoutSessionExternal) SetLineItems(v []CheckoutSessionLineItemExternal)`

SetLineItems sets LineItems field to given value.


### GetMode

`func (o *CheckoutSessionExternal) GetMode() CheckoutMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CheckoutSessionExternal) GetModeOk() (*CheckoutMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CheckoutSessionExternal) SetMode(v CheckoutMode)`

SetMode sets Mode field to given value.


### GetObject

`func (o *CheckoutSessionExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CheckoutSessionExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CheckoutSessionExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *CheckoutSessionExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetReturnUrl

`func (o *CheckoutSessionExternal) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CheckoutSessionExternal) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CheckoutSessionExternal) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### SetReturnUrlNil

`func (o *CheckoutSessionExternal) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *CheckoutSessionExternal) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetSecureToken

`func (o *CheckoutSessionExternal) GetSecureToken() string`

GetSecureToken returns the SecureToken field if non-nil, zero value otherwise.

### GetSecureTokenOk

`func (o *CheckoutSessionExternal) GetSecureTokenOk() (*string, bool)`

GetSecureTokenOk returns a tuple with the SecureToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureToken

`func (o *CheckoutSessionExternal) SetSecureToken(v string)`

SetSecureToken sets SecureToken field to given value.


### GetSetupIntent

`func (o *CheckoutSessionExternal) GetSetupIntent() SetupIntentExternal`

GetSetupIntent returns the SetupIntent field if non-nil, zero value otherwise.

### GetSetupIntentOk

`func (o *CheckoutSessionExternal) GetSetupIntentOk() (*SetupIntentExternal, bool)`

GetSetupIntentOk returns a tuple with the SetupIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupIntent

`func (o *CheckoutSessionExternal) SetSetupIntent(v SetupIntentExternal)`

SetSetupIntent sets SetupIntent field to given value.


### SetSetupIntentNil

`func (o *CheckoutSessionExternal) SetSetupIntentNil(b bool)`

 SetSetupIntentNil sets the value for SetupIntent to be an explicit nil

### UnsetSetupIntent
`func (o *CheckoutSessionExternal) UnsetSetupIntent()`

UnsetSetupIntent ensures that no value is present for SetupIntent, not even an explicit nil
### GetStatus

`func (o *CheckoutSessionExternal) GetStatus() CheckoutSessionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckoutSessionExternal) GetStatusOk() (*CheckoutSessionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckoutSessionExternal) SetStatus(v CheckoutSessionStatus)`

SetStatus sets Status field to given value.


### GetSubscriptionId

`func (o *CheckoutSessionExternal) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CheckoutSessionExternal) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CheckoutSessionExternal) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CheckoutSessionExternal) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *CheckoutSessionExternal) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *CheckoutSessionExternal) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetSubscriptionIds

`func (o *CheckoutSessionExternal) GetSubscriptionIds() []string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *CheckoutSessionExternal) GetSubscriptionIdsOk() (*[]string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *CheckoutSessionExternal) SetSubscriptionIds(v []string)`

SetSubscriptionIds sets SubscriptionIds field to given value.


### SetSubscriptionIdsNil

`func (o *CheckoutSessionExternal) SetSubscriptionIdsNil(b bool)`

 SetSubscriptionIdsNil sets the value for SubscriptionIds to be an explicit nil

### UnsetSubscriptionIds
`func (o *CheckoutSessionExternal) UnsetSubscriptionIds()`

UnsetSubscriptionIds ensures that no value is present for SubscriptionIds, not even an explicit nil
### GetSuccessUrl

`func (o *CheckoutSessionExternal) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CheckoutSessionExternal) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CheckoutSessionExternal) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.


### SetSuccessUrlNil

`func (o *CheckoutSessionExternal) SetSuccessUrlNil(b bool)`

 SetSuccessUrlNil sets the value for SuccessUrl to be an explicit nil

### UnsetSuccessUrl
`func (o *CheckoutSessionExternal) UnsetSuccessUrl()`

UnsetSuccessUrl ensures that no value is present for SuccessUrl, not even an explicit nil
### GetTaxAmountAtom

`func (o *CheckoutSessionExternal) GetTaxAmountAtom() int32`

GetTaxAmountAtom returns the TaxAmountAtom field if non-nil, zero value otherwise.

### GetTaxAmountAtomOk

`func (o *CheckoutSessionExternal) GetTaxAmountAtomOk() (*int32, bool)`

GetTaxAmountAtomOk returns a tuple with the TaxAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountAtom

`func (o *CheckoutSessionExternal) SetTaxAmountAtom(v int32)`

SetTaxAmountAtom sets TaxAmountAtom field to given value.


### GetTrialEnd

`func (o *CheckoutSessionExternal) GetTrialEnd() time.Time`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *CheckoutSessionExternal) GetTrialEndOk() (*time.Time, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *CheckoutSessionExternal) SetTrialEnd(v time.Time)`

SetTrialEnd sets TrialEnd field to given value.


### SetTrialEndNil

`func (o *CheckoutSessionExternal) SetTrialEndNil(b bool)`

 SetTrialEndNil sets the value for TrialEnd to be an explicit nil

### UnsetTrialEnd
`func (o *CheckoutSessionExternal) UnsetTrialEnd()`

UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
### GetTrialFromPrice

`func (o *CheckoutSessionExternal) GetTrialFromPrice() bool`

GetTrialFromPrice returns the TrialFromPrice field if non-nil, zero value otherwise.

### GetTrialFromPriceOk

`func (o *CheckoutSessionExternal) GetTrialFromPriceOk() (*bool, bool)`

GetTrialFromPriceOk returns a tuple with the TrialFromPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialFromPrice

`func (o *CheckoutSessionExternal) SetTrialFromPrice(v bool)`

SetTrialFromPrice sets TrialFromPrice field to given value.


### SetTrialFromPriceNil

`func (o *CheckoutSessionExternal) SetTrialFromPriceNil(b bool)`

 SetTrialFromPriceNil sets the value for TrialFromPrice to be an explicit nil

### UnsetTrialFromPrice
`func (o *CheckoutSessionExternal) UnsetTrialFromPrice()`

UnsetTrialFromPrice ensures that no value is present for TrialFromPrice, not even an explicit nil
### GetTrialPeriodDays

`func (o *CheckoutSessionExternal) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *CheckoutSessionExternal) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *CheckoutSessionExternal) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.


### SetTrialPeriodDaysNil

`func (o *CheckoutSessionExternal) SetTrialPeriodDaysNil(b bool)`

 SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil

### UnsetTrialPeriodDays
`func (o *CheckoutSessionExternal) UnsetTrialPeriodDays()`

UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil
### GetUpdatedAt

`func (o *CheckoutSessionExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CheckoutSessionExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CheckoutSessionExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *CheckoutSessionExternal) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckoutSessionExternal) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckoutSessionExternal) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


