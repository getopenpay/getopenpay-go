# PriceExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the price. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_PRICE]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**IsActive** | **bool** | Whether the price can be used for new purchases. Price can be activated later. | 
**ProductName** | **string** | Name of the product associated with this price. | 
**ProductId** | **string** | Unique identifier of the product. | 
**Product** | Pointer to [**NullableProductExternal**](ProductExternal.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**InternalDescription** | **NullableString** |  | 
**BillingScheme** | [**BillingSchemeEnum**](BillingSchemeEnum.md) |  | 
**UnitAmountAtom** | **NullableInt32** |  | 
**TransformQuantityDivideBy** | **float32** | This transformation will be applied on quantity before mulitplying by &#39;unit_amount_atom&#39;. | 
**PriceTiers** | [**[]PriceTierExternal**](PriceTierExternal.md) |  | 
**PriceType** | [**PriceTypeEnum**](PriceTypeEnum.md) |  | 
**RecurringDetails** | [**NullableRecurringDetails**](RecurringDetails.md) |  | 
**TiersMode** | [**NullablePricingTiersEnum**](PricingTiersEnum.md) |  | 
**BillingInterval** | [**NullableCalendarIntervalEnum**](CalendarIntervalEnum.md) |  | 
**BillingIntervalCount** | **NullableInt32** |  | 
**ContractTermMultiple** | **NullableInt32** |  | 
**ContractAutoRenew** | **NullableBool** |  | 
**InvoiceSettings** | [**NullableInvoiceSettings**](InvoiceSettings.md) |  | 
**IsLicensed** | **bool** | Whether the price is licensed or not. | 
**IsExclusive** | Pointer to **bool** | Whether the price is exclusive or not. | [optional] [default to false]
**ListedExclusivelyForCustomers** | **[]string** |  | 
**CanOnlyBePurchasedWith** | **[]string** |  | 
**OptionalAddOns** | **[]string** |  | 
**EligibleForUpdates** | **bool** | If the price can be updated or not. | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) |  | 
**IsDefault** | **NullableBool** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPriceExternal

`func NewPriceExternal(id string, createdAt time.Time, updatedAt time.Time, isActive bool, productName string, productId string, internalDescription NullableString, billingScheme BillingSchemeEnum, unitAmountAtom NullableInt32, transformQuantityDivideBy float32, priceTiers []PriceTierExternal, priceType PriceTypeEnum, recurringDetails NullableRecurringDetails, tiersMode NullablePricingTiersEnum, billingInterval NullableCalendarIntervalEnum, billingIntervalCount NullableInt32, contractTermMultiple NullableInt32, contractAutoRenew NullableBool, invoiceSettings NullableInvoiceSettings, isLicensed bool, listedExclusivelyForCustomers []string, canOnlyBePurchasedWith []string, optionalAddOns []string, eligibleForUpdates bool, currency CurrencyEnum, isDefault NullableBool, ) *PriceExternal`

NewPriceExternal instantiates a new PriceExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceExternalWithDefaults

`func NewPriceExternalWithDefaults() *PriceExternal`

NewPriceExternalWithDefaults instantiates a new PriceExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PriceExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PriceExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PriceExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PriceExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *PriceExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PriceExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PriceExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PriceExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PriceExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PriceExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PriceExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *PriceExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PriceExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PriceExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PriceExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsActive

`func (o *PriceExternal) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PriceExternal) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PriceExternal) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetProductName

`func (o *PriceExternal) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *PriceExternal) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *PriceExternal) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetProductId

`func (o *PriceExternal) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PriceExternal) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PriceExternal) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetProduct

`func (o *PriceExternal) GetProduct() ProductExternal`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PriceExternal) GetProductOk() (*ProductExternal, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PriceExternal) SetProduct(v ProductExternal)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PriceExternal) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *PriceExternal) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *PriceExternal) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetName

`func (o *PriceExternal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PriceExternal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PriceExternal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PriceExternal) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PriceExternal) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PriceExternal) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetInternalDescription

`func (o *PriceExternal) GetInternalDescription() string`

GetInternalDescription returns the InternalDescription field if non-nil, zero value otherwise.

### GetInternalDescriptionOk

`func (o *PriceExternal) GetInternalDescriptionOk() (*string, bool)`

GetInternalDescriptionOk returns a tuple with the InternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDescription

`func (o *PriceExternal) SetInternalDescription(v string)`

SetInternalDescription sets InternalDescription field to given value.


### SetInternalDescriptionNil

