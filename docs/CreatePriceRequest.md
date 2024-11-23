# CreatePriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsExclusive** | Pointer to **NullableBool** |  | [optional] 
**ListedExclusivelyForCustomers** | Pointer to **[]string** |  | [optional] 
**IsActive** | **bool** | Whether the price can be used for new purchases. | 
**Name** | Pointer to **NullableString** |  | [optional] 
**InternalDescription** | Pointer to **NullableString** |  | [optional] 
**PricingModel** | [**PricingModel**](PricingModel.md) |  | 
**UnitAmountAtom** | **NullableInt32** |  | 
**TransformQuantityDivideBy** | Pointer to **float32** | This transformation will be applied on quantity before multiplying by unit_amount_atom. | [optional] [default to 1.0]
**Currency** | Pointer to [**CurrencyEnum**](CurrencyEnum.md) |  | [optional] [default to CURRENCYENUM_USD]
**PriceTiers** | Pointer to [**[]PriceTierParams**](PriceTierParams.md) |  | [optional] 
**PriceType** | [**PriceTypeEnum**](PriceTypeEnum.md) |  | 
**BillingInterval** | Pointer to [**NullableCalendarIntervalEnum**](CalendarIntervalEnum.md) |  | [optional] 
**BillingIntervalCount** | Pointer to **NullableInt32** |  | [optional] 
**ContractTermMultiple** | Pointer to **NullableInt32** |  | [optional] 
**ContractAutoRenew** | Pointer to **NullableBool** |  | [optional] 
**TrialPeriodDays** | Pointer to **int32** | Number of trail days for this Price. | [optional] [default to 0]
**UsageType** | Pointer to [**UsageTypeEnum**](UsageTypeEnum.md) |  | [optional] [default to USAGETYPEENUM_LICENSED]
**AggregateUsage** | Pointer to [**UsageAggMethodEnum**](UsageAggMethodEnum.md) |  | [optional] [default to USAGEAGGMETHODENUM_SUM]
**DefaultNetD** | Pointer to **NullableInt32** |  | [optional] 
**CanOnlyBePurchasedWith** | Pointer to **[]string** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**ProductId** | **string** | Unique identifier of the product. | 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCreatePriceRequest

`func NewCreatePriceRequest(isActive bool, pricingModel PricingModel, unitAmountAtom NullableInt32, priceType PriceTypeEnum, productId string, ) *CreatePriceRequest`

NewCreatePriceRequest instantiates a new CreatePriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePriceRequestWithDefaults

`func NewCreatePriceRequestWithDefaults() *CreatePriceRequest`

NewCreatePriceRequestWithDefaults instantiates a new CreatePriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsExclusive

`func (o *CreatePriceRequest) GetIsExclusive() bool`

GetIsExclusive returns the IsExclusive field if non-nil, zero value otherwise.

### GetIsExclusiveOk

`func (o *CreatePriceRequest) GetIsExclusiveOk() (*bool, bool)`

GetIsExclusiveOk returns a tuple with the IsExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExclusive

`func (o *CreatePriceRequest) SetIsExclusive(v bool)`

SetIsExclusive sets IsExclusive field to given value.

### HasIsExclusive

`func (o *CreatePriceRequest) HasIsExclusive() bool`

HasIsExclusive returns a boolean if a field has been set.

### SetIsExclusiveNil

`func (o *CreatePriceRequest) SetIsExclusiveNil(b bool)`

 SetIsExclusiveNil sets the value for IsExclusive to be an explicit nil

### UnsetIsExclusive
`func (o *CreatePriceRequest) UnsetIsExclusive()`

UnsetIsExclusive ensures that no value is present for IsExclusive, not even an explicit nil
### GetListedExclusivelyForCustomers

`func (o *CreatePriceRequest) GetListedExclusivelyForCustomers() []string`

GetListedExclusivelyForCustomers returns the ListedExclusivelyForCustomers field if non-nil, zero value otherwise.

### GetListedExclusivelyForCustomersOk

`func (o *CreatePriceRequest) GetListedExclusivelyForCustomersOk() (*[]string, bool)`

GetListedExclusivelyForCustomersOk returns a tuple with the ListedExclusivelyForCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedExclusivelyForCustomers

`func (o *CreatePriceRequest) SetListedExclusivelyForCustomers(v []string)`

