# InvoiceExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the invoice. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_INVOICE]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**AccountId** | **string** | Unique identifier of the account. | 
**BillingReason** | [**BillingReasonEnum**](BillingReasonEnum.md) |  | 
**CollectionMethod** | [**CollectionMethodEnum**](CollectionMethodEnum.md) |  | 
**Comments** | Pointer to [**[]InvoiceComment**](InvoiceComment.md) |  | [optional] 
**Customer** | Pointer to [**NullableCustomerExternal**](CustomerExternal.md) |  | [optional] 
**CustomerId** | **string** | Unique identifier of the customer. | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) |  | 
**DefaultPaymentMethodId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DueDate** | **time.Time** | Due date to pay the invoice. | 
**LatestPaymentIntentId** | Pointer to **NullableString** |  | [optional] 
**PaidAt** | Pointer to **NullableTime** |  | [optional] 
**PeriodEnd** | **time.Time** | End of the usage period during which invoice_items were added to this invoice. It is in &#39;ISO 8601&#39; format. | 
**PeriodStart** | **time.Time** | Start of the usage period during which invoice_items were added to this invoice. It is in &#39;ISO 8601&#39; format. | 
**Status** | [**InvoiceStatusEnum**](InvoiceStatusEnum.md) |  | 
**PaidOutOfBand** | **NullableBool** |  | 
**Subscription** | Pointer to [**NullableSubscriptionExternal**](SubscriptionExternal.md) |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**TotalAmountAtom** | **int32** | Total amount of the invoice. Sum of invoice_items’ total amount. It is in atomic units (in USD this is cents). | 
**RemainingAmountAtom** | **int32** | Remaining amount of the invoice to be paid. It is in atomic units (in USD this is cents). | 
**TaxAmountAtom** | **int32** | Total tax amount in atomic units. | 
**Lines** | [**[]InvoiceItemExternal**](InvoiceItemExternal.md) |  | 
**PrePaymentCreditNotesAmount** | Pointer to **NullableInt32** |  | [optional] 
**PostPaymentCreditNotesAmount** | Pointer to **NullableInt32** |  | [optional] 
**CreditNoteIds** | Pointer to **[]string** |  | [optional] 
**AppliedBalanceAmountAtom** | **int32** | Amount applied from customer balance either from credit or debit. It is in atomic units (in USD this is cents). | 
**DueAmountAtom** | **int32** | Final amount due at this time for this invoice. It isin atomic units (in USD this is cents). | 
**PaidAmountAtom** | **int32** | The amount that was paid in atomic units (in USD this is cents). | 
**Discounts** | Pointer to [**[]DiscountExternal**](DiscountExternal.md) |  | [optional] 
**TotalDiscountAmountAtoms** | Pointer to [**[]InvoiceDiscountAmountsExternal**](InvoiceDiscountAmountsExternal.md) |  | [optional] 
**NetD** | Pointer to **int32** | Number of days the customer has to pay the invoice, from 0 to 365, where -1 indicates due immediately. | [optional] [default to -1]
**HostedInvoiceUrl** | Pointer to **NullableString** |  | [optional] 
**LatestPaymentAttemptFailureMessage** | Pointer to **NullableString** |  | [optional] 
**TaxProcessorUpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**RefundIds** | **[]string** |  | 
**RefundedAmountAtom** | **NullableInt32** |  | 
**InvoicePdfUrl** | **string** | The URL for the Invoice PDF | 
**ReceiptPdfUrl** | **string** | The URL for the Receipt PDF | 
**IsInitialInvoiceForTrialSub** | **bool** | Whether this is the first invoice for a trial subscription. | 
**InvoiceNumber** | **int32** | Invoice number | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**EmailInvoiceOnFinalization** | Pointer to **bool** | Whether invoice should be sent to the customer upon finalizing the invoice | [optional] [default to false]

## Methods

### NewInvoiceExternal

