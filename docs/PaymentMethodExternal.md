# PaymentMethodExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Identifier of the payment_method. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_PAYMENT_METHOD]
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

## Methods

### NewPaymentMethodExternal

`func NewPaymentMethodExternal(id string, createdAt time.Time, updatedAt time.Time, provider PaymentProviderType, displayName string, ) *PaymentMethodExternal`

NewPaymentMethodExternal instantiates a new PaymentMethodExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodExternalWithDefaults

`func NewPaymentMethodExternalWithDefaults() *PaymentMethodExternal`

NewPaymentMethodExternalWithDefaults instantiates a new PaymentMethodExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentMethodExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PaymentMethodExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethodExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethodExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *PaymentMethodExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentMethodExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethodExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethodExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PaymentMethodExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentMethodExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentMethodExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *PaymentMethodExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PaymentMethodExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PaymentMethodExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PaymentMethodExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetCustomerId

`func (o *PaymentMethodExternal) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PaymentMethodExternal) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PaymentMethodExternal) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PaymentMethodExternal) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *PaymentMethodExternal) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *PaymentMethodExternal) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetBillingAddress

`func (o *PaymentMethodExternal) GetBillingAddress() CompleteAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *PaymentMethodExternal) GetBillingAddressOk() (*CompleteAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *PaymentMethodExternal) SetBillingAddress(v CompleteAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *PaymentMethodExternal) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### SetBillingAddressNil

`func (o *PaymentMethodExternal) SetBillingAddressNil(b bool)`

 SetBillingAddressNil sets the value for BillingAddress to be an explicit nil

### UnsetBillingAddress
`func (o *PaymentMethodExternal) UnsetBillingAddress()`

UnsetBillingAddress ensures that no value is present for BillingAddress, not even an explicit nil
### GetMetadata

`func (o *PaymentMethodExternal) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentMethodExternal) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentMethodExternal) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentMethodExternal) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PaymentMethodExternal) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PaymentMethodExternal) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetProvider

`func (o *PaymentMethodExternal) GetProvider() PaymentProviderType`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PaymentMethodExternal) GetProviderOk() (*PaymentProviderType, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PaymentMethodExternal) SetProvider(v PaymentProviderType)`

SetProvider sets Provider field to given value.


### GetCardType

`func (o *PaymentMethodExternal) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *PaymentMethodExternal) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *PaymentMethodExternal) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *PaymentMethodExternal) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### SetCardTypeNil

`func (o *PaymentMethodExternal) SetCardTypeNil(b bool)`

 SetCardTypeNil sets the value for CardType to be an explicit nil

### UnsetCardType
`func (o *PaymentMethodExternal) UnsetCardType()`

UnsetCardType ensures that no value is present for CardType, not even an explicit nil
### GetLastFour

`func (o *PaymentMethodExternal) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *PaymentMethodExternal) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *PaymentMethodExternal) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *PaymentMethodExternal) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### SetLastFourNil

`func (o *PaymentMethodExternal) SetLastFourNil(b bool)`

 SetLastFourNil sets the value for LastFour to be an explicit nil

### UnsetLastFour
`func (o *PaymentMethodExternal) UnsetLastFour()`

UnsetLastFour ensures that no value is present for LastFour, not even an explicit nil
### GetDisplayName

`func (o *PaymentMethodExternal) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PaymentMethodExternal) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PaymentMethodExternal) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


