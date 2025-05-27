# InvoicePublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedBalanceAmountAtom** | **int32** | Amount applied from customer balance either from credit or debit. It is in atomic units (in USD this is cents). | 
**BilledTo** | Pointer to **NullableString** |  | [optional] 
**BilledToBusinessName** | Pointer to **NullableString** |  | [optional] 
**BillingAddress** | Pointer to **NullableString** |  | [optional] 
**Branding** | **map[string]interface{}** | The branding settings associated with the account | 
**CardType** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**Currency** | [**CurrencyEnum**](CurrencyEnum.md) | Three-letter ISO currency code, in lowercase. | 
**CustomerTaxIds** | Pointer to [**[]TaxIdSetting**](TaxIdSetting.md) | The tax ID settings of the customer. | [optional] 
**DueAmountAtom** | **int32** | Final amount due at this time for this invoice. It isin atomic units (in USD this is cents). | 
**DueDate** | Pointer to **NullableTime** |  | [optional] 
**Id** | **string** | Unique identifier of the invoice. | 
**InvoicePdfUrl** | **string** | The URL for the Invoice PDF | 
**IsInitialInvoiceForTrialSub** | **bool** | Whether this is the first invoice for a trial subscription. | 
**LastFour** | Pointer to **NullableString** |  | [optional] 
**Lines** | Pointer to [**[]InvoiceItemPublic**](InvoiceItemPublic.md) | List of individual line items that make up the invoice. | [optional] 
**MerchantBillingAddress** | [**NullableCompleteAddress**](CompleteAddress.md) |  | 
**MerchantTaxIds** | Pointer to [**[]MerchantTaxIdSetting**](MerchantTaxIdSetting.md) | The tax ID settings of the merchant. | [optional] 
**PaidAmountAtom** | **int32** | Total amount paid. It is in atomic units (in USD this is cents). | 
**PaidAt** | Pointer to **NullableTime** |  | [optional] 
**ProviderTypeFeeAmountAtom** | **int32** | Total fee amount charged for the payment provider type in atomic units. | 
**ReceiptPdfUrl** | **string** | The URL for the Receipt PDF | 
**RemainingAmountAtom** | **int32** | Remaining amount of the invoice to be paid. It is in atomic units (in USD this is cents). | 
**Status** | [**InvoiceStatusEnum**](InvoiceStatusEnum.md) | Status of the invoice. | 
**TaxAmountAtom** | **int32** | Total tax amount in atomic units. | 
**TotalAmountAtom** | **int32** | Total amount of the invoice. Sum of invoice_items&#39; total amount. It is in atomic units (in USD this is cents). | 
**TotalDiscountAmountAtoms** | Pointer to [**[]InvoiceDiscountAmountsExternal**](InvoiceDiscountAmountsExternal.md) | The aggregate amount_atoms calculated per discount across all line items. | [optional] 
**TotalExcludingTaxesAmountAtom** | **int32** | Total amount excluding taxes. It is in atomic units (in USD this is cents). | 
**TrialEndForSub** | **NullableTime** |  | 
**TrialStartForSub** | **NullableTime** |  | 

## Methods

### NewInvoicePublic

`func NewInvoicePublic(appliedBalanceAmountAtom int32, branding map[string]interface{}, createdAt time.Time, currency CurrencyEnum, dueAmountAtom int32, id string, invoicePdfUrl string, isInitialInvoiceForTrialSub bool, merchantBillingAddress NullableCompleteAddress, paidAmountAtom int32, providerTypeFeeAmountAtom int32, receiptPdfUrl string, remainingAmountAtom int32, status InvoiceStatusEnum, taxAmountAtom int32, totalAmountAtom int32, totalExcludingTaxesAmountAtom int32, trialEndForSub NullableTime, trialStartForSub NullableTime, ) *InvoicePublic`

NewInvoicePublic instantiates a new InvoicePublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePublicWithDefaults

`func NewInvoicePublicWithDefaults() *InvoicePublic`

NewInvoicePublicWithDefaults instantiates a new InvoicePublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedBalanceAmountAtom

`func (o *InvoicePublic) GetAppliedBalanceAmountAtom() int32`

GetAppliedBalanceAmountAtom returns the AppliedBalanceAmountAtom field if non-nil, zero value otherwise.

### GetAppliedBalanceAmountAtomOk

`func (o *InvoicePublic) GetAppliedBalanceAmountAtomOk() (*int32, bool)`