SetListedExclusivelyForCustomers sets ListedExclusivelyForCustomers field to given value.

### HasListedExclusivelyForCustomers

`func (o *CreatePriceRequest) HasListedExclusivelyForCustomers() bool`

HasListedExclusivelyForCustomers returns a boolean if a field has been set.

### GetIsActive

`func (o *CreatePriceRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreatePriceRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreatePriceRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetName

`func (o *CreatePriceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePriceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePriceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatePriceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreatePriceRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreatePriceRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetInternalDescription

`func (o *CreatePriceRequest) GetInternalDescription() string`

GetInternalDescription returns the InternalDescription field if non-nil, zero value otherwise.

### GetInternalDescriptionOk

`func (o *CreatePriceRequest) GetInternalDescriptionOk() (*string, bool)`

GetInternalDescriptionOk returns a tuple with the InternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDescription

`func (o *CreatePriceRequest) SetInternalDescription(v string)`

SetInternalDescription sets InternalDescription field to given value.

### HasInternalDescription

`func (o *CreatePriceRequest) HasInternalDescription() bool`

HasInternalDescription returns a boolean if a field has been set.

### SetInternalDescriptionNil

`func (o *CreatePriceRequest) SetInternalDescriptionNil(b bool)`

 SetInternalDescriptionNil sets the value for InternalDescription to be an explicit nil

### UnsetInternalDescription
`func (o *CreatePriceRequest) UnsetInternalDescription()`

UnsetInternalDescription ensures that no value is present for InternalDescription, not even an explicit nil
### GetPricingModel

`func (o *CreatePriceRequest) GetPricingModel() PricingModel`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *CreatePriceRequest) GetPricingModelOk() (*PricingModel, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *CreatePriceRequest) SetPricingModel(v PricingModel)`

SetPricingModel sets PricingModel field to given value.


### GetUnitAmountAtom

`func (o *CreatePriceRequest) GetUnitAmountAtom() int32`

GetUnitAmountAtom returns the UnitAmountAtom field if non-nil, zero value otherwise.

### GetUnitAmountAtomOk

`func (o *CreatePriceRequest) GetUnitAmountAtomOk() (*int32, bool)`

GetUnitAmountAtomOk returns a tuple with the UnitAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountAtom

`func (o *CreatePriceRequest) SetUnitAmountAtom(v int32)`

SetUnitAmountAtom sets UnitAmountAtom field to given value.


### SetUnitAmountAtomNil

`func (o *CreatePriceRequest) SetUnitAmountAtomNil(b bool)`

 SetUnitAmountAtomNil sets the value for UnitAmountAtom to be an explicit nil

### UnsetUnitAmountAtom
`func (o *CreatePriceRequest) UnsetUnitAmountAtom()`

UnsetUnitAmountAtom ensures that no value is present for UnitAmountAtom, not even an explicit nil
### GetTransformQuantityDivideBy

`func (o *CreatePriceRequest) GetTransformQuantityDivideBy() float32`

GetTransformQuantityDivideBy returns the TransformQuantityDivideBy field if non-nil, zero value otherwise.

### GetTransformQuantityDivideByOk

`func (o *CreatePriceRequest) GetTransformQuantityDivideByOk() (*float32, bool)`

GetTransformQuantityDivideByOk returns a tuple with the TransformQuantityDivideBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformQuantityDivideBy

`func (o *CreatePriceRequest) SetTransformQuantityDivideBy(v float32)`

SetTransformQuantityDivideBy sets TransformQuantityDivideBy field to given value.

### HasTransformQuantityDivideBy

`func (o *CreatePriceRequest) HasTransformQuantityDivideBy() bool`

HasTransformQuantityDivideBy returns a boolean if a field has been set.

### GetCurrency

`func (o *CreatePriceRequest) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreatePriceRequest) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreatePriceRequest) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreatePriceRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPriceTiers

`func (o *CreatePriceRequest) GetPriceTiers() []PriceTierParams`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *CreatePriceRequest) GetPriceTiersOk() (*[]PriceTierParams, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *CreatePriceRequest) SetPriceTiers(v []PriceTierParams)`

SetPriceTiers sets PriceTiers field to given value.

### HasPriceTiers

`func (o *CreatePriceRequest) HasPriceTiers() bool`

HasPriceTiers returns a boolean if a field has been set.

### GetPriceType

`func (o *CreatePriceRequest) GetPriceType() PriceTypeEnum`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *CreatePriceRequest) GetPriceTypeOk() (*PriceTypeEnum, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *CreatePriceRequest) SetPriceType(v PriceTypeEnum)`

