# PaymentIntentExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountAtom** | **int32** |  | 
**AmountAtomCapturable** | **NullableInt32** |  | 
**AmountAtomReceived** | **NullableInt32** |  | 
**ChargeIds** | **[]string** |  | 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomerId** | **string** |  | 
**DeclineReason** | **NullableString** |  | 
**Id** | **string** |  | 
**Invoice** | [**NullableInvoiceExternal**](InvoiceExternal.md) |  | 
**InvoiceId** | **NullableString** |  | 
**InvoicePaymentProviderTypeFee** | **NullableInt32** |  | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**LastRefundDate** | **NullableTime** |  | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**PaymentMethodId** | **NullableString** |  | 
**PaymentProcessorName** | **NullableString** |  | 
**ProcessorPaymentIntentId** | **NullableString** |  | 
**RefundIds** | **[]string** |  | 
**Status** | [**PaymentIntentStatus**](PaymentIntentStatus.md) |  | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewPaymentIntentExternal

`func NewPaymentIntentExternal(amountAtom int32, amountAtomCapturable NullableInt32, amountAtomReceived NullableInt32, chargeIds []string, createdAt time.Time, currency CurrencyEnum, customerId string, declineReason NullableString, id string, invoice NullableInvoiceExternal, invoiceId NullableString, invoicePaymentProviderTypeFee NullableInt32, lastRefundDate NullableTime, paymentMethodId NullableString, paymentProcessorName NullableString, processorPaymentIntentId NullableString, refundIds []string, status PaymentIntentStatus, updatedAt time.Time, ) *PaymentIntentExternal`

NewPaymentIntentExternal instantiates a new PaymentIntentExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentIntentExternalWithDefaults

`func NewPaymentIntentExternalWithDefaults() *PaymentIntentExternal`

NewPaymentIntentExternalWithDefaults instantiates a new PaymentIntentExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountAtom

`func (o *PaymentIntentExternal) GetAmountAtom() int32`

GetAmountAtom returns the AmountAtom field if non-nil, zero value otherwise.

### GetAmountAtomOk

`func (o *PaymentIntentExternal) GetAmountAtomOk() (*int32, bool)`

GetAmountAtomOk returns a tuple with the AmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtom

`func (o *PaymentIntentExternal) SetAmountAtom(v int32)`

SetAmountAtom sets AmountAtom field to given value.


### GetAmountAtomCapturable

`func (o *PaymentIntentExternal) GetAmountAtomCapturable() int32`

GetAmountAtomCapturable returns the AmountAtomCapturable field if non-nil, zero value otherwise.

### GetAmountAtomCapturableOk

`func (o *PaymentIntentExternal) GetAmountAtomCapturableOk() (*int32, bool)`

GetAmountAtomCapturableOk returns a tuple with the AmountAtomCapturable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtomCapturable

`func (o *PaymentIntentExternal) SetAmountAtomCapturable(v int32)`

SetAmountAtomCapturable sets AmountAtomCapturable field to given value.


### SetAmountAtomCapturableNil

`func (o *PaymentIntentExternal) SetAmountAtomCapturableNil(b bool)`

 SetAmountAtomCapturableNil sets the value for AmountAtomCapturable to be an explicit nil

### UnsetAmountAtomCapturable
`func (o *PaymentIntentExternal) UnsetAmountAtomCapturable()`

UnsetAmountAtomCapturable ensures that no value is present for AmountAtomCapturable, not even an explicit nil
### GetAmountAtomReceived

`func (o *PaymentIntentExternal) GetAmountAtomReceived() int32`

GetAmountAtomReceived returns the AmountAtomReceived field if non-nil, zero value otherwise.

### GetAmountAtomReceivedOk

`func (o *PaymentIntentExternal) GetAmountAtomReceivedOk() (*int32, bool)`

GetAmountAtomReceivedOk returns a tuple with the AmountAtomReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtomReceived

`func (o *PaymentIntentExternal) SetAmountAtomReceived(v int32)`

SetAmountAtomReceived sets AmountAtomReceived field to given value.


### SetAmountAtomReceivedNil

`func (o *PaymentIntentExternal) SetAmountAtomReceivedNil(b bool)`

 SetAmountAtomReceivedNil sets the value for AmountAtomReceived to be an explicit nil

