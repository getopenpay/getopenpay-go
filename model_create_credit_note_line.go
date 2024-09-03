/*
OpenPay API

super charge your subscription management.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateCreditNoteLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCreditNoteLine{}

// CreateCreditNoteLine struct for CreateCreditNoteLine
type CreateCreditNoteLine struct {
	// The integer amount representing the gross amount being credited for this line item.
	AmountAtom int32 `json:"amount_atom"`
	Currency CurrencyEnum `json:"currency"`
	Type CreditNoteLineType `json:"type"`
	Quantity *int32 `json:"quantity,omitempty"`
	InvoiceItemId NullableString `json:"invoice_item_id,omitempty"`
	UnitAmountAtom NullableInt32 `json:"unit_amount_atom,omitempty"`
}

type _CreateCreditNoteLine CreateCreditNoteLine

// NewCreateCreditNoteLine instantiates a new CreateCreditNoteLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCreditNoteLine(amountAtom int32, currency CurrencyEnum, type_ CreditNoteLineType) *CreateCreditNoteLine {
	this := CreateCreditNoteLine{}
	this.AmountAtom = amountAtom
	this.Currency = currency
	this.Type = type_
	var quantity int32 = 1
	this.Quantity = &quantity
	return &this
}

// NewCreateCreditNoteLineWithDefaults instantiates a new CreateCreditNoteLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCreditNoteLineWithDefaults() *CreateCreditNoteLine {
	this := CreateCreditNoteLine{}
	var quantity int32 = 1
	this.Quantity = &quantity
	return &this
}

// GetAmountAtom returns the AmountAtom field value
func (o *CreateCreditNoteLine) GetAmountAtom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountAtom
}

// GetAmountAtomOk returns a tuple with the AmountAtom field value
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteLine) GetAmountAtomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountAtom, true
}

// SetAmountAtom sets field value
func (o *CreateCreditNoteLine) SetAmountAtom(v int32) {
	o.AmountAtom = v
}

// GetCurrency returns the Currency field value
func (o *CreateCreditNoteLine) GetCurrency() CurrencyEnum {
	if o == nil {
		var ret CurrencyEnum
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteLine) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CreateCreditNoteLine) SetCurrency(v CurrencyEnum) {
	o.Currency = v
}

// GetType returns the Type field value
func (o *CreateCreditNoteLine) GetType() CreditNoteLineType {
	if o == nil {
		var ret CreditNoteLineType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteLine) GetTypeOk() (*CreditNoteLineType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateCreditNoteLine) SetType(v CreditNoteLineType) {
	o.Type = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *CreateCreditNoteLine) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteLine) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *CreateCreditNoteLine) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *CreateCreditNoteLine) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetInvoiceItemId returns the InvoiceItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCreditNoteLine) GetInvoiceItemId() string {
	if o == nil || IsNil(o.InvoiceItemId.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceItemId.Get()
}

// GetInvoiceItemIdOk returns a tuple with the InvoiceItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCreditNoteLine) GetInvoiceItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceItemId.Get(), o.InvoiceItemId.IsSet()
}

// HasInvoiceItemId returns a boolean if a field has been set.
func (o *CreateCreditNoteLine) HasInvoiceItemId() bool {
	if o != nil && o.InvoiceItemId.IsSet() {
		return true
	}

	return false
}

// SetInvoiceItemId gets a reference to the given NullableString and assigns it to the InvoiceItemId field.
func (o *CreateCreditNoteLine) SetInvoiceItemId(v string) {
	o.InvoiceItemId.Set(&v)
}
// SetInvoiceItemIdNil sets the value for InvoiceItemId to be an explicit nil
func (o *CreateCreditNoteLine) SetInvoiceItemIdNil() {
	o.InvoiceItemId.Set(nil)
}

// UnsetInvoiceItemId ensures that no value is present for InvoiceItemId, not even an explicit nil
func (o *CreateCreditNoteLine) UnsetInvoiceItemId() {
	o.InvoiceItemId.Unset()
}

// GetUnitAmountAtom returns the UnitAmountAtom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCreditNoteLine) GetUnitAmountAtom() int32 {
	if o == nil || IsNil(o.UnitAmountAtom.Get()) {
		var ret int32
		return ret
	}
	return *o.UnitAmountAtom.Get()
}

// GetUnitAmountAtomOk returns a tuple with the UnitAmountAtom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCreditNoteLine) GetUnitAmountAtomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitAmountAtom.Get(), o.UnitAmountAtom.IsSet()
}

// HasUnitAmountAtom returns a boolean if a field has been set.
func (o *CreateCreditNoteLine) HasUnitAmountAtom() bool {
	if o != nil && o.UnitAmountAtom.IsSet() {
		return true
	}

	return false
}

// SetUnitAmountAtom gets a reference to the given NullableInt32 and assigns it to the UnitAmountAtom field.
func (o *CreateCreditNoteLine) SetUnitAmountAtom(v int32) {
	o.UnitAmountAtom.Set(&v)
}
// SetUnitAmountAtomNil sets the value for UnitAmountAtom to be an explicit nil
func (o *CreateCreditNoteLine) SetUnitAmountAtomNil() {
	o.UnitAmountAtom.Set(nil)
}

// UnsetUnitAmountAtom ensures that no value is present for UnitAmountAtom, not even an explicit nil
func (o *CreateCreditNoteLine) UnsetUnitAmountAtom() {
	o.UnitAmountAtom.Unset()
}

func (o CreateCreditNoteLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCreditNoteLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount_atom"] = o.AmountAtom
	toSerialize["currency"] = o.Currency
	toSerialize["type"] = o.Type
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if o.InvoiceItemId.IsSet() {
		toSerialize["invoice_item_id"] = o.InvoiceItemId.Get()
	}
	if o.UnitAmountAtom.IsSet() {
		toSerialize["unit_amount_atom"] = o.UnitAmountAtom.Get()
	}
	return toSerialize, nil
}

func (o *CreateCreditNoteLine) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount_atom",
		"currency",
		"type",
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

	varCreateCreditNoteLine := _CreateCreditNoteLine{}

	err = json.Unmarshal(bytes, &varCreateCreditNoteLine)

	if err != nil {
		return err
	}

	*o = CreateCreditNoteLine(varCreateCreditNoteLine)

	return err
}

type NullableCreateCreditNoteLine struct {
	value *CreateCreditNoteLine
	isSet bool
}

func (v NullableCreateCreditNoteLine) Get() *CreateCreditNoteLine {
	return v.value
}

func (v *NullableCreateCreditNoteLine) Set(val *CreateCreditNoteLine) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCreditNoteLine) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCreditNoteLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCreditNoteLine(val *CreateCreditNoteLine) *NullableCreateCreditNoteLine {
	return &NullableCreateCreditNoteLine{value: val, isSet: true}
}

func (v NullableCreateCreditNoteLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCreditNoteLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