SetPriceType sets PriceType field to given value.


### GetBillingInterval

`func (o *CreatePriceRequest) GetBillingInterval() CalendarIntervalEnum`

GetBillingInterval returns the BillingInterval field if non-nil, zero value otherwise.

### GetBillingIntervalOk

`func (o *CreatePriceRequest) GetBillingIntervalOk() (*CalendarIntervalEnum, bool)`

GetBillingIntervalOk returns a tuple with the BillingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInterval

`func (o *CreatePriceRequest) SetBillingInterval(v CalendarIntervalEnum)`

SetBillingInterval sets BillingInterval field to given value.

### HasBillingInterval

`func (o *CreatePriceRequest) HasBillingInterval() bool`

HasBillingInterval returns a boolean if a field has been set.

### SetBillingIntervalNil

`func (o *CreatePriceRequest) SetBillingIntervalNil(b bool)`

 SetBillingIntervalNil sets the value for BillingInterval to be an explicit nil

### UnsetBillingInterval
`func (o *CreatePriceRequest) UnsetBillingInterval()`

UnsetBillingInterval ensures that no value is present for BillingInterval, not even an explicit nil
### GetBillingIntervalCount

`func (o *CreatePriceRequest) GetBillingIntervalCount() int32`

GetBillingIntervalCount returns the BillingIntervalCount field if non-nil, zero value otherwise.

### GetBillingIntervalCountOk

`func (o *CreatePriceRequest) GetBillingIntervalCountOk() (*int32, bool)`

GetBillingIntervalCountOk returns a tuple with the BillingIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingIntervalCount

`func (o *CreatePriceRequest) SetBillingIntervalCount(v int32)`

SetBillingIntervalCount sets BillingIntervalCount field to given value.

### HasBillingIntervalCount

`func (o *CreatePriceRequest) HasBillingIntervalCount() bool`

HasBillingIntervalCount returns a boolean if a field has been set.

### SetBillingIntervalCountNil

`func (o *CreatePriceRequest) SetBillingIntervalCountNil(b bool)`

 SetBillingIntervalCountNil sets the value for BillingIntervalCount to be an explicit nil

### UnsetBillingIntervalCount
`func (o *CreatePriceRequest) UnsetBillingIntervalCount()`

UnsetBillingIntervalCount ensures that no value is present for BillingIntervalCount, not even an explicit nil
### GetContractTermMultiple

`func (o *CreatePriceRequest) GetContractTermMultiple() int32`

GetContractTermMultiple returns the ContractTermMultiple field if non-nil, zero value otherwise.

### GetContractTermMultipleOk

`func (o *CreatePriceRequest) GetContractTermMultipleOk() (*int32, bool)`

GetContractTermMultipleOk returns a tuple with the ContractTermMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTermMultiple

`func (o *CreatePriceRequest) SetContractTermMultiple(v int32)`

SetContractTermMultiple sets ContractTermMultiple field to given value.

### HasContractTermMultiple

`func (o *CreatePriceRequest) HasContractTermMultiple() bool`

HasContractTermMultiple returns a boolean if a field has been set.

### SetContractTermMultipleNil

`func (o *CreatePriceRequest) SetContractTermMultipleNil(b bool)`

 SetContractTermMultipleNil sets the value for ContractTermMultiple to be an explicit nil

### UnsetContractTermMultiple
`func (o *CreatePriceRequest) UnsetContractTermMultiple()`

UnsetContractTermMultiple ensures that no value is present for ContractTermMultiple, not even an explicit nil
### GetContractAutoRenew

`func (o *CreatePriceRequest) GetContractAutoRenew() bool`

GetContractAutoRenew returns the ContractAutoRenew field if non-nil, zero value otherwise.

### GetContractAutoRenewOk

`func (o *CreatePriceRequest) GetContractAutoRenewOk() (*bool, bool)`

GetContractAutoRenewOk returns a tuple with the ContractAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAutoRenew

`func (o *CreatePriceRequest) SetContractAutoRenew(v bool)`