GetAppliedBalanceAmountAtomOk returns a tuple with the AppliedBalanceAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedBalanceAmountAtom

`func (o *InvoicePublic) SetAppliedBalanceAmountAtom(v int32)`

SetAppliedBalanceAmountAtom sets AppliedBalanceAmountAtom field to given value.


### GetBilledTo

`func (o *InvoicePublic) GetBilledTo() string`

GetBilledTo returns the BilledTo field if non-nil, zero value otherwise.

### GetBilledToOk

`func (o *InvoicePublic) GetBilledToOk() (*string, bool)`

GetBilledToOk returns a tuple with the BilledTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledTo

`func (o *InvoicePublic) SetBilledTo(v string)`

SetBilledTo sets BilledTo field to given value.

### HasBilledTo

`func (o *InvoicePublic) HasBilledTo() bool`

HasBilledTo returns a boolean if a field has been set.

### SetBilledToNil

`func (o *InvoicePublic) SetBilledToNil(b bool)`

 SetBilledToNil sets the value for BilledTo to be an explicit nil

### UnsetBilledTo
`func (o *InvoicePublic) UnsetBilledTo()`

UnsetBilledTo ensures that no value is present for BilledTo, not even an explicit nil
### GetBilledToBusinessName

`func (o *InvoicePublic) GetBilledToBusinessName() string`

GetBilledToBusinessName returns the BilledToBusinessName field if non-nil, zero value otherwise.

### GetBilledToBusinessNameOk

`func (o *InvoicePublic) GetBilledToBusinessNameOk() (*string, bool)`

GetBilledToBusinessNameOk returns a tuple with the BilledToBusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledToBusinessName

`func (o *InvoicePublic) SetBilledToBusinessName(v string)`

SetBilledToBusinessName sets BilledToBusinessName field to given value.

### HasBilledToBusinessName

`func (o *InvoicePublic) HasBilledToBusinessName() bool`

HasBilledToBusinessName returns a boolean if a field has been set.

### SetBilledToBusinessNameNil

`func (o *InvoicePublic) SetBilledToBusinessNameNil(b bool)`

 SetBilledToBusinessNameNil sets the value for BilledToBusinessName to be an explicit nil

### UnsetBilledToBusinessName
`func (o *InvoicePublic) UnsetBilledToBusinessName()`

UnsetBilledToBusinessName ensures that no value is present for BilledToBusinessName, not even an explicit nil
### GetBillingAddress

`func (o *InvoicePublic) GetBillingAddress() string`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *InvoicePublic) GetBillingAddressOk() (*string, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *InvoicePublic) SetBillingAddress(v string)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *InvoicePublic) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### SetBillingAddressNil

`func (o *InvoicePublic) SetBillingAddressNil(b bool)`

 SetBillingAddressNil sets the value for BillingAddress to be an explicit nil

### UnsetBillingAddress
`func (o *InvoicePublic) UnsetBillingAddress()`

UnsetBillingAddress ensures that no value is present for BillingAddress, not even an explicit nil
### GetBranding

`func (o *InvoicePublic) GetBranding() map[string]interface{}`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *InvoicePublic) GetBrandingOk() (*map[string]interface{}, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *InvoicePublic) SetBranding(v map[string]interface{})`

SetBranding sets Branding field to given value.


### GetCardType

`func (o *InvoicePublic) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *InvoicePublic) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *InvoicePublic) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *InvoicePublic) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### SetCardTypeNil

`func (o *InvoicePublic) SetCardTypeNil(b bool)`

 SetCardTypeNil sets the value for CardType to be an explicit nil

### UnsetCardType
`func (o *InvoicePublic) UnsetCardType()`

UnsetCardType ensures that no value is present for CardType, not even an explicit nil
### GetCreatedAt

