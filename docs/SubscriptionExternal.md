# SubscriptionExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Identifier of the subscription. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_SUBSCRIPTION]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Name** | **string** | Name for this subscription. | 
**AccountId** | **string** | Unique Identifier of the account. | 
**CustomerId** | **string** | Unique Identifier of the customer. | 
**Description** | **NullableString** |  | 
**Status** | [**SubscriptionStatusEnum**](SubscriptionStatusEnum.md) |  | 
**InitialStart** | **time.Time** | Very beginning of the subscription, in &#39;ISO 8601&#39; format. | 
**CurrentPeriodStart** | **time.Time** | Starting of the current billing period, in &#39;ISO 8601&#39; format. | 
**CurrentPeriodEnd** | **time.Time** | End of the current billing period, in &#39;ISO 8601&#39; format. | 
**DefaultPaymentMethodId** | **NullableString** |  | 
**CancelAtEnd** | **bool** | Whether this subscription should cancel at the end of contract or billing cycle. | 
**PauseAtEnd** | Pointer to **NullableBool** |  | [optional] 
**PauseForCycles** | Pointer to **NullableInt32** |  | [optional] 
**BillingInterval** | Pointer to [**NullableCalendarIntervalEnum**](CalendarIntervalEnum.md) |  | [optional] 
**BillingIntervalCount** | Pointer to **NullableInt32** |  | [optional] 
**ContractStart** | Pointer to **NullableTime** |  | [optional] 
**ContractAutoRenew** | Pointer to **NullableBool** |  | [optional] 
**Currency** | [**NullableCurrencyEnum**](CurrencyEnum.md) |  | 
**SubscriptionItems** | [**[]SubscriptionItemExternal**](SubscriptionItemExternal.md) |  | 
**CancellationDetails** | Pointer to [**NullableSubscriptionCancellationDetails**](SubscriptionCancellationDetails.md) |  | [optional] 
**CancelledAt** | **NullableTime** |  | 
**TrialStart** | Pointer to **NullableTime** |  | [optional] 
**TrialEnd** | Pointer to **NullableTime** |  | [optional] 
**Discount** | Pointer to [**NullableDiscountExternal**](DiscountExternal.md) |  | [optional] 
**TotalBillingCycles** | Pointer to **NullableInt32** |  | [optional] 
**RemainingBillingCycles** | Pointer to **NullableInt32** |  | [optional] 
**PausedAt** | Pointer to **NullableTime** |  | [optional] 
**ResumesAt** | Pointer to **NullableTime** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CollectionMethod** | [**CollectionMethodEnum**](CollectionMethodEnum.md) |  | 
**NetD** | **NullableInt32** |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**RenewsAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewSubscriptionExternal

`func NewSubscriptionExternal(id string, createdAt time.Time, updatedAt time.Time, name string, accountId string, customerId string, description NullableString, status SubscriptionStatusEnum, initialStart time.Time, currentPeriodStart time.Time, currentPeriodEnd time.Time, defaultPaymentMethodId NullableString, cancelAtEnd bool, currency NullableCurrencyEnum, subscriptionItems []SubscriptionItemExternal, cancelledAt NullableTime, collectionMethod CollectionMethodEnum, netD NullableInt32, ) *SubscriptionExternal`

NewSubscriptionExternal instantiates a new SubscriptionExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionExternalWithDefaults

`func NewSubscriptionExternalWithDefaults() *SubscriptionExternal`

NewSubscriptionExternalWithDefaults instantiates a new SubscriptionExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *SubscriptionExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscriptionExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscriptionExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *SubscriptionExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SubscriptionExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *SubscriptionExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SubscriptionExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SubscriptionExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *SubscriptionExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetName

`func (o *SubscriptionExternal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionExternal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionExternal) SetName(v string)`

SetName sets Name field to given value.


### GetAccountId

`func (o *SubscriptionExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SubscriptionExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SubscriptionExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCustomerId

`func (o *SubscriptionExternal) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SubscriptionExternal) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SubscriptionExternal) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDescription

`func (o *SubscriptionExternal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionExternal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionExternal) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *SubscriptionExternal) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SubscriptionExternal) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *SubscriptionExternal) GetStatus() SubscriptionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionExternal) GetStatusOk() (*SubscriptionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionExternal) SetStatus(v SubscriptionStatusEnum)`

SetStatus sets Status field to given value.


### GetInitialStart

`func (o *SubscriptionExternal) GetInitialStart() time.Time`

GetInitialStart returns the InitialStart field if non-nil, zero value otherwise.

### GetInitialStartOk

`func (o *SubscriptionExternal) GetInitialStartOk() (*time.Time, bool)`

GetInitialStartOk returns a tuple with the InitialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialStart

`func (o *SubscriptionExternal) SetInitialStart(v time.Time)`

SetInitialStart sets InitialStart field to given value.


### GetCurrentPeriodStart

`func (o *SubscriptionExternal) GetCurrentPeriodStart() time.Time`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *SubscriptionExternal) GetCurrentPeriodStartOk() (*time.Time, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *SubscriptionExternal) SetCurrentPeriodStart(v time.Time)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.


### GetCurrentPeriodEnd

`func (o *SubscriptionExternal) GetCurrentPeriodEnd() time.Time`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *SubscriptionExternal) GetCurrentPeriodEndOk() (*time.Time, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *SubscriptionExternal) SetCurrentPeriodEnd(v time.Time)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.


### GetDefaultPaymentMethodId

`func (o *SubscriptionExternal) GetDefaultPaymentMethodId() string`

GetDefaultPaymentMethodId returns the DefaultPaymentMethodId field if non-nil, zero value otherwise.

### GetDefaultPaymentMethodIdOk

`func (o *SubscriptionExternal) GetDefaultPaymentMethodIdOk() (*string, bool)`

GetDefaultPaymentMethodIdOk returns a tuple with the DefaultPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentMethodId

`func (o *SubscriptionExternal) SetDefaultPaymentMethodId(v string)`

SetDefaultPaymentMethodId sets DefaultPaymentMethodId field to given value.


### SetDefaultPaymentMethodIdNil

`func (o *SubscriptionExternal) SetDefaultPaymentMethodIdNil(b bool)`

 SetDefaultPaymentMethodIdNil sets the value for DefaultPaymentMethodId to be an explicit nil

### UnsetDefaultPaymentMethodId
`func (o *SubscriptionExternal) UnsetDefaultPaymentMethodId()`

UnsetDefaultPaymentMethodId ensures that no value is present for DefaultPaymentMethodId, not even an explicit nil
### GetCancelAtEnd

`func (o *SubscriptionExternal) GetCancelAtEnd() bool`

GetCancelAtEnd returns the CancelAtEnd field if non-nil, zero value otherwise.

### GetCancelAtEndOk

`func (o *SubscriptionExternal) GetCancelAtEndOk() (*bool, bool)`

GetCancelAtEndOk returns a tuple with the CancelAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtEnd

`func (o *SubscriptionExternal) SetCancelAtEnd(v bool)`

SetCancelAtEnd sets CancelAtEnd field to given value.


### GetPauseAtEnd

`func (o *SubscriptionExternal) GetPauseAtEnd() bool`

GetPauseAtEnd returns the PauseAtEnd field if non-nil, zero value otherwise.

### GetPauseAtEndOk

`func (o *SubscriptionExternal) GetPauseAtEndOk() (*bool, bool)`

GetPauseAtEndOk returns a tuple with the PauseAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseAtEnd

`func (o *SubscriptionExternal) SetPauseAtEnd(v bool)`

SetPauseAtEnd sets PauseAtEnd field to given value.

### HasPauseAtEnd

`func (o *SubscriptionExternal) HasPauseAtEnd() bool`

HasPauseAtEnd returns a boolean if a field has been set.

### SetPauseAtEndNil

`func (o *SubscriptionExternal) SetPauseAtEndNil(b bool)`

 SetPauseAtEndNil sets the value for PauseAtEnd to be an explicit nil

### UnsetPauseAtEnd
`func (o *SubscriptionExternal) UnsetPauseAtEnd()`

