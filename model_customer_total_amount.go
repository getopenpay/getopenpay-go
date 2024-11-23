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

// checks if the CustomerTotalAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerTotalAmount{}

// CustomerTotalAmount struct for CustomerTotalAmount
type CustomerTotalAmount struct {
	Currency CurrencyEnum `json:"currency"`
	// The total refunds/spent of the customer.
	AmountAtom *int32 `json:"amount_atom,omitempty"`
}

type _CustomerTotalAmount CustomerTotalAmount

// NewCustomerTotalAmount instantiates a new CustomerTotalAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerTotalAmount(currency CurrencyEnum) *CustomerTotalAmount {
	this := CustomerTotalAmount{}
	this.Currency = currency
	var amountAtom int32 = 0
	this.AmountAtom = &amountAtom
	return &this
}

// NewCustomerTotalAmountWithDefaults instantiates a new CustomerTotalAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerTotalAmountWithDefaults() *CustomerTotalAmount {
	this := CustomerTotalAmount{}
	var amountAtom int32 = 0
	this.AmountAtom = &amountAtom
	return &this
}

// GetCurrency returns the Currency field value
func (o *CustomerTotalAmount) GetCurrency() CurrencyEnum {
	if o == nil {
		var ret CurrencyEnum
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CustomerTotalAmount) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CustomerTotalAmount) SetCurrency(v CurrencyEnum) {
	o.Currency = v
}

// GetAmountAtom returns the AmountAtom field value if set, zero value otherwise.
func (o *CustomerTotalAmount) GetAmountAtom() int32 {
	if o == nil || IsNil(o.AmountAtom) {
		var ret int32
		return ret
	}
	return *o.AmountAtom
}

// GetAmountAtomOk returns a tuple with the AmountAtom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTotalAmount) GetAmountAtomOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountAtom) {
		return nil, false
	}
	return o.AmountAtom, true
}

// HasAmountAtom returns a boolean if a field has been set.
func (o *CustomerTotalAmount) HasAmountAtom() bool {
	if o != nil && !IsNil(o.AmountAtom) {
		return true
	}

	return false
}

// SetAmountAtom gets a reference to the given int32 and assigns it to the AmountAtom field.
func (o *CustomerTotalAmount) SetAmountAtom(v int32) {
	o.AmountAtom = &v
}

func (o CustomerTotalAmount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerTotalAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.AmountAtom) {
		toSerialize["amount_atom"] = o.AmountAtom
	}
	return toSerialize, nil
}

func (o *CustomerTotalAmount) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
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

	varCustomerTotalAmount := _CustomerTotalAmount{}

	err = json.Unmarshal(bytes, &varCustomerTotalAmount)

	if err != nil {
		return err
	}

	*o = CustomerTotalAmount(varCustomerTotalAmount)

	return err
}

type NullableCustomerTotalAmount struct {
	value *CustomerTotalAmount
	isSet bool
}

func (v NullableCustomerTotalAmount) Get() *CustomerTotalAmount {
	return v.value
}

func (v *NullableCustomerTotalAmount) Set(val *CustomerTotalAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerTotalAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerTotalAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerTotalAmount(val *CustomerTotalAmount) *NullableCustomerTotalAmount {
	return &NullableCustomerTotalAmount{value: val, isSet: true}
}

func (v NullableCustomerTotalAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerTotalAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