`func (o *PriceExternal) SetInternalDescriptionNil(b bool)`

 SetInternalDescriptionNil sets the value for InternalDescription to be an explicit nil

### UnsetInternalDescription
`func (o *PriceExternal) UnsetInternalDescription()`

UnsetInternalDescription ensures that no value is present for InternalDescription, not even an explicit nil
### GetBillingScheme

`func (o *PriceExternal) GetBillingScheme() BillingSchemeEnum`

GetBillingScheme returns the BillingScheme field if non-nil, zero value otherwise.

### GetBillingSchemeOk

`func (o *PriceExternal) GetBillingSchemeOk() (*BillingSchemeEnum, bool)`

GetBillingSchemeOk returns a tuple with the BillingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingScheme

`func (o *PriceExternal) SetBillingScheme(v BillingSchemeEnum)`

SetBillingScheme sets BillingScheme field to given value.


### GetUnitAmountAtom

`func (o *PriceExternal) GetUnitAmountAtom() int32`

GetUnitAmountAtom returns the UnitAmountAtom field if non-nil, zero value otherwise.

### GetUnitAmountAtomOk

`func (o *PriceExternal) GetUnitAmountAtomOk() (*int32, bool)`

GetUnitAmountAtomOk returns a tuple with the UnitAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountAtom

`func (o *PriceExternal) SetUnitAmountAtom(v int32)`

SetUnitAmountAtom sets UnitAmountAtom field to given value.


### SetUnitAmountAtomNil

`func (o *PriceExternal) SetUnitAmountAtomNil(b bool)`

 SetUnitAmountAtomNil sets the value for UnitAmountAtom to be an explicit nil

### UnsetUnitAmountAtom
`func (o *PriceExternal) UnsetUnitAmountAtom()`

UnsetUnitAmountAtom ensures that no value is present for UnitAmountAtom, not even an explicit nil
### GetTransformQuantityDivideBy

`func (o *PriceExternal) GetTransformQuantityDivideBy() float32`

GetTransformQuantityDivideBy returns the TransformQuantityDivideBy field if non-nil, zero value otherwise.

### GetTransformQuantityDivideByOk

`func (o *PriceExternal) GetTransformQuantityDivideByOk() (*float32, bool)`

GetTransformQuantityDivideByOk returns a tuple with the TransformQuantityDivideBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformQuantityDivideBy

`func (o *PriceExternal) SetTransformQuantityDivideBy(v float32)`

SetTransformQuantityDivideBy sets TransformQuantityDivideBy field to given value.


### GetPriceTiers

`func (o *PriceExternal) GetPriceTiers() []PriceTierExternal`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *PriceExternal) GetPriceTiersOk() (*[]PriceTierExternal, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *PriceExternal) SetPriceTiers(v []PriceTierExternal)`

SetPriceTiers sets PriceTiers field to given value.


### GetPriceType

`func (o *PriceExternal) GetPriceType() PriceTypeEnum`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *PriceExternal) GetPriceTypeOk() (*PriceTypeEnum, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *PriceExternal) SetPriceType(v PriceTypeEnum)`

SetPriceType sets PriceType field to given value.


### GetRecurringDetails

`func (o *PriceExternal) GetRecurringDetails() RecurringDetails`

GetRecurringDetails returns the RecurringDetails field if non-nil, zero value otherwise.

### GetRecurringDetailsOk

`func (o *PriceExternal) GetRecurringDetailsOk() (*RecurringDetails, bool)`

GetRecurringDetailsOk returns a tuple with the RecurringDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetails

`func (o *PriceExternal) SetRecurringDetails(v RecurringDetails)`

SetRecurringDetails sets RecurringDetails field to given value.


### SetRecurringDetailsNil

`func (o *PriceExternal) SetRecurringDetailsNil(b bool)`

 SetRecurringDetailsNil sets the value for RecurringDetails to be an explicit nil

### UnsetRecurringDetails
`func (o *PriceExternal) UnsetRecurringDetails()`

UnsetRecurringDetails ensures that no value is present for RecurringDetails, not even an explicit nil
### GetTiersMode

`func (o *PriceExternal) GetTiersMode() PricingTiersEnum`

GetTiersMode returns the TiersMode field if non-nil, zero value otherwise.

### GetTiersModeOk

`func (o *PriceExternal) GetTiersModeOk() (*PricingTiersEnum, bool)`

GetTiersModeOk returns a tuple with the TiersMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiersMode

`func (o *PriceExternal) SetTiersMode(v PricingTiersEnum)`

SetTiersMode sets TiersMode field to given value.


### SetTiersModeNil