`func NewInvoiceExternal(id string, createdAt time.Time, updatedAt time.Time, accountId string, billingReason BillingReasonEnum, collectionMethod CollectionMethodEnum, customerId string, currency CurrencyEnum, dueDate time.Time, periodEnd time.Time, periodStart time.Time, status InvoiceStatusEnum, paidOutOfBand NullableBool, totalAmountAtom int32, remainingAmountAtom int32, taxAmountAtom int32, lines []InvoiceItemExternal, appliedBalanceAmountAtom int32, dueAmountAtom int32, paidAmountAtom int32, refundIds []string, refundedAmountAtom NullableInt32, invoicePdfUrl string, receiptPdfUrl string, isInitialInvoiceForTrialSub bool, invoiceNumber int32, ) *InvoiceExternal`

NewInvoiceExternal instantiates a new InvoiceExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceExternalWithDefaults

`func NewInvoiceExternalWithDefaults() *InvoiceExternal`

NewInvoiceExternalWithDefaults instantiates a new InvoiceExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *InvoiceExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *InvoiceExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *InvoiceExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *InvoiceExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InvoiceExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InvoiceExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *InvoiceExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *InvoiceExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *InvoiceExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *InvoiceExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetAccountId

`func (o *InvoiceExternal) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InvoiceExternal) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InvoiceExternal) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBillingReason

`func (o *InvoiceExternal) GetBillingReason() BillingReasonEnum`

GetBillingReason returns the BillingReason field if non-nil, zero value otherwise.

### GetBillingReasonOk

`func (o *InvoiceExternal) GetBillingReasonOk() (*BillingReasonEnum, bool)`

GetBillingReasonOk returns a tuple with the BillingReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingReason

`func (o *InvoiceExternal) SetBillingReason(v BillingReasonEnum)`

SetBillingReason sets BillingReason field to given value.


### GetCollectionMethod

`func (o *InvoiceExternal) GetCollectionMethod() CollectionMethodEnum`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *InvoiceExternal) GetCollectionMethodOk() (*CollectionMethodEnum, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *InvoiceExternal) SetCollectionMethod(v CollectionMethodEnum)`

SetCollectionMethod sets CollectionMethod field to given value.


### GetComments

`func (o *InvoiceExternal) GetComments() []InvoiceComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *InvoiceExternal) GetCommentsOk() (*[]InvoiceComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *InvoiceExternal) SetComments(v []InvoiceComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *InvoiceExternal) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCustomer

`func (o *InvoiceExternal) GetCustomer() CustomerExternal`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InvoiceExternal) GetCustomerOk() (*CustomerExternal, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InvoiceExternal) SetCustomer(v CustomerExternal)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *InvoiceExternal) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### SetCustomerNil

`func (o *InvoiceExternal) SetCustomerNil(b bool)`

 SetCustomerNil sets the value for Customer to be an explicit nil

### UnsetCustomer
`func (o *InvoiceExternal) UnsetCustomer()`

UnsetCustomer ensures that no value is present for Customer, not even an explicit nil
### GetCustomerId

`func (o *InvoiceExternal) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InvoiceExternal) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InvoiceExternal) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetCurrency

`func (o *InvoiceExternal) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceExternal) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceExternal) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetDefaultPaymentMethodId

`func (o *InvoiceExternal) GetDefaultPaymentMethodId() string`

GetDefaultPaymentMethodId returns the DefaultPaymentMethodId field if non-nil, zero value otherwise.

### GetDefaultPaymentMethodIdOk

`func (o *InvoiceExternal) GetDefaultPaymentMethodIdOk() (*string, bool)`

GetDefaultPaymentMethodIdOk returns a tuple with the DefaultPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentMethodId

`func (o *InvoiceExternal) SetDefaultPaymentMethodId(v string)`

SetDefaultPaymentMethodId sets DefaultPaymentMethodId field to given value.

### HasDefaultPaymentMethodId

`func (o *InvoiceExternal) HasDefaultPaymentMethodId() bool`

HasDefaultPaymentMethodId returns a boolean if a field has been set.

### SetDefaultPaymentMethodIdNil

`func (o *InvoiceExternal) SetDefaultPaymentMethodIdNil(b bool)`

 SetDefaultPaymentMethodIdNil sets the value for DefaultPaymentMethodId to be an explicit nil

### UnsetDefaultPaymentMethodId
`func (o *InvoiceExternal) UnsetDefaultPaymentMethodId()`

UnsetDefaultPaymentMethodId ensures that no value is present for DefaultPaymentMethodId, not even an explicit nil
### GetDescription

`func (o *InvoiceExternal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceExternal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceExternal) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceExternal) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InvoiceExternal) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvoiceExternal) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDueDate

`func (o *InvoiceExternal) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceExternal) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceExternal) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.


### GetLatestPaymentIntentId

`func (o *InvoiceExternal) GetLatestPaymentIntentId() string`

GetLatestPaymentIntentId returns the LatestPaymentIntentId field if non-nil, zero value otherwise.

### GetLatestPaymentIntentIdOk

`func (o *InvoiceExternal) GetLatestPaymentIntentIdOk() (*string, bool)`

GetLatestPaymentIntentIdOk returns a tuple with the LatestPaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestPaymentIntentId

`func (o *InvoiceExternal) SetLatestPaymentIntentId(v string)`

SetLatestPaymentIntentId sets LatestPaymentIntentId field to given value.

### HasLatestPaymentIntentId

`func (o *InvoiceExternal) HasLatestPaymentIntentId() bool`

HasLatestPaymentIntentId returns a boolean if a field has been set.

### SetLatestPaymentIntentIdNil

`func (o *InvoiceExternal) SetLatestPaymentIntentIdNil(b bool)`

 SetLatestPaymentIntentIdNil sets the value for LatestPaymentIntentId to be an explicit nil

### UnsetLatestPaymentIntentId
`func (o *InvoiceExternal) UnsetLatestPaymentIntentId()`

UnsetLatestPaymentIntentId ensures that no value is present for LatestPaymentIntentId, not even an explicit nil
### GetPaidAt

`func (o *InvoiceExternal) GetPaidAt() time.Time`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *InvoiceExternal) GetPaidAtOk() (*time.Time, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *InvoiceExternal) SetPaidAt(v time.Time)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *InvoiceExternal) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.

### SetPaidAtNil

`func (o *InvoiceExternal) SetPaidAtNil(b bool)`

 SetPaidAtNil sets the value for PaidAt to be an explicit nil

### UnsetPaidAt
`func (o *InvoiceExternal) UnsetPaidAt()`

UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
### GetPeriodEnd

`func (o *InvoiceExternal) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *InvoiceExternal) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *InvoiceExternal) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetPeriodStart

`func (o *InvoiceExternal) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *InvoiceExternal) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *InvoiceExternal) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.


### GetStatus

`func (o *InvoiceExternal) GetStatus() InvoiceStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceExternal) GetStatusOk() (*InvoiceStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceExternal) SetStatus(v InvoiceStatusEnum)`

SetStatus sets Status field to given value.


### GetPaidOutOfBand

`func (o *InvoiceExternal) GetPaidOutOfBand() bool`

GetPaidOutOfBand returns the PaidOutOfBand field if non-nil, zero value otherwise.

### GetPaidOutOfBandOk

`func (o *InvoiceExternal) GetPaidOutOfBandOk() (*bool, bool)`

GetPaidOutOfBandOk returns a tuple with the PaidOutOfBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidOutOfBand

`func (o *InvoiceExternal) SetPaidOutOfBand(v bool)`

SetPaidOutOfBand sets PaidOutOfBand field to given value.


### SetPaidOutOfBandNil

`func (o *InvoiceExternal) SetPaidOutOfBandNil(b bool)`

 SetPaidOutOfBandNil sets the value for PaidOutOfBand to be an explicit nil

### UnsetPaidOutOfBand
`func (o *InvoiceExternal) UnsetPaidOutOfBand()`

UnsetPaidOutOfBand ensures that no value is present for PaidOutOfBand, not even an explicit nil
### GetSubscription

`func (o *InvoiceExternal) GetSubscription() SubscriptionExternal`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *InvoiceExternal) GetSubscriptionOk() (*SubscriptionExternal, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *InvoiceExternal) SetSubscription(v SubscriptionExternal)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *InvoiceExternal) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### SetSubscriptionNil

`func (o *InvoiceExternal) SetSubscriptionNil(b bool)`

 SetSubscriptionNil sets the value for Subscription to be an explicit nil

### UnsetSubscription
`func (o *InvoiceExternal) UnsetSubscription()`

UnsetSubscription ensures that no value is present for Subscription, not even an explicit nil
### GetSubscriptionId

`func (o *InvoiceExternal) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *InvoiceExternal) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *InvoiceExternal) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *InvoiceExternal) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *InvoiceExternal) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *InvoiceExternal) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetTotalAmountAtom

