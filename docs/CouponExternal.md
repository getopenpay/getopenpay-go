# CouponExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_COUPON]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Name** | **string** |  | 
**IsActive** | **bool** |  | 
**AmountAtomOff** | **NullableInt32** |  | 
**PercentOff** | **NullableInt32** |  | 
**TrialDaysOff** | Pointer to **NullableInt32** |  | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) |  | [optional] 
**Duration** | [**CouponDuration**](CouponDuration.md) |  | 
**DurationInMonths** | **NullableInt32** |  | 
**MaxRedemptions** | **NullableInt32** |  | 
**TimesRedeemed** | **NullableInt32** |  | 
**RedeemBy** | **NullableTime** |  | 
**Products** | Pointer to **[]string** |  | [optional] 
**ProductFamilies** | Pointer to **[]string** |  | [optional] 
**PromotionCodes** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCouponExternal

`func NewCouponExternal(id string, createdAt time.Time, updatedAt time.Time, name string, isActive bool, amountAtomOff NullableInt32, percentOff NullableInt32, duration CouponDuration, durationInMonths NullableInt32, maxRedemptions NullableInt32, timesRedeemed NullableInt32, redeemBy NullableTime, ) *CouponExternal`

NewCouponExternal instantiates a new CouponExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponExternalWithDefaults

`func NewCouponExternalWithDefaults() *CouponExternal`

NewCouponExternalWithDefaults instantiates a new CouponExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CouponExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CouponExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CouponExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CouponExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *CouponExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CouponExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CouponExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CouponExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CouponExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CouponExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CouponExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *CouponExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CouponExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CouponExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CouponExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetName

`func (o *CouponExternal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CouponExternal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CouponExternal) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *CouponExternal) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CouponExternal) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CouponExternal) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetAmountAtomOff

`func (o *CouponExternal) GetAmountAtomOff() int32`

GetAmountAtomOff returns the AmountAtomOff field if non-nil, zero value otherwise.

### GetAmountAtomOffOk

`func (o *CouponExternal) GetAmountAtomOffOk() (*int32, bool)`

GetAmountAtomOffOk returns a tuple with the AmountAtomOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtomOff

`func (o *CouponExternal) SetAmountAtomOff(v int32)`

SetAmountAtomOff sets AmountAtomOff field to given value.


### SetAmountAtomOffNil

`func (o *CouponExternal) SetAmountAtomOffNil(b bool)`

 SetAmountAtomOffNil sets the value for AmountAtomOff to be an explicit nil

### UnsetAmountAtomOff
`func (o *CouponExternal) UnsetAmountAtomOff()`

UnsetAmountAtomOff ensures that no value is present for AmountAtomOff, not even an explicit nil
### GetPercentOff

`func (o *CouponExternal) GetPercentOff() int32`

GetPercentOff returns the PercentOff field if non-nil, zero value otherwise.

### GetPercentOffOk

`func (o *CouponExternal) GetPercentOffOk() (*int32, bool)`

GetPercentOffOk returns a tuple with the PercentOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOff

`func (o *CouponExternal) SetPercentOff(v int32)`

SetPercentOff sets PercentOff field to given value.


### SetPercentOffNil

`func (o *CouponExternal) SetPercentOffNil(b bool)`

 SetPercentOffNil sets the value for PercentOff to be an explicit nil

### UnsetPercentOff
`func (o *CouponExternal) UnsetPercentOff()`

UnsetPercentOff ensures that no value is present for PercentOff, not even an explicit nil
### GetTrialDaysOff

`func (o *CouponExternal) GetTrialDaysOff() int32`

GetTrialDaysOff returns the TrialDaysOff field if non-nil, zero value otherwise.

### GetTrialDaysOffOk

`func (o *CouponExternal) GetTrialDaysOffOk() (*int32, bool)`

GetTrialDaysOffOk returns a tuple with the TrialDaysOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDaysOff

`func (o *CouponExternal) SetTrialDaysOff(v int32)`

SetTrialDaysOff sets TrialDaysOff field to given value.

### HasTrialDaysOff

`func (o *CouponExternal) HasTrialDaysOff() bool`

HasTrialDaysOff returns a boolean if a field has been set.

### SetTrialDaysOffNil

`func (o *CouponExternal) SetTrialDaysOffNil(b bool)`

 SetTrialDaysOffNil sets the value for TrialDaysOff to be an explicit nil

### UnsetTrialDaysOff
`func (o *CouponExternal) UnsetTrialDaysOff()`

UnsetTrialDaysOff ensures that no value is present for TrialDaysOff, not even an explicit nil
### GetCurrency

`func (o *CouponExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CouponExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CouponExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CouponExternal) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CouponExternal) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CouponExternal) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetDuration

`func (o *CouponExternal) GetDuration() CouponDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CouponExternal) GetDurationOk() (*CouponDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CouponExternal) SetDuration(v CouponDuration)`

