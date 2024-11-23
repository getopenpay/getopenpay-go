# TokenizedCreditCardExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Identifier of the payment_method. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_TOKENIZED_CARD_PAYMENT_METHOD]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**BillingAddress** | Pointer to [**NullableCompleteAddress**](CompleteAddress.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Provider** | [**PaymentProviderType**](PaymentProviderType.md) |  | 
**CardType** | Pointer to **NullableString** |  | [optional] 
**LastFour** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | **string** | Display name for the payment method to show on the UI. | 
**ExpiryDate** | **NullableString** |  | 
**CardFingerprint** | Pointer to **NullableString** |  | [optional] 
**CardIin** | Pointer to **NullableString** |  | [optional] 
**CardCountry** | Pointer to **NullableString** |  | [optional] 
**CardBrand** | Pointer to **NullableString** |  | [optional] 
**CardIssuer** | Pointer to **NullableString** |  | [optional] 
**IsFullDetailsKnown** | **bool** | Indicates if full details of the card are known | 
**CdeExternalId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTokenizedCreditCardExternal

`func NewTokenizedCreditCardExternal(id string, createdAt time.Time, updatedAt time.Time, provider PaymentProviderType, displayName string, expiryDate NullableString, isFullDetailsKnown bool, ) *TokenizedCreditCardExternal`

NewTokenizedCreditCardExternal instantiates a new TokenizedCreditCardExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizedCreditCardExternalWithDefaults

`func NewTokenizedCreditCardExternalWithDefaults() *TokenizedCreditCardExternal`

NewTokenizedCreditCardExternalWithDefaults instantiates a new TokenizedCreditCardExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenizedCreditCardExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenizedCreditCardExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenizedCreditCardExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *TokenizedCreditCardExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TokenizedCreditCardExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TokenizedCreditCardExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *TokenizedCreditCardExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TokenizedCreditCardExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TokenizedCreditCardExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TokenizedCreditCardExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TokenizedCreditCardExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TokenizedCreditCardExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TokenizedCreditCardExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *TokenizedCreditCardExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TokenizedCreditCardExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TokenizedCreditCardExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TokenizedCreditCardExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetCustomerId

`func (o *TokenizedCreditCardExternal) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TokenizedCreditCardExternal) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TokenizedCreditCardExternal) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *TokenizedCreditCardExternal) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *TokenizedCreditCardExternal) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *TokenizedCreditCardExternal) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetBillingAddress

`func (o *TokenizedCreditCardExternal) GetBillingAddress() CompleteAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *TokenizedCreditCardExternal) GetBillingAddressOk() (*CompleteAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *TokenizedCreditCardExternal) SetBillingAddress(v CompleteAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *TokenizedCreditCardExternal) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### SetBillingAddressNil

`func (o *TokenizedCreditCardExternal) SetBillingAddressNil(b bool)`

 SetBillingAddressNil sets the value for BillingAddress to be an explicit nil

### UnsetBillingAddress
`func (o *TokenizedCreditCardExternal) UnsetBillingAddress()`

UnsetBillingAddress ensures that no value is present for BillingAddress, not even an explicit nil
### GetMetadata

`func (o *TokenizedCreditCardExternal) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TokenizedCreditCardExternal) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TokenizedCreditCardExternal) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TokenizedCreditCardExternal) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *TokenizedCreditCardExternal) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *TokenizedCreditCardExternal) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetProvider

`func (o *TokenizedCreditCardExternal) GetProvider() PaymentProviderType`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TokenizedCreditCardExternal) GetProviderOk() (*PaymentProviderType, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TokenizedCreditCardExternal) SetProvider(v PaymentProviderType)`

SetProvider sets Provider field to given value.


### GetCardType

`func (o *TokenizedCreditCardExternal) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *TokenizedCreditCardExternal) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *TokenizedCreditCardExternal) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *TokenizedCreditCardExternal) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### SetCardTypeNil

`func (o *TokenizedCreditCardExternal) SetCardTypeNil(b bool)`

 SetCardTypeNil sets the value for CardType to be an explicit nil

### UnsetCardType
`func (o *TokenizedCreditCardExternal) UnsetCardType()`

UnsetCardType ensures that no value is present for CardType, not even an explicit nil
### GetLastFour

`func (o *TokenizedCreditCardExternal) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *TokenizedCreditCardExternal) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *TokenizedCreditCardExternal) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *TokenizedCreditCardExternal) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### SetLastFourNil

`func (o *TokenizedCreditCardExternal) SetLastFourNil(b bool)`

 SetLastFourNil sets the value for LastFour to be an explicit nil

### UnsetLastFour
`func (o *TokenizedCreditCardExternal) UnsetLastFour()`

UnsetLastFour ensures that no value is present for LastFour, not even an explicit nil
### GetDisplayName

`func (o *TokenizedCreditCardExternal) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TokenizedCreditCardExternal) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TokenizedCreditCardExternal) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetExpiryDate

`func (o *TokenizedCreditCardExternal) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *TokenizedCreditCardExternal) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *TokenizedCreditCardExternal) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.


### SetExpiryDateNil

