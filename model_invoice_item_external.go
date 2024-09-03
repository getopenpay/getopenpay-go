/*
OpenPay API

super charge your subscription management.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the InvoiceItemExternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceItemExternal{}

// InvoiceItemExternal struct for InvoiceItemExternal
type InvoiceItemExternal struct {
	// Unique identifier of the invoice_item.
	Id string `json:"id"`
	Object *ObjectName `json:"object,omitempty"`
	// DateTime at which the object was created, in 'ISO 8601' format.
	CreatedAt time.Time `json:"created_at"`
	// DateTime at which the object was updated, in 'ISO 8601' format.
	UpdatedAt time.Time `json:"updated_at"`
	// If true, indicates that this object has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	InvoiceId NullableString `json:"invoice_id"`
	// Unique identifier of the account.
	AccountId string `json:"account_id"`
	// Total amount of invoice_item in atomic units (in USD this is cents).
	AmountAtom int32 `json:"amount_atom"`
	Currency CurrencyEnum `json:"currency"`
	// Unique identifier of the customer.
	CustomerId string `json:"customer_id"`
	Description NullableString `json:"description"`
	// End of the usage period of the invoice_item. It is in 'ISO 8601' format.
	PeriodEnd time.Time `json:"period_end"`
	// Start of the usage period of the invoice_item. It is in 'ISO 8601' format.
	PeriodStart time.Time `json:"period_start"`
	// Unique identifier of the price.
	PriceId string `json:"price_id"`
	PriceType PriceTypeEnum `json:"price_type"`
	// Unique identifier of the product.
	ProductId string `json:"product_id"`
	// Whether the invoice item was created automatically as a proration adjustment when the customer switched plans.
	Proration bool `json:"proration"`
	// Quantity of the invoice_item.
	Quantity int32 `json:"quantity"`
	SubscriptionId NullableString `json:"subscription_id"`
	SubscriptionItemId NullableString `json:"subscription_item_id,omitempty"`
	SubscriptionCancelledAt NullableTime `json:"subscription_cancelled_at,omitempty"`
	Discounts []string `json:"discounts"`
	DiscountAmountAtoms []InvoiceItemDiscountAmountsExternal `json:"discount_amount_atoms"`
	// Total amount of invoice_item in atomic units considering discounts. Contains both invoice-level and invoice item-level discounts
	AmountAtomConsideringDiscountApplied int32 `json:"amount_atom_considering_discount_applied"`
	ProrationDetailsDebitInvoiceItem NullableString `json:"proration_details_debit_invoice_item"`
}

type _InvoiceItemExternal InvoiceItemExternal

// NewInvoiceItemExternal instantiates a new InvoiceItemExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceItemExternal(id string, createdAt time.Time, updatedAt time.Time, invoiceId NullableString, accountId string, amountAtom int32, currency CurrencyEnum, customerId string, description NullableString, periodEnd time.Time, periodStart time.Time, priceId string, priceType PriceTypeEnum, productId string, proration bool, quantity int32, subscriptionId NullableString, discounts []string, discountAmountAtoms []InvoiceItemDiscountAmountsExternal, amountAtomConsideringDiscountApplied int32, prorationDetailsDebitInvoiceItem NullableString) *InvoiceItemExternal {
	this := InvoiceItemExternal{}
	this.Id = id
	var object ObjectName = OBJECTNAME_INVOICE_ITEM
	this.Object = &object
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	this.InvoiceId = invoiceId
	this.AccountId = accountId
	this.AmountAtom = amountAtom
	this.Currency = currency
	this.CustomerId = customerId
	this.Description = description
	this.PeriodEnd = periodEnd
	this.PeriodStart = periodStart
	this.PriceId = priceId
	this.PriceType = priceType
	this.ProductId = productId
	this.Proration = proration
	this.Quantity = quantity
	this.SubscriptionId = subscriptionId
	this.Discounts = discounts
	this.DiscountAmountAtoms = discountAmountAtoms
	this.AmountAtomConsideringDiscountApplied = amountAtomConsideringDiscountApplied
	this.ProrationDetailsDebitInvoiceItem = prorationDetailsDebitInvoiceItem
	return &this
}

// NewInvoiceItemExternalWithDefaults instantiates a new InvoiceItemExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceItemExternalWithDefaults() *InvoiceItemExternal {
	this := InvoiceItemExternal{}
	var object ObjectName = OBJECTNAME_INVOICE_ITEM
	this.Object = &object
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	return &this
}

// GetId returns the Id field value
func (o *InvoiceItemExternal) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InvoiceItemExternal) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *InvoiceItemExternal) GetObject() ObjectName {
	if o == nil || IsNil(o.Object) {
		var ret ObjectName
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetObjectOk() (*ObjectName, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *InvoiceItemExternal) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given ObjectName and assigns it to the Object field.
func (o *InvoiceItemExternal) SetObject(v ObjectName) {
	o.Object = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *InvoiceItemExternal) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InvoiceItemExternal) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *InvoiceItemExternal) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *InvoiceItemExternal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *InvoiceItemExternal) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *InvoiceItemExternal) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *InvoiceItemExternal) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetInvoiceId returns the InvoiceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvoiceItemExternal) GetInvoiceId() string {
	if o == nil || o.InvoiceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceItemExternal) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// SetInvoiceId sets field value
func (o *InvoiceItemExternal) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}

// GetAccountId returns the AccountId field value
func (o *InvoiceItemExternal) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *InvoiceItemExternal) SetAccountId(v string) {
	o.AccountId = v
}

// GetAmountAtom returns the AmountAtom field value
func (o *InvoiceItemExternal) GetAmountAtom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountAtom
}

// GetAmountAtomOk returns a tuple with the AmountAtom field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetAmountAtomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountAtom, true
}

// SetAmountAtom sets field value
func (o *InvoiceItemExternal) SetAmountAtom(v int32) {
	o.AmountAtom = v
}

// GetCurrency returns the Currency field value
func (o *InvoiceItemExternal) GetCurrency() CurrencyEnum {
	if o == nil {
		var ret CurrencyEnum
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *InvoiceItemExternal) SetCurrency(v CurrencyEnum) {
	o.Currency = v
}

// GetCustomerId returns the CustomerId field value
func (o *InvoiceItemExternal) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *InvoiceItemExternal) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvoiceItemExternal) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceItemExternal) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *InvoiceItemExternal) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetPeriodEnd returns the PeriodEnd field value
func (o *InvoiceItemExternal) GetPeriodEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetPeriodEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodEnd, true
}

// SetPeriodEnd sets field value
func (o *InvoiceItemExternal) SetPeriodEnd(v time.Time) {
	o.PeriodEnd = v
}

// GetPeriodStart returns the PeriodStart field value
func (o *InvoiceItemExternal) GetPeriodStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetPeriodStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodStart, true
}

// SetPeriodStart sets field value
func (o *InvoiceItemExternal) SetPeriodStart(v time.Time) {
	o.PeriodStart = v
}

// GetPriceId returns the PriceId field value
func (o *InvoiceItemExternal) GetPriceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PriceId
}

// GetPriceIdOk returns a tuple with the PriceId field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceId, true
}

// SetPriceId sets field value
func (o *InvoiceItemExternal) SetPriceId(v string) {
	o.PriceId = v
}

// GetPriceType returns the PriceType field value
func (o *InvoiceItemExternal) GetPriceType() PriceTypeEnum {
	if o == nil {
		var ret PriceTypeEnum
		return ret
	}

	return o.PriceType
}

// GetPriceTypeOk returns a tuple with the PriceType field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetPriceTypeOk() (*PriceTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceType, true
}

// SetPriceType sets field value
func (o *InvoiceItemExternal) SetPriceType(v PriceTypeEnum) {
	o.PriceType = v
}

// GetProductId returns the ProductId field value
func (o *InvoiceItemExternal) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *InvoiceItemExternal) SetProductId(v string) {
	o.ProductId = v
}

// GetProration returns the Proration field value
func (o *InvoiceItemExternal) GetProration() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Proration
}

// GetProrationOk returns a tuple with the Proration field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetProrationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proration, true
}

// SetProration sets field value
func (o *InvoiceItemExternal) SetProration(v bool) {
	o.Proration = v
}

// GetQuantity returns the Quantity field value
func (o *InvoiceItemExternal) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *InvoiceItemExternal) SetQuantity(v int32) {
	o.Quantity = v
}

// GetSubscriptionId returns the SubscriptionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvoiceItemExternal) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceItemExternal) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// SetSubscriptionId sets field value
func (o *InvoiceItemExternal) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}

// GetSubscriptionItemId returns the SubscriptionItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceItemExternal) GetSubscriptionItemId() string {
	if o == nil || IsNil(o.SubscriptionItemId.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionItemId.Get()
}

// GetSubscriptionItemIdOk returns a tuple with the SubscriptionItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceItemExternal) GetSubscriptionItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionItemId.Get(), o.SubscriptionItemId.IsSet()
}

// HasSubscriptionItemId returns a boolean if a field has been set.
func (o *InvoiceItemExternal) HasSubscriptionItemId() bool {
	if o != nil && o.SubscriptionItemId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionItemId gets a reference to the given NullableString and assigns it to the SubscriptionItemId field.
func (o *InvoiceItemExternal) SetSubscriptionItemId(v string) {
	o.SubscriptionItemId.Set(&v)
}
// SetSubscriptionItemIdNil sets the value for SubscriptionItemId to be an explicit nil
func (o *InvoiceItemExternal) SetSubscriptionItemIdNil() {
	o.SubscriptionItemId.Set(nil)
}

// UnsetSubscriptionItemId ensures that no value is present for SubscriptionItemId, not even an explicit nil
func (o *InvoiceItemExternal) UnsetSubscriptionItemId() {
	o.SubscriptionItemId.Unset()
}

// GetSubscriptionCancelledAt returns the SubscriptionCancelledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceItemExternal) GetSubscriptionCancelledAt() time.Time {
	if o == nil || IsNil(o.SubscriptionCancelledAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SubscriptionCancelledAt.Get()
}

// GetSubscriptionCancelledAtOk returns a tuple with the SubscriptionCancelledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceItemExternal) GetSubscriptionCancelledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionCancelledAt.Get(), o.SubscriptionCancelledAt.IsSet()
}

// HasSubscriptionCancelledAt returns a boolean if a field has been set.
func (o *InvoiceItemExternal) HasSubscriptionCancelledAt() bool {
	if o != nil && o.SubscriptionCancelledAt.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionCancelledAt gets a reference to the given NullableTime and assigns it to the SubscriptionCancelledAt field.
func (o *InvoiceItemExternal) SetSubscriptionCancelledAt(v time.Time) {
	o.SubscriptionCancelledAt.Set(&v)
}
// SetSubscriptionCancelledAtNil sets the value for SubscriptionCancelledAt to be an explicit nil
func (o *InvoiceItemExternal) SetSubscriptionCancelledAtNil() {
	o.SubscriptionCancelledAt.Set(nil)
}

// UnsetSubscriptionCancelledAt ensures that no value is present for SubscriptionCancelledAt, not even an explicit nil
func (o *InvoiceItemExternal) UnsetSubscriptionCancelledAt() {
	o.SubscriptionCancelledAt.Unset()
}

// GetDiscounts returns the Discounts field value
func (o *InvoiceItemExternal) GetDiscounts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Discounts
}

// GetDiscountsOk returns a tuple with the Discounts field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetDiscountsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Discounts, true
}

// SetDiscounts sets field value
func (o *InvoiceItemExternal) SetDiscounts(v []string) {
	o.Discounts = v
}

// GetDiscountAmountAtoms returns the DiscountAmountAtoms field value
func (o *InvoiceItemExternal) GetDiscountAmountAtoms() []InvoiceItemDiscountAmountsExternal {
	if o == nil {
		var ret []InvoiceItemDiscountAmountsExternal
		return ret
	}

	return o.DiscountAmountAtoms
}

// GetDiscountAmountAtomsOk returns a tuple with the DiscountAmountAtoms field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetDiscountAmountAtomsOk() ([]InvoiceItemDiscountAmountsExternal, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountAmountAtoms, true
}

// SetDiscountAmountAtoms sets field value
func (o *InvoiceItemExternal) SetDiscountAmountAtoms(v []InvoiceItemDiscountAmountsExternal) {
	o.DiscountAmountAtoms = v
}

// GetAmountAtomConsideringDiscountApplied returns the AmountAtomConsideringDiscountApplied field value
func (o *InvoiceItemExternal) GetAmountAtomConsideringDiscountApplied() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountAtomConsideringDiscountApplied
}

// GetAmountAtomConsideringDiscountAppliedOk returns a tuple with the AmountAtomConsideringDiscountApplied field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemExternal) GetAmountAtomConsideringDiscountAppliedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountAtomConsideringDiscountApplied, true
}

// SetAmountAtomConsideringDiscountApplied sets field value
func (o *InvoiceItemExternal) SetAmountAtomConsideringDiscountApplied(v int32) {
	o.AmountAtomConsideringDiscountApplied = v
}

// GetProrationDetailsDebitInvoiceItem returns the ProrationDetailsDebitInvoiceItem field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvoiceItemExternal) GetProrationDetailsDebitInvoiceItem() string {
	if o == nil || o.ProrationDetailsDebitInvoiceItem.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProrationDetailsDebitInvoiceItem.Get()
}

// GetProrationDetailsDebitInvoiceItemOk returns a tuple with the ProrationDetailsDebitInvoiceItem field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceItemExternal) GetProrationDetailsDebitInvoiceItemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProrationDetailsDebitInvoiceItem.Get(), o.ProrationDetailsDebitInvoiceItem.IsSet()
}

// SetProrationDetailsDebitInvoiceItem sets field value
func (o *InvoiceItemExternal) SetProrationDetailsDebitInvoiceItem(v string) {
	o.ProrationDetailsDebitInvoiceItem.Set(&v)
}

func (o InvoiceItemExternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceItemExternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	toSerialize["invoice_id"] = o.InvoiceId.Get()
	toSerialize["account_id"] = o.AccountId
	toSerialize["amount_atom"] = o.AmountAtom
	toSerialize["currency"] = o.Currency
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["description"] = o.Description.Get()
	toSerialize["period_end"] = o.PeriodEnd
	toSerialize["period_start"] = o.PeriodStart
	toSerialize["price_id"] = o.PriceId
	toSerialize["price_type"] = o.PriceType
	toSerialize["product_id"] = o.ProductId
	toSerialize["proration"] = o.Proration
	toSerialize["quantity"] = o.Quantity
	toSerialize["subscription_id"] = o.SubscriptionId.Get()
	if o.SubscriptionItemId.IsSet() {
		toSerialize["subscription_item_id"] = o.SubscriptionItemId.Get()
	}
	if o.SubscriptionCancelledAt.IsSet() {
		toSerialize["subscription_cancelled_at"] = o.SubscriptionCancelledAt.Get()
	}
	toSerialize["discounts"] = o.Discounts
	toSerialize["discount_amount_atoms"] = o.DiscountAmountAtoms
	toSerialize["amount_atom_considering_discount_applied"] = o.AmountAtomConsideringDiscountApplied
	toSerialize["proration_details_debit_invoice_item"] = o.ProrationDetailsDebitInvoiceItem.Get()
	return toSerialize, nil
}

func (o *InvoiceItemExternal) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"invoice_id",
		"account_id",
		"amount_atom",
		"currency",
		"customer_id",
		"description",
		"period_end",
		"period_start",
		"price_id",
		"price_type",
		"product_id",
		"proration",
		"quantity",
		"subscription_id",
		"discounts",
		"discount_amount_atoms",
		"amount_atom_considering_discount_applied",
		"proration_details_debit_invoice_item",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInvoiceItemExternal := _InvoiceItemExternal{}

	err = json.Unmarshal(bytes, &varInvoiceItemExternal)

	if err != nil {
		return err
	}

	*o = InvoiceItemExternal(varInvoiceItemExternal)

	return err
}

type NullableInvoiceItemExternal struct {
	value *InvoiceItemExternal
	isSet bool
}

func (v NullableInvoiceItemExternal) Get() *InvoiceItemExternal {
	return v.value
}

func (v *NullableInvoiceItemExternal) Set(val *InvoiceItemExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceItemExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceItemExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceItemExternal(val *InvoiceItemExternal) *NullableInvoiceItemExternal {
	return &NullableInvoiceItemExternal{value: val, isSet: true}
}

func (v NullableInvoiceItemExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceItemExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