SetContractAutoRenew sets ContractAutoRenew field to given value.

### HasContractAutoRenew

`func (o *CreatePriceRequest) HasContractAutoRenew() bool`

HasContractAutoRenew returns a boolean if a field has been set.

### SetContractAutoRenewNil

`func (o *CreatePriceRequest) SetContractAutoRenewNil(b bool)`

 SetContractAutoRenewNil sets the value for ContractAutoRenew to be an explicit nil

### UnsetContractAutoRenew
`func (o *CreatePriceRequest) UnsetContractAutoRenew()`

UnsetContractAutoRenew ensures that no value is present for ContractAutoRenew, not even an explicit nil
### GetTrialPeriodDays

`func (o *CreatePriceRequest) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *CreatePriceRequest) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *CreatePriceRequest) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *CreatePriceRequest) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### GetUsageType

`func (o *CreatePriceRequest) GetUsageType() UsageTypeEnum`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *CreatePriceRequest) GetUsageTypeOk() (*UsageTypeEnum, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *CreatePriceRequest) SetUsageType(v UsageTypeEnum)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *CreatePriceRequest) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetAggregateUsage

`func (o *CreatePriceRequest) GetAggregateUsage() UsageAggMethodEnum`

GetAggregateUsage returns the AggregateUsage field if non-nil, zero value otherwise.

### GetAggregateUsageOk

`func (o *CreatePriceRequest) GetAggregateUsageOk() (*UsageAggMethodEnum, bool)`

GetAggregateUsageOk returns a tuple with the AggregateUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateUsage

`func (o *CreatePriceRequest) SetAggregateUsage(v UsageAggMethodEnum)`

SetAggregateUsage sets AggregateUsage field to given value.

### HasAggregateUsage

`func (o *CreatePriceRequest) HasAggregateUsage() bool`

HasAggregateUsage returns a boolean if a field has been set.

### GetDefaultNetD

`func (o *CreatePriceRequest) GetDefaultNetD() int32`

GetDefaultNetD returns the DefaultNetD field if non-nil, zero value otherwise.

### GetDefaultNetDOk

`func (o *CreatePriceRequest) GetDefaultNetDOk() (*int32, bool)`

GetDefaultNetDOk returns a tuple with the DefaultNetD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetD

`func (o *CreatePriceRequest) SetDefaultNetD(v int32)`

SetDefaultNetD sets DefaultNetD field to given value.

### HasDefaultNetD

`func (o *CreatePriceRequest) HasDefaultNetD() bool`

HasDefaultNetD returns a boolean if a field has been set.

### SetDefaultNetDNil

`func (o *CreatePriceRequest) SetDefaultNetDNil(b bool)`

 SetDefaultNetDNil sets the value for DefaultNetD to be an explicit nil

### UnsetDefaultNetD
`func (o *CreatePriceRequest) UnsetDefaultNetD()`

UnsetDefaultNetD ensures that no value is present for DefaultNetD, not even an explicit nil
### GetCanOnlyBePurchasedWith

`func (o *CreatePriceRequest) GetCanOnlyBePurchasedWith() []string`

GetCanOnlyBePurchasedWith returns the CanOnlyBePurchasedWith field if non-nil, zero value otherwise.

### GetCanOnlyBePurchasedWithOk

`func (o *CreatePriceRequest) GetCanOnlyBePurchasedWithOk() (*[]string, bool)`

GetCanOnlyBePurchasedWithOk returns a tuple with the CanOnlyBePurchasedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanOnlyBePurchasedWith

`func (o *CreatePriceRequest) SetCanOnlyBePurchasedWith(v []string)`

SetCanOnlyBePurchasedWith sets CanOnlyBePurchasedWith field to given value.

### HasCanOnlyBePurchasedWith

`func (o *CreatePriceRequest) HasCanOnlyBePurchasedWith() bool`

HasCanOnlyBePurchasedWith returns a boolean if a field has been set.

### GetMeta

`func (o *CreatePriceRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreatePriceRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreatePriceRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreatePriceRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *CreatePriceRequest) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *CreatePriceRequest) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetProductId

`func (o *CreatePriceRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreatePriceRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreatePriceRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetIsDefault

`func (o *CreatePriceRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreatePriceRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreatePriceRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CreatePriceRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *CreatePriceRequest) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *CreatePriceRequest) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


