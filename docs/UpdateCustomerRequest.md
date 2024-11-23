# UpdateCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**Line1** | Pointer to **NullableString** |  | [optional] 
**Line2** | Pointer to **NullableString** |  | [optional] 
**Line3** | Pointer to **NullableString** |  | [optional] 
**City** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to **NullableString** |  | [optional] 
**ZipCode** | Pointer to **NullableString** |  | [optional] 
**CouponId** | Pointer to **NullableString** |  | [optional] 
**PromotionCodeId** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**InvoiceSettings** | Pointer to [**NullableCustomerInvoiceSettings**](CustomerInvoiceSettings.md) |  | [optional] 
**BillingEmail** | Pointer to **NullableString** |  | [optional] 
**Language** | Pointer to [**NullableCustomerLanguage**](CustomerLanguage.md) |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateCustomerRequest

`func NewUpdateCustomerRequest() *UpdateCustomerRequest`

NewUpdateCustomerRequest instantiates a new UpdateCustomerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomerRequestWithDefaults

`func NewUpdateCustomerRequestWithDefaults() *UpdateCustomerRequest`

NewUpdateCustomerRequestWithDefaults instantiates a new UpdateCustomerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UpdateCustomerRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateCustomerRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateCustomerRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateCustomerRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UpdateCustomerRequest) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UpdateCustomerRequest) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UpdateCustomerRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateCustomerRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateCustomerRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateCustomerRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UpdateCustomerRequest) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UpdateCustomerRequest) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetLine1

`func (o *UpdateCustomerRequest) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *UpdateCustomerRequest) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *UpdateCustomerRequest) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *UpdateCustomerRequest) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### SetLine1Nil

`func (o *UpdateCustomerRequest) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *UpdateCustomerRequest) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *UpdateCustomerRequest) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *UpdateCustomerRequest) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *UpdateCustomerRequest) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *UpdateCustomerRequest) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *UpdateCustomerRequest) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *UpdateCustomerRequest) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetLine3

`func (o *UpdateCustomerRequest) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *UpdateCustomerRequest) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *UpdateCustomerRequest) SetLine3(v string)`

SetLine3 sets Line3 field to given value.

### HasLine3

`func (o *UpdateCustomerRequest) HasLine3() bool`

HasLine3 returns a boolean if a field has been set.

### SetLine3Nil

`func (o *UpdateCustomerRequest) SetLine3Nil(b bool)`

 SetLine3Nil sets the value for Line3 to be an explicit nil

### UnsetLine3
`func (o *UpdateCustomerRequest) UnsetLine3()`

UnsetLine3 ensures that no value is present for Line3, not even an explicit nil
### GetCity

`func (o *UpdateCustomerRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UpdateCustomerRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UpdateCustomerRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UpdateCustomerRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *UpdateCustomerRequest) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *UpdateCustomerRequest) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *UpdateCustomerRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateCustomerRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateCustomerRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateCustomerRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *UpdateCustomerRequest) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *UpdateCustomerRequest) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCountry

`func (o *UpdateCustomerRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateCustomerRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateCustomerRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UpdateCustomerRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *UpdateCustomerRequest) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *UpdateCustomerRequest) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetZipCode

`func (o *UpdateCustomerRequest) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UpdateCustomerRequest) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UpdateCustomerRequest) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *UpdateCustomerRequest) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *UpdateCustomerRequest) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *UpdateCustomerRequest) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetCouponId

`func (o *UpdateCustomerRequest) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *UpdateCustomerRequest) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *UpdateCustomerRequest) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *UpdateCustomerRequest) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponIdNil

`func (o *UpdateCustomerRequest) SetCouponIdNil(b bool)`

 SetCouponIdNil sets the value for CouponId to be an explicit nil

### UnsetCouponId
`func (o *UpdateCustomerRequest) UnsetCouponId()`

UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
### GetPromotionCodeId

`func (o *UpdateCustomerRequest) GetPromotionCodeId() string`

GetPromotionCodeId returns the PromotionCodeId field if non-nil, zero value otherwise.

### GetPromotionCodeIdOk

`func (o *UpdateCustomerRequest) GetPromotionCodeIdOk() (*string, bool)`

GetPromotionCodeIdOk returns a tuple with the PromotionCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionCodeId

`func (o *UpdateCustomerRequest) SetPromotionCodeId(v string)`

SetPromotionCodeId sets PromotionCodeId field to given value.

### HasPromotionCodeId

`func (o *UpdateCustomerRequest) HasPromotionCodeId() bool`

HasPromotionCodeId returns a boolean if a field has been set.

### SetPromotionCodeIdNil

`func (o *UpdateCustomerRequest) SetPromotionCodeIdNil(b bool)`

 SetPromotionCodeIdNil sets the value for PromotionCodeId to be an explicit nil

### UnsetPromotionCodeId
`func (o *UpdateCustomerRequest) UnsetPromotionCodeId()`

UnsetPromotionCodeId ensures that no value is present for PromotionCodeId, not even an explicit nil
### GetNotes

`func (o *UpdateCustomerRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateCustomerRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateCustomerRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateCustomerRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *UpdateCustomerRequest) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *UpdateCustomerRequest) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetCustomFields

`func (o *UpdateCustomerRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *UpdateCustomerRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *UpdateCustomerRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *UpdateCustomerRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *UpdateCustomerRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *UpdateCustomerRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetEmail

`func (o *UpdateCustomerRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateCustomerRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateCustomerRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateCustomerRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdateCustomerRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdateCustomerRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetInvoiceSettings

`func (o *UpdateCustomerRequest) GetInvoiceSettings() CustomerInvoiceSettings`

GetInvoiceSettings returns the InvoiceSettings field if non-nil, zero value otherwise.

### GetInvoiceSettingsOk

`func (o *UpdateCustomerRequest) GetInvoiceSettingsOk() (*CustomerInvoiceSettings, bool)`

GetInvoiceSettingsOk returns a tuple with the InvoiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceSettings

`func (o *UpdateCustomerRequest) SetInvoiceSettings(v CustomerInvoiceSettings)`

SetInvoiceSettings sets InvoiceSettings field to given value.

### HasInvoiceSettings

`func (o *UpdateCustomerRequest) HasInvoiceSettings() bool`

HasInvoiceSettings returns a boolean if a field has been set.

### SetInvoiceSettingsNil

`func (o *UpdateCustomerRequest) SetInvoiceSettingsNil(b bool)`

 SetInvoiceSettingsNil sets the value for InvoiceSettings to be an explicit nil

### UnsetInvoiceSettings
`func (o *UpdateCustomerRequest) UnsetInvoiceSettings()`

UnsetInvoiceSettings ensures that no value is present for InvoiceSettings, not even an explicit nil
### GetBillingEmail

`func (o *UpdateCustomerRequest) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *UpdateCustomerRequest) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *UpdateCustomerRequest) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *UpdateCustomerRequest) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *UpdateCustomerRequest) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *UpdateCustomerRequest) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetLanguage

`func (o *UpdateCustomerRequest) GetLanguage() CustomerLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UpdateCustomerRequest) GetLanguageOk() (*CustomerLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UpdateCustomerRequest) SetLanguage(v CustomerLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UpdateCustomerRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *UpdateCustomerRequest) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *UpdateCustomerRequest) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetPhoneNumber

`func (o *UpdateCustomerRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UpdateCustomerRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UpdateCustomerRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UpdateCustomerRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *UpdateCustomerRequest) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *UpdateCustomerRequest) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