`func (o *InvoicePublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoicePublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoicePublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *InvoicePublic) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoicePublic) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoicePublic) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.


### GetCustomerTaxIds

`func (o *InvoicePublic) GetCustomerTaxIds() []TaxIdSetting`

GetCustomerTaxIds returns the CustomerTaxIds field if non-nil, zero value otherwise.

### GetCustomerTaxIdsOk

`func (o *InvoicePublic) GetCustomerTaxIdsOk() (*[]TaxIdSetting, bool)`

GetCustomerTaxIdsOk returns a tuple with the CustomerTaxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTaxIds

`func (o *InvoicePublic) SetCustomerTaxIds(v []TaxIdSetting)`

SetCustomerTaxIds sets CustomerTaxIds field to given value.

### HasCustomerTaxIds

`func (o *InvoicePublic) HasCustomerTaxIds() bool`

HasCustomerTaxIds returns a boolean if a field has been set.

### GetDueAmountAtom

`func (o *InvoicePublic) GetDueAmountAtom() int32`

GetDueAmountAtom returns the DueAmountAtom field if non-nil, zero value otherwise.

### GetDueAmountAtomOk

`func (o *InvoicePublic) GetDueAmountAtomOk() (*int32, bool)`

GetDueAmountAtomOk returns a tuple with the DueAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAmountAtom

`func (o *InvoicePublic) SetDueAmountAtom(v int32)`

SetDueAmountAtom sets DueAmountAtom field to given value.


### GetDueDate

`func (o *InvoicePublic) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoicePublic) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoicePublic) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoicePublic) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *InvoicePublic) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *InvoicePublic) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetId

`func (o *InvoicePublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoicePublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoicePublic) SetId(v string)`

SetId sets Id field to given value.


### GetInvoicePdfUrl

`func (o *InvoicePublic) GetInvoicePdfUrl() string`

GetInvoicePdfUrl returns the InvoicePdfUrl field if non-nil, zero value otherwise.

### GetInvoicePdfUrlOk

`func (o *InvoicePublic) GetInvoicePdfUrlOk() (*string, bool)`

GetInvoicePdfUrlOk returns a tuple with the InvoicePdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdfUrl

`func (o *InvoicePublic) SetInvoicePdfUrl(v string)`

SetInvoicePdfUrl sets InvoicePdfUrl field to given value.


### GetIsInitialInvoiceForTrialSub

`func (o *InvoicePublic) GetIsInitialInvoiceForTrialSub() bool`

GetIsInitialInvoiceForTrialSub returns the IsInitialInvoiceForTrialSub field if non-nil, zero value otherwise.

### GetIsInitialInvoiceForTrialSubOk

`func (o *InvoicePublic) GetIsInitialInvoiceForTrialSubOk() (*bool, bool)`

GetIsInitialInvoiceForTrialSubOk returns a tuple with the IsInitialInvoiceForTrialSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitialInvoiceForTrialSub

`func (o *InvoicePublic) SetIsInitialInvoiceForTrialSub(v bool)`

SetIsInitialInvoiceForTrialSub sets IsInitialInvoiceForTrialSub field to given value.


### GetLastFour

`func (o *InvoicePublic) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *InvoicePublic) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *InvoicePublic) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *InvoicePublic) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### SetLastFourNil

`func (o *InvoicePublic) SetLastFourNil(b bool)`

 SetLastFourNil sets the value for LastFour to be an explicit nil

### UnsetLastFour
`func (o *InvoicePublic) UnsetLastFour()`

UnsetLastFour ensures that no value is present for LastFour, not even an explicit nil
### GetLines

`func (o *InvoicePublic) GetLines() []InvoiceItemPublic`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *InvoicePublic) GetLinesOk() (*[]InvoiceItemPublic, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *InvoicePublic) SetLines(v []InvoiceItemPublic)`

SetLines sets Lines field to given value.

### HasLines