UnsetPauseAtEnd ensures that no value is present for PauseAtEnd, not even an explicit nil
### GetPauseForCycles

`func (o *SubscriptionExternal) GetPauseForCycles() int32`

GetPauseForCycles returns the PauseForCycles field if non-nil, zero value otherwise.

### GetPauseForCyclesOk

`func (o *SubscriptionExternal) GetPauseForCyclesOk() (*int32, bool)`

GetPauseForCyclesOk returns a tuple with the PauseForCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseForCycles

`func (o *SubscriptionExternal) SetPauseForCycles(v int32)`

SetPauseForCycles sets PauseForCycles field to given value.

### HasPauseForCycles

`func (o *SubscriptionExternal) HasPauseForCycles() bool`

HasPauseForCycles returns a boolean if a field has been set.

### SetPauseForCyclesNil

`func (o *SubscriptionExternal) SetPauseForCyclesNil(b bool)`

 SetPauseForCyclesNil sets the value for PauseForCycles to be an explicit nil

### UnsetPauseForCycles
`func (o *SubscriptionExternal) UnsetPauseForCycles()`

UnsetPauseForCycles ensures that no value is present for PauseForCycles, not even an explicit nil
### GetBillingInterval

`func (o *SubscriptionExternal) GetBillingInterval() CalendarIntervalEnum`

GetBillingInterval returns the BillingInterval field if non-nil, zero value otherwise.

### GetBillingIntervalOk

`func (o *SubscriptionExternal) GetBillingIntervalOk() (*CalendarIntervalEnum, bool)`

GetBillingIntervalOk returns a tuple with the BillingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInterval

`func (o *SubscriptionExternal) SetBillingInterval(v CalendarIntervalEnum)`

SetBillingInterval sets BillingInterval field to given value.

### HasBillingInterval

`func (o *SubscriptionExternal) HasBillingInterval() bool`

HasBillingInterval returns a boolean if a field has been set.

### SetBillingIntervalNil

`func (o *SubscriptionExternal) SetBillingIntervalNil(b bool)`

 SetBillingIntervalNil sets the value for BillingInterval to be an explicit nil

### UnsetBillingInterval
`func (o *SubscriptionExternal) UnsetBillingInterval()`

UnsetBillingInterval ensures that no value is present for BillingInterval, not even an explicit nil
### GetBillingIntervalCount

`func (o *SubscriptionExternal) GetBillingIntervalCount() int32`

GetBillingIntervalCount returns the BillingIntervalCount field if non-nil, zero value otherwise.

### GetBillingIntervalCountOk

`func (o *SubscriptionExternal) GetBillingIntervalCountOk() (*int32, bool)`

GetBillingIntervalCountOk returns a tuple with the BillingIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingIntervalCount

`func (o *SubscriptionExternal) SetBillingIntervalCount(v int32)`

SetBillingIntervalCount sets BillingIntervalCount field to given value.

### HasBillingIntervalCount

`func (o *SubscriptionExternal) HasBillingIntervalCount() bool`

HasBillingIntervalCount returns a boolean if a field has been set.

### SetBillingIntervalCountNil

`func (o *SubscriptionExternal) SetBillingIntervalCountNil(b bool)`

 SetBillingIntervalCountNil sets the value for BillingIntervalCount to be an explicit nil

### UnsetBillingIntervalCount
`func (o *SubscriptionExternal) UnsetBillingIntervalCount()`

UnsetBillingIntervalCount ensures that no value is present for BillingIntervalCount, not even an explicit nil
### GetContractStart

`func (o *SubscriptionExternal) GetContractStart() time.Time`

GetContractStart returns the ContractStart field if non-nil, zero value otherwise.

### GetContractStartOk

`func (o *SubscriptionExternal) GetContractStartOk() (*time.Time, bool)`

GetContractStartOk returns a tuple with the ContractStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStart

`func (o *SubscriptionExternal) SetContractStart(v time.Time)`

SetContractStart sets ContractStart field to given value.

### HasContractStart

`func (o *SubscriptionExternal) HasContractStart() bool`

HasContractStart returns a boolean if a field has been set.