`func (o *TokenizedCreditCardExternal) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *TokenizedCreditCardExternal) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetCardFingerprint

`func (o *TokenizedCreditCardExternal) GetCardFingerprint() string`

GetCardFingerprint returns the CardFingerprint field if non-nil, zero value otherwise.

### GetCardFingerprintOk

`func (o *TokenizedCreditCardExternal) GetCardFingerprintOk() (*string, bool)`

GetCardFingerprintOk returns a tuple with the CardFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFingerprint

`func (o *TokenizedCreditCardExternal) SetCardFingerprint(v string)`

SetCardFingerprint sets CardFingerprint field to given value.

### HasCardFingerprint

`func (o *TokenizedCreditCardExternal) HasCardFingerprint() bool`

HasCardFingerprint returns a boolean if a field has been set.

### SetCardFingerprintNil

`func (o *TokenizedCreditCardExternal) SetCardFingerprintNil(b bool)`

 SetCardFingerprintNil sets the value for CardFingerprint to be an explicit nil

### UnsetCardFingerprint
`func (o *TokenizedCreditCardExternal) UnsetCardFingerprint()`

UnsetCardFingerprint ensures that no value is present for CardFingerprint, not even an explicit nil
### GetCardIin

`func (o *TokenizedCreditCardExternal) GetCardIin() string`

GetCardIin returns the CardIin field if non-nil, zero value otherwise.

### GetCardIinOk

`func (o *TokenizedCreditCardExternal) GetCardIinOk() (*string, bool)`

GetCardIinOk returns a tuple with the CardIin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIin

`func (o *TokenizedCreditCardExternal) SetCardIin(v string)`

SetCardIin sets CardIin field to given value.

### HasCardIin

`func (o *TokenizedCreditCardExternal) HasCardIin() bool`

HasCardIin returns a boolean if a field has been set.

### SetCardIinNil

`func (o *TokenizedCreditCardExternal) SetCardIinNil(b bool)`

 SetCardIinNil sets the value for CardIin to be an explicit nil

### UnsetCardIin
`func (o *TokenizedCreditCardExternal) UnsetCardIin()`

UnsetCardIin ensures that no value is present for CardIin, not even an explicit nil
### GetCardCountry

`func (o *TokenizedCreditCardExternal) GetCardCountry() string`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *TokenizedCreditCardExternal) GetCardCountryOk() (*string, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *TokenizedCreditCardExternal) SetCardCountry(v string)`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *TokenizedCreditCardExternal) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.

### SetCardCountryNil

`func (o *TokenizedCreditCardExternal) SetCardCountryNil(b bool)`

 SetCardCountryNil sets the value for CardCountry to be an explicit nil

### UnsetCardCountry
`func (o *TokenizedCreditCardExternal) UnsetCardCountry()`

UnsetCardCountry ensures that no value is present for CardCountry, not even an explicit nil
### GetCardBrand

`func (o *TokenizedCreditCardExternal) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *TokenizedCreditCardExternal) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *TokenizedCreditCardExternal) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *TokenizedCreditCardExternal) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### SetCardBrandNil

`func (o *TokenizedCreditCardExternal) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *TokenizedCreditCardExternal) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetCardIssuer

`func (o *TokenizedCreditCardExternal) GetCardIssuer() string`

GetCardIssuer returns the CardIssuer field if non-nil, zero value otherwise.

### GetCardIssuerOk

`func (o *TokenizedCreditCardExternal) GetCardIssuerOk() (*string, bool)`

GetCardIssuerOk returns a tuple with the CardIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIssuer

`func (o *TokenizedCreditCardExternal) SetCardIssuer(v string)`

SetCardIssuer sets CardIssuer field to given value.

### HasCardIssuer

`func (o *TokenizedCreditCardExternal) HasCardIssuer() bool`

HasCardIssuer returns a boolean if a field has been set.

### SetCardIssuerNil

`func (o *TokenizedCreditCardExternal) SetCardIssuerNil(b bool)`

 SetCardIssuerNil sets the value for CardIssuer to be an explicit nil

### UnsetCardIssuer
`func (o *TokenizedCreditCardExternal) UnsetCardIssuer()`

UnsetCardIssuer ensures that no value is present for CardIssuer, not even an explicit nil
### GetIsFullDetailsKnown

`func (o *TokenizedCreditCardExternal) GetIsFullDetailsKnown() bool`

GetIsFullDetailsKnown returns the IsFullDetailsKnown field if non-nil, zero value otherwise.

### GetIsFullDetailsKnownOk

`func (o *TokenizedCreditCardExternal) GetIsFullDetailsKnownOk() (*bool, bool)`

GetIsFullDetailsKnownOk returns a tuple with the IsFullDetailsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullDetailsKnown

`func (o *TokenizedCreditCardExternal) SetIsFullDetailsKnown(v bool)`

SetIsFullDetailsKnown sets IsFullDetailsKnown field to given value.


### GetCdeExternalId

`func (o *TokenizedCreditCardExternal) GetCdeExternalId() string`

GetCdeExternalId returns the CdeExternalId field if non-nil, zero value otherwise.

### GetCdeExternalIdOk

`func (o *TokenizedCreditCardExternal) GetCdeExternalIdOk() (*string, bool)`

GetCdeExternalIdOk returns a tuple with the CdeExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdeExternalId

`func (o *TokenizedCreditCardExternal) SetCdeExternalId(v string)`

SetCdeExternalId sets CdeExternalId field to given value.

### HasCdeExternalId

`func (o *TokenizedCreditCardExternal) HasCdeExternalId() bool`

HasCdeExternalId returns a boolean if a field has been set.

### SetCdeExternalIdNil

`func (o *TokenizedCreditCardExternal) SetCdeExternalIdNil(b bool)`

 SetCdeExternalIdNil sets the value for CdeExternalId to be an explicit nil

### UnsetCdeExternalId
`func (o *TokenizedCreditCardExternal) UnsetCdeExternalId()`

UnsetCdeExternalId ensures that no value is present for CdeExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


