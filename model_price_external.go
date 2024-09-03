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

// checks if the PriceExternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceExternal{}

// PriceExternal struct for PriceExternal
type PriceExternal struct {
	// Unique identifier of the price.
	Id string `json:"id"`
	Object *ObjectName `json:"object,omitempty"`
	// DateTime at which the object was created, in 'ISO 8601' format.
	CreatedAt time.Time `json:"created_at"`
	// DateTime at which the object was updated, in 'ISO 8601' format.
	UpdatedAt time.Time `json:"updated_at"`
	// If true, indicates that this object has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// Whether the price can be used for new purchases. Price can be activated later.
	IsActive bool `json:"is_active"`
	// Name of the product associated with this price.
	ProductName string `json:"product_name"`
	// Unique identifier of the product.
	ProductId string `json:"product_id"`
	Product NullableProductExternal `json:"product,omitempty"`
	InternalDescription NullableString `json:"internal_description"`
	BillingScheme BillingSchemeEnum `json:"billing_scheme"`
	UnitAmountAtom NullableInt32 `json:"unit_amount_atom"`
	// This transformation will be applied on quantity before mulitplying by 'unit_amount_atom'.
	TransformQuantityDivideBy float32 `json:"transform_quantity_divide_by"`
	PriceTiers []PriceTierExternal `json:"price_tiers"`
	PriceType PriceTypeEnum `json:"price_type"`
	RecurringDetails NullableRecurringDetails `json:"recurring_details"`
	TiersMode NullablePricingTiersEnum `json:"tiers_mode"`
	BillingInterval NullableCalendarIntervalEnum `json:"billing_interval"`
	BillingIntervalCount NullableInt32 `json:"billing_interval_count"`
	ContractTermMultiple NullableInt32 `json:"contract_term_multiple"`
	ContractAutoRenew NullableBool `json:"contract_auto_renew"`
	InvoiceSettings NullableInvoiceSettings `json:"invoice_settings"`
	// Whether the price is licensed or not.
	IsLicensed bool `json:"is_licensed"`
	// Whether the price is exclusive or not.
	IsExclusive *bool `json:"is_exclusive,omitempty"`
	ListedExclusivelyForCustomers []string `json:"listed_exclusively_for_customers"`
	CanOnlyBePurchasedWith []string `json:"can_only_be_purchased_with"`
	OptionalAddOns []string `json:"optional_add_ons"`
	// If the price can be updated or not.
	EligibleForUpdates bool `json:"eligible_for_updates"`
	Currency CurrencyEnum `json:"currency"`
	IsDefault NullableBool `json:"is_default"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _PriceExternal PriceExternal

// NewPriceExternal instantiates a new PriceExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceExternal(id string, createdAt time.Time, updatedAt time.Time, isActive bool, productName string, productId string, internalDescription NullableString, billingScheme BillingSchemeEnum, unitAmountAtom NullableInt32, transformQuantityDivideBy float32, priceTiers []PriceTierExternal, priceType PriceTypeEnum, recurringDetails NullableRecurringDetails, tiersMode NullablePricingTiersEnum, billingInterval NullableCalendarIntervalEnum, billingIntervalCount NullableInt32, contractTermMultiple NullableInt32, contractAutoRenew NullableBool, invoiceSettings NullableInvoiceSettings, isLicensed bool, listedExclusivelyForCustomers []string, canOnlyBePurchasedWith []string, optionalAddOns []string, eligibleForUpdates bool, currency CurrencyEnum, isDefault NullableBool) *PriceExternal {
	this := PriceExternal{}
	this.Id = id
	var object ObjectName = OBJECTNAME_PRICE
	this.Object = &object
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	this.IsActive = isActive
	this.ProductName = productName
	this.ProductId = productId
	this.InternalDescription = internalDescription
	this.BillingScheme = billingScheme
	this.UnitAmountAtom = unitAmountAtom
	this.TransformQuantityDivideBy = transformQuantityDivideBy
	this.PriceTiers = priceTiers
	this.PriceType = priceType
	this.RecurringDetails = recurringDetails
	this.TiersMode = tiersMode
	this.BillingInterval = billingInterval
	this.BillingIntervalCount = billingIntervalCount
	this.ContractTermMultiple = contractTermMultiple
	this.ContractAutoRenew = contractAutoRenew
	this.InvoiceSettings = invoiceSettings
	this.IsLicensed = isLicensed
	var isExclusive bool = false
	this.IsExclusive = &isExclusive
	this.ListedExclusivelyForCustomers = listedExclusivelyForCustomers
	this.CanOnlyBePurchasedWith = canOnlyBePurchasedWith
	this.OptionalAddOns = optionalAddOns
	this.EligibleForUpdates = eligibleForUpdates
	this.Currency = currency
	this.IsDefault = isDefault
	return &this
}

// NewPriceExternalWithDefaults instantiates a new PriceExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceExternalWithDefaults() *PriceExternal {
	this := PriceExternal{}
	var object ObjectName = OBJECTNAME_PRICE
	this.Object = &object
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	var isExclusive bool = false
	this.IsExclusive = &isExclusive
	return &this
}

// GetId returns the Id field value
func (o *PriceExternal) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PriceExternal) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *PriceExternal) GetObject() ObjectName {
	if o == nil || IsNil(o.Object) {
		var ret ObjectName
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetObjectOk() (*ObjectName, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *PriceExternal) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given ObjectName and assigns it to the Object field.
func (o *PriceExternal) SetObject(v ObjectName) {
	o.Object = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PriceExternal) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PriceExternal) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PriceExternal) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PriceExternal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *PriceExternal) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *PriceExternal) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *PriceExternal) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetIsActive returns the IsActive field value
func (o *PriceExternal) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *PriceExternal) SetIsActive(v bool) {
	o.IsActive = v
}

// GetProductName returns the ProductName field value
func (o *PriceExternal) GetProductName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductName, true
}

// SetProductName sets field value
func (o *PriceExternal) SetProductName(v string) {
	o.ProductName = v
}

// GetProductId returns the ProductId field value
func (o *PriceExternal) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *PriceExternal) SetProductId(v string) {
	o.ProductId = v
}

// GetProduct returns the Product field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceExternal) GetProduct() ProductExternal {
	if o == nil || IsNil(o.Product.Get()) {
		var ret ProductExternal
		return ret
	}
	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceExternal) GetProductOk() (*ProductExternal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// HasProduct returns a boolean if a field has been set.
func (o *PriceExternal) HasProduct() bool {
	if o != nil && o.Product.IsSet() {
		return true
	}

	return false
}

// SetProduct gets a reference to the given NullableProductExternal and assigns it to the Product field.
func (o *PriceExternal) SetProduct(v ProductExternal) {
	o.Product.Set(&v)
}
// SetProductNil sets the value for Product to be an explicit nil
func (o *PriceExternal) SetProductNil() {
	o.Product.Set(nil)
}

// UnsetProduct ensures that no value is present for Product, not even an explicit nil
func (o *PriceExternal) UnsetProduct() {
	o.Product.Unset()
}

// GetInternalDescription returns the InternalDescription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PriceExternal) GetInternalDescription() string {
	if o == nil || o.InternalDescription.Get() == nil {
		var ret string
		return ret
	}

	return *o.InternalDescription.Get()
}

// GetInternalDescriptionOk returns a tuple with the InternalDescription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceExternal) GetInternalDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternalDescription.Get(), o.InternalDescription.IsSet()
}

// SetInternalDescription sets field value
func (o *PriceExternal) SetInternalDescription(v string) {
	o.InternalDescription.Set(&v)
}

// GetBillingScheme returns the BillingScheme field value
func (o *PriceExternal) GetBillingScheme() BillingSchemeEnum {
	if o == nil {
		var ret BillingSchemeEnum
		return ret
	}

	return o.BillingScheme
}

// GetBillingSchemeOk returns a tuple with the BillingScheme field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetBillingSchemeOk() (*BillingSchemeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingScheme, true
}

// SetBillingScheme sets field value
func (o *PriceExternal) SetBillingScheme(v BillingSchemeEnum) {
	o.BillingScheme = v
}

// GetUnitAmountAtom returns the UnitAmountAtom field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PriceExternal) GetUnitAmountAtom() int32 {
	if o == nil || o.UnitAmountAtom.Get() == nil {
		var ret int32
		return ret
	}

	return *o.UnitAmountAtom.Get()
}

// GetUnitAmountAtomOk returns a tuple with the UnitAmountAtom field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceExternal) GetUnitAmountAtomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitAmountAtom.Get(), o.UnitAmountAtom.IsSet()
}

// SetUnitAmountAtom sets field value
func (o *PriceExternal) SetUnitAmountAtom(v int32) {
	o.UnitAmountAtom.Set(&v)
}

// GetTransformQuantityDivideBy returns the TransformQuantityDivideBy field value
func (o *PriceExternal) GetTransformQuantityDivideBy() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TransformQuantityDivideBy
}

// GetTransformQuantityDivideByOk returns a tuple with the TransformQuantityDivideBy field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetTransformQuantityDivideByOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransformQuantityDivideBy, true
}

// SetTransformQuantityDivideBy sets field value
func (o *PriceExternal) SetTransformQuantityDivideBy(v float32) {
	o.TransformQuantityDivideBy = v
}

// GetPriceTiers returns the PriceTiers field value
func (o *PriceExternal) GetPriceTiers() []PriceTierExternal {
	if o == nil {
		var ret []PriceTierExternal
		return ret
	}

	return o.PriceTiers
}

// GetPriceTiersOk returns a tuple with the PriceTiers field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetPriceTiersOk() ([]PriceTierExternal, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceTiers, true
}

// SetPriceTiers sets field value
func (o *PriceExternal) SetPriceTiers(v []PriceTierExternal) {
	o.PriceTiers = v
}

// GetPriceType returns the PriceType field value
func (o *PriceExternal) GetPriceType() PriceTypeEnum {
	if o == nil {
		var ret PriceTypeEnum
		return ret
	}

	return o.PriceType
}

// GetPriceTypeOk returns a tuple with the PriceType field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetPriceTypeOk() (*PriceTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceType, true
}

// SetPriceType sets field value
func (o *PriceExternal) SetPriceType(v PriceTypeEnum) {
	o.PriceType = v
}

// GetRecurringDetails returns the RecurringDetails field value
// If the value is explicit nil, the zero value for RecurringDetails will be returned
func (o *PriceExternal) GetRecurringDetails() RecurringDetails {
	if o == nil || o.RecurringDetails.Get() == nil {
		var ret RecurringDetails
		return ret
	}

	return *o.RecurringDetails.Get()
}

// GetRecurringDetailsOk returns a tuple with the RecurringDetails field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceExternal) GetRecurringDetailsOk() (*RecurringDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringDetails.Get(), o.RecurringDetails.IsSet()
}

// SetRecurringDetails sets field value
func (o *PriceExternal) SetRecurringDetails(v RecurringDetails) {
	o.RecurringDetails.Set(&v)
}

// GetTiersMode returns the TiersMode field value
// If the value is explicit nil, the zero value for PricingTiersEnum will be returned
func (o *PriceExternal) GetTiersMode() PricingTiersEnum {
	if o == nil || o.TiersMode.Get() == nil {
		var ret PricingTiersEnum
		return ret
	}

	return *o.TiersMode.Get()
}

// GetTiersModeOk returns a tuple with the TiersMode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceExternal) GetTiersModeOk() (*PricingTiersEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.TiersMode.Get(), o.TiersMode.IsSet()
}

// SetTiersMode sets field value
func (o *PriceExternal) SetTiersMode(v PricingTiersEnum) {
	o.TiersMode.Set(&v)
}

// GetBillingInterval returns the BillingInterval field value
// If the value is explicit nil, the zero value for CalendarIntervalEnum will be returned
func (o *PriceExternal) GetBillingInterval() CalendarIntervalEnum {
	if o == nil || o.BillingInterval.Get() == nil {
		var ret CalendarIntervalEnum
		return ret
	}

	return *o.BillingInterval.Get()
}

// GetBillingIntervalOk returns a tuple with the BillingInterval field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceExternal) GetBillingIntervalOk() (*CalendarIntervalEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingInterval.Get(), o.BillingInterval.IsSet()
}

// SetBillingInterval sets field value
func (o *PriceExternal) SetBillingInterval(v CalendarIntervalEnum) {
	o.BillingInterval.Set(&v)
}

// GetBillingIntervalCount returns the BillingIntervalCount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PriceExternal) GetBillingIntervalCount() int32 {
	if o == nil || o.BillingIntervalCount.Get() == nil {
		var ret int32
		return ret
	}

	return *o.BillingIntervalCount.Get()
}

// GetBillingIntervalCountOk returns a tuple with the BillingIntervalCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceExternal) GetBillingIntervalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingIntervalCount.Get(), o.BillingIntervalCount.IsSet()
}

// SetBillingIntervalCount sets field value
func (o *PriceExternal) SetBillingIntervalCount(v int32) {
	o.BillingIntervalCount.Set(&v)
}

// GetContractTermMultiple returns the ContractTermMultiple field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PriceExternal) GetContractTermMultiple() int32 {
	if o == nil || o.ContractTermMultiple.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ContractTermMultiple.Get()
}

// GetContractTermMultipleOk returns a tuple with the ContractTermMultiple field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceExternal) GetContractTermMultipleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractTermMultiple.Get(), o.ContractTermMultiple.IsSet()
}

// SetContractTermMultiple sets field value
func (o *PriceExternal) SetContractTermMultiple(v int32) {
	o.ContractTermMultiple.Set(&v)
}

// GetContractAutoRenew returns the ContractAutoRenew field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *PriceExternal) GetContractAutoRenew() bool {
	if o == nil || o.ContractAutoRenew.Get() == nil {
		var ret bool
		return ret
	}

	return *o.ContractAutoRenew.Get()
}

// GetContractAutoRenewOk returns a tuple with the ContractAutoRenew field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceExternal) GetContractAutoRenewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAutoRenew.Get(), o.ContractAutoRenew.IsSet()
}

// SetContractAutoRenew sets field value
func (o *PriceExternal) SetContractAutoRenew(v bool) {
	o.ContractAutoRenew.Set(&v)
}

// GetInvoiceSettings returns the InvoiceSettings field value
// If the value is explicit nil, the zero value for InvoiceSettings will be returned
func (o *PriceExternal) GetInvoiceSettings() InvoiceSettings {
	if o == nil || o.InvoiceSettings.Get() == nil {
		var ret InvoiceSettings
		return ret
	}

	return *o.InvoiceSettings.Get()
}

// GetInvoiceSettingsOk returns a tuple with the InvoiceSettings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceExternal) GetInvoiceSettingsOk() (*InvoiceSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceSettings.Get(), o.InvoiceSettings.IsSet()
}

// SetInvoiceSettings sets field value
func (o *PriceExternal) SetInvoiceSettings(v InvoiceSettings) {
	o.InvoiceSettings.Set(&v)
}

// GetIsLicensed returns the IsLicensed field value
func (o *PriceExternal) GetIsLicensed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLicensed
}

// GetIsLicensedOk returns a tuple with the IsLicensed field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetIsLicensedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLicensed, true
}

// SetIsLicensed sets field value
func (o *PriceExternal) SetIsLicensed(v bool) {
	o.IsLicensed = v
}

// GetIsExclusive returns the IsExclusive field value if set, zero value otherwise.
func (o *PriceExternal) GetIsExclusive() bool {
	if o == nil || IsNil(o.IsExclusive) {
		var ret bool
		return ret
	}
	return *o.IsExclusive
}

// GetIsExclusiveOk returns a tuple with the IsExclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetIsExclusiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExclusive) {
		return nil, false
	}
	return o.IsExclusive, true
}

// HasIsExclusive returns a boolean if a field has been set.
func (o *PriceExternal) HasIsExclusive() bool {
	if o != nil && !IsNil(o.IsExclusive) {
		return true
	}

	return false
}

// SetIsExclusive gets a reference to the given bool and assigns it to the IsExclusive field.
func (o *PriceExternal) SetIsExclusive(v bool) {
	o.IsExclusive = &v
}

// GetListedExclusivelyForCustomers returns the ListedExclusivelyForCustomers field value
func (o *PriceExternal) GetListedExclusivelyForCustomers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ListedExclusivelyForCustomers
}

// GetListedExclusivelyForCustomersOk returns a tuple with the ListedExclusivelyForCustomers field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetListedExclusivelyForCustomersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ListedExclusivelyForCustomers, true
}

// SetListedExclusivelyForCustomers sets field value
func (o *PriceExternal) SetListedExclusivelyForCustomers(v []string) {
	o.ListedExclusivelyForCustomers = v
}

// GetCanOnlyBePurchasedWith returns the CanOnlyBePurchasedWith field value
func (o *PriceExternal) GetCanOnlyBePurchasedWith() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CanOnlyBePurchasedWith
}

// GetCanOnlyBePurchasedWithOk returns a tuple with the CanOnlyBePurchasedWith field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetCanOnlyBePurchasedWithOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanOnlyBePurchasedWith, true
}

// SetCanOnlyBePurchasedWith sets field value
func (o *PriceExternal) SetCanOnlyBePurchasedWith(v []string) {
	o.CanOnlyBePurchasedWith = v
}

// GetOptionalAddOns returns the OptionalAddOns field value
func (o *PriceExternal) GetOptionalAddOns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OptionalAddOns
}

// GetOptionalAddOnsOk returns a tuple with the OptionalAddOns field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetOptionalAddOnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptionalAddOns, true
}

// SetOptionalAddOns sets field value
func (o *PriceExternal) SetOptionalAddOns(v []string) {
	o.OptionalAddOns = v
}

// GetEligibleForUpdates returns the EligibleForUpdates field value
func (o *PriceExternal) GetEligibleForUpdates() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EligibleForUpdates
}

// GetEligibleForUpdatesOk returns a tuple with the EligibleForUpdates field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetEligibleForUpdatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EligibleForUpdates, true
}

// SetEligibleForUpdates sets field value
func (o *PriceExternal) SetEligibleForUpdates(v bool) {
	o.EligibleForUpdates = v
}

// GetCurrency returns the Currency field value
func (o *PriceExternal) GetCurrency() CurrencyEnum {
	if o == nil {
		var ret CurrencyEnum
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PriceExternal) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PriceExternal) SetCurrency(v CurrencyEnum) {
	o.Currency = v
}

// GetIsDefault returns the IsDefault field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *PriceExternal) GetIsDefault() bool {
	if o == nil || o.IsDefault.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceExternal) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// SetIsDefault sets field value
func (o *PriceExternal) SetIsDefault(v bool) {
	o.IsDefault.Set(&v)
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceExternal) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceExternal) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PriceExternal) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PriceExternal) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PriceExternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceExternal) ToMap() (map[string]interface{}, error) {
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
	toSerialize["is_active"] = o.IsActive
	toSerialize["product_name"] = o.ProductName
	toSerialize["product_id"] = o.ProductId
	if o.Product.IsSet() {
		toSerialize["product"] = o.Product.Get()
	}
	toSerialize["internal_description"] = o.InternalDescription.Get()
	toSerialize["billing_scheme"] = o.BillingScheme
	toSerialize["unit_amount_atom"] = o.UnitAmountAtom.Get()
	toSerialize["transform_quantity_divide_by"] = o.TransformQuantityDivideBy
	toSerialize["price_tiers"] = o.PriceTiers
	toSerialize["price_type"] = o.PriceType
	toSerialize["recurring_details"] = o.RecurringDetails.Get()
	toSerialize["tiers_mode"] = o.TiersMode.Get()
	toSerialize["billing_interval"] = o.BillingInterval.Get()
	toSerialize["billing_interval_count"] = o.BillingIntervalCount.Get()
	toSerialize["contract_term_multiple"] = o.ContractTermMultiple.Get()
	toSerialize["contract_auto_renew"] = o.ContractAutoRenew.Get()
	toSerialize["invoice_settings"] = o.InvoiceSettings.Get()
	toSerialize["is_licensed"] = o.IsLicensed
	if !IsNil(o.IsExclusive) {
		toSerialize["is_exclusive"] = o.IsExclusive
	}
	toSerialize["listed_exclusively_for_customers"] = o.ListedExclusivelyForCustomers
	toSerialize["can_only_be_purchased_with"] = o.CanOnlyBePurchasedWith
	toSerialize["optional_add_ons"] = o.OptionalAddOns
	toSerialize["eligible_for_updates"] = o.EligibleForUpdates
	toSerialize["currency"] = o.Currency
	toSerialize["is_default"] = o.IsDefault.Get()
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *PriceExternal) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"is_active",
		"product_name",
		"product_id",
		"internal_description",
		"billing_scheme",
		"unit_amount_atom",
		"transform_quantity_divide_by",
		"price_tiers",
		"price_type",
		"recurring_details",
		"tiers_mode",
		"billing_interval",
		"billing_interval_count",
		"contract_term_multiple",
		"contract_auto_renew",
		"invoice_settings",
		"is_licensed",
		"listed_exclusively_for_customers",
		"can_only_be_purchased_with",
		"optional_add_ons",
		"eligible_for_updates",
		"currency",
		"is_default",
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

	varPriceExternal := _PriceExternal{}

	err = json.Unmarshal(bytes, &varPriceExternal)

	if err != nil {
		return err
	}

	*o = PriceExternal(varPriceExternal)

	return err
}

type NullablePriceExternal struct {
	value *PriceExternal
	isSet bool
}

func (v NullablePriceExternal) Get() *PriceExternal {
	return v.value
}

func (v *NullablePriceExternal) Set(val *PriceExternal) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceExternal) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceExternal(val *PriceExternal) *NullablePriceExternal {
	return &NullablePriceExternal{value: val, isSet: true}
}

func (v NullablePriceExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

