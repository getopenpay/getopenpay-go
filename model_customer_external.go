/*
OpenPay API

super charge your subscription management.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the CustomerExternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerExternal{}

// CustomerExternal struct for CustomerExternal
type CustomerExternal struct {
	// Unique identifier for the account.
	AccountId string `json:"account_id"`
	Address NullableCompleteAddress `json:"address"`
	// List of the customer's balance in the smallest unit of currency (e.g., cents for USD). Positive values indicate the amount owed by the customer, to be applied to the next invoice. Negative values represent credit to apply to future invoices.
	BalanceAtoms []CustomerBalanceExternal `json:"balance_atoms,omitempty"`
	BillingEmail NullableString `json:"billing_email,omitempty"`
	BusinessName NullableString `json:"business_name"`
	// DateTime at which the object was created, in 'ISO 8601' format.
	CreatedAt time.Time `json:"created_at"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	CustomerBillingAddress NullableCompleteAddress `json:"customer_billing_address,omitempty"`
	Discount NullableDiscountExternal `json:"discount,omitempty"`
	// Customer's email address.
	Email string `json:"email"`
	FirstName NullableString `json:"first_name"`
	// Unique identifier of the customer.
	Id string `json:"id"`
	InvoiceSettings NullableCustomerInvoiceSettings `json:"invoice_settings,omitempty"`
	// If true, indicates that this object has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// Language of the customer.
	Language *CustomerLanguage `json:"language,omitempty"`
	LastName NullableString `json:"last_name"`
	LastSuccessfulPaymentIntent NullablePaymentIntentExternal `json:"last_successful_payment_intent,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The monthly recurring revenue of the customer broken down by currency.
	Mrr []CustomerTotalAmount `json:"mrr,omitempty"`
	Notes NullableString `json:"notes"`
	Object *ObjectName `json:"object,omitempty"`
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// List of customer's shipping addresses.
	ShippingAddresses []CompleteAddress `json:"shipping_addresses,omitempty"`
	// Whether email should be sent or not on payment.
	ShouldSendPaymentReceipt bool `json:"should_send_payment_receipt"`
	Status NullableCustomerStatus `json:"status,omitempty"`
	// List of products that the customer is subscribed to.
	SubscribedToProducts []ProductExternal `json:"subscribed_to_products,omitempty"`
	// List of customer's subscriptions.
	Subscriptions []SubscriptionExternal `json:"subscriptions,omitempty"`
	// The total amount refunded to the customer.
	TotalRefunds []CustomerTotalAmount `json:"total_refunds,omitempty"`
	// The total amount spent by the customer.
	TotalSpent []CustomerTotalAmount `json:"total_spent,omitempty"`
	// DateTime at which the object was updated, in 'ISO 8601' format.
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _CustomerExternal CustomerExternal

// NewCustomerExternal instantiates a new CustomerExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerExternal(accountId string, address NullableCompleteAddress, businessName NullableString, createdAt time.Time, email string, firstName NullableString, id string, lastName NullableString, notes NullableString, shouldSendPaymentReceipt bool, updatedAt time.Time) *CustomerExternal {
	this := CustomerExternal{}
	this.AccountId = accountId
	this.Address = address
	this.BusinessName = businessName
	this.CreatedAt = createdAt
	this.Email = email
	this.FirstName = firstName
	this.Id = id
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	this.LastName = lastName
	this.Notes = notes
	this.ShouldSendPaymentReceipt = shouldSendPaymentReceipt
	this.UpdatedAt = updatedAt
	return &this
}

// NewCustomerExternalWithDefaults instantiates a new CustomerExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerExternalWithDefaults() *CustomerExternal {
	this := CustomerExternal{}
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	return &this
}

// GetAccountId returns the AccountId field value
func (o *CustomerExternal) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CustomerExternal) SetAccountId(v string) {
	o.AccountId = v
}

// GetAddress returns the Address field value
// If the value is explicit nil, the zero value for CompleteAddress will be returned
func (o *CustomerExternal) GetAddress() CompleteAddress {
	if o == nil || o.Address.Get() == nil {
		var ret CompleteAddress
		return ret
	}

	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetAddressOk() (*CompleteAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// SetAddress sets field value
func (o *CustomerExternal) SetAddress(v CompleteAddress) {
	o.Address.Set(&v)
}

// GetBalanceAtoms returns the BalanceAtoms field value if set, zero value otherwise.
func (o *CustomerExternal) GetBalanceAtoms() []CustomerBalanceExternal {
	if o == nil || IsNil(o.BalanceAtoms) {
		var ret []CustomerBalanceExternal
		return ret
	}
	return o.BalanceAtoms
}

// GetBalanceAtomsOk returns a tuple with the BalanceAtoms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetBalanceAtomsOk() ([]CustomerBalanceExternal, bool) {
	if o == nil || IsNil(o.BalanceAtoms) {
		return nil, false
	}
	return o.BalanceAtoms, true
}

// HasBalanceAtoms returns a boolean if a field has been set.
func (o *CustomerExternal) HasBalanceAtoms() bool {
	if o != nil && !IsNil(o.BalanceAtoms) {
		return true
	}

	return false
}

// SetBalanceAtoms gets a reference to the given []CustomerBalanceExternal and assigns it to the BalanceAtoms field.
func (o *CustomerExternal) SetBalanceAtoms(v []CustomerBalanceExternal) {
	o.BalanceAtoms = v
}

// GetBillingEmail returns the BillingEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerExternal) GetBillingEmail() string {
	if o == nil || IsNil(o.BillingEmail.Get()) {
		var ret string
		return ret
	}
	return *o.BillingEmail.Get()
}

// GetBillingEmailOk returns a tuple with the BillingEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetBillingEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingEmail.Get(), o.BillingEmail.IsSet()
}

// HasBillingEmail returns a boolean if a field has been set.
func (o *CustomerExternal) HasBillingEmail() bool {
	if o != nil && o.BillingEmail.IsSet() {
		return true
	}

	return false
}

// SetBillingEmail gets a reference to the given NullableString and assigns it to the BillingEmail field.
func (o *CustomerExternal) SetBillingEmail(v string) {
	o.BillingEmail.Set(&v)
}
// SetBillingEmailNil sets the value for BillingEmail to be an explicit nil
func (o *CustomerExternal) SetBillingEmailNil() {
	o.BillingEmail.Set(nil)
}

// UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
func (o *CustomerExternal) UnsetBillingEmail() {
	o.BillingEmail.Unset()
}

// GetBusinessName returns the BusinessName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomerExternal) GetBusinessName() string {
	if o == nil || o.BusinessName.Get() == nil {
		var ret string
		return ret
	}

	return *o.BusinessName.Get()
}

// GetBusinessNameOk returns a tuple with the BusinessName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetBusinessNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessName.Get(), o.BusinessName.IsSet()
}

// SetBusinessName sets field value
func (o *CustomerExternal) SetBusinessName(v string) {
	o.BusinessName.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *CustomerExternal) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CustomerExternal) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerExternal) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CustomerExternal) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CustomerExternal) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCustomerBillingAddress returns the CustomerBillingAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerExternal) GetCustomerBillingAddress() CompleteAddress {
	if o == nil || IsNil(o.CustomerBillingAddress.Get()) {
		var ret CompleteAddress
		return ret
	}
	return *o.CustomerBillingAddress.Get()
}

// GetCustomerBillingAddressOk returns a tuple with the CustomerBillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetCustomerBillingAddressOk() (*CompleteAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerBillingAddress.Get(), o.CustomerBillingAddress.IsSet()
}

// HasCustomerBillingAddress returns a boolean if a field has been set.
func (o *CustomerExternal) HasCustomerBillingAddress() bool {
	if o != nil && o.CustomerBillingAddress.IsSet() {
		return true
	}

	return false
}

// SetCustomerBillingAddress gets a reference to the given NullableCompleteAddress and assigns it to the CustomerBillingAddress field.
func (o *CustomerExternal) SetCustomerBillingAddress(v CompleteAddress) {
	o.CustomerBillingAddress.Set(&v)
}
// SetCustomerBillingAddressNil sets the value for CustomerBillingAddress to be an explicit nil
func (o *CustomerExternal) SetCustomerBillingAddressNil() {
	o.CustomerBillingAddress.Set(nil)
}

// UnsetCustomerBillingAddress ensures that no value is present for CustomerBillingAddress, not even an explicit nil
func (o *CustomerExternal) UnsetCustomerBillingAddress() {
	o.CustomerBillingAddress.Unset()
}

// GetDiscount returns the Discount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerExternal) GetDiscount() DiscountExternal {
	if o == nil || IsNil(o.Discount.Get()) {
		var ret DiscountExternal
		return ret
	}
	return *o.Discount.Get()
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetDiscountOk() (*DiscountExternal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Discount.Get(), o.Discount.IsSet()
}

// HasDiscount returns a boolean if a field has been set.
func (o *CustomerExternal) HasDiscount() bool {
	if o != nil && o.Discount.IsSet() {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given NullableDiscountExternal and assigns it to the Discount field.
func (o *CustomerExternal) SetDiscount(v DiscountExternal) {
	o.Discount.Set(&v)
}
// SetDiscountNil sets the value for Discount to be an explicit nil
func (o *CustomerExternal) SetDiscountNil() {
	o.Discount.Set(nil)
}

// UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
func (o *CustomerExternal) UnsetDiscount() {
	o.Discount.Unset()
}

// GetEmail returns the Email field value
func (o *CustomerExternal) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CustomerExternal) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomerExternal) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}

	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// SetFirstName sets field value
func (o *CustomerExternal) SetFirstName(v string) {
	o.FirstName.Set(&v)
}

// GetId returns the Id field value
func (o *CustomerExternal) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerExternal) SetId(v string) {
	o.Id = v
}

// GetInvoiceSettings returns the InvoiceSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerExternal) GetInvoiceSettings() CustomerInvoiceSettings {
	if o == nil || IsNil(o.InvoiceSettings.Get()) {
		var ret CustomerInvoiceSettings
		return ret
	}
	return *o.InvoiceSettings.Get()
}

// GetInvoiceSettingsOk returns a tuple with the InvoiceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetInvoiceSettingsOk() (*CustomerInvoiceSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceSettings.Get(), o.InvoiceSettings.IsSet()
}

// HasInvoiceSettings returns a boolean if a field has been set.
func (o *CustomerExternal) HasInvoiceSettings() bool {
	if o != nil && o.InvoiceSettings.IsSet() {
		return true
	}

	return false
}

// SetInvoiceSettings gets a reference to the given NullableCustomerInvoiceSettings and assigns it to the InvoiceSettings field.
func (o *CustomerExternal) SetInvoiceSettings(v CustomerInvoiceSettings) {
	o.InvoiceSettings.Set(&v)
}
// SetInvoiceSettingsNil sets the value for InvoiceSettings to be an explicit nil
func (o *CustomerExternal) SetInvoiceSettingsNil() {
	o.InvoiceSettings.Set(nil)
}

// UnsetInvoiceSettings ensures that no value is present for InvoiceSettings, not even an explicit nil
func (o *CustomerExternal) UnsetInvoiceSettings() {
	o.InvoiceSettings.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *CustomerExternal) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *CustomerExternal) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *CustomerExternal) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *CustomerExternal) GetLanguage() CustomerLanguage {
	if o == nil || IsNil(o.Language) {
		var ret CustomerLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetLanguageOk() (*CustomerLanguage, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *CustomerExternal) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given CustomerLanguage and assigns it to the Language field.
func (o *CustomerExternal) SetLanguage(v CustomerLanguage) {
	o.Language = &v
}

// GetLastName returns the LastName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomerExternal) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// SetLastName sets field value
func (o *CustomerExternal) SetLastName(v string) {
	o.LastName.Set(&v)
}

// GetLastSuccessfulPaymentIntent returns the LastSuccessfulPaymentIntent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerExternal) GetLastSuccessfulPaymentIntent() PaymentIntentExternal {
	if o == nil || IsNil(o.LastSuccessfulPaymentIntent.Get()) {
		var ret PaymentIntentExternal
		return ret
	}
	return *o.LastSuccessfulPaymentIntent.Get()
}

// GetLastSuccessfulPaymentIntentOk returns a tuple with the LastSuccessfulPaymentIntent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetLastSuccessfulPaymentIntentOk() (*PaymentIntentExternal, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSuccessfulPaymentIntent.Get(), o.LastSuccessfulPaymentIntent.IsSet()
}

// HasLastSuccessfulPaymentIntent returns a boolean if a field has been set.
func (o *CustomerExternal) HasLastSuccessfulPaymentIntent() bool {
	if o != nil && o.LastSuccessfulPaymentIntent.IsSet() {
		return true
	}

	return false
}

// SetLastSuccessfulPaymentIntent gets a reference to the given NullablePaymentIntentExternal and assigns it to the LastSuccessfulPaymentIntent field.
func (o *CustomerExternal) SetLastSuccessfulPaymentIntent(v PaymentIntentExternal) {
	o.LastSuccessfulPaymentIntent.Set(&v)
}
// SetLastSuccessfulPaymentIntentNil sets the value for LastSuccessfulPaymentIntent to be an explicit nil
func (o *CustomerExternal) SetLastSuccessfulPaymentIntentNil() {
	o.LastSuccessfulPaymentIntent.Set(nil)
}

// UnsetLastSuccessfulPaymentIntent ensures that no value is present for LastSuccessfulPaymentIntent, not even an explicit nil
func (o *CustomerExternal) UnsetLastSuccessfulPaymentIntent() {
	o.LastSuccessfulPaymentIntent.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerExternal) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CustomerExternal) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CustomerExternal) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMrr returns the Mrr field value if set, zero value otherwise.
func (o *CustomerExternal) GetMrr() []CustomerTotalAmount {
	if o == nil || IsNil(o.Mrr) {
		var ret []CustomerTotalAmount
		return ret
	}
	return o.Mrr
}

// GetMrrOk returns a tuple with the Mrr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetMrrOk() ([]CustomerTotalAmount, bool) {
	if o == nil || IsNil(o.Mrr) {
		return nil, false
	}
	return o.Mrr, true
}

// HasMrr returns a boolean if a field has been set.
func (o *CustomerExternal) HasMrr() bool {
	if o != nil && !IsNil(o.Mrr) {
		return true
	}

	return false
}

// SetMrr gets a reference to the given []CustomerTotalAmount and assigns it to the Mrr field.
func (o *CustomerExternal) SetMrr(v []CustomerTotalAmount) {
	o.Mrr = v
}

// GetNotes returns the Notes field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomerExternal) GetNotes() string {
	if o == nil || o.Notes.Get() == nil {
		var ret string
		return ret
	}

	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// SetNotes sets field value
func (o *CustomerExternal) SetNotes(v string) {
	o.Notes.Set(&v)
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CustomerExternal) GetObject() ObjectName {
	if o == nil || IsNil(o.Object) {
		var ret ObjectName
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetObjectOk() (*ObjectName, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CustomerExternal) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given ObjectName and assigns it to the Object field.
func (o *CustomerExternal) SetObject(v ObjectName) {
	o.Object = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerExternal) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber.Get()) {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CustomerExternal) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *CustomerExternal) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *CustomerExternal) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *CustomerExternal) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetShippingAddresses returns the ShippingAddresses field value if set, zero value otherwise.
func (o *CustomerExternal) GetShippingAddresses() []CompleteAddress {
	if o == nil || IsNil(o.ShippingAddresses) {
		var ret []CompleteAddress
		return ret
	}
	return o.ShippingAddresses
}

// GetShippingAddressesOk returns a tuple with the ShippingAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetShippingAddressesOk() ([]CompleteAddress, bool) {
	if o == nil || IsNil(o.ShippingAddresses) {
		return nil, false
	}
	return o.ShippingAddresses, true
}

// HasShippingAddresses returns a boolean if a field has been set.
func (o *CustomerExternal) HasShippingAddresses() bool {
	if o != nil && !IsNil(o.ShippingAddresses) {
		return true
	}

	return false
}

// SetShippingAddresses gets a reference to the given []CompleteAddress and assigns it to the ShippingAddresses field.
func (o *CustomerExternal) SetShippingAddresses(v []CompleteAddress) {
	o.ShippingAddresses = v
}

// GetShouldSendPaymentReceipt returns the ShouldSendPaymentReceipt field value
func (o *CustomerExternal) GetShouldSendPaymentReceipt() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShouldSendPaymentReceipt
}

// GetShouldSendPaymentReceiptOk returns a tuple with the ShouldSendPaymentReceipt field value
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetShouldSendPaymentReceiptOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShouldSendPaymentReceipt, true
}

// SetShouldSendPaymentReceipt sets field value
func (o *CustomerExternal) SetShouldSendPaymentReceipt(v bool) {
	o.ShouldSendPaymentReceipt = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerExternal) GetStatus() CustomerStatus {
	if o == nil || IsNil(o.Status.Get()) {
		var ret CustomerStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerExternal) GetStatusOk() (*CustomerStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomerExternal) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableCustomerStatus and assigns it to the Status field.
func (o *CustomerExternal) SetStatus(v CustomerStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *CustomerExternal) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *CustomerExternal) UnsetStatus() {
	o.Status.Unset()
}

// GetSubscribedToProducts returns the SubscribedToProducts field value if set, zero value otherwise.
func (o *CustomerExternal) GetSubscribedToProducts() []ProductExternal {
	if o == nil || IsNil(o.SubscribedToProducts) {
		var ret []ProductExternal
		return ret
	}
	return o.SubscribedToProducts
}

// GetSubscribedToProductsOk returns a tuple with the SubscribedToProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetSubscribedToProductsOk() ([]ProductExternal, bool) {
	if o == nil || IsNil(o.SubscribedToProducts) {
		return nil, false
	}
	return o.SubscribedToProducts, true
}

// HasSubscribedToProducts returns a boolean if a field has been set.
func (o *CustomerExternal) HasSubscribedToProducts() bool {
	if o != nil && !IsNil(o.SubscribedToProducts) {
		return true
	}

	return false
}

// SetSubscribedToProducts gets a reference to the given []ProductExternal and assigns it to the SubscribedToProducts field.
func (o *CustomerExternal) SetSubscribedToProducts(v []ProductExternal) {
	o.SubscribedToProducts = v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *CustomerExternal) GetSubscriptions() []SubscriptionExternal {
	if o == nil || IsNil(o.Subscriptions) {
		var ret []SubscriptionExternal
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetSubscriptionsOk() ([]SubscriptionExternal, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *CustomerExternal) HasSubscriptions() bool {
	if o != nil && !IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []SubscriptionExternal and assigns it to the Subscriptions field.
func (o *CustomerExternal) SetSubscriptions(v []SubscriptionExternal) {
	o.Subscriptions = v
}

// GetTotalRefunds returns the TotalRefunds field value if set, zero value otherwise.
func (o *CustomerExternal) GetTotalRefunds() []CustomerTotalAmount {
	if o == nil || IsNil(o.TotalRefunds) {
		var ret []CustomerTotalAmount
		return ret
	}
	return o.TotalRefunds
}

// GetTotalRefundsOk returns a tuple with the TotalRefunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetTotalRefundsOk() ([]CustomerTotalAmount, bool) {
	if o == nil || IsNil(o.TotalRefunds) {
		return nil, false
	}
	return o.TotalRefunds, true
}

// HasTotalRefunds returns a boolean if a field has been set.
func (o *CustomerExternal) HasTotalRefunds() bool {
	if o != nil && !IsNil(o.TotalRefunds) {
		return true
	}

	return false
}

// SetTotalRefunds gets a reference to the given []CustomerTotalAmount and assigns it to the TotalRefunds field.
func (o *CustomerExternal) SetTotalRefunds(v []CustomerTotalAmount) {
	o.TotalRefunds = v
}

// GetTotalSpent returns the TotalSpent field value if set, zero value otherwise.
func (o *CustomerExternal) GetTotalSpent() []CustomerTotalAmount {
	if o == nil || IsNil(o.TotalSpent) {
		var ret []CustomerTotalAmount
		return ret
	}
	return o.TotalSpent
}

// GetTotalSpentOk returns a tuple with the TotalSpent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetTotalSpentOk() ([]CustomerTotalAmount, bool) {
	if o == nil || IsNil(o.TotalSpent) {
		return nil, false
	}
	return o.TotalSpent, true
}

// HasTotalSpent returns a boolean if a field has been set.
func (o *CustomerExternal) HasTotalSpent() bool {
	if o != nil && !IsNil(o.TotalSpent) {
		return true
	}

	return false
}

// SetTotalSpent gets a reference to the given []CustomerTotalAmount and assigns it to the TotalSpent field.
func (o *CustomerExternal) SetTotalSpent(v []CustomerTotalAmount) {
	o.TotalSpent = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CustomerExternal) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomerExternal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CustomerExternal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o CustomerExternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerExternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["address"] = o.Address.Get()
	if !IsNil(o.BalanceAtoms) {
		toSerialize["balance_atoms"] = o.BalanceAtoms
	}
	if o.BillingEmail.IsSet() {
		toSerialize["billing_email"] = o.BillingEmail.Get()
	}
	toSerialize["business_name"] = o.BusinessName.Get()
	toSerialize["created_at"] = o.CreatedAt
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.CustomerBillingAddress.IsSet() {
		toSerialize["customer_billing_address"] = o.CustomerBillingAddress.Get()
	}
	if o.Discount.IsSet() {
		toSerialize["discount"] = o.Discount.Get()
	}
	toSerialize["email"] = o.Email
	toSerialize["first_name"] = o.FirstName.Get()
	toSerialize["id"] = o.Id
	if o.InvoiceSettings.IsSet() {
		toSerialize["invoice_settings"] = o.InvoiceSettings.Get()
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	toSerialize["last_name"] = o.LastName.Get()
	if o.LastSuccessfulPaymentIntent.IsSet() {
		toSerialize["last_successful_payment_intent"] = o.LastSuccessfulPaymentIntent.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Mrr) {
		toSerialize["mrr"] = o.Mrr
	}
	toSerialize["notes"] = o.Notes.Get()
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if !IsNil(o.ShippingAddresses) {
		toSerialize["shipping_addresses"] = o.ShippingAddresses
	}
	toSerialize["should_send_payment_receipt"] = o.ShouldSendPaymentReceipt
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !IsNil(o.SubscribedToProducts) {
		toSerialize["subscribed_to_products"] = o.SubscribedToProducts
	}
	if !IsNil(o.Subscriptions) {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if !IsNil(o.TotalRefunds) {
		toSerialize["total_refunds"] = o.TotalRefunds
	}
	if !IsNil(o.TotalSpent) {
		toSerialize["total_spent"] = o.TotalSpent
	}
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerExternal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"address",
		"business_name",
		"created_at",
		"email",
		"first_name",
		"id",
		"last_name",
		"notes",
		"should_send_payment_receipt",
		"updated_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCustomerExternal := _CustomerExternal{}

	err = json.Unmarshal(data, &varCustomerExternal)

	if err != nil {
		return err
	}

	*o = CustomerExternal(varCustomerExternal)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "address")
		delete(additionalProperties, "balance_atoms")
		delete(additionalProperties, "billing_email")
		delete(additionalProperties, "business_name")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "customer_billing_address")
		delete(additionalProperties, "discount")
		delete(additionalProperties, "email")
		delete(additionalProperties, "first_name")
		delete(additionalProperties, "id")
		delete(additionalProperties, "invoice_settings")
		delete(additionalProperties, "is_deleted")
		delete(additionalProperties, "language")
		delete(additionalProperties, "last_name")
		delete(additionalProperties, "last_successful_payment_intent")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "mrr")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "object")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "shipping_addresses")
		delete(additionalProperties, "should_send_payment_receipt")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subscribed_to_products")
		delete(additionalProperties, "subscriptions")
		delete(additionalProperties, "total_refunds")
		delete(additionalProperties, "total_spent")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerExternal struct {
	value *CustomerExternal
	isSet bool
}

func (v NullableCustomerExternal) Get() *CustomerExternal {
	return v.value
}

func (v *NullableCustomerExternal) Set(val *CustomerExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerExternal(val *CustomerExternal) *NullableCustomerExternal {
	return &NullableCustomerExternal{value: val, isSet: true}
}

func (v NullableCustomerExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