`func (o *InvoicePublic) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetMerchantBillingAddress

`func (o *InvoicePublic) GetMerchantBillingAddress() CompleteAddress`

GetMerchantBillingAddress returns the MerchantBillingAddress field if non-nil, zero value otherwise.

### GetMerchantBillingAddressOk

`func (o *InvoicePublic) GetMerchantBillingAddressOk() (*CompleteAddress, bool)`

GetMerchantBillingAddressOk returns a tuple with the MerchantBillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantBillingAddress

`func (o *InvoicePublic) SetMerchantBillingAddress(v CompleteAddress)`

SetMerchantBillingAddress sets MerchantBillingAddress field to given value.


### SetMerchantBillingAddressNil

`func (o *InvoicePublic) SetMerchantBillingAddressNil(b bool)`

 SetMerchantBillingAddressNil sets the value for MerchantBillingAddress to be an explicit nil

### UnsetMerchantBillingAddress
`func (o *InvoicePublic) UnsetMerchantBillingAddress()`

UnsetMerchantBillingAddress ensures that no value is present for MerchantBillingAddress, not even an explicit nil
### GetMerchantTaxIds

`func (o *InvoicePublic) GetMerchantTaxIds() []MerchantTaxIdSetting`

GetMerchantTaxIds returns the MerchantTaxIds field if non-nil, zero value otherwise.

### GetMerchantTaxIdsOk

`func (o *InvoicePublic) GetMerchantTaxIdsOk() (*[]MerchantTaxIdSetting, bool)`

GetMerchantTaxIdsOk returns a tuple with the MerchantTaxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTaxIds

`func (o *InvoicePublic) SetMerchantTaxIds(v []MerchantTaxIdSetting)`

SetMerchantTaxIds sets MerchantTaxIds field to given value.

### HasMerchantTaxIds

`func (o *InvoicePublic) HasMerchantTaxIds() bool`

HasMerchantTaxIds returns a boolean if a field has been set.

### GetPaidAmountAtom

`func (o *InvoicePublic) GetPaidAmountAtom() int32`

GetPaidAmountAtom returns the PaidAmountAtom field if non-nil, zero value otherwise.

### GetPaidAmountAtomOk

`func (o *InvoicePublic) GetPaidAmountAtomOk() (*int32, bool)`

GetPaidAmountAtomOk returns a tuple with the PaidAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAmountAtom

`func (o *InvoicePublic) SetPaidAmountAtom(v int32)`

SetPaidAmountAtom sets PaidAmountAtom field to given value.


### GetPaidAt

`func (o *InvoicePublic) GetPaidAt() time.Time`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *InvoicePublic) GetPaidAtOk() (*time.Time, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *InvoicePublic) SetPaidAt(v time.Time)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *InvoicePublic) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.

### SetPaidAtNil

`func (o *InvoicePublic) SetPaidAtNil(b bool)`

 SetPaidAtNil sets the value for PaidAt to be an explicit nil

### UnsetPaidAt
`func (o *InvoicePublic) UnsetPaidAt()`

UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
### GetProviderTypeFeeAmountAtom

`func (o *InvoicePublic) GetProviderTypeFeeAmountAtom() int32`

GetProviderTypeFeeAmountAtom returns the ProviderTypeFeeAmountAtom field if non-nil, zero value otherwise.

### GetProviderTypeFeeAmountAtomOk

`func (o *InvoicePublic) GetProviderTypeFeeAmountAtomOk() (*int32, bool)`

GetProviderTypeFeeAmountAtomOk returns a tuple with the ProviderTypeFeeAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTypeFeeAmountAtom

`func (o *InvoicePublic) SetProviderTypeFeeAmountAtom(v int32)`

SetProviderTypeFeeAmountAtom sets ProviderTypeFeeAmountAtom field to given value.


### GetReceiptPdfUrl

`func (o *InvoicePublic) GetReceiptPdfUrl() string`

GetReceiptPdfUrl returns the ReceiptPdfUrl field if non-nil, zero value otherwise.

### GetReceiptPdfUrlOk

`func (o *InvoicePublic) GetReceiptPdfUrlOk() (*string, bool)`

GetReceiptPdfUrlOk returns a tuple with the ReceiptPdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptPdfUrl

`func (o *InvoicePublic) SetReceiptPdfUrl(v string)`

SetReceiptPdfUrl sets ReceiptPdfUrl field to given value.


### GetRemainingAmountAtom

`func (o *InvoicePublic) GetRemainingAmountAtom() int32`

GetRemainingAmountAtom returns the RemainingAmountAtom field if non-nil, zero value otherwise.

### GetRemainingAmountAtomOk

`func (o *InvoicePublic) GetRemainingAmountAtomOk() (*int32, bool)`

GetRemainingAmountAtomOk returns a tuple with the RemainingAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingAmountAtom

`func (o *InvoicePublic) SetRemainingAmountAtom(v int32)`

SetRemainingAmountAtom sets RemainingAmountAtom field to given value.


### GetStatus

`func (o *InvoicePublic) GetStatus() InvoiceStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoicePublic) GetStatusOk() (*InvoiceStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoicePublic) SetStatus(v InvoiceStatusEnum)`

SetStatus sets Status field to given value.


### GetTaxAmountAtom

`func (o *InvoicePublic) GetTaxAmountAtom() int32`

GetTaxAmountAtom returns the TaxAmountAtom field if non-nil, zero value otherwise.

### GetTaxAmountAtomOk

`func (o *InvoicePublic) GetTaxAmountAtomOk() (*int32, bool)`

GetTaxAmountAtomOk returns a tuple with the TaxAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountAtom

`func (o *InvoicePublic) SetTaxAmountAtom(v int32)`

SetTaxAmountAtom sets TaxAmountAtom field to given value.


### GetTotalAmountAtom

`func (o *InvoicePublic) GetTotalAmountAtom() int32`

GetTotalAmountAtom returns the TotalAmountAtom field if non-nil, zero value otherwise.

### GetTotalAmountAtomOk

`func (o *InvoicePublic) GetTotalAmountAtomOk() (*int32, bool)`

GetTotalAmountAtomOk returns a tuple with the TotalAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountAtom

`func (o *InvoicePublic) SetTotalAmountAtom(v int32)`

SetTotalAmountAtom sets TotalAmountAtom field to given value.


### GetTotalDiscountAmountAtoms

`func (o *InvoicePublic) GetTotalDiscountAmountAtoms() []InvoiceDiscountAmountsExternal`

GetTotalDiscountAmountAtoms returns the TotalDiscountAmountAtoms field if non-nil, zero value otherwise.

### GetTotalDiscountAmountAtomsOk

`func (o *InvoicePublic) GetTotalDiscountAmountAtomsOk() (*[]InvoiceDiscountAmountsExternal, bool)`

GetTotalDiscountAmountAtomsOk returns a tuple with the TotalDiscountAmountAtoms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountAmountAtoms

`func (o *InvoicePublic) SetTotalDiscountAmountAtoms(v []InvoiceDiscountAmountsExternal)`

SetTotalDiscountAmountAtoms sets TotalDiscountAmountAtoms field to given value.

### HasTotalDiscountAmountAtoms

`func (o *InvoicePublic) HasTotalDiscountAmountAtoms() bool`

HasTotalDiscountAmountAtoms returns a boolean if a field has been set.

### GetTotalExcludingTaxesAmountAtom

`func (o *InvoicePublic) GetTotalExcludingTaxesAmountAtom() int32`

GetTotalExcludingTaxesAmountAtom returns the TotalExcludingTaxesAmountAtom field if non-nil, zero value otherwise.

### GetTotalExcludingTaxesAmountAtomOk

`func (o *InvoicePublic) GetTotalExcludingTaxesAmountAtomOk() (*int32, bool)`

GetTotalExcludingTaxesAmountAtomOk returns a tuple with the TotalExcludingTaxesAmountAtom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExcludingTaxesAmountAtom

`func (o *InvoicePublic) SetTotalExcludingTaxesAmountAtom(v int32)`

SetTotalExcludingTaxesAmountAtom sets TotalExcludingTaxesAmountAtom field to given value.


### GetTrialEndForSub

`func (o *InvoicePublic) GetTrialEndForSub() time.Time`

GetTrialEndForSub returns the TrialEndForSub field if non-nil, zero value otherwise.

### GetTrialEndForSubOk

`func (o *InvoicePublic) GetTrialEndForSubOk() (*time.Time, bool)`

GetTrialEndForSubOk returns a tuple with the TrialEndForSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndForSub

`func (o *InvoicePublic) SetTrialEndForSub(v time.Time)`

SetTrialEndForSub sets TrialEndForSub field to given value.


### SetTrialEndForSubNil

`func (o *InvoicePublic) SetTrialEndForSubNil(b bool)`

 SetTrialEndForSubNil sets the value for TrialEndForSub to be an explicit nil

### UnsetTrialEndForSub
`func (o *InvoicePublic) UnsetTrialEndForSub()`

UnsetTrialEndForSub ensures that no value is present for TrialEndForSub, not even an explicit nil
### GetTrialStartForSub

`func (o *InvoicePublic) GetTrialStartForSub() time.Time`

GetTrialStartForSub returns the TrialStartForSub field if non-nil, zero value otherwise.

### GetTrialStartForSubOk

`func (o *InvoicePublic) GetTrialStartForSubOk() (*time.Time, bool)`

GetTrialStartForSubOk returns a tuple with the TrialStartForSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStartForSub

`func (o *InvoicePublic) SetTrialStartForSub(v time.Time)`

SetTrialStartForSub sets TrialStartForSub field to given value.


### SetTrialStartForSubNil

`func (o *InvoicePublic) SetTrialStartForSubNil(b bool)`

 SetTrialStartForSubNil sets the value for TrialStartForSub to be an explicit nil

### UnsetTrialStartForSub
`func (o *InvoicePublic) UnsetTrialStartForSub()`

UnsetTrialStartForSub ensures that no value is present for TrialStartForSub, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


