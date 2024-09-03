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

// checks if the InvoiceItemDiscountAmountsExternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceItemDiscountAmountsExternal{}

// InvoiceItemDiscountAmountsExternal struct for InvoiceItemDiscountAmountsExternal
type InvoiceItemDiscountAmountsExternal struct {
	Object *ObjectName `json:"object,omitempty"`
	// Id of the discount that was applied to get this discount amount.
	DiscountId string `json:"discount_id"`
	// The amount_atom, of the discount.
	AmountAtom int32 `json:"amount_atom"`
}

type _InvoiceItemDiscountAmountsExternal InvoiceItemDiscountAmountsExternal

// NewInvoiceItemDiscountAmountsExternal instantiates a new InvoiceItemDiscountAmountsExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceItemDiscountAmountsExternal(discountId string, amountAtom int32) *InvoiceItemDiscountAmountsExternal {
	this := InvoiceItemDiscountAmountsExternal{}
	var object ObjectName = OBJECTNAME_INVOICE_ITEM_DISCOUNT
	this.Object = &object
	this.DiscountId = discountId
	this.AmountAtom = amountAtom
	return &this
}

// NewInvoiceItemDiscountAmountsExternalWithDefaults instantiates a new InvoiceItemDiscountAmountsExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceItemDiscountAmountsExternalWithDefaults() *InvoiceItemDiscountAmountsExternal {
	this := InvoiceItemDiscountAmountsExternal{}
	var object ObjectName = OBJECTNAME_INVOICE_ITEM_DISCOUNT
	this.Object = &object
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *InvoiceItemDiscountAmountsExternal) GetObject() ObjectName {
	if o == nil || IsNil(o.Object) {
		var ret ObjectName
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItemDiscountAmountsExternal) GetObjectOk() (*ObjectName, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *InvoiceItemDiscountAmountsExternal) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given ObjectName and assigns it to the Object field.
func (o *InvoiceItemDiscountAmountsExternal) SetObject(v ObjectName) {
	o.Object = &v
}

// GetDiscountId returns the DiscountId field value
func (o *InvoiceItemDiscountAmountsExternal) GetDiscountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiscountId
}

// GetDiscountIdOk returns a tuple with the DiscountId field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemDiscountAmountsExternal) GetDiscountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountId, true
}

// SetDiscountId sets field value
func (o *InvoiceItemDiscountAmountsExternal) SetDiscountId(v string) {
	o.DiscountId = v
}

// GetAmountAtom returns the AmountAtom field value
func (o *InvoiceItemDiscountAmountsExternal) GetAmountAtom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountAtom
}

// GetAmountAtomOk returns a tuple with the AmountAtom field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemDiscountAmountsExternal) GetAmountAtomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountAtom, true
}

// SetAmountAtom sets field value
func (o *InvoiceItemDiscountAmountsExternal) SetAmountAtom(v int32) {
	o.AmountAtom = v
}

func (o InvoiceItemDiscountAmountsExternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceItemDiscountAmountsExternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	toSerialize["discount_id"] = o.DiscountId
	toSerialize["amount_atom"] = o.AmountAtom
	return toSerialize, nil
}

func (o *InvoiceItemDiscountAmountsExternal) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"discount_id",
		"amount_atom",
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

	varInvoiceItemDiscountAmountsExternal := _InvoiceItemDiscountAmountsExternal{}

	err = json.Unmarshal(bytes, &varInvoiceItemDiscountAmountsExternal)

	if err != nil {
		return err
	}

	*o = InvoiceItemDiscountAmountsExternal(varInvoiceItemDiscountAmountsExternal)

	return err
}

type NullableInvoiceItemDiscountAmountsExternal struct {
	value *InvoiceItemDiscountAmountsExternal
	isSet bool
}

func (v NullableInvoiceItemDiscountAmountsExternal) Get() *InvoiceItemDiscountAmountsExternal {
	return v.value
}

func (v *NullableInvoiceItemDiscountAmountsExternal) Set(val *InvoiceItemDiscountAmountsExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceItemDiscountAmountsExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceItemDiscountAmountsExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceItemDiscountAmountsExternal(val *InvoiceItemDiscountAmountsExternal) *NullableInvoiceItemDiscountAmountsExternal {
	return &NullableInvoiceItemDiscountAmountsExternal{value: val, isSet: true}
}

func (v NullableInvoiceItemDiscountAmountsExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceItemDiscountAmountsExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


