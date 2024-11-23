# CreateSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | Unique Identifier of the customer. | 
**PaymentMethodId** | Pointer to **NullableString** |  | [optional] 
**SelectedProductPriceQuantity** | [**[]SelectedPriceQuantity**](SelectedPriceQuantity.md) |  | 
**TotalAmountAtom** | **int32** | Total amount of this subscription. It is in atomic units (in USD this is cents). | 
**CancelAtEnd** | Pointer to **bool** | Boolean indicating whether this subscription should cancel at the end of the current period. | [optional] [default to false]
**Description** | Pointer to **NullableString** |  | [optional] 
**TrialEnd** | Pointer to **NullableTime** |  | [optional] 
**TrialPeriodDays** | Pointer to **NullableInt32** |  | [optional] 
**TrialFromPrice** | Pointer to **bool** | Indicates if a price’s trial_period_days should be applied to the subscription. Setting trial_end per subscription is preferred, and this defaults to false. Setting this flag to true together with trial_end is not allowed. In case of subscription containing multiple prices and the trial period of them are not same, minimum of trial days will be used. | [optional] [default to true]
**CouponId** | Pointer to **NullableString** |  | [optional] 
**PromotionCode** | Pointer to **NullableString** |  | [optional] 
**CollectionMethod** | Pointer to [**CollectionMethodEnum**](CollectionMethodEnum.md) |  | [optional] [default to COLLECTIONMETHODENUM_CHARGE_AUTOMATICALLY]
**NetD** | Pointer to **NullableInt32** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**IsPreview** | Pointer to **bool** | Whether the request is in preview mode (subscriptions won&#39;t actually be created) | [optional] [default to false]

## Methods

### NewCreateSubscriptionRequest

`func NewCreateSubscriptionRequest(customerId string, selectedProductPriceQuantity []SelectedPriceQuantity, totalAmountAtom int32, ) *CreateSubscriptionRequest`

NewCreateSubscriptionRequest instantiates a new CreateSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionRequestWithDefaults

`func NewCreateSubscriptionRequestWithDefaults() *CreateSubscriptionRequest`

NewCreateSubscriptionRequestWithDefaults instantiates a new CreateSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CreateSubscriptionRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateSubscriptionRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateSubscriptionRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetPaymentMethodId

`func (o *CreateSubscriptionRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *CreateSubscriptionRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *CreateSubscriptionRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *CreateSubscriptionRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### SetPaymentMethodIdNil

`func (o *CreateSubscriptionRequest) SetPaymentMethodIdNil(b bool)`

 SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil

### UnsetPaymentMethodId
`func (o *CreateSubscriptionRequest) UnsetPaymentMethodId()`

UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
### GetSelectedProductPriceQuantity

`func (o *CreateSubscriptionRequest) GetSelectedProductPriceQuantity() []SelectedPriceQuantity`

GetSelectedProductPriceQuantity returns the SelectedProductPriceQuantity field if non-nil, zero value otherwise.

### GetSelectedProductPriceQuantityOk

`func (o *CreateSubscriptionRequest) GetSelectedProductPriceQuantityOk() (*[]SelectedPriceQuantity, bool)`

GetSelectedProductPriceQuantityOk returns a tuple with the SelectedProductPriceQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedProductPriceQuantity

`func (o *CreateSubscriptionRequest) SetSelectedProductPriceQuantity(v []SelectedPriceQuantity)`

SetSelectedProductPriceQuantity sets SelectedProductPriceQuantity field to given value.


### GetTotalAmountAtom

`func (o *CreateSubscriptionRequest) GetTotalAmountAtom() int32`

GetTotalAmountAtom returns the TotalAmountAtom field if non-nil, zero value otherwise.

### GetTotalAmountAtomOk

`func (o *CreateSubscriptionRequest) GetTotalAmountAtomOk() (*int32, bool)`

GetTotalAmountAtomOk returns a tuple with the TotalAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountAtom

`func (o *CreateSubscriptionRequest) SetTotalAmountAtom(v int32)`

SetTotalAmountAtom sets TotalAmountAtom field to given value.


### GetCancelAtEnd

`func (o *CreateSubscriptionRequest) GetCancelAtEnd() bool`

GetCancelAtEnd returns the CancelAtEnd field if non-nil, zero value otherwise.

### GetCancelAtEndOk

`func (o *CreateSubscriptionRequest) GetCancelAtEndOk() (*bool, bool)`

GetCancelAtEndOk returns a tuple with the CancelAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtEnd

`func (o *CreateSubscriptionRequest) SetCancelAtEnd(v bool)`

SetCancelAtEnd sets CancelAtEnd field to given value.

### HasCancelAtEnd

`func (o *CreateSubscriptionRequest) HasCancelAtEnd() bool`

HasCancelAtEnd returns a boolean if a field has been set.

### GetDescription

`func (o *CreateSubscriptionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSubscriptionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSubscriptionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSubscriptionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateSubscriptionRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateSubscriptionRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTrialEnd

`func (o *CreateSubscriptionRequest) GetTrialEnd() time.Time`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *CreateSubscriptionRequest) GetTrialEndOk() (*time.Time, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *CreateSubscriptionRequest) SetTrialEnd(v time.Time)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *CreateSubscriptionRequest) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### SetTrialEndNil