### UnsetAmountAtomReceived
`func (o *PaymentIntentExternal) UnsetAmountAtomReceived()`

UnsetAmountAtomReceived ensures that no value is present for AmountAtomReceived, not even an explicit nil
### GetChargeIds

`func (o *PaymentIntentExternal) GetChargeIds() []string`

GetChargeIds returns the ChargeIds field if non-nil, zero value otherwise.

### GetChargeIdsOk

`func (o *PaymentIntentExternal) GetChargeIdsOk() (*[]string, bool)`

GetChargeIdsOk returns a tuple with the ChargeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeIds

`func (o *PaymentIntentExternal) SetChargeIds(v []string)`

SetChargeIds sets ChargeIds field to given value.


### GetCreatedAt

`func (o *PaymentIntentExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentIntentExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentIntentExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *PaymentIntentExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentIntentExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentIntentExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetCustomFields

`func (o *PaymentIntentExternal) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PaymentIntentExternal) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PaymentIntentExternal) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PaymentIntentExternal) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *PaymentIntentExternal) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *PaymentIntentExternal) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetCustomerId

`func (o *PaymentIntentExternal) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PaymentIntentExternal) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PaymentIntentExternal) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDeclineReason

`func (o *PaymentIntentExternal) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *PaymentIntentExternal) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *PaymentIntentExternal) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.


### SetDeclineReasonNil

`func (o *PaymentIntentExternal) SetDeclineReasonNil(b bool)`

 SetDeclineReasonNil sets the value for DeclineReason to be an explicit nil

### UnsetDeclineReason
`func (o *PaymentIntentExternal) UnsetDeclineReason()`

UnsetDeclineReason ensures that no value is present for DeclineReason, not even an explicit nil
### GetId

`func (o *PaymentIntentExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentIntentExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentIntentExternal) SetId(v string)`

SetId sets Id field to given value.


### GetInvoice

`func (o *PaymentIntentExternal) GetInvoice() InvoiceExternal`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *PaymentIntentExternal) GetInvoiceOk() (*InvoiceExternal, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *PaymentIntentExternal) SetInvoice(v InvoiceExternal)`

SetInvoice sets Invoice field to given value.


### SetInvoiceNil

`func (o *PaymentIntentExternal) SetInvoiceNil(b bool)`

 SetInvoiceNil sets the value for Invoice to be an explicit nil

### UnsetInvoice
`func (o *PaymentIntentExternal) UnsetInvoice()`

UnsetInvoice ensures that no value is present for Invoice, not even an explicit nil
### GetInvoiceId

`func (o *PaymentIntentExternal) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PaymentIntentExternal) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PaymentIntentExternal) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *PaymentIntentExternal) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *PaymentIntentExternal) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetInvoicePaymentProviderTypeFee

`func (o *PaymentIntentExternal) GetInvoicePaymentProviderTypeFee() int32`

GetInvoicePaymentProviderTypeFee returns the InvoicePaymentProviderTypeFee field if non-nil, zero value otherwise.

### GetInvoicePaymentProviderTypeFeeOk

`func (o *PaymentIntentExternal) GetInvoicePaymentProviderTypeFeeOk() (*int32, bool)`

GetInvoicePaymentProviderTypeFeeOk returns a tuple with the InvoicePaymentProviderTypeFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePaymentProviderTypeFee

`func (o *PaymentIntentExternal) SetInvoicePaymentProviderTypeFee(v int32)`

SetInvoicePaymentProviderTypeFee sets InvoicePaymentProviderTypeFee field to given value.


### SetInvoicePaymentProviderTypeFeeNil

`func (o *PaymentIntentExternal) SetInvoicePaymentProviderTypeFeeNil(b bool)`

 SetInvoicePaymentProviderTypeFeeNil sets the value for InvoicePaymentProviderTypeFee to be an explicit nil

### UnsetInvoicePaymentProviderTypeFee
`func (o *PaymentIntentExternal) UnsetInvoicePaymentProviderTypeFee()`

UnsetInvoicePaymentProviderTypeFee ensures that no value is present for InvoicePaymentProviderTypeFee, not even an explicit nil
### GetIsDeleted

