# UpdateSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodId** | Pointer to **NullableString** |  | [optional] 
**CancelAtEnd** | Pointer to **NullableBool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProrationBehavior** | Pointer to [**ProrationEnum**](ProrationEnum.md) |  | [optional] [default to PRORATIONENUM_ALWAYS_INVOICE]
**TrialEnd** | Pointer to **NullableTime** |  | [optional] 
**TrialPeriodDays** | Pointer to **NullableInt32** |  | [optional] 
**CouponId** | Pointer to **NullableString** |  | [optional] 
**PromotionCode** | Pointer to **NullableString** |  | [optional] 
**CollectionMethod** | Pointer to [**NullableCollectionMethodEnum**](CollectionMethodEnum.md) |  | [optional] 
**NetD** | Pointer to **NullableInt32** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**NewPeriodEnd** | Pointer to **NullableTime** |  | [optional] 
**IsPreview** | Pointer to **bool** | Whether the request is in preview mode (subscriptions won&#39;t actually be updated) | [optional] [default to false]
**PreviewRenewalInvoices** | Pointer to **bool** | Whether to include the preview of the renewal invoices for the subscriptions in the response. | [optional] [default to false]
**Items** | Pointer to [**[]InlineSubscriptionItemUpdate**](InlineSubscriptionItemUpdate.md) |  | [optional] 

## Methods

### NewUpdateSubscriptionRequest

`func NewUpdateSubscriptionRequest() *UpdateSubscriptionRequest`

NewUpdateSubscriptionRequest instantiates a new UpdateSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptionRequestWithDefaults

`func NewUpdateSubscriptionRequestWithDefaults() *UpdateSubscriptionRequest`

NewUpdateSubscriptionRequestWithDefaults instantiates a new UpdateSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodId

`func (o *UpdateSubscriptionRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *UpdateSubscriptionRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *UpdateSubscriptionRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *UpdateSubscriptionRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### SetPaymentMethodIdNil

`func (o *UpdateSubscriptionRequest) SetPaymentMethodIdNil(b bool)`

 SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil

### UnsetPaymentMethodId
`func (o *UpdateSubscriptionRequest) UnsetPaymentMethodId()`

UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
### GetCancelAtEnd

`func (o *UpdateSubscriptionRequest) GetCancelAtEnd() bool`

GetCancelAtEnd returns the CancelAtEnd field if non-nil, zero value otherwise.

### GetCancelAtEndOk

`func (o *UpdateSubscriptionRequest) GetCancelAtEndOk() (*bool, bool)`

GetCancelAtEndOk returns a tuple with the CancelAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtEnd

`func (o *UpdateSubscriptionRequest) SetCancelAtEnd(v bool)`

SetCancelAtEnd sets CancelAtEnd field to given value.

### HasCancelAtEnd

`func (o *UpdateSubscriptionRequest) HasCancelAtEnd() bool`

HasCancelAtEnd returns a boolean if a field has been set.

### SetCancelAtEndNil

`func (o *UpdateSubscriptionRequest) SetCancelAtEndNil(b bool)`

 SetCancelAtEndNil sets the value for CancelAtEnd to be an explicit nil

### UnsetCancelAtEnd
`func (o *UpdateSubscriptionRequest) UnsetCancelAtEnd()`

UnsetCancelAtEnd ensures that no value is present for CancelAtEnd, not even an explicit nil
### GetDescription

`func (o *UpdateSubscriptionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSubscriptionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSubscriptionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSubscriptionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateSubscriptionRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateSubscriptionRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProrationBehavior

`func (o *UpdateSubscriptionRequest) GetProrationBehavior() ProrationEnum`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *UpdateSubscriptionRequest) GetProrationBehaviorOk() (*ProrationEnum, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *UpdateSubscriptionRequest) SetProrationBehavior(v ProrationEnum)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *UpdateSubscriptionRequest) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.

### GetTrialEnd