`func (o *InvoiceExternal) GetTotalAmountAtom() int32`

GetTotalAmountAtom returns the TotalAmountAtom field if non-nil, zero value otherwise.

### GetTotalAmountAtomOk

`func (o *InvoiceExternal) GetTotalAmountAtomOk() (*int32, bool)`

GetTotalAmountAtomOk returns a tuple with the TotalAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountAtom

`func (o *InvoiceExternal) SetTotalAmountAtom(v int32)`

SetTotalAmountAtom sets TotalAmountAtom field to given value.


### GetRemainingAmountAtom

`func (o *InvoiceExternal) GetRemainingAmountAtom() int32`

GetRemainingAmountAtom returns the RemainingAmountAtom field if non-nil, zero value otherwise.

### GetRemainingAmountAtomOk

`func (o *InvoiceExternal) GetRemainingAmountAtomOk() (*int32, bool)`

GetRemainingAmountAtomOk returns a tuple with the RemainingAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingAmountAtom

`func (o *InvoiceExternal) SetRemainingAmountAtom(v int32)`

SetRemainingAmountAtom sets RemainingAmountAtom field to given value.


### GetTaxAmountAtom

`func (o *InvoiceExternal) GetTaxAmountAtom() int32`

GetTaxAmountAtom returns the TaxAmountAtom field if non-nil, zero value otherwise.

### GetTaxAmountAtomOk

`func (o *InvoiceExternal) GetTaxAmountAtomOk() (*int32, bool)`

GetTaxAmountAtomOk returns a tuple with the TaxAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountAtom

`func (o *InvoiceExternal) SetTaxAmountAtom(v int32)`

SetTaxAmountAtom sets TaxAmountAtom field to given value.


### GetLines