`func (o *PriceExternal) SetTiersModeNil(b bool)`

 SetTiersModeNil sets the value for TiersMode to be an explicit nil

### UnsetTiersMode
`func (o *PriceExternal) UnsetTiersMode()`

UnsetTiersMode ensures that no value is present for TiersMode, not even an explicit nil
### GetBillingInterval

`func (o *PriceExternal) GetBillingInterval() CalendarIntervalEnum`

GetBillingInterval returns the BillingInterval field if non-nil, zero value otherwise.

### GetBillingIntervalOk

`func (o *PriceExternal) GetBillingIntervalOk() (*CalendarIntervalEnum, bool)`

GetBillingIntervalOk returns a tuple with the BillingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInterval

`func (o *PriceExternal) SetBillingInterval(v CalendarIntervalEnum)`

SetBillingInterval sets BillingInterval field to given value.


### SetBillingIntervalNil

`func (o *PriceExternal) SetBillingIntervalNil(b bool)`

 SetBillingIntervalNil sets the value for BillingInterval to be an explicit nil

### UnsetBillingInterval
`func (o *PriceExternal) UnsetBillingInterval()`

UnsetBillingInterval ensures that no value is present for BillingInterval, not even an explicit nil
### GetBillingIntervalCount

`func (o *PriceExternal) GetBillingIntervalCount() int32`

GetBillingIntervalCount returns the BillingIntervalCount field if non-nil, zero value otherwise.

### GetBillingIntervalCountOk

`func (o *PriceExternal) GetBillingIntervalCountOk() (*int32, bool)`

GetBillingIntervalCountOk returns a tuple with the BillingIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingIntervalCount

`func (o *PriceExternal) SetBillingIntervalCount(v int32)`

SetBillingIntervalCount sets BillingIntervalCount field to given value.


### SetBillingIntervalCountNil

`func (o *PriceExternal) SetBillingIntervalCountNil(b bool)`

 SetBillingIntervalCountNil sets the value for BillingIntervalCount to be an explicit nil

### UnsetBillingIntervalCount
`func (o *PriceExternal) UnsetBillingIntervalCount()`

UnsetBillingIntervalCount ensures that no value is present for BillingIntervalCount, not even an explicit nil
### GetContractTermMultiple

`func (o *PriceExternal) GetContractTermMultiple() int32`

GetContractTermMultiple returns the ContractTermMultiple field if non-nil, zero value otherwise.

### GetContractTermMultipleOk

`func (o *PriceExternal) GetContractTermMultipleOk() (*int32, bool)`

GetContractTermMultipleOk returns a tuple with the ContractTermMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTermMultiple

`func (o *PriceExternal) SetContractTermMultiple(v int32)`

SetContractTermMultiple sets ContractTermMultiple field to given value.


### SetContractTermMultipleNil

`func (o *PriceExternal) SetContractTermMultipleNil(b bool)`

 SetContractTermMultipleNil sets the value for ContractTermMultiple to be an explicit nil

### UnsetContractTermMultiple
`func (o *PriceExternal) UnsetContractTermMultiple()`

UnsetContractTermMultiple ensures that no value is present for ContractTermMultiple, not even an explicit nil
### GetContractAutoRenew

`func (o *PriceExternal) GetContractAutoRenew() bool`

GetContractAutoRenew returns the ContractAutoRenew field if non-nil, zero value otherwise.

### GetContractAutoRenewOk

`func (o *PriceExternal) GetContractAutoRenewOk() (*bool, bool)`

GetContractAutoRenewOk returns a tuple with the ContractAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAutoRenew

`func (o *PriceExternal) SetContractAutoRenew(v bool)`

SetContractAutoRenew sets ContractAutoRenew field to given value.


### SetContractAutoRenewNil

`func (o *PriceExternal) SetContractAutoRenewNil(b bool)`

 SetContractAutoRenewNil sets the value for ContractAutoRenew to be an explicit nil

### UnsetContractAutoRenew
`func (o *PriceExternal) UnsetContractAutoRenew()`

UnsetContractAutoRenew ensures that no value is present for ContractAutoRenew, not even an explicit nil
### GetInvoiceSettings

`func (o *PriceExternal) GetInvoiceSettings() InvoiceSettings`

GetInvoiceSettings returns the InvoiceSettings field if non-nil, zero value otherwise.

### GetInvoiceSettingsOk

`func (o *PriceExternal) GetInvoiceSettingsOk() (*InvoiceSettings, bool)`

GetInvoiceSettingsOk returns a tuple with the InvoiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceSettings

`func (o *PriceExternal) SetInvoiceSettings(v InvoiceSettings)`