SetDuration sets Duration field to given value.


### GetDurationInMonths

`func (o *CouponExternal) GetDurationInMonths() int32`

GetDurationInMonths returns the DurationInMonths field if non-nil, zero value otherwise.

### GetDurationInMonthsOk

`func (o *CouponExternal) GetDurationInMonthsOk() (*int32, bool)`

GetDurationInMonthsOk returns a tuple with the DurationInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMonths

`func (o *CouponExternal) SetDurationInMonths(v int32)`

SetDurationInMonths sets DurationInMonths field to given value.


### SetDurationInMonthsNil

`func (o *CouponExternal) SetDurationInMonthsNil(b bool)`

 SetDurationInMonthsNil sets the value for DurationInMonths to be an explicit nil

### UnsetDurationInMonths
`func (o *CouponExternal) UnsetDurationInMonths()`

UnsetDurationInMonths ensures that no value is present for DurationInMonths, not even an explicit nil
### GetMaxRedemptions

`func (o *CouponExternal) GetMaxRedemptions() int32`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *CouponExternal) GetMaxRedemptionsOk() (*int32, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *CouponExternal) SetMaxRedemptions(v int32)`

SetMaxRedemptions sets MaxRedemptions field to given value.


### SetMaxRedemptionsNil

`func (o *CouponExternal) SetMaxRedemptionsNil(b bool)`

 SetMaxRedemptionsNil sets the value for MaxRedemptions to be an explicit nil

### UnsetMaxRedemptions
`func (o *CouponExternal) UnsetMaxRedemptions()`

UnsetMaxRedemptions ensures that no value is present for MaxRedemptions, not even an explicit nil
### GetTimesRedeemed

`func (o *CouponExternal) GetTimesRedeemed() int32`

GetTimesRedeemed returns the TimesRedeemed field if non-nil, zero value otherwise.

### GetTimesRedeemedOk

`func (o *CouponExternal) GetTimesRedeemedOk() (*int32, bool)`

GetTimesRedeemedOk returns a tuple with the TimesRedeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesRedeemed

`func (o *CouponExternal) SetTimesRedeemed(v int32)`

SetTimesRedeemed sets TimesRedeemed field to given value.


### SetTimesRedeemedNil

`func (o *CouponExternal) SetTimesRedeemedNil(b bool)`

 SetTimesRedeemedNil sets the value for TimesRedeemed to be an explicit nil

### UnsetTimesRedeemed
`func (o *CouponExternal) UnsetTimesRedeemed()`

UnsetTimesRedeemed ensures that no value is present for TimesRedeemed, not even an explicit nil
### GetRedeemBy

`func (o *CouponExternal) GetRedeemBy() time.Time`

GetRedeemBy returns the RedeemBy field if non-nil, zero value otherwise.

### GetRedeemByOk

`func (o *CouponExternal) GetRedeemByOk() (*time.Time, bool)`

GetRedeemByOk returns a tuple with the RedeemBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemBy

`func (o *CouponExternal) SetRedeemBy(v time.Time)`

SetRedeemBy sets RedeemBy field to given value.


### SetRedeemByNil

`func (o *CouponExternal) SetRedeemByNil(b bool)`

 SetRedeemByNil sets the value for RedeemBy to be an explicit nil

### UnsetRedeemBy
`func (o *CouponExternal) UnsetRedeemBy()`

UnsetRedeemBy ensures that no value is present for RedeemBy, not even an explicit nil
### GetProducts

`func (o *CouponExternal) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CouponExternal) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CouponExternal) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *CouponExternal) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetProductFamilies

`func (o *CouponExternal) GetProductFamilies() []string`

GetProductFamilies returns the ProductFamilies field if non-nil, zero value otherwise.

### GetProductFamiliesOk

`func (o *CouponExternal) GetProductFamiliesOk() (*[]string, bool)`

GetProductFamiliesOk returns a tuple with the ProductFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductFamilies

`func (o *CouponExternal) SetProductFamilies(v []string)`

SetProductFamilies sets ProductFamilies field to given value.

### HasProductFamilies

`func (o *CouponExternal) HasProductFamilies() bool`

HasProductFamilies returns a boolean if a field has been set.

### GetPromotionCodes

`func (o *CouponExternal) GetPromotionCodes() []string`

GetPromotionCodes returns the PromotionCodes field if non-nil, zero value otherwise.

### GetPromotionCodesOk

`func (o *CouponExternal) GetPromotionCodesOk() (*[]string, bool)`

GetPromotionCodesOk returns a tuple with the PromotionCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionCodes

`func (o *CouponExternal) SetPromotionCodes(v []string)`

SetPromotionCodes sets PromotionCodes field to given value.

### HasPromotionCodes

`func (o *CouponExternal) HasPromotionCodes() bool`

HasPromotionCodes returns a boolean if a field has been set.

### GetMetadata

`func (o *CouponExternal) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CouponExternal) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CouponExternal) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CouponExternal) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CouponExternal) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CouponExternal) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


