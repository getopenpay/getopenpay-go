# UpdatePriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsExclusive** | Pointer to **NullableBool** |  | [optional] 
**ListedExclusivelyForCustomers** | Pointer to **[]string** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**InternalDescription** | Pointer to **NullableString** |  | [optional] 
**PricingModel** | Pointer to [**NullablePricingModel**](PricingModel.md) |  | [optional] 
**UnitAmountAtom** | Pointer to **NullableInt32** |  | [optional] 
**TransformQuantityDivideBy** | Pointer to **NullableFloat32** |  | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) |  | [optional] 
**PriceTiers** | Pointer to [**[]PriceTierParams**](PriceTierParams.md) |  | [optional] 
**PriceType** | Pointer to [**NullablePriceTypeEnum**](PriceTypeEnum.md) |  | [optional] 
**BillingInterval** | Pointer to [**NullableCalendarIntervalEnum**](CalendarIntervalEnum.md) |  | [optional] 
**BillingIntervalCount** | Pointer to **NullableInt32** |  | [optional] 
**ContractTermMultiple** | Pointer to **NullableInt32** |  | [optional] 
**ContractAutoRenew** | Pointer to **NullableBool** |  | [optional] 
**TrialPeriodDays** | Pointer to **NullableInt32** |  | [optional] 
**UsageType** | Pointer to [**NullableUsageTypeEnum**](UsageTypeEnum.md) |  | [optional] 
**AggregateUsage** | Pointer to [**NullableUsageAggMethodEnum**](UsageAggMethodEnum.md) |  | [optional] 
**DefaultNetD** | Pointer to **NullableInt32** |  | [optional] 
**CanOnlyBePurchasedWith** | Pointer to **[]string** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdatePriceRequest

`func NewUpdatePriceRequest() *UpdatePriceRequest`

NewUpdatePriceRequest instantiates a new UpdatePriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePriceRequestWithDefaults

`func NewUpdatePriceRequestWithDefaults() *UpdatePriceRequest`

NewUpdatePriceRequestWithDefaults instantiates a new UpdatePriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsExclusive

`func (o *UpdatePriceRequest) GetIsExclusive() bool`

GetIsExclusive returns the IsExclusive field if non-nil, zero value otherwise.

### GetIsExclusiveOk

`func (o *UpdatePriceRequest) GetIsExclusiveOk() (*bool, bool)`

GetIsExclusiveOk returns a tuple with the IsExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExclusive

`func (o *UpdatePriceRequest) SetIsExclusive(v bool)`

SetIsExclusive sets IsExclusive field to given value.

### HasIsExclusive

`func (o *UpdatePriceRequest) HasIsExclusive() bool`

HasIsExclusive returns a boolean if a field has been set.

### SetIsExclusiveNil

`func (o *UpdatePriceRequest) SetIsExclusiveNil(b bool)`

 SetIsExclusiveNil sets the value for IsExclusive to be an explicit nil

### UnsetIsExclusive
`func (o *UpdatePriceRequest) UnsetIsExclusive()`

UnsetIsExclusive ensures that no value is present for IsExclusive, not even an explicit nil
### GetListedExclusivelyForCustomers

`func (o *UpdatePriceRequest) GetListedExclusivelyForCustomers() []string`

GetListedExclusivelyForCustomers returns the ListedExclusivelyForCustomers field if non-nil, zero value otherwise.

### GetListedExclusivelyForCustomersOk

`func (o *UpdatePriceRequest) GetListedExclusivelyForCustomersOk() (*[]string, bool)`

GetListedExclusivelyForCustomersOk returns a tuple with the ListedExclusivelyForCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedExclusivelyForCustomers

`func (o *UpdatePriceRequest) SetListedExclusivelyForCustomers(v []string)`

SetListedExclusivelyForCustomers sets ListedExclusivelyForCustomers field to given value.

### HasListedExclusivelyForCustomers

`func (o *UpdatePriceRequest) HasListedExclusivelyForCustomers() bool`

HasListedExclusivelyForCustomers returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdatePriceRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdatePriceRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdatePriceRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdatePriceRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UpdatePriceRequest) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UpdatePriceRequest) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetName

`func (o *UpdatePriceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePriceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePriceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePriceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdatePriceRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdatePriceRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetInternalDescription

`func (o *UpdatePriceRequest) GetInternalDescription() string`

GetInternalDescription returns the InternalDescription field if non-nil, zero value otherwise.

### GetInternalDescriptionOk

`func (o *UpdatePriceRequest) GetInternalDescriptionOk() (*string, bool)`

GetInternalDescriptionOk returns a tuple with the InternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDescription

`func (o *UpdatePriceRequest) SetInternalDescription(v string)`

SetInternalDescription sets InternalDescription field to given value.

### HasInternalDescription

`func (o *UpdatePriceRequest) HasInternalDescription() bool`

HasInternalDescription returns a boolean if a field has been set.

### SetInternalDescriptionNil

`func (o *UpdatePriceRequest) SetInternalDescriptionNil(b bool)`

 SetInternalDescriptionNil sets the value for InternalDescription to be an explicit nil

### UnsetInternalDescription
`func (o *UpdatePriceRequest) UnsetInternalDescription()`

UnsetInternalDescription ensures that no value is present for InternalDescription, not even an explicit nil
### GetPricingModel

`func (o *UpdatePriceRequest) GetPricingModel() PricingModel`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *UpdatePriceRequest) GetPricingModelOk() (*PricingModel, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *UpdatePriceRequest) SetPricingModel(v PricingModel)`

SetPricingModel sets PricingModel field to given value.

### HasPricingModel

`func (o *UpdatePriceRequest) HasPricingModel() bool`

HasPricingModel returns a boolean if a field has been set.

### SetPricingModelNil

`func (o *UpdatePriceRequest) SetPricingModelNil(b bool)`

 SetPricingModelNil sets the value for PricingModel to be an explicit nil

### UnsetPricingModel
`func (o *UpdatePriceRequest) UnsetPricingModel()`

UnsetPricingModel ensures that no value is present for PricingModel, not even an explicit nil
### GetUnitAmountAtom

`func (o *UpdatePriceRequest) GetUnitAmountAtom() int32`

GetUnitAmountAtom returns the UnitAmountAtom field if non-nil, zero value otherwise.

### GetUnitAmountAtomOk

`func (o *UpdatePriceRequest) GetUnitAmountAtomOk() (*int32, bool)`

GetUnitAmountAtomOk returns a tuple with the UnitAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountAtom

`func (o *UpdatePriceRequest) SetUnitAmountAtom(v int32)`

SetUnitAmountAtom sets UnitAmountAtom field to given value.

### HasUnitAmountAtom

`func (o *UpdatePriceRequest) HasUnitAmountAtom() bool`

HasUnitAmountAtom returns a boolean if a field has been set.

### SetUnitAmountAtomNil

`func (o *UpdatePriceRequest) SetUnitAmountAtomNil(b bool)`

 SetUnitAmountAtomNil sets the value for UnitAmountAtom to be an explicit nil

### UnsetUnitAmountAtom
`func (o *UpdatePriceRequest) UnsetUnitAmountAtom()`

UnsetUnitAmountAtom ensures that no value is present for UnitAmountAtom, not even an explicit nil
### GetTransformQuantityDivideBy

`func (o *UpdatePriceRequest) GetTransformQuantityDivideBy() float32`

GetTransformQuantityDivideBy returns the TransformQuantityDivideBy field if non-nil, zero value otherwise.

### GetTransformQuantityDivideByOk

`func (o *UpdatePriceRequest) GetTransformQuantityDivideByOk() (*float32, bool)`

GetTransformQuantityDivideByOk returns a tuple with the TransformQuantityDivideBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformQuantityDivideBy

`func (o *UpdatePriceRequest) SetTransformQuantityDivideBy(v float32)`

SetTransformQuantityDivideBy sets TransformQuantityDivideBy field to given value.

### HasTransformQuantityDivideBy

`func (o *UpdatePriceRequest) HasTransformQuantityDivideBy() bool`

HasTransformQuantityDivideBy returns a boolean if a field has been set.

### SetTransformQuantityDivideByNil

`func (o *UpdatePriceRequest) SetTransformQuantityDivideByNil(b bool)`

 SetTransformQuantityDivideByNil sets the value for TransformQuantityDivideBy to be an explicit nil

### UnsetTransformQuantityDivideBy
`func (o *UpdatePriceRequest) UnsetTransformQuantityDivideBy()`

UnsetTransformQuantityDivideBy ensures that no value is present for TransformQuantityDivideBy, not even an explicit nil
### GetCurrency

`func (o *UpdatePriceRequest) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdatePriceRequest) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdatePriceRequest) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdatePriceRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *UpdatePriceRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *UpdatePriceRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetPriceTiers

`func (o *UpdatePriceRequest) GetPriceTiers() []PriceTierParams`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *UpdatePriceRequest) GetPriceTiersOk() (*[]PriceTierParams, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *UpdatePriceRequest) SetPriceTiers(v []PriceTierParams)`

SetPriceTiers sets PriceTiers field to given value.

### HasPriceTiers

`func (o *UpdatePriceRequest) HasPriceTiers() bool`

HasPriceTiers returns a boolean if a field has been set.

### GetPriceType