SetInvoiceSettings sets InvoiceSettings field to given value.


### SetInvoiceSettingsNil

`func (o *PriceExternal) SetInvoiceSettingsNil(b bool)`

 SetInvoiceSettingsNil sets the value for InvoiceSettings to be an explicit nil

### UnsetInvoiceSettings
`func (o *PriceExternal) UnsetInvoiceSettings()`

UnsetInvoiceSettings ensures that no value is present for InvoiceSettings, not even an explicit nil
### GetIsLicensed

`func (o *PriceExternal) GetIsLicensed() bool`

GetIsLicensed returns the IsLicensed field if non-nil, zero value otherwise.

### GetIsLicensedOk

`func (o *PriceExternal) GetIsLicensedOk() (*bool, bool)`

GetIsLicensedOk returns a tuple with the IsLicensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLicensed

`func (o *PriceExternal) SetIsLicensed(v bool)`

SetIsLicensed sets IsLicensed field to given value.


### GetIsExclusive

`func (o *PriceExternal) GetIsExclusive() bool`

GetIsExclusive returns the IsExclusive field if non-nil, zero value otherwise.

### GetIsExclusiveOk

`func (o *PriceExternal) GetIsExclusiveOk() (*bool, bool)`

GetIsExclusiveOk returns a tuple with the IsExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExclusive

`func (o *PriceExternal) SetIsExclusive(v bool)`

SetIsExclusive sets IsExclusive field to given value.

### HasIsExclusive

`func (o *PriceExternal) HasIsExclusive() bool`

HasIsExclusive returns a boolean if a field has been set.

### GetListedExclusivelyForCustomers

`func (o *PriceExternal) GetListedExclusivelyForCustomers() []string`

GetListedExclusivelyForCustomers returns the ListedExclusivelyForCustomers field if non-nil, zero value otherwise.

### GetListedExclusivelyForCustomersOk

`func (o *PriceExternal) GetListedExclusivelyForCustomersOk() (*[]string, bool)`

GetListedExclusivelyForCustomersOk returns a tuple with the ListedExclusivelyForCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedExclusivelyForCustomers

`func (o *PriceExternal) SetListedExclusivelyForCustomers(v []string)`

SetListedExclusivelyForCustomers sets ListedExclusivelyForCustomers field to given value.


### GetCanOnlyBePurchasedWith

`func (o *PriceExternal) GetCanOnlyBePurchasedWith() []string`

GetCanOnlyBePurchasedWith returns the CanOnlyBePurchasedWith field if non-nil, zero value otherwise.

### GetCanOnlyBePurchasedWithOk

`func (o *PriceExternal) GetCanOnlyBePurchasedWithOk() (*[]string, bool)`

GetCanOnlyBePurchasedWithOk returns a tuple with the CanOnlyBePurchasedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanOnlyBePurchasedWith

`func (o *PriceExternal) SetCanOnlyBePurchasedWith(v []string)`

SetCanOnlyBePurchasedWith sets CanOnlyBePurchasedWith field to given value.


### GetOptionalAddOns

`func (o *PriceExternal) GetOptionalAddOns() []string`

GetOptionalAddOns returns the OptionalAddOns field if non-nil, zero value otherwise.

### GetOptionalAddOnsOk

`func (o *PriceExternal) GetOptionalAddOnsOk() (*[]string, bool)`

GetOptionalAddOnsOk returns a tuple with the OptionalAddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalAddOns

`func (o *PriceExternal) SetOptionalAddOns(v []string)`

SetOptionalAddOns sets OptionalAddOns field to given value.


### GetEligibleForUpdates

`func (o *PriceExternal) GetEligibleForUpdates() bool`

GetEligibleForUpdates returns the EligibleForUpdates field if non-nil, zero value otherwise.

### GetEligibleForUpdatesOk

`func (o *PriceExternal) GetEligibleForUpdatesOk() (*bool, bool)`

GetEligibleForUpdatesOk returns a tuple with the EligibleForUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleForUpdates

`func (o *PriceExternal) SetEligibleForUpdates(v bool)`

SetEligibleForUpdates sets EligibleForUpdates field to given value.


### GetCurrency

`func (o *PriceExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PriceExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PriceExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetIsDefault

`func (o *PriceExternal) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PriceExternal) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PriceExternal) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### SetIsDefaultNil

`func (o *PriceExternal) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *PriceExternal) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetMetadata

`func (o *PriceExternal) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PriceExternal) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PriceExternal) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PriceExternal) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PriceExternal) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PriceExternal) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


