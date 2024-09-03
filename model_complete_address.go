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

// checks if the CompleteAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompleteAddress{}

// CompleteAddress struct for CompleteAddress
type CompleteAddress struct {
	Line1 NullableString `json:"line1,omitempty"`
	Line2 NullableString `json:"line2,omitempty"`
	Line3 NullableString `json:"line3,omitempty"`
	City NullableString `json:"city,omitempty"`
	State NullableString `json:"state,omitempty"`
	ZipCode NullableString `json:"zip_code,omitempty"`
	Country NullableString `json:"country,omitempty"`
}

// NewCompleteAddress instantiates a new CompleteAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteAddress() *CompleteAddress {
	this := CompleteAddress{}
	return &this
}

// NewCompleteAddressWithDefaults instantiates a new CompleteAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteAddressWithDefaults() *CompleteAddress {
	this := CompleteAddress{}
	return &this
}

// GetLine1 returns the Line1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompleteAddress) GetLine1() string {
	if o == nil || IsNil(o.Line1.Get()) {
		var ret string
		return ret
	}
	return *o.Line1.Get()
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompleteAddress) GetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line1.Get(), o.Line1.IsSet()
}

// HasLine1 returns a boolean if a field has been set.
func (o *CompleteAddress) HasLine1() bool {
	if o != nil && o.Line1.IsSet() {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given NullableString and assigns it to the Line1 field.
func (o *CompleteAddress) SetLine1(v string) {
	o.Line1.Set(&v)
}
// SetLine1Nil sets the value for Line1 to be an explicit nil
func (o *CompleteAddress) SetLine1Nil() {
	o.Line1.Set(nil)
}

// UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
func (o *CompleteAddress) UnsetLine1() {
	o.Line1.Unset()
}

// GetLine2 returns the Line2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompleteAddress) GetLine2() string {
	if o == nil || IsNil(o.Line2.Get()) {
		var ret string
		return ret
	}
	return *o.Line2.Get()
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompleteAddress) GetLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line2.Get(), o.Line2.IsSet()
}

// HasLine2 returns a boolean if a field has been set.
func (o *CompleteAddress) HasLine2() bool {
	if o != nil && o.Line2.IsSet() {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given NullableString and assigns it to the Line2 field.
func (o *CompleteAddress) SetLine2(v string) {
	o.Line2.Set(&v)
}
// SetLine2Nil sets the value for Line2 to be an explicit nil
func (o *CompleteAddress) SetLine2Nil() {
	o.Line2.Set(nil)
}

// UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
func (o *CompleteAddress) UnsetLine2() {
	o.Line2.Unset()
}

// GetLine3 returns the Line3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompleteAddress) GetLine3() string {
	if o == nil || IsNil(o.Line3.Get()) {
		var ret string
		return ret
	}
	return *o.Line3.Get()
}

// GetLine3Ok returns a tuple with the Line3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompleteAddress) GetLine3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line3.Get(), o.Line3.IsSet()
}

// HasLine3 returns a boolean if a field has been set.
func (o *CompleteAddress) HasLine3() bool {
	if o != nil && o.Line3.IsSet() {
		return true
	}

	return false
}

// SetLine3 gets a reference to the given NullableString and assigns it to the Line3 field.
func (o *CompleteAddress) SetLine3(v string) {
	o.Line3.Set(&v)
}
// SetLine3Nil sets the value for Line3 to be an explicit nil
func (o *CompleteAddress) SetLine3Nil() {
	o.Line3.Set(nil)
}

// UnsetLine3 ensures that no value is present for Line3, not even an explicit nil
func (o *CompleteAddress) UnsetLine3() {
	o.Line3.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompleteAddress) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompleteAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *CompleteAddress) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *CompleteAddress) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *CompleteAddress) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *CompleteAddress) UnsetCity() {
	o.City.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompleteAddress) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompleteAddress) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *CompleteAddress) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *CompleteAddress) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *CompleteAddress) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *CompleteAddress) UnsetState() {
	o.State.Unset()
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompleteAddress) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode.Get()) {
		var ret string
		return ret
	}
	return *o.ZipCode.Get()
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompleteAddress) GetZipCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZipCode.Get(), o.ZipCode.IsSet()
}

// HasZipCode returns a boolean if a field has been set.
func (o *CompleteAddress) HasZipCode() bool {
	if o != nil && o.ZipCode.IsSet() {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given NullableString and assigns it to the ZipCode field.
func (o *CompleteAddress) SetZipCode(v string) {
	o.ZipCode.Set(&v)
}
// SetZipCodeNil sets the value for ZipCode to be an explicit nil
func (o *CompleteAddress) SetZipCodeNil() {
	o.ZipCode.Set(nil)
}

// UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
func (o *CompleteAddress) UnsetZipCode() {
	o.ZipCode.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompleteAddress) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompleteAddress) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *CompleteAddress) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *CompleteAddress) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *CompleteAddress) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *CompleteAddress) UnsetCountry() {
	o.Country.Unset()
}

func (o CompleteAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompleteAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Line1.IsSet() {
		toSerialize["line1"] = o.Line1.Get()
	}
	if o.Line2.IsSet() {
		toSerialize["line2"] = o.Line2.Get()
	}
	if o.Line3.IsSet() {
		toSerialize["line3"] = o.Line3.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.ZipCode.IsSet() {
		toSerialize["zip_code"] = o.ZipCode.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	return toSerialize, nil
}

type NullableCompleteAddress struct {
	value *CompleteAddress
	isSet bool
}

func (v NullableCompleteAddress) Get() *CompleteAddress {
	return v.value
}

func (v *NullableCompleteAddress) Set(val *CompleteAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteAddress(val *CompleteAddress) *NullableCompleteAddress {
	return &NullableCompleteAddress{value: val, isSet: true}
}

func (v NullableCompleteAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