`func (o *CreateSubscriptionRequest) SetTrialEndNil(b bool)`

 SetTrialEndNil sets the value for TrialEnd to be an explicit nil

### UnsetTrialEnd
`func (o *CreateSubscriptionRequest) UnsetTrialEnd()`

UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
### GetTrialPeriodDays

`func (o *CreateSubscriptionRequest) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *CreateSubscriptionRequest) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *CreateSubscriptionRequest) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *CreateSubscriptionRequest) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### SetTrialPeriodDaysNil

`func (o *CreateSubscriptionRequest) SetTrialPeriodDaysNil(b bool)`

 SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil

### UnsetTrialPeriodDays
`func (o *CreateSubscriptionRequest) UnsetTrialPeriodDays()`

UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil
### GetTrialFromPrice

`func (o *CreateSubscriptionRequest) GetTrialFromPrice() bool`

GetTrialFromPrice returns the TrialFromPrice field if non-nil, zero value otherwise.

### GetTrialFromPriceOk

`func (o *CreateSubscriptionRequest) GetTrialFromPriceOk() (*bool, bool)`

GetTrialFromPriceOk returns a tuple with the TrialFromPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialFromPrice

`func (o *CreateSubscriptionRequest) SetTrialFromPrice(v bool)`

SetTrialFromPrice sets TrialFromPrice field to given value.

### HasTrialFromPrice

`func (o *CreateSubscriptionRequest) HasTrialFromPrice() bool`

HasTrialFromPrice returns a boolean if a field has been set.

### GetCouponId

`func (o *CreateSubscriptionRequest) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *CreateSubscriptionRequest) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *CreateSubscriptionRequest) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *CreateSubscriptionRequest) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponIdNil

`func (o *CreateSubscriptionRequest) SetCouponIdNil(b bool)`

 SetCouponIdNil sets the value for CouponId to be an explicit nil

### UnsetCouponId
`func (o *CreateSubscriptionRequest) UnsetCouponId()`

UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
### GetPromotionCode

`func (o *CreateSubscriptionRequest) GetPromotionCode() string`

GetPromotionCode returns the PromotionCode field if non-nil, zero value otherwise.

### GetPromotionCodeOk

`func (o *CreateSubscriptionRequest) GetPromotionCodeOk() (*string, bool)`

GetPromotionCodeOk returns a tuple with the PromotionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionCode

`func (o *CreateSubscriptionRequest) SetPromotionCode(v string)`

SetPromotionCode sets PromotionCode field to given value.

### HasPromotionCode

`func (o *CreateSubscriptionRequest) HasPromotionCode() bool`

HasPromotionCode returns a boolean if a field has been set.

### SetPromotionCodeNil

`func (o *CreateSubscriptionRequest) SetPromotionCodeNil(b bool)`

 SetPromotionCodeNil sets the value for PromotionCode to be an explicit nil

### UnsetPromotionCode
`func (o *CreateSubscriptionRequest) UnsetPromotionCode()`

UnsetPromotionCode ensures that no value is present for PromotionCode, not even an explicit nil
### GetCollectionMethod

`func (o *CreateSubscriptionRequest) GetCollectionMethod() CollectionMethodEnum`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *CreateSubscriptionRequest) GetCollectionMethodOk() (*CollectionMethodEnum, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *CreateSubscriptionRequest) SetCollectionMethod(v CollectionMethodEnum)`

SetCollectionMethod sets CollectionMethod field to given value.

### HasCollectionMethod

`func (o *CreateSubscriptionRequest) HasCollectionMethod() bool`

HasCollectionMethod returns a boolean if a field has been set.

### GetNetD

`func (o *CreateSubscriptionRequest) GetNetD() int32`

GetNetD returns the NetD field if non-nil, zero value otherwise.

### GetNetDOk

`func (o *CreateSubscriptionRequest) GetNetDOk() (*int32, bool)`

GetNetDOk returns a tuple with the NetD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetD

`func (o *CreateSubscriptionRequest) SetNetD(v int32)`

SetNetD sets NetD field to given value.

### HasNetD

`func (o *CreateSubscriptionRequest) HasNetD() bool`

HasNetD returns a boolean if a field has been set.

### SetNetDNil

`func (o *CreateSubscriptionRequest) SetNetDNil(b bool)`

 SetNetDNil sets the value for NetD to be an explicit nil

### UnsetNetD
`func (o *CreateSubscriptionRequest) UnsetNetD()`

UnsetNetD ensures that no value is present for NetD, not even an explicit nil
### GetCustomFields

`func (o *CreateSubscriptionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateSubscriptionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateSubscriptionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateSubscriptionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CreateSubscriptionRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CreateSubscriptionRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetIsPreview

`func (o *CreateSubscriptionRequest) GetIsPreview() bool`

GetIsPreview returns the IsPreview field if non-nil, zero value otherwise.

### GetIsPreviewOk

`func (o *CreateSubscriptionRequest) GetIsPreviewOk() (*bool, bool)`

GetIsPreviewOk returns a tuple with the IsPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreview

`func (o *CreateSubscriptionRequest) SetIsPreview(v bool)`

SetIsPreview sets IsPreview field to given value.

### HasIsPreview

`func (o *CreateSubscriptionRequest) HasIsPreview() bool`

HasIsPreview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