### SetContractStartNil

`func (o *SubscriptionExternal) SetContractStartNil(b bool)`

 SetContractStartNil sets the value for ContractStart to be an explicit nil

### UnsetContractStart
`func (o *SubscriptionExternal) UnsetContractStart()`

UnsetContractStart ensures that no value is present for ContractStart, not even an explicit nil
### GetContractAutoRenew

`func (o *SubscriptionExternal) GetContractAutoRenew() bool`

GetContractAutoRenew returns the ContractAutoRenew field if non-nil, zero value otherwise.

### GetContractAutoRenewOk

`func (o *SubscriptionExternal) GetContractAutoRenewOk() (*bool, bool)`

GetContractAutoRenewOk returns a tuple with the ContractAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAutoRenew

`func (o *SubscriptionExternal) SetContractAutoRenew(v bool)`

SetContractAutoRenew sets ContractAutoRenew field to given value.

### HasContractAutoRenew

`func (o *SubscriptionExternal) HasContractAutoRenew() bool`

HasContractAutoRenew returns a boolean if a field has been set.

### SetContractAutoRenewNil

`func (o *SubscriptionExternal) SetContractAutoRenewNil(b bool)`

 SetContractAutoRenewNil sets the value for ContractAutoRenew to be an explicit nil

### UnsetContractAutoRenew
`func (o *SubscriptionExternal) UnsetContractAutoRenew()`

UnsetContractAutoRenew ensures that no value is present for ContractAutoRenew, not even an explicit nil
### GetCurrency

`func (o *SubscriptionExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *SubscriptionExternal) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *SubscriptionExternal) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetSubscriptionItems

`func (o *SubscriptionExternal) GetSubscriptionItems() []SubscriptionItemExternal`

GetSubscriptionItems returns the SubscriptionItems field if non-nil, zero value otherwise.

### GetSubscriptionItemsOk

`func (o *SubscriptionExternal) GetSubscriptionItemsOk() (*[]SubscriptionItemExternal, bool)`

GetSubscriptionItemsOk returns a tuple with the SubscriptionItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionItems

`func (o *SubscriptionExternal) SetSubscriptionItems(v []SubscriptionItemExternal)`

SetSubscriptionItems sets SubscriptionItems field to given value.


### GetCancellationDetails

`func (o *SubscriptionExternal) GetCancellationDetails() SubscriptionCancellationDetails`

GetCancellationDetails returns the CancellationDetails field if non-nil, zero value otherwise.

### GetCancellationDetailsOk

`func (o *SubscriptionExternal) GetCancellationDetailsOk() (*SubscriptionCancellationDetails, bool)`

GetCancellationDetailsOk returns a tuple with the CancellationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationDetails

`func (o *SubscriptionExternal) SetCancellationDetails(v SubscriptionCancellationDetails)`

SetCancellationDetails sets CancellationDetails field to given value.

### HasCancellationDetails

`func (o *SubscriptionExternal) HasCancellationDetails() bool`

HasCancellationDetails returns a boolean if a field has been set.

### SetCancellationDetailsNil

`func (o *SubscriptionExternal) SetCancellationDetailsNil(b bool)`

 SetCancellationDetailsNil sets the value for CancellationDetails to be an explicit nil

### UnsetCancellationDetails
`func (o *SubscriptionExternal) UnsetCancellationDetails()`

UnsetCancellationDetails ensures that no value is present for CancellationDetails, not even an explicit nil
### GetCancelledAt

`func (o *SubscriptionExternal) GetCancelledAt() time.Time`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *SubscriptionExternal) GetCancelledAtOk() (*time.Time, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *SubscriptionExternal) SetCancelledAt(v time.Time)`

SetCancelledAt sets CancelledAt field to given value.


### SetCancelledAtNil

