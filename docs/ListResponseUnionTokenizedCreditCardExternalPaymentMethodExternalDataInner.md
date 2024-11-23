# ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner

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

### NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner

`func NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner(id interface{}, createdAt interface{}, updatedAt interface{}, provider PaymentProviderType, displayName interface{}, expiryDate interface{}, isFullDetailsKnown interface{}, ) *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner`

NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner instantiates a new ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInnerWithDefaults

`func NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInnerWithDefaults() *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner`

NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInnerWithDefaults instantiates a new ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObject

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetIsDeleted

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetIsDeleted() interface{}`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetIsDeletedOk() (*interface{}, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetIsDeleted(v interface{})`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetCustomerId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCustomerId() interface{}`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCustomerIdOk() (*interface{}, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCustomerId(v interface{})`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetBillingAddress

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetBillingAddress() CompleteAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetBillingAddressOk() (*CompleteAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetBillingAddress(v CompleteAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetProvider

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetProvider() PaymentProviderType`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetProviderOk() (*PaymentProviderType, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetProvider(v PaymentProviderType)`

SetProvider sets Provider field to given value.


### GetCardType

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardType() interface{}`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardTypeOk() (*interface{}, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardType(v interface{})`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### SetCardTypeNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardTypeNil(b bool)`

 SetCardTypeNil sets the value for CardType to be an explicit nil

### UnsetCardType
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetCardType()`

UnsetCardType ensures that no value is present for CardType, not even an explicit nil
### GetLastFour

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetLastFour() interface{}`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetLastFourOk() (*interface{}, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetLastFour(v interface{})`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### SetLastFourNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetLastFourNil(b bool)`

 SetLastFourNil sets the value for LastFour to be an explicit nil

### UnsetLastFour
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetLastFour()`

UnsetLastFour ensures that no value is present for LastFour, not even an explicit nil
### GetDisplayName

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetDisplayName() interface{}`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetDisplayNameOk() (*interface{}, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetDisplayName(v interface{})`

SetDisplayName sets DisplayName field to given value.


### SetDisplayNameNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExpiryDate

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetExpiryDate() interface{}`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetExpiryDateOk() (*interface{}, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetExpiryDate(v interface{})`

SetExpiryDate sets ExpiryDate field to given value.


### SetExpiryDateNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetCardFingerprint

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardFingerprint() interface{}`

GetCardFingerprint returns the CardFingerprint field if non-nil, zero value otherwise.

### GetCardFingerprintOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardFingerprintOk() (*interface{}, bool)`

GetCardFingerprintOk returns a tuple with the CardFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFingerprint

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardFingerprint(v interface{})`

SetCardFingerprint sets CardFingerprint field to given value.

### HasCardFingerprint

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCardFingerprint() bool`

HasCardFingerprint returns a boolean if a field has been set.

### SetCardFingerprintNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardFingerprintNil(b bool)`

 SetCardFingerprintNil sets the value for CardFingerprint to be an explicit nil

### UnsetCardFingerprint
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetCardFingerprint()`

UnsetCardFingerprint ensures that no value is present for CardFingerprint, not even an explicit nil
### GetCardIin

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardIin() interface{}`

GetCardIin returns the CardIin field if non-nil, zero value otherwise.

### GetCardIinOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardIinOk() (*interface{}, bool)`

GetCardIinOk returns a tuple with the CardIin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIin

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardIin(v interface{})`

SetCardIin sets CardIin field to given value.

### HasCardIin

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCardIin() bool`

HasCardIin returns a boolean if a field has been set.

### SetCardIinNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardIinNil(b bool)`

 SetCardIinNil sets the value for CardIin to be an explicit nil

### UnsetCardIin
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetCardIin()`

UnsetCardIin ensures that no value is present for CardIin, not even an explicit nil
### GetCardCountry

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardCountry() interface{}`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardCountryOk() (*interface{}, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardCountry(v interface{})`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.

### SetCardCountryNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardCountryNil(b bool)`

 SetCardCountryNil sets the value for CardCountry to be an explicit nil

### UnsetCardCountry
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetCardCountry()`

UnsetCardCountry ensures that no value is present for CardCountry, not even an explicit nil
### GetCardBrand

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardBrand() interface{}`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardBrandOk() (*interface{}, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardBrand(v interface{})`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### SetCardBrandNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetCardIssuer

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardIssuer() interface{}`

GetCardIssuer returns the CardIssuer field if non-nil, zero value otherwise.

### GetCardIssuerOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCardIssuerOk() (*interface{}, bool)`

GetCardIssuerOk returns a tuple with the CardIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIssuer

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardIssuer(v interface{})`

SetCardIssuer sets CardIssuer field to given value.

### HasCardIssuer

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCardIssuer() bool`

HasCardIssuer returns a boolean if a field has been set.

### SetCardIssuerNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCardIssuerNil(b bool)`

 SetCardIssuerNil sets the value for CardIssuer to be an explicit nil

### UnsetCardIssuer
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetCardIssuer()`

UnsetCardIssuer ensures that no value is present for CardIssuer, not even an explicit nil
### GetIsFullDetailsKnown

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetIsFullDetailsKnown() interface{}`

GetIsFullDetailsKnown returns the IsFullDetailsKnown field if non-nil, zero value otherwise.

### GetIsFullDetailsKnownOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetIsFullDetailsKnownOk() (*interface{}, bool)`

GetIsFullDetailsKnownOk returns a tuple with the IsFullDetailsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullDetailsKnown

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetIsFullDetailsKnown(v interface{})`

SetIsFullDetailsKnown sets IsFullDetailsKnown field to given value.


### SetIsFullDetailsKnownNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetIsFullDetailsKnownNil(b bool)`

 SetIsFullDetailsKnownNil sets the value for IsFullDetailsKnown to be an explicit nil

### UnsetIsFullDetailsKnown
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetIsFullDetailsKnown()`

UnsetIsFullDetailsKnown ensures that no value is present for IsFullDetailsKnown, not even an explicit nil
### GetCdeExternalId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCdeExternalId() interface{}`

GetCdeExternalId returns the CdeExternalId field if non-nil, zero value otherwise.

### GetCdeExternalIdOk

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) GetCdeExternalIdOk() (*interface{}, bool)`

GetCdeExternalIdOk returns a tuple with the CdeExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdeExternalId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCdeExternalId(v interface{})`

SetCdeExternalId sets CdeExternalId field to given value.

### HasCdeExternalId

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) HasCdeExternalId() bool`

HasCdeExternalId returns a boolean if a field has been set.

### SetCdeExternalIdNil

`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) SetCdeExternalIdNil(b bool)`

 SetCdeExternalIdNil sets the value for CdeExternalId to be an explicit nil

### UnsetCdeExternalId
`func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnsetCdeExternalId()`

UnsetCdeExternalId ensures that no value is present for CdeExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