`func (o *PaymentIntentExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PaymentIntentExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PaymentIntentExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PaymentIntentExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLastRefundDate

`func (o *PaymentIntentExternal) GetLastRefundDate() time.Time`

GetLastRefundDate returns the LastRefundDate field if non-nil, zero value otherwise.

### GetLastRefundDateOk

`func (o *PaymentIntentExternal) GetLastRefundDateOk() (*time.Time, bool)`

GetLastRefundDateOk returns a tuple with the LastRefundDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefundDate

`func (o *PaymentIntentExternal) SetLastRefundDate(v time.Time)`

SetLastRefundDate sets LastRefundDate field to given value.


### SetLastRefundDateNil

`func (o *PaymentIntentExternal) SetLastRefundDateNil(b bool)`

 SetLastRefundDateNil sets the value for LastRefundDate to be an explicit nil

### UnsetLastRefundDate
`func (o *PaymentIntentExternal) UnsetLastRefundDate()`

UnsetLastRefundDate ensures that no value is present for LastRefundDate, not even an explicit nil
### GetObject

`func (o *PaymentIntentExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentIntentExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentIntentExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *PaymentIntentExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *PaymentIntentExternal) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *PaymentIntentExternal) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *PaymentIntentExternal) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.


### SetPaymentMethodIdNil

`func (o *PaymentIntentExternal) SetPaymentMethodIdNil(b bool)`

 SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil

### UnsetPaymentMethodId
`func (o *PaymentIntentExternal) UnsetPaymentMethodId()`

UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
### GetPaymentProcessorName

`func (o *PaymentIntentExternal) GetPaymentProcessorName() string`

GetPaymentProcessorName returns the PaymentProcessorName field if non-nil, zero value otherwise.

### GetPaymentProcessorNameOk

`func (o *PaymentIntentExternal) GetPaymentProcessorNameOk() (*string, bool)`

GetPaymentProcessorNameOk returns a tuple with the PaymentProcessorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProcessorName

`func (o *PaymentIntentExternal) SetPaymentProcessorName(v string)`

SetPaymentProcessorName sets PaymentProcessorName field to given value.


### SetPaymentProcessorNameNil

`func (o *PaymentIntentExternal) SetPaymentProcessorNameNil(b bool)`

 SetPaymentProcessorNameNil sets the value for PaymentProcessorName to be an explicit nil

### UnsetPaymentProcessorName
`func (o *PaymentIntentExternal) UnsetPaymentProcessorName()`

UnsetPaymentProcessorName ensures that no value is present for PaymentProcessorName, not even an explicit nil
### GetProcessorPaymentIntentId

`func (o *PaymentIntentExternal) GetProcessorPaymentIntentId() string`

GetProcessorPaymentIntentId returns the ProcessorPaymentIntentId field if non-nil, zero value otherwise.

### GetProcessorPaymentIntentIdOk

`func (o *PaymentIntentExternal) GetProcessorPaymentIntentIdOk() (*string, bool)`

GetProcessorPaymentIntentIdOk returns a tuple with the ProcessorPaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorPaymentIntentId

`func (o *PaymentIntentExternal) SetProcessorPaymentIntentId(v string)`

SetProcessorPaymentIntentId sets ProcessorPaymentIntentId field to given value.


### SetProcessorPaymentIntentIdNil

`func (o *PaymentIntentExternal) SetProcessorPaymentIntentIdNil(b bool)`

 SetProcessorPaymentIntentIdNil sets the value for ProcessorPaymentIntentId to be an explicit nil

### UnsetProcessorPaymentIntentId
`func (o *PaymentIntentExternal) UnsetProcessorPaymentIntentId()`

UnsetProcessorPaymentIntentId ensures that no value is present for ProcessorPaymentIntentId, not even an explicit nil
### GetRefundIds

`func (o *PaymentIntentExternal) GetRefundIds() []string`

GetRefundIds returns the RefundIds field if non-nil, zero value otherwise.

### GetRefundIdsOk

`func (o *PaymentIntentExternal) GetRefundIdsOk() (*[]string, bool)`

GetRefundIdsOk returns a tuple with the RefundIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundIds

`func (o *PaymentIntentExternal) SetRefundIds(v []string)`

SetRefundIds sets RefundIds field to given value.


### GetStatus

`func (o *PaymentIntentExternal) GetStatus() PaymentIntentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentIntentExternal) GetStatusOk() (*PaymentIntentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentIntentExternal) SetStatus(v PaymentIntentStatus)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *PaymentIntentExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentIntentExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentIntentExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


