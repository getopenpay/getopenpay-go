/*
OpenPay API

super charge your subscription management.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
)

// checks if the ListActiveSubParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListActiveSubParams{}

// ListActiveSubParams struct for ListActiveSubParams
type ListActiveSubParams struct {
	ProductId NullableString `json:"product_id,omitempty"`
	PriceId NullableString `json:"price_id,omitempty"`
}

// NewListActiveSubParams instantiates a new ListActiveSubParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListActiveSubParams() *ListActiveSubParams {
	this := ListActiveSubParams{}
	return &this
}

// NewListActiveSubParamsWithDefaults instantiates a new ListActiveSubParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListActiveSubParamsWithDefaults() *ListActiveSubParams {
	this := ListActiveSubParams{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListActiveSubParams) GetProductId() string {
	if o == nil || IsNil(o.ProductId.Get()) {
		var ret string
		return ret
	}
	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListActiveSubParams) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// HasProductId returns a boolean if a field has been set.
func (o *ListActiveSubParams) HasProductId() bool {
	if o != nil && o.ProductId.IsSet() {
		return true
	}

	return false
}

// SetProductId gets a reference to the given NullableString and assigns it to the ProductId field.
func (o *ListActiveSubParams) SetProductId(v string) {
	o.ProductId.Set(&v)
}
// SetProductIdNil sets the value for ProductId to be an explicit nil
func (o *ListActiveSubParams) SetProductIdNil() {
	o.ProductId.Set(nil)
}

// UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
func (o *ListActiveSubParams) UnsetProductId() {
	o.ProductId.Unset()
}

// GetPriceId returns the PriceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListActiveSubParams) GetPriceId() string {
	if o == nil || IsNil(o.PriceId.Get()) {
		var ret string
		return ret
	}
	return *o.PriceId.Get()
}

// GetPriceIdOk returns a tuple with the PriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListActiveSubParams) GetPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceId.Get(), o.PriceId.IsSet()
}

// HasPriceId returns a boolean if a field has been set.
func (o *ListActiveSubParams) HasPriceId() bool {
	if o != nil && o.PriceId.IsSet() {
		return true
	}

	return false
}

// SetPriceId gets a reference to the given NullableString and assigns it to the PriceId field.
func (o *ListActiveSubParams) SetPriceId(v string) {
	o.PriceId.Set(&v)
}
// SetPriceIdNil sets the value for PriceId to be an explicit nil
func (o *ListActiveSubParams) SetPriceIdNil() {
	o.PriceId.Set(nil)
}

// UnsetPriceId ensures that no value is present for PriceId, not even an explicit nil
func (o *ListActiveSubParams) UnsetPriceId() {
	o.PriceId.Unset()
}

func (o ListActiveSubParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListActiveSubParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProductId.IsSet() {
		toSerialize["product_id"] = o.ProductId.Get()
	}
	if o.PriceId.IsSet() {
		toSerialize["price_id"] = o.PriceId.Get()
	}
	return toSerialize, nil
}

type NullableListActiveSubParams struct {
	value *ListActiveSubParams
	isSet bool
}

func (v NullableListActiveSubParams) Get() *ListActiveSubParams {
	return v.value
}

func (v *NullableListActiveSubParams) Set(val *ListActiveSubParams) {
	v.value = val
	v.isSet = true
}

func (v NullableListActiveSubParams) IsSet() bool {
	return v.isSet
}

func (v *NullableListActiveSubParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListActiveSubParams(val *ListActiveSubParams) *NullableListActiveSubParams {
	return &NullableListActiveSubParams{value: val, isSet: true}
}

func (v NullableListActiveSubParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListActiveSubParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