`func (o *UpdateSubscriptionRequest) GetTrialEnd() time.Time`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *UpdateSubscriptionRequest) GetTrialEndOk() (*time.Time, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *UpdateSubscriptionRequest) SetTrialEnd(v time.Time)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *UpdateSubscriptionRequest) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### SetTrialEndNil

`func (o *UpdateSubscriptionRequest) SetTrialEndNil(b bool)`

 SetTrialEndNil sets the value for TrialEnd to be an explicit nil

### UnsetTrialEnd
`func (o *UpdateSubscriptionRequest) UnsetTrialEnd()`

UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
### GetTrialPeriodDays

`func (o *UpdateSubscriptionRequest) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *UpdateSubscriptionRequest) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *UpdateSubscriptionRequest) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *UpdateSubscriptionRequest) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### SetTrialPeriodDaysNil

`func (o *UpdateSubscriptionRequest) SetTrialPeriodDaysNil(b bool)`

 SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil

### UnsetTrialPeriodDays
`func (o *UpdateSubscriptionRequest) UnsetTrialPeriodDays()`

UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil
### GetCouponId

`func (o *UpdateSubscriptionRequest) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *UpdateSubscriptionRequest) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *UpdateSubscriptionRequest) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *UpdateSubscriptionRequest) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponIdNil

`func (o *UpdateSubscriptionRequest) SetCouponIdNil(b bool)`

 SetCouponIdNil sets the value for CouponId to be an explicit nil

### UnsetCouponId
`func (o *UpdateSubscriptionRequest) UnsetCouponId()`

UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
### GetPromotionCode

`func (o *UpdateSubscriptionRequest) GetPromotionCode() string`

GetPromotionCode returns the PromotionCode field if non-nil, zero value otherwise.

### GetPromotionCodeOk

`func (o *UpdateSubscriptionRequest) GetPromotionCodeOk() (*string, bool)`

GetPromotionCodeOk returns a tuple with the PromotionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionCode

`func (o *UpdateSubscriptionRequest) SetPromotionCode(v string)`

SetPromotionCode sets PromotionCode field to given value.

### HasPromotionCode

`func (o *UpdateSubscriptionRequest) HasPromotionCode() bool`

HasPromotionCode returns a boolean if a field has been set.

### SetPromotionCodeNil

`func (o *UpdateSubscriptionRequest) SetPromotionCodeNil(b bool)`

 SetPromotionCodeNil sets the value for PromotionCode to be an explicit nil

### UnsetPromotionCode
`func (o *UpdateSubscriptionRequest) UnsetPromotionCode()`

UnsetPromotionCode ensures that no value is present for PromotionCode, not even an explicit nil
### GetCollectionMethod