`func (o *UpdatePriceRequest) GetPriceType() PriceTypeEnum`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *UpdatePriceRequest) GetPriceTypeOk() (*PriceTypeEnum, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *UpdatePriceRequest) SetPriceType(v PriceTypeEnum)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *UpdatePriceRequest) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### SetPriceTypeNil

`func (o *UpdatePriceRequest) SetPriceTypeNil(b bool)`

 SetPriceTypeNil sets the value for PriceType to be an explicit nil

### UnsetPriceType
`func (o *UpdatePriceRequest) UnsetPriceType()`

UnsetPriceType ensures that no value is present for PriceType, not even an explicit nil
### GetBillingInterval

`func (o *UpdatePriceRequest) GetBillingInterval() CalendarIntervalEnum`

GetBillingInterval returns the BillingInterval field if non-nil, zero value otherwise.

### GetBillingIntervalOk

`func (o *UpdatePriceRequest) GetBillingIntervalOk() (*CalendarIntervalEnum, bool)`

GetBillingIntervalOk returns a tuple with the BillingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInterval

`func (o *UpdatePriceRequest) SetBillingInterval(v CalendarIntervalEnum)`

SetBillingInterval sets BillingInterval field to given value.

### HasBillingInterval

`func (o *UpdatePriceRequest) HasBillingInterval() bool`

HasBillingInterval returns a boolean if a field has been set.

### SetBillingIntervalNil

`func (o *UpdatePriceRequest) SetBillingIntervalNil(b bool)`

 SetBillingIntervalNil sets the value for BillingInterval to be an explicit nil

### UnsetBillingInterval
`func (o *UpdatePriceRequest) UnsetBillingInterval()`

UnsetBillingInterval ensures that no value is present for BillingInterval, not even an explicit nil
### GetBillingIntervalCount

`func (o *UpdatePriceRequest) GetBillingIntervalCount() int32`

GetBillingIntervalCount returns the BillingIntervalCount field if non-nil, zero value otherwise.

### GetBillingIntervalCountOk

`func (o *UpdatePriceRequest) GetBillingIntervalCountOk() (*int32, bool)`

GetBillingIntervalCountOk returns a tuple with the BillingIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingIntervalCount

`func (o *UpdatePriceRequest) SetBillingIntervalCount(v int32)`

SetBillingIntervalCount sets BillingIntervalCount field to given value.

### HasBillingIntervalCount

`func (o *UpdatePriceRequest) HasBillingIntervalCount() bool`

HasBillingIntervalCount returns a boolean if a field has been set.

### SetBillingIntervalCountNil

`func (o *UpdatePriceRequest) SetBillingIntervalCountNil(b bool)`

 SetBillingIntervalCountNil sets the value for BillingIntervalCount to be an explicit nil

### UnsetBillingIntervalCount
`func (o *UpdatePriceRequest) UnsetBillingIntervalCount()`

UnsetBillingIntervalCount ensures that no value is present for BillingIntervalCount, not even an explicit nil
### GetContractTermMultiple

`func (o *UpdatePriceRequest) GetContractTermMultiple() int32`

GetContractTermMultiple returns the ContractTermMultiple field if non-nil, zero value otherwise.

### GetContractTermMultipleOk

`func (o *UpdatePriceRequest) GetContractTermMultipleOk() (*int32, bool)`

GetContractTermMultipleOk returns a tuple with the ContractTermMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTermMultiple

`func (o *UpdatePriceRequest) SetContractTermMultiple(v int32)`

SetContractTermMultiple sets ContractTermMultiple field to given value.

### HasContractTermMultiple

`func (o *UpdatePriceRequest) HasContractTermMultiple() bool`

HasContractTermMultiple returns a boolean if a field has been set.

### SetContractTermMultipleNil

`func (o *UpdatePriceRequest) SetContractTermMultipleNil(b bool)`

 SetContractTermMultipleNil sets the value for ContractTermMultiple to be an explicit nil

### UnsetContractTermMultiple
`func (o *UpdatePriceRequest) UnsetContractTermMultiple()`

UnsetContractTermMultiple ensures that no value is present for ContractTermMultiple, not even an explicit nil
### GetContractAutoRenew

`func (o *UpdatePriceRequest) GetContractAutoRenew() bool`

GetContractAutoRenew returns the ContractAutoRenew field if non-nil, zero value otherwise.

### GetContractAutoRenewOk

`func (o *UpdatePriceRequest) GetContractAutoRenewOk() (*bool, bool)`

GetContractAutoRenewOk returns a tuple with the ContractAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAutoRenew

`func (o *UpdatePriceRequest) SetContractAutoRenew(v bool)`

SetContractAutoRenew sets ContractAutoRenew field to given value.

### HasContractAutoRenew

`func (o *UpdatePriceRequest) HasContractAutoRenew() bool`

