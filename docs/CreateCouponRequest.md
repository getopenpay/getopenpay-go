# CreateCouponRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the coupon displayed to customers on, for instance invoices, or receipts. | 
**AmountAtomOff** | Pointer to **NullableInt32** |  | [optional] 
**PercentOff** | Pointer to **NullableInt32** |  | [optional] 
**TrialDaysOff** | Pointer to **NullableInt32** |  | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) |  | [optional] 
**Duration** | Pointer to [**CouponDuration**](CouponDuration.md) |  | [optional] [default to COUPONDURATION_ONCE]
**DurationInMonths** | Pointer to **NullableInt32** |  | [optional] 
**ProductIds** | Pointer to **[]string** |  | [optional] 
**ProductFamilyIds** | Pointer to **[]string** |  | [optional] 
**MaxRedemptions** | Pointer to **NullableInt32** |  | [optional] 
**RedeemBy** | Pointer to **NullableTime** |  | [optional] 
**IsActive** | Pointer to **bool** | Whether the coupon is available to be redeemed. | [optional] [default to true]

## Methods

### NewCreateCouponRequest

`func NewCreateCouponRequest(name string, ) *CreateCouponRequest`

NewCreateCouponRequest instantiates a new CreateCouponRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCouponRequestWithDefaults

`func NewCreateCouponRequestWithDefaults() *CreateCouponRequest`

NewCreateCouponRequestWithDefaults instantiates a new CreateCouponRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCouponRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCouponRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCouponRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAmountAtomOff

`func (o *CreateCouponRequest) GetAmountAtomOff() int32`

GetAmountAtomOff returns the AmountAtomOff field if non-nil, zero value otherwise.

### GetAmountAtomOffOk

`func (o *CreateCouponRequest) GetAmountAtomOffOk() (*int32, bool)`

GetAmountAtomOffOk returns a tuple with the AmountAtomOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtomOff

`func (o *CreateCouponRequest) SetAmountAtomOff(v int32)`

SetAmountAtomOff sets AmountAtomOff field to given value.

### HasAmountAtomOff

`func (o *CreateCouponRequest) HasAmountAtomOff() bool`

HasAmountAtomOff returns a boolean if a field has been set.

### SetAmountAtomOffNil

`func (o *CreateCouponRequest) SetAmountAtomOffNil(b bool)`

 SetAmountAtomOffNil sets the value for AmountAtomOff to be an explicit nil

### UnsetAmountAtomOff
`func (o *CreateCouponRequest) UnsetAmountAtomOff()`

UnsetAmountAtomOff ensures that no value is present for AmountAtomOff, not even an explicit nil
### GetPercentOff

`func (o *CreateCouponRequest) GetPercentOff() int32`

GetPercentOff returns the PercentOff field if non-nil, zero value otherwise.

### GetPercentOffOk

`func (o *CreateCouponRequest) GetPercentOffOk() (*int32, bool)`

GetPercentOffOk returns a tuple with the PercentOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOff

`func (o *CreateCouponRequest) SetPercentOff(v int32)`

SetPercentOff sets PercentOff field to given value.

### HasPercentOff

`func (o *CreateCouponRequest) HasPercentOff() bool`

HasPercentOff returns a boolean if a field has been set.

### SetPercentOffNil

`func (o *CreateCouponRequest) SetPercentOffNil(b bool)`

 SetPercentOffNil sets the value for PercentOff to be an explicit nil

### UnsetPercentOff
`func (o *CreateCouponRequest) UnsetPercentOff()`

UnsetPercentOff ensures that no value is present for PercentOff, not even an explicit nil
### GetTrialDaysOff

`func (o *CreateCouponRequest) GetTrialDaysOff() int32`

GetTrialDaysOff returns the TrialDaysOff field if non-nil, zero value otherwise.

### GetTrialDaysOffOk

`func (o *CreateCouponRequest) GetTrialDaysOffOk() (*int32, bool)`

GetTrialDaysOffOk returns a tuple with the TrialDaysOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDaysOff

`func (o *CreateCouponRequest) SetTrialDaysOff(v int32)`

SetTrialDaysOff sets TrialDaysOff field to given value.

### HasTrialDaysOff

`func (o *CreateCouponRequest) HasTrialDaysOff() bool`

HasTrialDaysOff returns a boolean if a field has been set.

### SetTrialDaysOffNil

`func (o *CreateCouponRequest) SetTrialDaysOffNil(b bool)`

 SetTrialDaysOffNil sets the value for TrialDaysOff to be an explicit nil

