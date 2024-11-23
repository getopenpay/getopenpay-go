# ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **interface{}** | Unique Identifier of the payment_method. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_PAYMENT_METHOD]
**CreatedAt** | **interface{}** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **interface{}** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **interface{}** | If true, indicates that this object has been deleted | [optional] [default to false]
**CustomerId** | Pointer to **interface{}** |  | [optional] 
**BillingAddress** | Pointer to [**CompleteAddress**](CompleteAddress.md) |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**Provider** | [**PaymentProviderType**](PaymentProviderType.md) |  | 
**CardType** | Pointer to **interface{}** |  | [optional] 
**LastFour** | Pointer to **interface{}** |  | [optional] 
**DisplayName** | **interface{}** | Display name for the payment method to show on the UI. | 
**ExpiryDate** | **interface{}** |  | 
**CardFingerprint** | Pointer to **interface{}** |  | [optional] 
**CardIin** | Pointer to **interface{}** |  | [optional] 
**CardCountry** | Pointer to **interface{}** |  | [optional] 
**CardBrand** | Pointer to **interface{}** |  | [optional] 
**CardIssuer** | Pointer to **interface{}** |  | [optional] 
**IsFullDetailsKnown** | **interface{}** | Indicates if full details of the card are known | 
**CdeExternalId** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet

`func NewResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet(id interface{}, createdAt interface{}, updatedAt interface{}, provider PaymentProviderType, displayName interface{}, expiryDate interface{}, isFullDetailsKnown interface{}, ) *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet`

NewResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet instantiates a new ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGetWithDefaults

`func NewResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGetWithDefaults() *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet`

NewResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGetWithDefaults instantiates a new ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObject

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetIsDeleted

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetIsDeleted() interface{}`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetIsDeletedOk() (*interface{}, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetIsDeleted(v interface{})`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetCustomerId

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCustomerId() interface{}`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCustomerIdOk() (*interface{}, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCustomerId(v interface{})`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetBillingAddress

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetBillingAddress() CompleteAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetBillingAddressOk() (*CompleteAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetBillingAddress(v CompleteAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetProvider

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetProvider() PaymentProviderType`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetProviderOk() (*PaymentProviderType, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetProvider(v PaymentProviderType)`

SetProvider sets Provider field to given value.


### GetCardType

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCardType() interface{}`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCardTypeOk() (*interface{}, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCardType(v interface{})`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### SetCardTypeNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCardTypeNil(b bool)`

 SetCardTypeNil sets the value for CardType to be an explicit nil

### UnsetCardType
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetCardType()`

UnsetCardType ensures that no value is present for CardType, not even an explicit nil
### GetLastFour

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetLastFour() interface{}`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetLastFourOk() (*interface{}, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetLastFour(v interface{})`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### SetLastFourNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetLastFourNil(b bool)`

 SetLastFourNil sets the value for LastFour to be an explicit nil

### UnsetLastFour
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetLastFour()`

UnsetLastFour ensures that no value is present for LastFour, not even an explicit nil
### GetDisplayName

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetDisplayName() interface{}`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetDisplayNameOk() (*interface{}, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetDisplayName(v interface{})`

SetDisplayName sets DisplayName field to given value.


### SetDisplayNameNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExpiryDate

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetExpiryDate() interface{}`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetExpiryDateOk() (*interface{}, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetExpiryDate(v interface{})`

SetExpiryDate sets ExpiryDate field to given value.


### SetExpiryDateNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetCardFingerprint

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCardFingerprint() interface{}`

GetCardFingerprint returns the CardFingerprint field if non-nil, zero value otherwise.

### GetCardFingerprintOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCardFingerprintOk() (*interface{}, bool)`

GetCardFingerprintOk returns a tuple with the CardFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFingerprint

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCardFingerprint(v interface{})`

SetCardFingerprint sets CardFingerprint field to given value.

### HasCardFingerprint

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) HasCardFingerprint() bool`

HasCardFingerprint returns a boolean if a field has been set.

### SetCardFingerprintNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCardFingerprintNil(b bool)`

 SetCardFingerprintNil sets the value for CardFingerprint to be an explicit nil

### UnsetCardFingerprint
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetCardFingerprint()`

UnsetCardFingerprint ensures that no value is present for CardFingerprint, not even an explicit nil
### GetCardIin

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCardIin() interface{}`

GetCardIin returns the CardIin field if non-nil, zero value otherwise.

### GetCardIinOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCardIinOk() (*interface{}, bool)`

GetCardIinOk returns a tuple with the CardIin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIin

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCardIin(v interface{})`

SetCardIin sets CardIin field to given value.

### HasCardIin

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) HasCardIin() bool`

HasCardIin returns a boolean if a field has been set.

### SetCardIinNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCardIinNil(b bool)`

 SetCardIinNil sets the value for CardIin to be an explicit nil

### UnsetCardIin
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetCardIin()`

UnsetCardIin ensures that no value is present for CardIin, not even an explicit nil
### GetCardCountry

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCardCountry() interface{}`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCardCountryOk() (*interface{}, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCardCountry(v interface{})`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.

### SetCardCountryNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCardCountryNil(b bool)`

 SetCardCountryNil sets the value for CardCountry to be an explicit nil

### UnsetCardCountry
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetCardCountry()`

UnsetCardCountry ensures that no value is present for CardCountry, not even an explicit nil
### GetCardBrand

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCardBrand() interface{}`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCardBrandOk() (*interface{}, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCardBrand(v interface{})`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### SetCardBrandNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetCardIssuer

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCardIssuer() interface{}`

GetCardIssuer returns the CardIssuer field if non-nil, zero value otherwise.

### GetCardIssuerOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCardIssuerOk() (*interface{}, bool)`

GetCardIssuerOk returns a tuple with the CardIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIssuer

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCardIssuer(v interface{})`

SetCardIssuer sets CardIssuer field to given value.

### HasCardIssuer

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) HasCardIssuer() bool`

HasCardIssuer returns a boolean if a field has been set.

### SetCardIssuerNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCardIssuerNil(b bool)`

 SetCardIssuerNil sets the value for CardIssuer to be an explicit nil

### UnsetCardIssuer
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetCardIssuer()`

UnsetCardIssuer ensures that no value is present for CardIssuer, not even an explicit nil
### GetIsFullDetailsKnown

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetIsFullDetailsKnown() interface{}`

GetIsFullDetailsKnown returns the IsFullDetailsKnown field if non-nil, zero value otherwise.

### GetIsFullDetailsKnownOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetIsFullDetailsKnownOk() (*interface{}, bool)`

GetIsFullDetailsKnownOk returns a tuple with the IsFullDetailsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullDetailsKnown

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetIsFullDetailsKnown(v interface{})`

SetIsFullDetailsKnown sets IsFullDetailsKnown field to given value.


### SetIsFullDetailsKnownNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetIsFullDetailsKnownNil(b bool)`

 SetIsFullDetailsKnownNil sets the value for IsFullDetailsKnown to be an explicit nil

### UnsetIsFullDetailsKnown
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetIsFullDetailsKnown()`

UnsetIsFullDetailsKnown ensures that no value is present for IsFullDetailsKnown, not even an explicit nil
### GetCdeExternalId

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCdeExternalId() interface{}`

GetCdeExternalId returns the CdeExternalId field if non-nil, zero value otherwise.

### GetCdeExternalIdOk

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) GetCdeExternalIdOk() (*interface{}, bool)`

GetCdeExternalIdOk returns a tuple with the CdeExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdeExternalId

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCdeExternalId(v interface{})`

SetCdeExternalId sets CdeExternalId field to given value.

### HasCdeExternalId

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) HasCdeExternalId() bool`

HasCdeExternalId returns a boolean if a field has been set.

### SetCdeExternalIdNil

`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) SetCdeExternalIdNil(b bool)`

 SetCdeExternalIdNil sets the value for CdeExternalId to be an explicit nil

### UnsetCdeExternalId
`func (o *ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet) UnsetCdeExternalId()`

UnsetCdeExternalId ensures that no value is present for CdeExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