`func (o *UpdateSubscriptionRequest) GetCollectionMethod() CollectionMethodEnum`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *UpdateSubscriptionRequest) GetCollectionMethodOk() (*CollectionMethodEnum, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *UpdateSubscriptionRequest) SetCollectionMethod(v CollectionMethodEnum)`

SetCollectionMethod sets CollectionMethod field to given value.

### HasCollectionMethod

`func (o *UpdateSubscriptionRequest) HasCollectionMethod() bool`

HasCollectionMethod returns a boolean if a field has been set.

### SetCollectionMethodNil

`func (o *UpdateSubscriptionRequest) SetCollectionMethodNil(b bool)`

 SetCollectionMethodNil sets the value for CollectionMethod to be an explicit nil

### UnsetCollectionMethod
`func (o *UpdateSubscriptionRequest) UnsetCollectionMethod()`

UnsetCollectionMethod ensures that no value is present for CollectionMethod, not even an explicit nil
### GetNetD

`func (o *UpdateSubscriptionRequest) GetNetD() int32`

GetNetD returns the NetD field if non-nil, zero value otherwise.

### GetNetDOk

`func (o *UpdateSubscriptionRequest) GetNetDOk() (*int32, bool)`

GetNetDOk returns a tuple with the NetD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetD

`func (o *UpdateSubscriptionRequest) SetNetD(v int32)`

SetNetD sets NetD field to given value.

### HasNetD

`func (o *UpdateSubscriptionRequest) HasNetD() bool`

HasNetD returns a boolean if a field has been set.

### SetNetDNil

`func (o *UpdateSubscriptionRequest) SetNetDNil(b bool)`

 SetNetDNil sets the value for NetD to be an explicit nil

### UnsetNetD
`func (o *UpdateSubscriptionRequest) UnsetNetD()`

UnsetNetD ensures that no value is present for NetD, not even an explicit nil
### GetCustomFields

`func (o *UpdateSubscriptionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *UpdateSubscriptionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *UpdateSubscriptionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *UpdateSubscriptionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *UpdateSubscriptionRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *UpdateSubscriptionRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetNewPeriodEnd

`func (o *UpdateSubscriptionRequest) GetNewPeriodEnd() time.Time`

GetNewPeriodEnd returns the NewPeriodEnd field if non-nil, zero value otherwise.

### GetNewPeriodEndOk

`func (o *UpdateSubscriptionRequest) GetNewPeriodEndOk() (*time.Time, bool)`

GetNewPeriodEndOk returns a tuple with the NewPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPeriodEnd

`func (o *UpdateSubscriptionRequest) SetNewPeriodEnd(v time.Time)`

SetNewPeriodEnd sets NewPeriodEnd field to given value.

### HasNewPeriodEnd

`func (o *UpdateSubscriptionRequest) HasNewPeriodEnd() bool`

HasNewPeriodEnd returns a boolean if a field has been set.

### SetNewPeriodEndNil

`func (o *UpdateSubscriptionRequest) SetNewPeriodEndNil(b bool)`

 SetNewPeriodEndNil sets the value for NewPeriodEnd to be an explicit nil

### UnsetNewPeriodEnd
`func (o *UpdateSubscriptionRequest) UnsetNewPeriodEnd()`

UnsetNewPeriodEnd ensures that no value is present for NewPeriodEnd, not even an explicit nil
### GetIsPreview

`func (o *UpdateSubscriptionRequest) GetIsPreview() bool`

GetIsPreview returns the IsPreview field if non-nil, zero value otherwise.

### GetIsPreviewOk

`func (o *UpdateSubscriptionRequest) GetIsPreviewOk() (*bool, bool)`

GetIsPreviewOk returns a tuple with the IsPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreview

`func (o *UpdateSubscriptionRequest) SetIsPreview(v bool)`

SetIsPreview sets IsPreview field to given value.

### HasIsPreview

`func (o *UpdateSubscriptionRequest) HasIsPreview() bool`

HasIsPreview returns a boolean if a field has been set.

### GetPreviewRenewalInvoices

`func (o *UpdateSubscriptionRequest) GetPreviewRenewalInvoices() bool`

GetPreviewRenewalInvoices returns the PreviewRenewalInvoices field if non-nil, zero value otherwise.

### GetPreviewRenewalInvoicesOk

`func (o *UpdateSubscriptionRequest) GetPreviewRenewalInvoicesOk() (*bool, bool)`

GetPreviewRenewalInvoicesOk returns a tuple with the PreviewRenewalInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewRenewalInvoices

`func (o *UpdateSubscriptionRequest) SetPreviewRenewalInvoices(v bool)`

SetPreviewRenewalInvoices sets PreviewRenewalInvoices field to given value.

### HasPreviewRenewalInvoices

`func (o *UpdateSubscriptionRequest) HasPreviewRenewalInvoices() bool`

HasPreviewRenewalInvoices returns a boolean if a field has been set.

### GetItems

`func (o *UpdateSubscriptionRequest) GetItems() []InlineSubscriptionItemUpdate`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UpdateSubscriptionRequest) GetItemsOk() (*[]InlineSubscriptionItemUpdate, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UpdateSubscriptionRequest) SetItems(v []InlineSubscriptionItemUpdate)`

SetItems sets Items field to given value.

### HasItems

`func (o *UpdateSubscriptionRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