### UnsetTrialDaysOff
`func (o *CreateCouponRequest) UnsetTrialDaysOff()`

UnsetTrialDaysOff ensures that no value is present for TrialDaysOff, not even an explicit nil
### GetCurrency

`func (o *CreateCouponRequest) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCouponRequest) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCouponRequest) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateCouponRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CreateCouponRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CreateCouponRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetDuration

`func (o *CreateCouponRequest) GetDuration() CouponDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CreateCouponRequest) GetDurationOk() (*CouponDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CreateCouponRequest) SetDuration(v CouponDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CreateCouponRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationInMonths

`func (o *CreateCouponRequest) GetDurationInMonths() int32`

GetDurationInMonths returns the DurationInMonths field if non-nil, zero value otherwise.

### GetDurationInMonthsOk

`func (o *CreateCouponRequest) GetDurationInMonthsOk() (*int32, bool)`

GetDurationInMonthsOk returns a tuple with the DurationInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMonths

`func (o *CreateCouponRequest) SetDurationInMonths(v int32)`

SetDurationInMonths sets DurationInMonths field to given value.

### HasDurationInMonths

`func (o *CreateCouponRequest) HasDurationInMonths() bool`

HasDurationInMonths returns a boolean if a field has been set.

### SetDurationInMonthsNil

`func (o *CreateCouponRequest) SetDurationInMonthsNil(b bool)`

 SetDurationInMonthsNil sets the value for DurationInMonths to be an explicit nil

### UnsetDurationInMonths
`func (o *CreateCouponRequest) UnsetDurationInMonths()`

UnsetDurationInMonths ensures that no value is present for DurationInMonths, not even an explicit nil
### GetProductIds

`func (o *CreateCouponRequest) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *CreateCouponRequest) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *CreateCouponRequest) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *CreateCouponRequest) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetProductFamilyIds

`func (o *CreateCouponRequest) GetProductFamilyIds() []string`

GetProductFamilyIds returns the ProductFamilyIds field if non-nil, zero value otherwise.

### GetProductFamilyIdsOk

`func (o *CreateCouponRequest) GetProductFamilyIdsOk() (*[]string, bool)`

GetProductFamilyIdsOk returns a tuple with the ProductFamilyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductFamilyIds

`func (o *CreateCouponRequest) SetProductFamilyIds(v []string)`

SetProductFamilyIds sets ProductFamilyIds field to given value.

### HasProductFamilyIds

`func (o *CreateCouponRequest) HasProductFamilyIds() bool`

HasProductFamilyIds returns a boolean if a field has been set.

### GetMaxRedemptions

`func (o *CreateCouponRequest) GetMaxRedemptions() int32`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *CreateCouponRequest) GetMaxRedemptionsOk() (*int32, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *CreateCouponRequest) SetMaxRedemptions(v int32)`

SetMaxRedemptions sets MaxRedemptions field to given value.

### HasMaxRedemptions

`func (o *CreateCouponRequest) HasMaxRedemptions() bool`

HasMaxRedemptions returns a boolean if a field has been set.

### SetMaxRedemptionsNil

`func (o *CreateCouponRequest) SetMaxRedemptionsNil(b bool)`

 SetMaxRedemptionsNil sets the value for MaxRedemptions to be an explicit nil

### UnsetMaxRedemptions
`func (o *CreateCouponRequest) UnsetMaxRedemptions()`

UnsetMaxRedemptions ensures that no value is present for MaxRedemptions, not even an explicit nil
### GetRedeemBy

`func (o *CreateCouponRequest) GetRedeemBy() time.Time`

GetRedeemBy returns the RedeemBy field if non-nil, zero value otherwise.

### GetRedeemByOk

`func (o *CreateCouponRequest) GetRedeemByOk() (*time.Time, bool)`

GetRedeemByOk returns a tuple with the RedeemBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemBy

`func (o *CreateCouponRequest) SetRedeemBy(v time.Time)`

SetRedeemBy sets RedeemBy field to given value.

### HasRedeemBy

`func (o *CreateCouponRequest) HasRedeemBy() bool`

HasRedeemBy returns a boolean if a field has been set.

### SetRedeemByNil

`func (o *CreateCouponRequest) SetRedeemByNil(b bool)`

 SetRedeemByNil sets the value for RedeemBy to be an explicit nil

### UnsetRedeemBy
`func (o *CreateCouponRequest) UnsetRedeemBy()`

UnsetRedeemBy ensures that no value is present for RedeemBy, not even an explicit nil
### GetIsActive

`func (o *CreateCouponRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateCouponRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateCouponRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CreateCouponRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