HasContractAutoRenew returns a boolean if a field has been set.

### SetContractAutoRenewNil

`func (o *UpdatePriceRequest) SetContractAutoRenewNil(b bool)`

 SetContractAutoRenewNil sets the value for ContractAutoRenew to be an explicit nil

### UnsetContractAutoRenew
`func (o *UpdatePriceRequest) UnsetContractAutoRenew()`

UnsetContractAutoRenew ensures that no value is present for ContractAutoRenew, not even an explicit nil
### GetTrialPeriodDays

`func (o *UpdatePriceRequest) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *UpdatePriceRequest) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *UpdatePriceRequest) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *UpdatePriceRequest) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### SetTrialPeriodDaysNil

`func (o *UpdatePriceRequest) SetTrialPeriodDaysNil(b bool)`

 SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil

### UnsetTrialPeriodDays
`func (o *UpdatePriceRequest) UnsetTrialPeriodDays()`

UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil
### GetUsageType

`func (o *UpdatePriceRequest) GetUsageType() UsageTypeEnum`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *UpdatePriceRequest) GetUsageTypeOk() (*UsageTypeEnum, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *UpdatePriceRequest) SetUsageType(v UsageTypeEnum)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *UpdatePriceRequest) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *UpdatePriceRequest) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *UpdatePriceRequest) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil
### GetAggregateUsage

`func (o *UpdatePriceRequest) GetAggregateUsage() UsageAggMethodEnum`

GetAggregateUsage returns the AggregateUsage field if non-nil, zero value otherwise.

### GetAggregateUsageOk

`func (o *UpdatePriceRequest) GetAggregateUsageOk() (*UsageAggMethodEnum, bool)`

GetAggregateUsageOk returns a tuple with the AggregateUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateUsage

`func (o *UpdatePriceRequest) SetAggregateUsage(v UsageAggMethodEnum)`

SetAggregateUsage sets AggregateUsage field to given value.

### HasAggregateUsage

`func (o *UpdatePriceRequest) HasAggregateUsage() bool`

HasAggregateUsage returns a boolean if a field has been set.

### SetAggregateUsageNil

`func (o *UpdatePriceRequest) SetAggregateUsageNil(b bool)`

 SetAggregateUsageNil sets the value for AggregateUsage to be an explicit nil

### UnsetAggregateUsage
`func (o *UpdatePriceRequest) UnsetAggregateUsage()`

UnsetAggregateUsage ensures that no value is present for AggregateUsage, not even an explicit nil
### GetDefaultNetD

`func (o *UpdatePriceRequest) GetDefaultNetD() int32`

GetDefaultNetD returns the DefaultNetD field if non-nil, zero value otherwise.

### GetDefaultNetDOk

`func (o *UpdatePriceRequest) GetDefaultNetDOk() (*int32, bool)`

GetDefaultNetDOk returns a tuple with the DefaultNetD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetD

`func (o *UpdatePriceRequest) SetDefaultNetD(v int32)`

SetDefaultNetD sets DefaultNetD field to given value.

### HasDefaultNetD

`func (o *UpdatePriceRequest) HasDefaultNetD() bool`

HasDefaultNetD returns a boolean if a field has been set.

### SetDefaultNetDNil

`func (o *UpdatePriceRequest) SetDefaultNetDNil(b bool)`

 SetDefaultNetDNil sets the value for DefaultNetD to be an explicit nil

### UnsetDefaultNetD
`func (o *UpdatePriceRequest) UnsetDefaultNetD()`

UnsetDefaultNetD ensures that no value is present for DefaultNetD, not even an explicit nil
### GetCanOnlyBePurchasedWith

`func (o *UpdatePriceRequest) GetCanOnlyBePurchasedWith() []string`

GetCanOnlyBePurchasedWith returns the CanOnlyBePurchasedWith field if non-nil, zero value otherwise.

### GetCanOnlyBePurchasedWithOk

`func (o *UpdatePriceRequest) GetCanOnlyBePurchasedWithOk() (*[]string, bool)`

GetCanOnlyBePurchasedWithOk returns a tuple with the CanOnlyBePurchasedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanOnlyBePurchasedWith

`func (o *UpdatePriceRequest) SetCanOnlyBePurchasedWith(v []string)`

SetCanOnlyBePurchasedWith sets CanOnlyBePurchasedWith field to given value.

### HasCanOnlyBePurchasedWith

`func (o *UpdatePriceRequest) HasCanOnlyBePurchasedWith() bool`

HasCanOnlyBePurchasedWith returns a boolean if a field has been set.

### GetMeta

`func (o *UpdatePriceRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UpdatePriceRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UpdatePriceRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UpdatePriceRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *UpdatePriceRequest) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *UpdatePriceRequest) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