`func (o *InvoiceExternal) GetLines() []InvoiceItemExternal`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *InvoiceExternal) GetLinesOk() (*[]InvoiceItemExternal, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *InvoiceExternal) SetLines(v []InvoiceItemExternal)`

SetLines sets Lines field to given value.


### GetPrePaymentCreditNotesAmount

`func (o *InvoiceExternal) GetPrePaymentCreditNotesAmount() int32`

GetPrePaymentCreditNotesAmount returns the PrePaymentCreditNotesAmount field if non-nil, zero value otherwise.

### GetPrePaymentCreditNotesAmountOk

`func (o *InvoiceExternal) GetPrePaymentCreditNotesAmountOk() (*int32, bool)`

GetPrePaymentCreditNotesAmountOk returns a tuple with the PrePaymentCreditNotesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaymentCreditNotesAmount

`func (o *InvoiceExternal) SetPrePaymentCreditNotesAmount(v int32)`

SetPrePaymentCreditNotesAmount sets PrePaymentCreditNotesAmount field to given value.

### HasPrePaymentCreditNotesAmount

`func (o *InvoiceExternal) HasPrePaymentCreditNotesAmount() bool`

HasPrePaymentCreditNotesAmount returns a boolean if a field has been set.

### SetPrePaymentCreditNotesAmountNil

`func (o *InvoiceExternal) SetPrePaymentCreditNotesAmountNil(b bool)`

 SetPrePaymentCreditNotesAmountNil sets the value for PrePaymentCreditNotesAmount to be an explicit nil

### UnsetPrePaymentCreditNotesAmount
`func (o *InvoiceExternal) UnsetPrePaymentCreditNotesAmount()`

UnsetPrePaymentCreditNotesAmount ensures that no value is present for PrePaymentCreditNotesAmount, not even an explicit nil
### GetPostPaymentCreditNotesAmount

`func (o *InvoiceExternal) GetPostPaymentCreditNotesAmount() int32`

GetPostPaymentCreditNotesAmount returns the PostPaymentCreditNotesAmount field if non-nil, zero value otherwise.

### GetPostPaymentCreditNotesAmountOk

`func (o *InvoiceExternal) GetPostPaymentCreditNotesAmountOk() (*int32, bool)`

GetPostPaymentCreditNotesAmountOk returns a tuple with the PostPaymentCreditNotesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPaymentCreditNotesAmount

`func (o *InvoiceExternal) SetPostPaymentCreditNotesAmount(v int32)`

SetPostPaymentCreditNotesAmount sets PostPaymentCreditNotesAmount field to given value.

### HasPostPaymentCreditNotesAmount

`func (o *InvoiceExternal) HasPostPaymentCreditNotesAmount() bool`

HasPostPaymentCreditNotesAmount returns a boolean if a field has been set.

### SetPostPaymentCreditNotesAmountNil

`func (o *InvoiceExternal) SetPostPaymentCreditNotesAmountNil(b bool)`

 SetPostPaymentCreditNotesAmountNil sets the value for PostPaymentCreditNotesAmount to be an explicit nil

### UnsetPostPaymentCreditNotesAmount
`func (o *InvoiceExternal) UnsetPostPaymentCreditNotesAmount()`

UnsetPostPaymentCreditNotesAmount ensures that no value is present for PostPaymentCreditNotesAmount, not even an explicit nil
### GetCreditNoteIds

`func (o *InvoiceExternal) GetCreditNoteIds() []*string`

GetCreditNoteIds returns the CreditNoteIds field if non-nil, zero value otherwise.

### GetCreditNoteIdsOk

`func (o *InvoiceExternal) GetCreditNoteIdsOk() (*[]*string, bool)`

GetCreditNoteIdsOk returns a tuple with the CreditNoteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteIds

`func (o *InvoiceExternal) SetCreditNoteIds(v []*string)`

SetCreditNoteIds sets CreditNoteIds field to given value.

### HasCreditNoteIds

`func (o *InvoiceExternal) HasCreditNoteIds() bool`

HasCreditNoteIds returns a boolean if a field has been set.

### GetAppliedBalanceAmountAtom

`func (o *InvoiceExternal) GetAppliedBalanceAmountAtom() int32`

GetAppliedBalanceAmountAtom returns the AppliedBalanceAmountAtom field if non-nil, zero value otherwise.

### GetAppliedBalanceAmountAtomOk

`func (o *InvoiceExternal) GetAppliedBalanceAmountAtomOk() (*int32, bool)`

GetAppliedBalanceAmountAtomOk returns a tuple with the AppliedBalanceAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedBalanceAmountAtom

`func (o *InvoiceExternal) SetAppliedBalanceAmountAtom(v int32)`

SetAppliedBalanceAmountAtom sets AppliedBalanceAmountAtom field to given value.


### GetDueAmountAtom

`func (o *InvoiceExternal) GetDueAmountAtom() int32`

GetDueAmountAtom returns the DueAmountAtom field if non-nil, zero value otherwise.

### GetDueAmountAtomOk

`func (o *InvoiceExternal) GetDueAmountAtomOk() (*int32, bool)`

GetDueAmountAtomOk returns a tuple with the DueAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAmountAtom

`func (o *InvoiceExternal) SetDueAmountAtom(v int32)`

SetDueAmountAtom sets DueAmountAtom field to given value.


### GetPaidAmountAtom

`func (o *InvoiceExternal) GetPaidAmountAtom() int32`

GetPaidAmountAtom returns the PaidAmountAtom field if non-nil, zero value otherwise.

### GetPaidAmountAtomOk

`func (o *InvoiceExternal) GetPaidAmountAtomOk() (*int32, bool)`

GetPaidAmountAtomOk returns a tuple with the PaidAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAmountAtom

`func (o *InvoiceExternal) SetPaidAmountAtom(v int32)`

SetPaidAmountAtom sets PaidAmountAtom field to given value.


### GetDiscounts

`func (o *InvoiceExternal) GetDiscounts() []DiscountExternal`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *InvoiceExternal) GetDiscountsOk() (*[]DiscountExternal, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *InvoiceExternal) SetDiscounts(v []DiscountExternal)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *InvoiceExternal) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetTotalDiscountAmountAtoms

`func (o *InvoiceExternal) GetTotalDiscountAmountAtoms() []InvoiceDiscountAmountsExternal`

GetTotalDiscountAmountAtoms returns the TotalDiscountAmountAtoms field if non-nil, zero value otherwise.

### GetTotalDiscountAmountAtomsOk

`func (o *InvoiceExternal) GetTotalDiscountAmountAtomsOk() (*[]InvoiceDiscountAmountsExternal, bool)`

GetTotalDiscountAmountAtomsOk returns a tuple with the TotalDiscountAmountAtoms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountAmountAtoms

`func (o *InvoiceExternal) SetTotalDiscountAmountAtoms(v []InvoiceDiscountAmountsExternal)`

SetTotalDiscountAmountAtoms sets TotalDiscountAmountAtoms field to given value.

### HasTotalDiscountAmountAtoms

`func (o *InvoiceExternal) HasTotalDiscountAmountAtoms() bool`

HasTotalDiscountAmountAtoms returns a boolean if a field has been set.

### GetNetD

`func (o *InvoiceExternal) GetNetD() int32`

GetNetD returns the NetD field if non-nil, zero value otherwise.

### GetNetDOk

`func (o *InvoiceExternal) GetNetDOk() (*int32, bool)`

GetNetDOk returns a tuple with the NetD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetD

`func (o *InvoiceExternal) SetNetD(v int32)`

SetNetD sets NetD field to given value.

### HasNetD

`func (o *InvoiceExternal) HasNetD() bool`

HasNetD returns a boolean if a field has been set.

### GetHostedInvoiceUrl

`func (o *InvoiceExternal) GetHostedInvoiceUrl() string`

GetHostedInvoiceUrl returns the HostedInvoiceUrl field if non-nil, zero value otherwise.

### GetHostedInvoiceUrlOk

`func (o *InvoiceExternal) GetHostedInvoiceUrlOk() (*string, bool)`

GetHostedInvoiceUrlOk returns a tuple with the HostedInvoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedInvoiceUrl

`func (o *InvoiceExternal) SetHostedInvoiceUrl(v string)`

SetHostedInvoiceUrl sets HostedInvoiceUrl field to given value.

### HasHostedInvoiceUrl

`func (o *InvoiceExternal) HasHostedInvoiceUrl() bool`

HasHostedInvoiceUrl returns a boolean if a field has been set.

### SetHostedInvoiceUrlNil

`func (o *InvoiceExternal) SetHostedInvoiceUrlNil(b bool)`

 SetHostedInvoiceUrlNil sets the value for HostedInvoiceUrl to be an explicit nil

### UnsetHostedInvoiceUrl
`func (o *InvoiceExternal) UnsetHostedInvoiceUrl()`

UnsetHostedInvoiceUrl ensures that no value is present for HostedInvoiceUrl, not even an explicit nil
### GetLatestPaymentAttemptFailureMessage

`func (o *InvoiceExternal) GetLatestPaymentAttemptFailureMessage() string`

GetLatestPaymentAttemptFailureMessage returns the LatestPaymentAttemptFailureMessage field if non-nil, zero value otherwise.

### GetLatestPaymentAttemptFailureMessageOk

`func (o *InvoiceExternal) GetLatestPaymentAttemptFailureMessageOk() (*string, bool)`

GetLatestPaymentAttemptFailureMessageOk returns a tuple with the LatestPaymentAttemptFailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestPaymentAttemptFailureMessage

`func (o *InvoiceExternal) SetLatestPaymentAttemptFailureMessage(v string)`

SetLatestPaymentAttemptFailureMessage sets LatestPaymentAttemptFailureMessage field to given value.

### HasLatestPaymentAttemptFailureMessage

`func (o *InvoiceExternal) HasLatestPaymentAttemptFailureMessage() bool`

HasLatestPaymentAttemptFailureMessage returns a boolean if a field has been set.

### SetLatestPaymentAttemptFailureMessageNil

`func (o *InvoiceExternal) SetLatestPaymentAttemptFailureMessageNil(b bool)`

 SetLatestPaymentAttemptFailureMessageNil sets the value for LatestPaymentAttemptFailureMessage to be an explicit nil

### UnsetLatestPaymentAttemptFailureMessage
`func (o *InvoiceExternal) UnsetLatestPaymentAttemptFailureMessage()`

UnsetLatestPaymentAttemptFailureMessage ensures that no value is present for LatestPaymentAttemptFailureMessage, not even an explicit nil
### GetTaxProcessorUpdatedAt

`func (o *InvoiceExternal) GetTaxProcessorUpdatedAt() time.Time`

GetTaxProcessorUpdatedAt returns the TaxProcessorUpdatedAt field if non-nil, zero value otherwise.

### GetTaxProcessorUpdatedAtOk

`func (o *InvoiceExternal) GetTaxProcessorUpdatedAtOk() (*time.Time, bool)`

GetTaxProcessorUpdatedAtOk returns a tuple with the TaxProcessorUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxProcessorUpdatedAt

`func (o *InvoiceExternal) SetTaxProcessorUpdatedAt(v time.Time)`

SetTaxProcessorUpdatedAt sets TaxProcessorUpdatedAt field to given value.

### HasTaxProcessorUpdatedAt

`func (o *InvoiceExternal) HasTaxProcessorUpdatedAt() bool`

HasTaxProcessorUpdatedAt returns a boolean if a field has been set.

### SetTaxProcessorUpdatedAtNil

`func (o *InvoiceExternal) SetTaxProcessorUpdatedAtNil(b bool)`

 SetTaxProcessorUpdatedAtNil sets the value for TaxProcessorUpdatedAt to be an explicit nil

### UnsetTaxProcessorUpdatedAt
`func (o *InvoiceExternal) UnsetTaxProcessorUpdatedAt()`

UnsetTaxProcessorUpdatedAt ensures that no value is present for TaxProcessorUpdatedAt, not even an explicit nil
### GetRefundIds

`func (o *InvoiceExternal) GetRefundIds() []string`

GetRefundIds returns the RefundIds field if non-nil, zero value otherwise.

### GetRefundIdsOk

`func (o *InvoiceExternal) GetRefundIdsOk() (*[]string, bool)`

GetRefundIdsOk returns a tuple with the RefundIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundIds

`func (o *InvoiceExternal) SetRefundIds(v []string)`

SetRefundIds sets RefundIds field to given value.


### GetRefundedAmountAtom

`func (o *InvoiceExternal) GetRefundedAmountAtom() int32`

GetRefundedAmountAtom returns the RefundedAmountAtom field if non-nil, zero value otherwise.

### GetRefundedAmountAtomOk

`func (o *InvoiceExternal) GetRefundedAmountAtomOk() (*int32, bool)`

GetRefundedAmountAtomOk returns a tuple with the RefundedAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmountAtom

`func (o *InvoiceExternal) SetRefundedAmountAtom(v int32)`

SetRefundedAmountAtom sets RefundedAmountAtom field to given value.


### SetRefundedAmountAtomNil

`func (o *InvoiceExternal) SetRefundedAmountAtomNil(b bool)`

 SetRefundedAmountAtomNil sets the value for RefundedAmountAtom to be an explicit nil

### UnsetRefundedAmountAtom
`func (o *InvoiceExternal) UnsetRefundedAmountAtom()`

UnsetRefundedAmountAtom ensures that no value is present for RefundedAmountAtom, not even an explicit nil
### GetInvoicePdfUrl

`func (o *InvoiceExternal) GetInvoicePdfUrl() string`

GetInvoicePdfUrl returns the InvoicePdfUrl field if non-nil, zero value otherwise.

### GetInvoicePdfUrlOk

`func (o *InvoiceExternal) GetInvoicePdfUrlOk() (*string, bool)`

GetInvoicePdfUrlOk returns a tuple with the InvoicePdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdfUrl

`func (o *InvoiceExternal) SetInvoicePdfUrl(v string)`

SetInvoicePdfUrl sets InvoicePdfUrl field to given value.


### GetReceiptPdfUrl

`func (o *InvoiceExternal) GetReceiptPdfUrl() string`

GetReceiptPdfUrl returns the ReceiptPdfUrl field if non-nil, zero value otherwise.

### GetReceiptPdfUrlOk

`func (o *InvoiceExternal) GetReceiptPdfUrlOk() (*string, bool)`

GetReceiptPdfUrlOk returns a tuple with the ReceiptPdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptPdfUrl

`func (o *InvoiceExternal) SetReceiptPdfUrl(v string)`

SetReceiptPdfUrl sets ReceiptPdfUrl field to given value.


### GetIsInitialInvoiceForTrialSub

`func (o *InvoiceExternal) GetIsInitialInvoiceForTrialSub() bool`

GetIsInitialInvoiceForTrialSub returns the IsInitialInvoiceForTrialSub field if non-nil, zero value otherwise.

### GetIsInitialInvoiceForTrialSubOk

`func (o *InvoiceExternal) GetIsInitialInvoiceForTrialSubOk() (*bool, bool)`

GetIsInitialInvoiceForTrialSubOk returns a tuple with the IsInitialInvoiceForTrialSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitialInvoiceForTrialSub

`func (o *InvoiceExternal) SetIsInitialInvoiceForTrialSub(v bool)`

SetIsInitialInvoiceForTrialSub sets IsInitialInvoiceForTrialSub field to given value.


### GetInvoiceNumber

`func (o *InvoiceExternal) GetInvoiceNumber() int32`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoiceExternal) GetInvoiceNumberOk() (*int32, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoiceExternal) SetInvoiceNumber(v int32)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetMetadata

`func (o *InvoiceExternal) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InvoiceExternal) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InvoiceExternal) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InvoiceExternal) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *InvoiceExternal) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *InvoiceExternal) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCustomFields

`func (o *InvoiceExternal) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InvoiceExternal) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InvoiceExternal) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InvoiceExternal) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *InvoiceExternal) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *InvoiceExternal) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetEmailInvoiceOnFinalization

`func (o *InvoiceExternal) GetEmailInvoiceOnFinalization() bool`

GetEmailInvoiceOnFinalization returns the EmailInvoiceOnFinalization field if non-nil, zero value otherwise.

### GetEmailInvoiceOnFinalizationOk

`func (o *InvoiceExternal) GetEmailInvoiceOnFinalizationOk() (*bool, bool)`

GetEmailInvoiceOnFinalizationOk returns a tuple with the EmailInvoiceOnFinalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailInvoiceOnFinalization

`func (o *InvoiceExternal) SetEmailInvoiceOnFinalization(v bool)`

SetEmailInvoiceOnFinalization sets EmailInvoiceOnFinalization field to given value.

### HasEmailInvoiceOnFinalization

`func (o *InvoiceExternal) HasEmailInvoiceOnFinalization() bool`

HasEmailInvoiceOnFinalization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


