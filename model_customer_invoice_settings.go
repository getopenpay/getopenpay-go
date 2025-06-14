/*
OpenPay API

super charge your subscription management.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
)

// checks if the CustomerInvoiceSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerInvoiceSettings{}

// CustomerInvoiceSettings struct for CustomerInvoiceSettings
type CustomerInvoiceSettings struct {
	// The tax IDs of the customer.
	CustomerTaxIds []TaxIdSetting `json:"customer_tax_ids,omitempty"`
	DefaultNetD NullableInt32 `json:"default_net_d,omitempty"`
	EmailReceiptOnPaid NullableBool `json:"email_receipt_on_paid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerInvoiceSettings CustomerInvoiceSettings

// NewCustomerInvoiceSettings instantiates a new CustomerInvoiceSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerInvoiceSettings() *CustomerInvoiceSettings {
	this := CustomerInvoiceSettings{}
	return &this
}

// NewCustomerInvoiceSettingsWithDefaults instantiates a new CustomerInvoiceSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerInvoiceSettingsWithDefaults() *CustomerInvoiceSettings {
	this := CustomerInvoiceSettings{}
	return &this
}

// GetCustomerTaxIds returns the CustomerTaxIds field value if set, zero value otherwise.
func (o *CustomerInvoiceSettings) GetCustomerTaxIds() []TaxIdSetting {
	if o == nil || IsNil(o.CustomerTaxIds) {
		var ret []TaxIdSetting
		return ret
	}
	return o.CustomerTaxIds
}

// GetCustomerTaxIdsOk returns a tuple with the CustomerTaxIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInvoiceSettings) GetCustomerTaxIdsOk() ([]TaxIdSetting, bool) {
	if o == nil || IsNil(o.CustomerTaxIds) {
		return nil, false
	}
	return o.CustomerTaxIds, true
}

// HasCustomerTaxIds returns a boolean if a field has been set.
func (o *CustomerInvoiceSettings) HasCustomerTaxIds() bool {
	if o != nil && !IsNil(o.CustomerTaxIds) {
		return true
	}

	return false
}

// SetCustomerTaxIds gets a reference to the given []TaxIdSetting and assigns it to the CustomerTaxIds field.
func (o *CustomerInvoiceSettings) SetCustomerTaxIds(v []TaxIdSetting) {
	o.CustomerTaxIds = v
}

// GetDefaultNetD returns the DefaultNetD field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerInvoiceSettings) GetDefaultNetD() int32 {
	if o == nil || IsNil(o.DefaultNetD.Get()) {
		var ret int32
		return ret
	}
	return *o.DefaultNetD.Get()
}

// GetDefaultNetDOk returns a tuple with the DefaultNetD field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerInvoiceSettings) GetDefaultNetDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultNetD.Get(), o.DefaultNetD.IsSet()
}

// HasDefaultNetD returns a boolean if a field has been set.
func (o *CustomerInvoiceSettings) HasDefaultNetD() bool {
	if o != nil && o.DefaultNetD.IsSet() {
		return true
	}

	return false
}

// SetDefaultNetD gets a reference to the given NullableInt32 and assigns it to the DefaultNetD field.
func (o *CustomerInvoiceSettings) SetDefaultNetD(v int32) {
	o.DefaultNetD.Set(&v)
}
// SetDefaultNetDNil sets the value for DefaultNetD to be an explicit nil
func (o *CustomerInvoiceSettings) SetDefaultNetDNil() {
	o.DefaultNetD.Set(nil)
}

// UnsetDefaultNetD ensures that no value is present for DefaultNetD, not even an explicit nil
func (o *CustomerInvoiceSettings) UnsetDefaultNetD() {
	o.DefaultNetD.Unset()
}

// GetEmailReceiptOnPaid returns the EmailReceiptOnPaid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerInvoiceSettings) GetEmailReceiptOnPaid() bool {
	if o == nil || IsNil(o.EmailReceiptOnPaid.Get()) {
		var ret bool
		return ret
	}
	return *o.EmailReceiptOnPaid.Get()
}

// GetEmailReceiptOnPaidOk returns a tuple with the EmailReceiptOnPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerInvoiceSettings) GetEmailReceiptOnPaidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailReceiptOnPaid.Get(), o.EmailReceiptOnPaid.IsSet()
}

// HasEmailReceiptOnPaid returns a boolean if a field has been set.
func (o *CustomerInvoiceSettings) HasEmailReceiptOnPaid() bool {
	if o != nil && o.EmailReceiptOnPaid.IsSet() {
		return true
	}

	return false
}

// SetEmailReceiptOnPaid gets a reference to the given NullableBool and assigns it to the EmailReceiptOnPaid field.
func (o *CustomerInvoiceSettings) SetEmailReceiptOnPaid(v bool) {
	o.EmailReceiptOnPaid.Set(&v)
}
// SetEmailReceiptOnPaidNil sets the value for EmailReceiptOnPaid to be an explicit nil
func (o *CustomerInvoiceSettings) SetEmailReceiptOnPaidNil() {
	o.EmailReceiptOnPaid.Set(nil)
}

// UnsetEmailReceiptOnPaid ensures that no value is present for EmailReceiptOnPaid, not even an explicit nil
func (o *CustomerInvoiceSettings) UnsetEmailReceiptOnPaid() {
	o.EmailReceiptOnPaid.Unset()
}

func (o CustomerInvoiceSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerInvoiceSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerTaxIds) {
		toSerialize["customer_tax_ids"] = o.CustomerTaxIds
	}
	if o.DefaultNetD.IsSet() {
		toSerialize["default_net_d"] = o.DefaultNetD.Get()
	}
	if o.EmailReceiptOnPaid.IsSet() {
		toSerialize["email_receipt_on_paid"] = o.EmailReceiptOnPaid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerInvoiceSettings) UnmarshalJSON(data []byte) (err error) {
	varCustomerInvoiceSettings := _CustomerInvoiceSettings{}

	err = json.Unmarshal(data, &varCustomerInvoiceSettings)

	if err != nil {
		return err
	}

	*o = CustomerInvoiceSettings(varCustomerInvoiceSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_tax_ids")
		delete(additionalProperties, "default_net_d")
		delete(additionalProperties, "email_receipt_on_paid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerInvoiceSettings struct {
	value *CustomerInvoiceSettings
	isSet bool
}

func (v NullableCustomerInvoiceSettings) Get() *CustomerInvoiceSettings {
	return v.value
}

func (v *NullableCustomerInvoiceSettings) Set(val *CustomerInvoiceSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerInvoiceSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerInvoiceSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerInvoiceSettings(val *CustomerInvoiceSettings) *NullableCustomerInvoiceSettings {
	return &NullableCustomerInvoiceSettings{value: val, isSet: true}
}

func (v NullableCustomerInvoiceSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerInvoiceSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