`func (o *SubscriptionExternal) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *SubscriptionExternal) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetTrialStart

`func (o *SubscriptionExternal) GetTrialStart() time.Time`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *SubscriptionExternal) GetTrialStartOk() (*time.Time, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *SubscriptionExternal) SetTrialStart(v time.Time)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *SubscriptionExternal) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.

### SetTrialStartNil

`func (o *SubscriptionExternal) SetTrialStartNil(b bool)`

 SetTrialStartNil sets the value for TrialStart to be an explicit nil

### UnsetTrialStart
`func (o *SubscriptionExternal) UnsetTrialStart()`

UnsetTrialStart ensures that no value is present for TrialStart, not even an explicit nil
### GetTrialEnd

`func (o *SubscriptionExternal) GetTrialEnd() time.Time`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *SubscriptionExternal) GetTrialEndOk() (*time.Time, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *SubscriptionExternal) SetTrialEnd(v time.Time)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *SubscriptionExternal) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### SetTrialEndNil

`func (o *SubscriptionExternal) SetTrialEndNil(b bool)`

 SetTrialEndNil sets the value for TrialEnd to be an explicit nil

### UnsetTrialEnd
`func (o *SubscriptionExternal) UnsetTrialEnd()`

UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
### GetDiscount

`func (o *SubscriptionExternal) GetDiscount() DiscountExternal`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *SubscriptionExternal) GetDiscountOk() (*DiscountExternal, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *SubscriptionExternal) SetDiscount(v DiscountExternal)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *SubscriptionExternal) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *SubscriptionExternal) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *SubscriptionExternal) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
### GetTotalBillingCycles

`func (o *SubscriptionExternal) GetTotalBillingCycles() int32`

GetTotalBillingCycles returns the TotalBillingCycles field if non-nil, zero value otherwise.

### GetTotalBillingCyclesOk

`func (o *SubscriptionExternal) GetTotalBillingCyclesOk() (*int32, bool)`

GetTotalBillingCyclesOk returns a tuple with the TotalBillingCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBillingCycles

`func (o *SubscriptionExternal) SetTotalBillingCycles(v int32)`

SetTotalBillingCycles sets TotalBillingCycles field to given value.

### HasTotalBillingCycles

`func (o *SubscriptionExternal) HasTotalBillingCycles() bool`

HasTotalBillingCycles returns a boolean if a field has been set.

### SetTotalBillingCyclesNil

`func (o *SubscriptionExternal) SetTotalBillingCyclesNil(b bool)`

 SetTotalBillingCyclesNil sets the value for TotalBillingCycles to be an explicit nil

### UnsetTotalBillingCycles
`func (o *SubscriptionExternal) UnsetTotalBillingCycles()`

UnsetTotalBillingCycles ensures that no value is present for TotalBillingCycles, not even an explicit nil
### GetRemainingBillingCycles

`func (o *SubscriptionExternal) GetRemainingBillingCycles() int32`

GetRemainingBillingCycles returns the RemainingBillingCycles field if non-nil, zero value otherwise.

### GetRemainingBillingCyclesOk

`func (o *SubscriptionExternal) GetRemainingBillingCyclesOk() (*int32, bool)`

GetRemainingBillingCyclesOk returns a tuple with the RemainingBillingCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingBillingCycles

`func (o *SubscriptionExternal) SetRemainingBillingCycles(v int32)`

SetRemainingBillingCycles sets RemainingBillingCycles field to given value.

### HasRemainingBillingCycles

`func (o *SubscriptionExternal) HasRemainingBillingCycles() bool`

HasRemainingBillingCycles returns a boolean if a field has been set.

### SetRemainingBillingCyclesNil

`func (o *SubscriptionExternal) SetRemainingBillingCyclesNil(b bool)`

 SetRemainingBillingCyclesNil sets the value for RemainingBillingCycles to be an explicit nil

### UnsetRemainingBillingCycles
`func (o *SubscriptionExternal) UnsetRemainingBillingCycles()`

UnsetRemainingBillingCycles ensures that no value is present for RemainingBillingCycles, not even an explicit nil
### GetPausedAt

`func (o *SubscriptionExternal) GetPausedAt() time.Time`

GetPausedAt returns the PausedAt field if non-nil, zero value otherwise.

### GetPausedAtOk

`func (o *SubscriptionExternal) GetPausedAtOk() (*time.Time, bool)`

GetPausedAtOk returns a tuple with the PausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedAt

`func (o *SubscriptionExternal) SetPausedAt(v time.Time)`

SetPausedAt sets PausedAt field to given value.

### HasPausedAt

`func (o *SubscriptionExternal) HasPausedAt() bool`

HasPausedAt returns a boolean if a field has been set.

### SetPausedAtNil

`func (o *SubscriptionExternal) SetPausedAtNil(b bool)`

 SetPausedAtNil sets the value for PausedAt to be an explicit nil

### UnsetPausedAt
`func (o *SubscriptionExternal) UnsetPausedAt()`

UnsetPausedAt ensures that no value is present for PausedAt, not even an explicit nil
### GetResumesAt

`func (o *SubscriptionExternal) GetResumesAt() time.Time`

GetResumesAt returns the ResumesAt field if non-nil, zero value otherwise.

### GetResumesAtOk

`func (o *SubscriptionExternal) GetResumesAtOk() (*time.Time, bool)`

GetResumesAtOk returns a tuple with the ResumesAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumesAt

`func (o *SubscriptionExternal) SetResumesAt(v time.Time)`

SetResumesAt sets ResumesAt field to given value.

### HasResumesAt

`func (o *SubscriptionExternal) HasResumesAt() bool`

HasResumesAt returns a boolean if a field has been set.

### SetResumesAtNil

`func (o *SubscriptionExternal) SetResumesAtNil(b bool)`

 SetResumesAtNil sets the value for ResumesAt to be an explicit nil

### UnsetResumesAt
`func (o *SubscriptionExternal) UnsetResumesAt()`

UnsetResumesAt ensures that no value is present for ResumesAt, not even an explicit nil
### GetMetadata

`func (o *SubscriptionExternal) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionExternal) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionExternal) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscriptionExternal) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *SubscriptionExternal) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *SubscriptionExternal) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCollectionMethod

`func (o *SubscriptionExternal) GetCollectionMethod() CollectionMethodEnum`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *SubscriptionExternal) GetCollectionMethodOk() (*CollectionMethodEnum, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *SubscriptionExternal) SetCollectionMethod(v CollectionMethodEnum)`

SetCollectionMethod sets CollectionMethod field to given value.


### GetNetD

`func (o *SubscriptionExternal) GetNetD() int32`

GetNetD returns the NetD field if non-nil, zero value otherwise.

### GetNetDOk

`func (o *SubscriptionExternal) GetNetDOk() (*int32, bool)`

GetNetDOk returns a tuple with the NetD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetD

`func (o *SubscriptionExternal) SetNetD(v int32)`

SetNetD sets NetD field to given value.


### SetNetDNil

`func (o *SubscriptionExternal) SetNetDNil(b bool)`

 SetNetDNil sets the value for NetD to be an explicit nil

### UnsetNetD
`func (o *SubscriptionExternal) UnsetNetD()`

UnsetNetD ensures that no value is present for NetD, not even an explicit nil
### GetCustomFields

`func (o *SubscriptionExternal) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SubscriptionExternal) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SubscriptionExternal) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SubscriptionExternal) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *SubscriptionExternal) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *SubscriptionExternal) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetRenewsAt

`func (o *SubscriptionExternal) GetRenewsAt() time.Time`

GetRenewsAt returns the RenewsAt field if non-nil, zero value otherwise.

### GetRenewsAtOk

`func (o *SubscriptionExternal) GetRenewsAtOk() (*time.Time, bool)`

GetRenewsAtOk returns a tuple with the RenewsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewsAt

`func (o *SubscriptionExternal) SetRenewsAt(v time.Time)`

SetRenewsAt sets RenewsAt field to given value.

### HasRenewsAt

`func (o *SubscriptionExternal) HasRenewsAt() bool`

HasRenewsAt returns a boolean if a field has been set.

### SetRenewsAtNil

`func (o *SubscriptionExternal) SetRenewsAtNil(b bool)`

 SetRenewsAtNil sets the value for RenewsAt to be an explicit nil

### UnsetRenewsAt
`func (o *SubscriptionExternal) UnsetRenewsAt()`

UnsetRenewsAt ensures that no value is present for RenewsAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


